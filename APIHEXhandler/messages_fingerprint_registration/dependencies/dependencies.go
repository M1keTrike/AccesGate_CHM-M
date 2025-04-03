package dependencies

import (
	"log"
	"os"

	"github.com/M1keTrike/EventDriven/messages_fingerprint_registration/application"
	"github.com/M1keTrike/EventDriven/messages_fingerprint_registration/infrastructure/controllers"
	"github.com/M1keTrike/EventDriven/messages_fingerprint_registration/infrastructure/routers"
	service "github.com/M1keTrike/EventDriven/services/websocket/infraestructure"
	"github.com/gin-gonic/gin"
)

type MessageFingerprintRegistrationDependencies struct{}

func NewMessageFingerprintRegistrationDependencies() MessageFingerprintRegistrationDependencies {
	return MessageFingerprintRegistrationDependencies{}
}

func (d *MessageFingerprintRegistrationDependencies) Execute(router *gin.Engine) {
	webSocketURL := os.Getenv("WEBSOCKET_URL")
	if webSocketURL == "" {
		log.Fatal("WEBSOCKET_URL environment variable is not set")
	}

	webSocketEmitter, err := service.NewWebSocketFingerprintRegistrationEmitter(webSocketURL)
	if err != nil {
		log.Fatalf("Failed to create WebSocket emitter: %v", err)
	}

	sendFingerprintRegistrationUseCase := application.NewSendFingerprintRegistrationUseCase(webSocketEmitter)
	receiveFingerprintRegistrationUseCase := application.NewReceiveFingerprintRegistrationUseCase(&sendFingerprintRegistrationUseCase)
	receiveFingerprintRegistrationController := controllers.NewReceiveFingerprintRegistrationController(receiveFingerprintRegistrationUseCase)

	routers.AttachReceiveFingerprintRegistrationRoutes(router, receiveFingerprintRegistrationController)
}
