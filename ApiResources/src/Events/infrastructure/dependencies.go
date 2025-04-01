package infrastructure

import (
	"api_resources/src/Events/application"
	"api_resources/src/Events/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	// Adaptador secundario PostgreSQL
	eventRepo := NewPostgreSQL()

	// Casos de uso
	createEventUC := application.NewCreateEvent(eventRepo)
	deleteEventUC := application.NewDeleteEvent(eventRepo)
	getAllEventsUC := application.NewGetAllEvents(eventRepo)
	getEventByIDUC := application.NewGetEventByID(eventRepo)
	updateEventUC := application.NewUpdateEvent(eventRepo)

	// Controladores HTTP
	createEventCtrl := controllers.NewCreateEventController(createEventUC)
	deleteEventCtrl := controllers.NewDeleteEventController(deleteEventUC)
	getAllEventsCtrl := controllers.NewGetAllEventsController(getAllEventsUC)
	getEventByIDCtrl := controllers.NewGetEventByIDController(getEventByIDUC)
	updateEventCtrl := controllers.NewUpdateEventController(updateEventUC)

	EventsRoutes(router, EventsHandlers{
		create:  createEventCtrl,
		delete:  deleteEventCtrl,
		getAll:  getAllEventsCtrl,
		getByID: getEventByIDCtrl,
		update:  updateEventCtrl,
	})
}
