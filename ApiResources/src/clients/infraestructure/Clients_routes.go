package infraestructure

import (
	"api_resources/src/clients/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

type ClientsHandlers struct {
	create *controllers.CreateClientController
	getAll *controllers.GetAllClientsController
}

func ClientsRoutes(router *gin.Engine, handlers ClientsHandlers) {
	router.POST("/clients", handlers.create.Execute)
	router.GET("/clients", handlers.getAll.Execute)
}
