package entities

import "time"

type User struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	PasswordHash  string    `json:"password_hash"`
	Role          string    `json:"role"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	FingerprintID int16     `json:"fingerprint_id"`
	BiometricAuth bool      `json:"biometric_auth"`
	CreatedBy     int       `json:"created_by"`
}
