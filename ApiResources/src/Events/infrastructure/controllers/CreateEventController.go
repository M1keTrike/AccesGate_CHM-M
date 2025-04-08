package controllers

import (
	"api_resources/src/Events/application"
	"api_resources/src/Events/domain/entities"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateEventController struct {
	useCase *application.CreateEvent
}

func NewCreateEventController(useCase *application.CreateEvent) *CreateEventController {
	return &CreateEventController{useCase: useCase}
}

// CreateEvent godoc
// @Summary Crea un nuevo evento
// @Tags Events
// @Accept json
// @Produce json
// @Param event body entities.Event true "Evento a crear"
// @Success 201 {object} entities.Event
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /events [post]
func (c *CreateEventController) Execute(ctx *gin.Context) {
	var event entities.Event
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	event.CreatedAt = time.Now()

	if err := c.useCase.Execute(&event); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear el evento"})
		return
	}

	ctx.JSON(http.StatusCreated, event)
}
