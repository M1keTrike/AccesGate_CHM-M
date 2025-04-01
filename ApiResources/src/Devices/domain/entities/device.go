package entities

import "time"

type Device struct {
	ID           int       `json:"id"`
	HardwareID   string    `json:"hardware_id"`
	Type         string    `json:"type"`
	Status       string    `json:"status"`
	Location     string    `json:"location"`
	RegisteredAt time.Time `json:"registered_at"`
	AssignedTo   int       `json:"assigned_to"`
	UpdatedAt    time.Time `json:"updated_at"`
}
