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

// DeleteEvent godoc
// @Summary Elimina un evento
// @Tags Events
// @Produce json
// @Param id path int true "ID del evento"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /events/{id} [delete]
func (c *DeleteEventController) Execute(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := c.useCase.Execute(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar el evento"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Evento eliminado correctamente"})
}
