package dependencies

import (
	"log"
	"os"

	"github.com/M1keTrike/EventDriven/messages_alarm/application"
	"github.com/M1keTrike/EventDriven/messages_alarm/infrastructure/controllers"
	"github.com/M1keTrike/EventDriven/messages_alarm/infrastructure/routers"
	service "github.com/M1keTrike/EventDriven/services/websocket/infraestructure"
	"github.com/gin-gonic/gin"
)

type MessageAlarmDependencies struct{}

func NewMessageAlarmDependencies() MessageAlarmDependencies {
	return MessageAlarmDependencies{}
}

func (d *MessageAlarmDependencies) Execute(router *gin.Engine) {
	webSocketURL := os.Getenv("WEBSOCKET_URL")
	if webSocketURL == "" {
		log.Fatal("WEBSOCKET_URL environment variable is not set")
	}

	webSocketEmitter, err := service.NewWebSocketAlarmEmitter(webSocketURL)
	if err != nil {
		log.Fatalf("Failed to create WebSocket emitter: %v", err)
	}

	// Initialize use cases with proper message formatting
	sendAlarmUseCase := application.NewSendAlarmUseCase(webSocketEmitter)
	receiveAlarmUseCase := application.NewReceiveAlarmUseCase(&sendAlarmUseCase)

	// Create controller with message formatting support
	receiveAlarmController := controllers.NewReceiveAlarmController(receiveAlarmUseCase)

	// Attach routes
	routers.AttachReceiveAlarmRoutes(router, receiveAlarmController)
}
