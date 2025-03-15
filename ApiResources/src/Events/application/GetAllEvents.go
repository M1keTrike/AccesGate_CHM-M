package application

import (
	"api_resources/src/Events/domain"
	"api_resources/src/Events/domain/entities"
)

type GetAllEvents struct {
	repo domain.EventRepository
}

func NewGetAllEvents(repo domain.EventRepository) *GetAllEvents {
	return &GetAllEvents{repo: repo}
}

func (uc *GetAllEvents) Execute() ([]entities.Event, error) {
	return uc.repo.GetAllEvents()
}
