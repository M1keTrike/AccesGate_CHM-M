package application

import (
	"api_resources/src/Events/domain"
	"api_resources/src/Events/domain/entities"
)

type GetEventByID struct {
	repo domain.EventRepository
}

func NewGetEventByID(repo domain.EventRepository) *GetEventByID {
	return &GetEventByID{repo: repo}
}

func (uc *GetEventByID) Execute(id int) (entities.Event, error) {
	return uc.repo.GetEventByID(id)
}
