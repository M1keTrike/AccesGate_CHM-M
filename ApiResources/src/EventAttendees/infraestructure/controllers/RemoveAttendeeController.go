package controllers

import (
    "api_resources/src/EventAttendees/application"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type RemoveAttendeeController struct {
    useCase *application.RemoveAttendee
}

func NewRemoveAttendeeController(useCase *application.RemoveAttendee) *RemoveAttendeeController {
    return &RemoveAttendeeController{useCase: useCase}
}

// RemoveAttendee godoc
// @Summary Remove an attendee from an event
// @Description Removes a user's registration from a specific event
// @Tags Event Attendees
// @Produce json
// @Param eventId path int true "Event ID"
// @Param userId path int true "User ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /api/event-attendees/event/{eventId}/user/{userId} [delete]
func (c *RemoveAttendeeController) Execute(ctx *gin.Context) {
    eventID, err := strconv.Atoi(ctx.Param("eventId"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
        return
    }

    userID, err := strconv.Atoi(ctx.Param("userId"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    if err := c.useCase.Execute(eventID, userID); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.Status(http.StatusOK)
}