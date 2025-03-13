package core

import (
	"github.com/M1keTrike/EventDriven/internal/models"
	"github.com/M1keTrike/EventDriven/internal/ports"
)

type MessageService struct {
	repo ports.MessageRepositoryPort
}

func NewMessageService(repo ports.MessageRepositoryPort) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) ProcessMessage(msg *models.Message) (*models.Message, error) {

	err := s.repo.SaveMessage(msg)
	if err != nil {
		return nil, err
	}
	return msg, nil
}
