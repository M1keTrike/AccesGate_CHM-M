package entities

import (
    "time"
)

type EventAttendee struct {
    ID             int       `json:"id"`
    UserID         int       `json:"user_id"`
    EventID        int       `json:"event_id"`
    RegisteredAt   time.Time `json:"registered_at"`
    Attended       bool      `json:"attended" gorm:"default:false"`
}