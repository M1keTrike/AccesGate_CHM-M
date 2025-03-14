package dependencies

import (
	"os"

	"github.com/M1keTrike/EventDriven/messages_aglomeration/application"
	"github.com/M1keTrike/EventDriven/messages_aglomeration/infraestructure/controllers"
	"github.com/M1keTrike/EventDriven/messages_aglomeration/infraestructure/routers"
	service "github.com/M1keTrike/EventDriven/services/websocket/application"
	infrastructure "github.com/M1keTrike/EventDriven/services/websocket/infraestructure"
	"github.com/gin-gonic/gin"
)

type MessageAglomerationDependencies struct{}

func NewMessageAgloemrationDependencies() MessageAglomerationDependencies {
	return MessageAglomerationDependencies{}
}

func (d *MessageAglomerationDependencies) Execute(r *gin.Engine) {

	wsUrl := os.Getenv("WEBSOCKET_URL")

	webSocketService, err := infrastructure.NewWebSocketAglomerationEmitter(wsUrl)
	if err != nil {
		panic(err)

	}
	publishAglomerationUseCase := service.NewSendAglomerationService(webSocketService)
	sendAglomerationUseCase := application.NewSendAglomerationUseCase(publishAglomerationUseCase)
	receiveAglomerationUseCase := application.NewRecieveAglomerationUseCase(sendAglomerationUseCase)
	recieveAglomerationControllers := controllers.NewRecieveAglomerationController(receiveAglomerationUseCase)
	routers.AttachReceiveAglomerationRoutes(r, recieveAglomerationControllers)

}
