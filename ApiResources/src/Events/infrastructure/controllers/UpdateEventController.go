package controllers

import (
	"api_resources/src/Events/application"
	"api_resources/src/Events/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateEventController struct {
	useCase *application.UpdateEvent
}

func NewUpdateEventController(useCase *application.UpdateEvent) *UpdateEventController {
	return &UpdateEventController{useCase: useCase}
}

func (c *UpdateEventController) Handle(ctx *gin.Context) {
	var event entities.Event
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.useCase.Execute(event); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
