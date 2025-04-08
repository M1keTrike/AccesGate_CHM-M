package application

import (
	"api_resources/src/AccessEvents/domain"
)

type DeleteAccessEvent struct {
	Repo domain.AccessEventRepository
}

func NewDeleteAccessEvent(repo domain.AccessEventRepository) *DeleteAccessEvent {
	return &DeleteAccessEvent{Repo: repo}
}

func (d *DeleteAccessEvent) Execute(id int) error {
	return d.Repo.Delete(id)
}
