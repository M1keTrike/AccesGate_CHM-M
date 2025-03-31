package infraestructure

import (
	"api_resources/src/Events/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

type EventsHandlers struct {
	create  *controllers.CreateEventController
	delete  *controllers.DeleteEventController
	getAll  *controllers.GetAllEventsController
	getByID *controllers.GetEventByIDController
	update  *controllers.UpdateEventController
}

func EventsRoutes(router *gin.Engine, handlers EventsHandlers) {
	router.POST("/events", handlers.create.Handle)
	router.GET("/events", handlers.getAll.Handle)
	router.GET("/events/:id", handlers.getByID.Handle)
	router.PUT("/events/:id", handlers.update.Handle)
	router.DELETE("/events/:id", handlers.delete.Handle)
}
