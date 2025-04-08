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

// GetEventAttendees godoc
// @Summary Get all attendees for an event
// @Description Retrieves all users registered to attend a specific event
// @Tags Event Attendees
// @Produce json
// @Param eventId path int true "Event ID"
// @Success 200 {array} entities.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /event-attendees/events/{eventId}/attendees [get]
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