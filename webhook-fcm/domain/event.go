package domain

// Event representa un evento recibido desde el webhook
type Event struct {
	Event   string `json:"event"`
	Message string `json:"message"`
}
