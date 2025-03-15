package ports

import "github.com/M1keTrike/EventDriven/internal/models"

type MessageRepositoryPort interface {
	SaveMessage(msg *models.Message) error
}
