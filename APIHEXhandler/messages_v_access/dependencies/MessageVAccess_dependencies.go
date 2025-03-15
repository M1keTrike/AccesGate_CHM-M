package dependencies

import (
	"os"

	"github.com/M1keTrike/EventDriven/messages_v_access/application"
	"github.com/M1keTrike/EventDriven/messages_v_access/infraestructure/controllers"
	"github.com/M1keTrike/EventDriven/messages_v_access/infraestructure/routers"
	service "github.com/M1keTrike/EventDriven/services/websocket/application"
	infrastructure "github.com/M1keTrike/EventDriven/services/websocket/infraestructure"
	"github.com/gin-gonic/gin"
)

type MessageVAccessDependencies struct{}

func NewMessageAgloemrationDependencies() MessageVAccessDependencies {
	return MessageVAccessDependencies{}
}

func (d *MessageVAccessDependencies) Execute(r *gin.Engine) {

	wsUrl := os.Getenv("WEBSOCKET_URL")

	webSocketService, err := infrastructure.NewWebSocketVAccessEmitter(wsUrl)
	if err != nil {
		panic(err)

	}

	publishVAccessUseCase := service.NewSendVAccessService(webSocketService)
	sendVAccessUseCase := application.NewSendVAccessUseCase(publishVAccessUseCase)
	receiveVAccessUseCase := application.NewRecieveVAccessUseCase(sendVAccessUseCase)
	recieveVAccessControllers := controllers.NewRecieveVAccessController(receiveVAccessUseCase)
	routers.AttachReceiveVAccessRoutes(r, recieveVAccessControllers)

}
