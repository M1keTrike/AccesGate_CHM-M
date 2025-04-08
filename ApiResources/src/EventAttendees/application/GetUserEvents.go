package application

import (
    "api_resources/src/EventAttendees/domain"
    "api_resources/src/EventAttendees/domain/entities"
)

type GetUserEvents struct {
    repo domain.EventAttendeeRepository
}

func NewGetUserEvents(repo domain.EventAttendeeRepository) *GetUserEvents {
    return &GetUserEvents{repo: repo}
}

func (uc *GetUserEvents) Execute(userID int) ([]entities.EventAttendee, error) {
    return uc.repo.GetUserEvents(userID)
}