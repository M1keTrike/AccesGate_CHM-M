package infraestructure

import (
    "api_resources/src/EventAttendees/application"
    "api_resources/src/EventAttendees/infraestructure/controllers"
    "github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
    // Initialize PostgreSQL repository
    repository := NewPostgreSQL()

    // Initialize use cases
    registerUseCase := application.NewRegisterAttendee(repository)
    removeUseCase := application.NewRemoveAttendee(repository)
    getAttendeesUseCase := application.NewGetEventAttendees(repository)
    getUserEventsUseCase := application.NewGetUserEvents(repository)
    isRegisteredUseCase := application.NewIsUserRegistered(repository)
    UpdateAttendenceUseCase :=application.NewUpdateAttendanceStatus(repository)

    // Initialize controllers
    registerController := controllers.NewRegisterAttendeeController(registerUseCase)
    removeController := controllers.NewRemoveAttendeeController(removeUseCase)
    getAttendeesController := controllers.NewGetEventAttendeesController(getAttendeesUseCase)
    getUserEventsController := controllers.NewGetUserEventsController(getUserEventsUseCase)
    isRegisteredController := controllers.NewIsUserRegisteredController(isRegisteredUseCase)
    UpdateAttendenceController :=controllers.NewUpdateAttendanceController(UpdateAttendenceUseCase)
    // Initialize routes
    EventAttendeesRoutes(router, EventAttendeesHandlers{
        register:      registerController,
        remove:        removeController,
        getAttendees:  getAttendeesController,
        getUserEvents: getUserEventsController,
        isRegistered:  isRegisteredController,
        updateAttendance: UpdateAttendenceController,
    })
}