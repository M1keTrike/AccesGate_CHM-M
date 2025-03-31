package infrastructure

import (
    "api_resources/src/Events/application"
    "api_resources/src/Events/infrastructure/controllers"
    "github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
    os := NewPostgreSQL()
    
    createEventService := application.NewCreateEvent(os)
    getEventService := application.NewGetEventByID(os)
    getAllEventsService := application.NewGetAllEvents(os)
    updateEventService := application.NewUpdateEvent(os)
    deleteEventService := application.NewDeleteEvent(os)

    createEventController := controllers.NewCreateEventController(createEventService)
    getEventController := controllers.NewGetEventByIDController(getEventService)
    getAllEventsController := controllers.NewGetAllEventsController(getAllEventsService)
    updateEventController := controllers.NewUpdateEventController(updateEventService)
    deleteEventController := controllers.NewDeleteEventController(deleteEventService)

    EventsRoutes(router, EventsHandlers{
        create: createEventController,
        get:    getEventController,
        getAll: getAllEventsController,
        update: updateEventController,
        delete: deleteEventController,
    })
}