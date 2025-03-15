package application

import (
	"api_resources/src/Events/domain"
	"api_resources/src/Events/domain/entities"
)

type CreateEvent struct {
	repo domain.EventRepository
}

func NewCreateEvent(repo domain.EventRepository) *CreateEvent {
	return &CreateEvent{repo: repo}
}

func (uc *CreateEvent) Execute(event entities.Event) error {
	return uc.repo.CreateEvent(event)
}
