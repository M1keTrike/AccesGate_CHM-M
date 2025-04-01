package application

import (
	"api_resources/src/Events/domain"
	"api_resources/src/Events/domain/entities"
)

type GetEventsByCreator struct {
	repo domain.EventRepository
}

func NewGetEventsByCreator(repo domain.EventRepository) *GetEventsByCreator {
	return &GetEventsByCreator{repo: repo}
}

func (uc *GetEventsByCreator) Execute(userID int) ([]entities.Event, error) {
	return uc.repo.GetEventsByCreator(userID)
}
