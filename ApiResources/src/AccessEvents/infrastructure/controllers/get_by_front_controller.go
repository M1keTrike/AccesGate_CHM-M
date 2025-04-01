package controllers

import (
	"api_resources/src/AccessEvents/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetAccessEventsByFrontController struct {
	useCase application.GetAccessEventsByFront
}

func NewGetAccessEventsByFrontController(useCase application.GetAccessEventsByFront) *GetAccessEventsByFrontController {
	return &GetAccessEventsByFrontController{useCase: useCase}
}

// Execute godoc
// @Summary Obtiene eventos de acceso por puerta (front)
// @Tags Access Events
// @Produce json
// @Param frontId path int true "ID del front (puerta)"
// @Success 200 {array} entities.AccessEvent
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /access-events/front/{frontId} [get]
func (c *GetAccessEventsByFrontController) Execute(ctx *gin.Context) {
	frontID, err := strconv.Atoi(ctx.Param("frontId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "frontId inválido"})
		return
	}

	events, err := c.useCase.Execute(frontID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, events)
}
