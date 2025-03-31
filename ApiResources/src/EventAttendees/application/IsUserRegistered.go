package application

import (
    "api_resources/src/EventAttendees/domain"
)

type IsUserRegistered struct {
    repo domain.EventAttendeeRepository
}

func NewIsUserRegistered(repo domain.EventAttendeeRepository) *IsUserRegistered {
    return &IsUserRegistered{repo: repo}
}

func (uc *IsUserRegistered) Execute(eventID, userID int) (bool, error) {
    return uc.repo.IsUserRegistered(eventID, userID)
}