package chat

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type User struct {
	Username string
	Conn     *websocket.Conn
	Global   *Chat
}

// read any message
func (u *User) Read() {
	for {
		if _, message, err := u.Conn.ReadMessage(); err != nil {
			log.Printf("Error on read message: %v\n", err.Error())
			break
		} else {
			u.Global.messages <- NewMessage(string(message), u.Username)
		}
	}
	u.Global.leave <- u
}

// write any message to users connected
func (u *User) Write(message *Message) {
	b, _ := json.Marshal(message)
	if err := u.Conn.WriteMessage(websocket.TextMessage, b); err != nil {
		log.Printf("Error on write message: %v\n", err.Error())
	}
}
