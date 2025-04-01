package domain

import (
    "api_resources/src/EventAttendees/domain/entities"
)

type EventAttendeeRepository interface {
    RegisterAttendee(attendee *entities.EventAttendee) error
    RemoveAttendee(eventID, userID int) error
    GetEventAttendees(eventID int) ([]entities.EventAttendee, error)
    GetUserEvents(userID int) ([]entities.EventAttendee, error)
    IsUserRegistered(eventID, userID int) (bool, error)
}