package dependencies

import (
	"os"

	"github.com/M1keTrike/EventDriven/messages_nfc/application"
	"github.com/M1keTrike/EventDriven/messages_nfc/infraestructure/controllers"
	"github.com/M1keTrike/EventDriven/messages_nfc/infraestructure/routers"
	service "github.com/M1keTrike/EventDriven/services/websocket/application"
	infrastructure "github.com/M1keTrike/EventDriven/services/websocket/infraestructure"
	"github.com/gin-gonic/gin"
)

type MessageDependencies struct{}

func NewMessageDependencies() *MessageDependencies {
	return &MessageDependencies{}
}

func (d *MessageDependencies) Execute(r *gin.Engine) {

	wsUrl := os.Getenv("WEBSOCKET_URL")

	webSocketService, err := infrastructure.NewWebSocketEmitter(wsUrl)
	if err != nil {
		panic(err)

	}
	publishNFCUseCase := service.NewSendWSNFCService(webSocketService)
	sendNFCUseCase := application.NewSendMessageUseCase(publishNFCUseCase)
	receiveNFCUseCase := application.NewRecieveNFCMessageUseCase(sendNFCUseCase)
	recieveNFCController := controllers.NewReceiveNFCController(receiveNFCUseCase)
	routers.AttachReceiveRoutes(r, recieveNFCController)

}
