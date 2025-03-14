package routers

import (
	"github.com/M1keTrike/EventDriven/messages_aglomeration/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func AttachReceiveAglomerationRoutes(r *gin.Engine, receiveAglomerationController *controllers.ReceiveAglomerationController) {
	aglomeration := r.Group("/aglomeration")
	{
		aglomeration.POST("/recieve_aglomeration", receiveAglomerationController.Execute)
	}

}
