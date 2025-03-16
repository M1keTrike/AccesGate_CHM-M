package application

import (
	"fmt"
	"log"

	"github.com/M1keTrike/Accessgate/infrastructure/firebase"
)

type SuscibeToTokenService struct{}

func NewSuscribeToTokenService() *WebhookService {
	return &WebhookService{}
}

func (ws *WebhookService) Execute(token string) error {
	if token == "" {
		return fmt.Errorf("🚨 Token de registro vacío")
	}

	err := firebase.SubscribeToTopic(token, "all")
	if err != nil {
		return fmt.Errorf("🚨 Error al suscribirse al tema 'all': %v", err)
	}

	log.Println("✅ Token suscrito al tema 'all'")
	return nil
}
