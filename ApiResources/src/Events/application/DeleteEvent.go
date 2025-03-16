package application

import (
  "api_resources/src/Events/domain"
)

type DeleteEvent struct {
    repo domain.EventRepository
}

func NewDeleteEvent(repo domain.EventRepository) *DeleteEvent {
    return &DeleteEvent{repo: repo}
}

func (uc *DeleteEvent) Execute(id int) error {
    return uc.repo.DeleteEvent(id)
}