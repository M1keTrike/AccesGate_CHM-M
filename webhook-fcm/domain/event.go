package domain

type Event struct {
	Event string `json:"event"`
}

type SubscribeRequest struct {
	Token string `json:"token"`
}
