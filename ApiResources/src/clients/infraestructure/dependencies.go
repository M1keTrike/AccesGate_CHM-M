package infraestructure

import (
	"api_resources/src/clients/application"
	"api_resources/src/clients/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	db := NewPostgreSQL()

	createClientService := application.NewCreateClientUseCase(db)
	getAllClientsService := application.NewGetAllClientsUseCase(db)

	createClientController := controllers.NewCreateClientController(createClientService)
	getAllClientsController := controllers.NewGetAllClientsController(getAllClientsService)

	ClientsRoutes(router, ClientsHandlers{
		create: createClientController,
		getAll: getAllClientsController,
	})
}
