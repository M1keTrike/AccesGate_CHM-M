package application

import (
	"fmt"
	"log"

	"github.com/M1keTrike/Accessgate/domain"
	"github.com/M1keTrike/Accessgate/infrastructure/firebase"
)

type WebhookService struct{}

func NewWebhookService() *WebhookService {
	return &WebhookService{}
}

func (ws *WebhookService) ProcessEvent(event domain.Event) error {
	log.Println("📩 Procesando evento:", event.Event, "-", event.Message)

	if event.Event == "" || event.Message == "" {
		return fmt.Errorf("evento inválido: falta información")
	}

	err := firebase.SendNotification(event.Event, event.Message)
	if err != nil {
		return fmt.Errorf("error enviando notificación: %v", err)
	}

	log.Println("✅ Evento procesado y notificación enviada con éxito")
	return nil
}
