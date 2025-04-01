package infraestructure

import (
	"api_resources/src/Devices/infrastructure/controllers"
	"api_resources/src/core"

	"github.com/gin-gonic/gin"
)

type DeviceHandlers struct {
	create    *controllers.CreateDeviceController
	getAll    *controllers.GetAllDevicesController
	getByID   *controllers.GetDeviceByIDController
	getByHWID *controllers.GetDeviceByHardwareIDController
	update    *controllers.UpdateDeviceController
	delete    *controllers.DeleteDeviceController
}

func DeviceRoutes(router *gin.Engine, handlers DeviceHandlers) {
	protected := router.Group("/devices")
	protected.Use(core.AuthMiddleware())

	protected.POST("", handlers.create.Execute)
	protected.GET("", handlers.getAll.Execute)
	protected.GET("/:id", handlers.getByID.Execute)
	protected.GET("/hardware/:hardware_id", handlers.getByHWID.Execute)
	protected.PUT("/:id", handlers.update.Execute)
	protected.DELETE("/:id", handlers.delete.Execute)
}
