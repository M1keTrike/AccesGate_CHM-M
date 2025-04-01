package controllers

import (
	"api_resources/src/Events/application"
	"api_resources/src/Events/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateEventController struct {
	useCase *application.UpdateEvent
}

func NewUpdateEventController(useCase *application.UpdateEvent) *UpdateEventController {
	return &UpdateEventController{useCase: useCase}
}

// UpdateEvent godoc
// @Summary Actualiza un evento
// @Tags Events
// @Accept json
// @Produce json
// @Param id path int true "ID del evento"
// @Param event body entities.Event true "Datos del evento"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /events/{id} [put]
func (c *UpdateEventController) Execute(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var event entities.Event
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	event.ID = id

	if err := c.useCase.Execute(&event); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el evento"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Evento actualizado correctamente"})
}
