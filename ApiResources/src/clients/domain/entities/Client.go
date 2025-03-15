package entities

type Client struct {
	ClientID   int    `json:"client_id"`
	WorkerID   string `json:"worker_id"`
	FrontID    string `json:"front_id"`
	HardwareID string `json:"hardware_id"`
}
