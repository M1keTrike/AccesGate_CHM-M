package controllers

import (
	"api_resources/src/AccessEvents/application"
	"api_resources/src/AccessEvents/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateAccessEventController struct {
	useCase application.UpdateAccessEvent
}

func NewUpdateAccessEventController(useCase application.UpdateAccessEvent) *UpdateAccessEventController {
	return &UpdateAccessEventController{useCase: useCase}
}

// Execute godoc
// @Summary Actualiza un evento de acceso
// @Tags Access Events
// @Accept json
// @Produce json
// @Param event body entities.AccessEvent true "Evento de acceso actualizado"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /access-events [put]
func (c *UpdateAccessEventController) Execute(ctx *gin.Context) {
	var event entities.AccessEvent
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.useCase.Execute(&event); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "actualizado correctamente"})
}
