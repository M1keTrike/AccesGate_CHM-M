package application

import (
	"fmt"
	"log"
	"strings"

	"github.com/M1keTrike/Accessgate/domain"
	"github.com/M1keTrike/Accessgate/infrastructure/firebase"
)

type WebhookService struct{}

func NewWebhookService() *WebhookService {
	return &WebhookService{}
}

func (ws *WebhookService) ProcessEvent(event domain.Event) error {
	log.Println("📩 Procesando evento:", event.Event)

	if event.Event == "" {
		log.Println("🚨 Evento inválido, faltan datos.")
		return fmt.Errorf("evento inválido: falta información")
	}

	alertMessage := ws.detectFailure(event.Event)

	err := firebase.SendNotification(alertMessage)
	if err != nil {
		log.Println("🚨 Error enviando notificación:", err)
		return fmt.Errorf("error enviando notificación: %v", err)
	}

	log.Println("✅ Evento procesado y notificación enviada con éxito")
	return nil
}

func (ws *WebhookService) detectFailure(message string) string {
	loweredMessage := strings.ToLower(message)

	if strings.Contains(loweredMessage, "database") {
		return "🚨 Alerta: Fallo detectado en la base de datos 🚨"
	} else if strings.Contains(loweredMessage, "websocket") {
		return "⚠️ Atención: El servidor de de envío datos ha fallado ⚠️"
	} else if strings.Contains(loweredMessage, "broker") {
		return "🔥 Urgente: El servidor del hardware a fallado 🔥"
	}

	return message
}
