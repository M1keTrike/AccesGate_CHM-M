package application

import (
	"api_resources/src/AccessEvents/domain"
	"api_resources/src/AccessEvents/domain/entities"
	"time"
)

type CreateAccessEvent struct {
	Repo domain.AccessEventRepository
}

func NewCreateAccessEvent(repo domain.AccessEventRepository) *CreateAccessEvent {
	return &CreateAccessEvent{Repo: repo}
}

func (c *CreateAccessEvent) Execute(event *entities.AccessEvent) error {
	event.Timestamp = time.Now()
	return c.Repo.Create(event)
}
