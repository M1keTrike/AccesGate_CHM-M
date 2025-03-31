package application

import (
    "api_resources/src/EventAttendees/domain"
    "api_resources/src/EventAttendees/domain/entities"
)

type GetEventAttendees struct {
    repo domain.EventAttendeeRepository
}

func NewGetEventAttendees(repo domain.EventAttendeeRepository) *GetEventAttendees {
    return &GetEventAttendees{repo: repo}
}

func (uc *GetEventAttendees) Execute(eventID int) ([]entities.EventAttendee, error) {
    return uc.repo.GetEventAttendees(eventID)
}