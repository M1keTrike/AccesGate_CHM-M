package routers

import (
	"github.com/M1keTrike/EventDriven/messages_nfc/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func AttachReceiveRoutes(r *gin.Engine, receiveController *controllers.ReceiveNFCMessageController) {
	nfc := r.Group("/nfc")
	{
		nfc.POST("/receive_nfc", receiveController.Execute)
	}

}
