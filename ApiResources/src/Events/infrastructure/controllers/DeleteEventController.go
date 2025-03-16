package controllers

import (
	"api_resources/src/Events/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteEventController struct {
	useCase *application.DeleteEvent
}

func NewDeleteEventController(useCase *application.DeleteEvent) *DeleteEventController {
	return &DeleteEventController{useCase: useCase}
}

func (c *DeleteEventController) Handle(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.useCase.Execute(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
