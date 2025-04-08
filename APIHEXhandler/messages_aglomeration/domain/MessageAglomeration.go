package domain

import "encoding/json"

type MessageAglomeration struct {
    MAC    string `json:"mac"`
    Evento string `json:"evento"`
}

func ToJSON(m MessageAglomeration) ([]byte, error) {
    return json.Marshal(m)
}

func NewMessage(message string) MessageAglomeration {
    var msg map[string]string
    json.Unmarshal([]byte(message), &msg)
    return MessageAglomeration{
        MAC:    msg["mac"],
        Evento: msg["evento"],
    }
}
