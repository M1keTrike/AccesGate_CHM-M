package entities

import "time"

type AccessEvent struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	FrontID   int       `json:"front_id"`
	DeviceID  int       `json:"device_id"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}
