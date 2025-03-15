package routers

import (
	"github.com/M1keTrike/EventDriven/messages_v_access/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func AttachReceiveVAccessRoutes(r *gin.Engine, receiveVAccessController *controllers.ReceiveVAccessController) {
	aglomeration := r.Group("/v_access")
	{
		aglomeration.POST("/recieve_vaccess", receiveVAccessController.Execute)
	}

}
