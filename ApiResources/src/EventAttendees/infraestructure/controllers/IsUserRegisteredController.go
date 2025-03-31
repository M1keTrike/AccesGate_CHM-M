package controllers

import (
    "api_resources/src/EventAttendees/application"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type IsUserRegisteredController struct {
    useCase *application.IsUserRegistered
}

func NewIsUserRegisteredController(useCase *application.IsUserRegistered) *IsUserRegisteredController {
    return &IsUserRegisteredController{useCase: useCase}
}

func (c *IsUserRegisteredController) Execute(ctx *gin.Context) {
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

    isRegistered, err := c.useCase.Execute(eventID, userID)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"isRegistered": isRegistered})
}