package domain

import "encoding/json"

type Message struct {
    MAC    string `json:"mac"`
    Evento string `json:"evento"`
    ID     int    `json:"id"`
}

func NewMessage(msg string) Message {
    var message Message
    json.Unmarshal([]byte(msg), &message)
    return message
}

func ToJSON(message Message) ([]byte, error) {
    return json.Marshal(message)
}