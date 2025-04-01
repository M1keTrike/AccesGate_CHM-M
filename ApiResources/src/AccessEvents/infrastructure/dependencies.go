package infrastructure

import (
	"api_resources/src/AccessEvents/application"
	"api_resources/src/AccessEvents/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	repo := NewPostgreSQLAccessEvents()

	// Casos de uso
	createUC := application.NewCreateAccessEvent(repo)
	getAllUC := application.NewGetAllAccessEvents(repo)
	getByIDUC := application.NewGetAccessEventByID(repo)
	getByUserUC := application.NewGetAccessEventsByUser(repo)
	getByDeviceUC := application.NewGetAccessEventsByDevice(repo)
	getByFrontUC := application.NewGetAccessEventsByFront(repo)
	updateUC := application.NewUpdateAccessEvent(repo)
	deleteUC := application.NewDeleteAccessEvent(repo)

	// Controladores
	createCtrl := controllers.NewCreateAccessEventController(*createUC)
	getAllCtrl := controllers.NewGetAllAccessEventsController(*getAllUC)
	getByIDCtrl := controllers.NewGetAccessEventByIDController(*getByIDUC)
	getByUserCtrl := controllers.NewGetAccessEventsByUserController(*getByUserUC)
	getByDeviceCtrl := controllers.NewGetAccessEventsByDeviceController(*getByDeviceUC)
	getByFrontCtrl := controllers.NewGetAccessEventsByFrontController(*getByFrontUC)
	updateCtrl := controllers.NewUpdateAccessEventController(*updateUC)
	deleteCtrl := controllers.NewDeleteAccessEventController(*deleteUC)

	// Registrar rutas
	AccessEventsRoutes(router, AccessEventsHandlers{
		create:      createCtrl,
		getAll:      getAllCtrl,
		getByID:     getByIDCtrl,
		getByUser:   getByUserCtrl,
		getByDevice: getByDeviceCtrl,
		getByFront:  getByFrontCtrl,
		update:      updateCtrl,
		delete:      deleteCtrl,
	})
}
