package controllers

import (
    "api_resources/src/EventAttendees/application"
    "api_resources/src/EventAttendees/domain/entities"
    "github.com/gin-gonic/gin"
    "net/http"
)

type RegisterAttendeeController struct {
    useCase *application.RegisterAttendee
}

func NewRegisterAttendeeController(useCase *application.RegisterAttendee) *RegisterAttendeeController {
    return &RegisterAttendeeController{useCase: useCase}
}

func (c *RegisterAttendeeController) Execute(ctx *gin.Context) {
    var attendee entities.EventAttendee
    if err := ctx.ShouldBindJSON(&attendee); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.useCase.Execute(&attendee); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, attendee)
}