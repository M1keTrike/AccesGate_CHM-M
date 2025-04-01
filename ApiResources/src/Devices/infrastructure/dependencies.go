package infraestructure

import (
	"api_resources/src/Devices/application"
	"api_resources/src/Devices/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	db := NewPostgreSQLEvents()

	// Casos de uso
	create := application.NewCreateDevice(db)
	getAll := application.NewGetAllDevices(db)
	getByID := application.NewGetDeviceByID(db)
	getByHWID := application.NewGetDeviceByHardwareID(db)
	update := application.NewUpdateDevice(db)
	delete := application.NewDeleteDevice(db)

	// Controladores
	createCtrl := controllers.NewCreateDeviceController(create)
	getAllCtrl := controllers.NewGetAllDevicesController(getAll)
	getByIDCtrl := controllers.NewGetDeviceByIDController(getByID)
	getByHWIDCtrl := controllers.NewGetDeviceByHardwareIDController(getByHWID)
	updateCtrl := controllers.NewUpdateDeviceController(update)
	deleteCtrl := controllers.NewDeleteDeviceController(delete)

	// Rutas
	DeviceRoutes(router, DeviceHandlers{
		create:    createCtrl,
		getAll:    getAllCtrl,
		getByID:   getByIDCtrl,
		getByHWID: getByHWIDCtrl,
		update:    updateCtrl,
		delete:    deleteCtrl,
	})
}
