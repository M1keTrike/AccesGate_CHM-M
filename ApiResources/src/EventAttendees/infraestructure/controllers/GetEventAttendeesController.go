package controllers

import (
    "api_resources/src/EventAttendees/application"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type GetEventAttendeesController struct {
    useCase *application.GetEventAttendees
}

func NewGetEventAttendeesController(useCase *application.GetEventAttendees) *GetEventAttendeesController {
    return &GetEventAttendeesController{useCase: useCase}
}

func (c *GetEventAttendeesController) Execute(ctx *gin.Context) {
    eventID, err := strconv.Atoi(ctx.Param("eventId"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
        return
    }

    attendees, err := c.useCase.Execute(eventID)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, attendees)
}