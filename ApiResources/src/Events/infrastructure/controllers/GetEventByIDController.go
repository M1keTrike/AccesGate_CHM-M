package controllers

import (
	"api_resources/src/Events/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetEventByIDController struct {
	useCase *application.GetEventByID
}

func NewGetEventByIDController(useCase *application.GetEventByID) *GetEventByIDController {
	return &GetEventByIDController{useCase: useCase}
}

// GetEventByID godoc
// @Summary Obtiene un evento por ID
// @Tags Events
// @Produce json
// @Param id path int true "ID del evento"
// @Success 200 {object} entities.Event
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Security BearerAuth
// @Router /events/{id} [get]
func (c *GetEventByIDController) Execute(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	event, err := c.useCase.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Evento no encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, event)
}
