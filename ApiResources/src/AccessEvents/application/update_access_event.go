package application

import (
	"api_resources/src/AccessEvents/domain"
	"api_resources/src/AccessEvents/domain/entities"
)

type UpdateAccessEvent struct {
	Repo domain.AccessEventRepository
}

func NewUpdateAccessEvent(repo domain.AccessEventRepository) *UpdateAccessEvent {
	return &UpdateAccessEvent{Repo: repo}
}

func (u *UpdateAccessEvent) Execute(event *entities.AccessEvent) error {
	return u.Repo.Update(event)
}
