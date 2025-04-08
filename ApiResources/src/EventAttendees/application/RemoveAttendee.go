package application

import (
    "api_resources/src/EventAttendees/domain"
)

type RemoveAttendee struct {
    repo domain.EventAttendeeRepository
}

func NewRemoveAttendee(repo domain.EventAttendeeRepository) *RemoveAttendee {
    return &RemoveAttendee{repo: repo}
}

func (uc *RemoveAttendee) Execute(eventID, userID int) error {
    return uc.repo.RemoveAttendee(eventID, userID)
}