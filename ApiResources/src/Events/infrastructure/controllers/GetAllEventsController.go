package controllers

import (
	"api_resources/src/Events/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllEventsController struct {
	useCase *application.GetAllEvents
}

func NewGetAllEventsController(useCase *application.GetAllEvents) *GetAllEventsController {
	return &GetAllEventsController{useCase: useCase}
}

func (c *GetAllEventsController) Handle(ctx *gin.Context) {
	events, err := c.useCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, events)
}
