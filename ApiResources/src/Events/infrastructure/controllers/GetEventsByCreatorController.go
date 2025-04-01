package controllers

import (
	"api_resources/src/Events/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetEventsByCreatorController struct {
	useCase *application.GetEventsByCreator
}

func NewGetEventsByCreatorController(useCase *application.GetEventsByCreator) *GetEventsByCreatorController {
	return &GetEventsByCreatorController{useCase: useCase}
}

// GetEventsByCreator godoc
// @Summary Lista eventos por creador
// @Tags Events
// @Produce json
// @Param user_id path int true "ID del creador"
// @Success 200 {array} entities.Event
// @Failure 400 {object} map[string]string
// @Security BearerAuth
// @Router /events/creator/{user_id} [get]
func (c *GetEventsByCreatorController) Execute(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	events, err := c.useCase.Execute(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener eventos"})
		return
	}

	ctx.JSON(http.StatusOK, events)
}
