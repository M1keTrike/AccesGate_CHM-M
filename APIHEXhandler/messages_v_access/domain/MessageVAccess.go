package domain

import "encoding/json"

type MessageVAccess struct {
    MAC    string `json:"mac"`
    Evento string `json:"evento"`
}

func ToJSON(m MessageVAccess) ([]byte, error) {
    return json.Marshal(m)
}

func NewMessage(message string) MessageVAccess {
    var msg map[string]string
    json.Unmarshal([]byte(message), &msg)
    return MessageVAccess{
        MAC:    msg["mac"],
        Evento: msg["evento"],
    }
}
