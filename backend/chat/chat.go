package chat

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"server/utils"

	"github.com/gorilla/websocket"
)

type Chat struct {
	users    map[string]*User
	messages chan *Message
	join     chan *User
	leave    chan *User
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  512,
	WriteBufferSize: 512,
	CheckOrigin: func(r *http.Request) bool {
		log.Printf("%s %s%s %v\n", r.Method, r.Host, r.RequestURI, r.Proto)
		return r.Method == http.MethodGet
	},
}

func (c *Chat) Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Error on websocket connection: ", err.Error())
		return
	}

	// get username
	keys := r.URL.Query()
	username := keys.Get("username")

	// verify username
	if strings.TrimSpace(username) == "" {
		username = fmt.Sprintf("anom-%d", utils.GetRandomInt())
	}

	// add new user
	newUser := &User{
		Username: username,
		Conn:     conn,
		Global:   c,
	}
	c.join <- newUser

	// init read user
	newUser.Read()
}

// get channels
func (c *Chat) Run() {
	for {
		select {
		case user := <-c.join:
			c.add(user)
		case message := <-c.messages:
			c.broadcast(message)
		case user := <-c.leave:
			c.disconnect(user)
		}
	}
}

// add new user
func (c *Chat) add(user *User) {
	// add
	if _, ok := c.users[user.Username]; !ok {
		c.users[user.Username] = user
		log.Printf("Added user: %s, Total: %d\n", user.Username, len(c.users))
	}
}

// send all users
func (c *Chat) broadcast(message *Message) {
	log.Printf("Broadcast message: %v\n", message)
	for _, user := range c.users {
		user.Write(message)
	}
}

// handle when is disconnected
func (c *Chat) disconnect(user *User) {
	if _, ok := c.users[user.Username]; ok {
		defer user.Conn.Close()
		delete(c.users, user.Username)
		log.Printf("User left the chat: %s, Total: %d\n", user.Username, len(c.users))
	}
}

// start websocket server
func Start(host string, port string) {
	log.Printf("Chat listening on %s%s\n", host, port)
	c := &Chat{
		users:    make(map[string]*User),
		messages: make(chan *Message),
		join:     make(chan *User),
		leave:    make(chan *User),
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to go Webchat!"))
	})

	http.HandleFunc("/chat", c.Handler)

	go c.Run()

	log.Fatal(http.ListenAndServe(port, nil))
}
