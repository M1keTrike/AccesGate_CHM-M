package domain

import "encoding/json"

type MessageAglomeration struct {
	Message string
}

func ToJSON(m MessageAglomeration) ([]byte, error) {
	return json.Marshal(m)
}

func NewMessage(message string) MessageAglomeration {
	return MessageAglomeration{
		Message: message,
	}
}
