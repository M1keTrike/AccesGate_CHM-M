package infraestructure

import (
    "api_resources/src/EventAttendees/infraestructure/controllers"
    "github.com/gin-gonic/gin"
)

type EventAttendeesHandlers struct {
    register       *controllers.RegisterAttendeeController
    remove         *controllers.RemoveAttendeeController
    getAttendees   *controllers.GetEventAttendeesController
    getUserEvents  *controllers.GetUserEventsController
    isRegistered   *controllers.IsUserRegisteredController
}

func EventAttendeesRoutes(router *gin.Engine, handlers EventAttendeesHandlers) {
    group := router.Group("/event-attendees")
    {
        group.POST("/register", handlers.register.Execute)
        group.DELETE("/events/:eventId/users/:userId", handlers.remove.Execute)
        group.GET("/events/:eventId/attendees", handlers.getAttendees.Execute)
        group.GET("/users/:userId/events", handlers.getUserEvents.Execute)
        group.GET("/events/:eventId/users/:userId/check", handlers.isRegistered.Execute)
    }
}