



package infrastructure


import (
    "api_resources/src/Events/infrastructure/controllers"
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

    create  *controllers.CreateEventController
    getByID     *controllers.GetEventByIDController
    getAll  *controllers.GetAllEventsController
    update  *controllers.UpdateEventController
    delete  *controllers.DeleteEventController
}

func EventsRoutes(router *gin.Engine, handlers EventsHandlers) {
    router.POST("/events", handlers.create.Handle)
    router.GET("/events/:id", handlers.getByID.Handle)
    router.GET("/events", handlers.getAll.Handle)
    router.PUT("/events/:id", handlers.update.Handle)
    router.DELETE("/events/:id", handlers.delete.Handle)

}


