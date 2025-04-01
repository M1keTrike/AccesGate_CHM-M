package controllers

import (
    "api_resources/src/EventAttendees/application"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type GetUserEventsController struct {
    useCase *application.GetUserEvents
}

func NewGetUserEventsController(useCase *application.GetUserEvents) *GetUserEventsController {
    return &GetUserEventsController{useCase: useCase}
}

// GetUserEvents godoc
// @Summary Get all events for a user
// @Description Retrieves all events that a user is registered to attend
// @Tags Event Attendees
// @Produce json
// @Param userId path int true "User ID"
// @Success 200 {array} entities.Event
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /api/event-attendees/user/{userId}/events [get]
func (c *GetUserEventsController) Execute(ctx *gin.Context) {
    userID, err := strconv.Atoi(ctx.Param("userId"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    events, err := c.useCase.Execute(userID)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, events)
}