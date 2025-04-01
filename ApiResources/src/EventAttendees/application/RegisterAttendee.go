package application

import (
    "api_resources/src/EventAttendees/domain"
    "api_resources/src/EventAttendees/domain/entities"
)

type RegisterAttendee struct {
    repo domain.EventAttendeeRepository
}

func NewRegisterAttendee(repo domain.EventAttendeeRepository) *RegisterAttendee {
    return &RegisterAttendee{repo: repo}
}

func (uc *RegisterAttendee) Execute(attendee *entities.EventAttendee) error {
    exists, err := uc.repo.IsUserRegistered(attendee.EventID, attendee.UserID)
    if err != nil {
        return err
    }
    if exists {
        return nil
    }
    return uc.repo.RegisterAttendee(attendee)
}