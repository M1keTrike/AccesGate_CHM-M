package entities

import "time"

type NfcCardAssignment struct {
    ID         int       `json:"id"`
    UserID     int       `json:"user_id"`
    CardUID    string    `json:"card_uid"`
    AssignedAt time.Time `json:"assigned_at"`
    IsActive   bool      `json:"is_active"`
}