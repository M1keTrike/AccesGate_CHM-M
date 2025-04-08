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

// GetAllEvents godoc
// @Summary Lista todos los eventos
// @Tags Events
// @Produce json
// @Success 200 {array} entities.Event
// @Security BearerAuth
// @Router /events [get]
func (c *GetAllEventsController) Execute(ctx *gin.Context) {
	events, err := c.useCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener los eventos"})
		return
	}

	ctx.JSON(http.StatusOK, events)
}
