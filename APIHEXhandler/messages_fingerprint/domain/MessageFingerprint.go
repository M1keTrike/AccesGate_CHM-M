package domain

import "encoding/json"

type MessageFingerprint struct {
    MAC    string `json:"mac"`
    Evento string `json:"evento"`
    ID     string `json:"id"`
}

func ToJSON(m MessageFingerprint) ([]byte, error) {
    return json.Marshal(m)
}

func NewMessage(message string) MessageFingerprint {
    var msg map[string]string
    json.Unmarshal([]byte(message), &msg)
    return MessageFingerprint{
        MAC:    msg["mac"],
        Evento: msg["evento"],
        ID:     msg["id"],
    }
}