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