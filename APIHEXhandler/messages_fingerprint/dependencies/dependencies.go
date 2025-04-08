package dependencies

import (
	"log"
	"os"

	"github.com/M1keTrike/EventDriven/messages_fingerprint/application"
	"github.com/M1keTrike/EventDriven/messages_fingerprint/infrastructure/controllers"
	"github.com/M1keTrike/EventDriven/messages_fingerprint/infrastructure/routers"
	infrastructure "github.com/M1keTrike/EventDriven/services/websocket/infraestructure"
	"github.com/gin-gonic/gin"
)

type MessageFingerprintDependencies struct{}

func NewMessageFingerprintDependencies() MessageFingerprintDependencies {
	return MessageFingerprintDependencies{}
}

func (d *MessageFingerprintDependencies) Execute(router *gin.Engine) {
	webSocketURL := os.Getenv("WEBSOCKET_URL")
	if webSocketURL == "" {
		log.Fatal("WEBSOCKET_URL environment variable is not set")
	}

	webSocketEmitter, err := infrastructure.NewWebSocketFingerprintEmitter(webSocketURL)
	if err != nil {
		log.Fatalf("Failed to create WebSocket emitter: %v", err)
	}

	sendFingerprintUseCase := application.NewSendFingerprintUseCase(webSocketEmitter)
	receiveFingerprintUseCase := application.NewReceiveFingerprintUseCase(&sendFingerprintUseCase)
	receiveFingerprintController := controllers.NewReceiveFingerprintController(receiveFingerprintUseCase)

	routers.AttachReceiveFingerprintRoutes(router, receiveFingerprintController)
}
