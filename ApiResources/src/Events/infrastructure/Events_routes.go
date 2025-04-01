package infraestructure

import (
	"api_resources/src/Events/infrastructure/controllers"
	"api_resources/src/core"

	"github.com/gin-gonic/gin"
)

type EventsHandlers struct {
	create            *controllers.CreateEventController
	getAll            *controllers.GetAllEventsController
	getByID           *controllers.GetEventByIDController
	getByCreator      *controllers.GetEventsByCreatorController
	update            *controllers.UpdateEventController
	delete            *controllers.DeleteEventController
}

func EventsRoutes(router *gin.Engine, handlers EventsHandlers) {
	protected := router.Group("/events")
	protected.Use(core.AuthMiddleware())

	protected.POST("", handlers.create.Execute)
	protected.GET("", handlers.getAll.Execute)
	protected.GET("/:id", handlers.getByID.Execute)
	protected.GET("/creator/:user_id", handlers.getByCreator.Execute)
	protected.PUT("/:id", handlers.update.Execute)
	protected.DELETE("/:id", handlers.delete.Execute)
}
