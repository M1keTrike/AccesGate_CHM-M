package ports

import "github.com/M1keTrike/EventDriven/internal/models"

type WebSocketPort interface {
	SendMessage(msg *models.Message) error
}
