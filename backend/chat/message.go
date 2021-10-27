package chat

import "server/utils"

type Message struct {
	ID     int    `json:"id"`
	Body   string `json:"body"`
	Sender string `json:"sender"`
}

func NewMessage(body string, sender string) *Message {
	return &Message{
		ID:     utils.GetRandom().Int(),
		Body:   body,
		Sender: sender,
	}
}
