package infrastructure

import (
	"api_resources/src/AccessEvents/infrastructure/controllers"
	"api_resources/src/core"

	"github.com/gin-gonic/gin"
)

type AccessEventsHandlers struct {
	create      *controllers.CreateAccessEventController
	getAll      *controllers.GetAllAccessEventsController
	getByID     *controllers.GetAccessEventByIDController
	getByUser   *controllers.GetAccessEventsByUserController
	getByDevice *controllers.GetAccessEventsByDeviceController
	getByFront  *controllers.GetAccessEventsByFrontController
	update      *controllers.UpdateAccessEventController
	delete      *controllers.DeleteAccessEventController
}

func AccessEventsRoutes(router *gin.Engine, handlers AccessEventsHandlers) {
	protected := router.Group("/access-events")
	protected.Use(core.AuthMiddleware())

	protected.POST("", handlers.create.Execute)
	protected.GET("", handlers.getAll.Execute)
	protected.GET("/:id", handlers.getByID.Execute)
	protected.GET("/user/:userId", handlers.getByUser.Execute)
	protected.GET("/device/:deviceId", handlers.getByDevice.Execute)
	protected.GET("/front/:frontId", handlers.getByFront.Execute)
	protected.PUT("", handlers.update.Execute)
	protected.DELETE("/:id", handlers.delete.Execute)
}
