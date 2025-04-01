package controllers

import (
	"api_resources/src/AccessEvents/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteAccessEventController struct {
	useCase application.DeleteAccessEvent
}

func NewDeleteAccessEventController(useCase application.DeleteAccessEvent) *DeleteAccessEventController {
	return &DeleteAccessEventController{useCase: useCase}
}

// Execute godoc
// @Summary Elimina un evento de acceso
// @Tags Access Events
// @Produce json
// @Param id path int true "ID del evento de acceso"
// @Success 200 {string} string "OK"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /access-events/{id} [delete]
func (c *DeleteAccessEventController) Execute(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := c.useCase.Execute(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
