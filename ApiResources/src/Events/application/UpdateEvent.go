package application

import (
	"api_resources/src/Events/domain"
	"api_resources/src/Events/domain/entities"
)

type UpdateEvent struct {
	repo domain.EventRepository
}

func NewUpdateEvent(repo domain.EventRepository) *UpdateEvent {
	return &UpdateEvent{repo: repo}
}

func (uc *UpdateEvent) Execute(event *entities.Event) error {
	return uc.repo.UpdateEvent(event)
}
