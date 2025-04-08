package domain

import "encoding/json"

type Message struct {
    MAC    string `json:"mac"`
    Evento string `json:"evento"`
}

func ToJSON(m Message) ([]byte, error) {
	return json.Marshal(m)
}

func NewMessage(message string) Message {
    var msg map[string]string
    json.Unmarshal([]byte(message), &msg)
    return Message{
        MAC:    msg["mac"],
        Evento: msg["evento"],
    }
}
