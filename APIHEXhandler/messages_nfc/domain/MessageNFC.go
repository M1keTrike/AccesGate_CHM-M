package domain

import "encoding/json"

type Message struct {
	Uid string `json:"uid"`
}

func ToJSON(m Message) ([]byte, error) {
	return json.Marshal(m)
}

func NewMessage(uid string) Message {
	return Message{
		Uid: uid,
	}
}
