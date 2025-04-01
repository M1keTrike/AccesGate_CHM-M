package infraestructure

import (
	"api_resources/src/Events/application"
	"api_resources/src/Events/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	repo := NewPostgreSQLEvents()

	// Casos de uso
	createUC := application.NewCreateEvent(repo)
	getAllUC := application.NewGetAllEvents(repo)
	getByIDUC := application.NewGetEventByID(repo)
	getByCreatorUC := application.NewGetEventsByCreator(repo)
	updateUC := application.NewUpdateEvent(repo)
	deleteUC := application.NewDeleteEvent(repo)

	// Controladores
	createCtrl := controllers.NewCreateEventController(createUC)
	getAllCtrl := controllers.NewGetAllEventsController(getAllUC)
	getByIDCtrl := controllers.NewGetEventByIDController(getByIDUC)
	getByCreatorCtrl := controllers.NewGetEventsByCreatorController(getByCreatorUC)
	updateCtrl := controllers.NewUpdateEventController(updateUC)
	deleteCtrl := controllers.NewDeleteEventController(deleteUC)

	// Registrar rutas
	EventsRoutes(router, EventsHandlers{
		create:       createCtrl,
		getAll:       getAllCtrl,
		getByID:      getByIDCtrl,
		getByCreator: getByCreatorCtrl,
		update:       updateCtrl,
		delete:       deleteCtrl,
	})
}
