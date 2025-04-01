package controllers

import (
	"api_resources/src/AccessEvents/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetAccessEventByIDController struct {
	useCase application.GetAccessEventByID
}

func NewGetAccessEventByIDController(useCase application.GetAccessEventByID) *GetAccessEventByIDController {
	return &GetAccessEventByIDController{useCase: useCase}
}


// Execute godoc
// @Summary Obtiene un evento de acceso por ID
// @Tags Access Events
// @Produce json
// @Param id path int true "ID del evento de acceso"
// @Success 200 {object} entities.AccessEvent
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Security BearerAuth
// @Router /access-events/{id} [get]
func (c *GetAccessEventByIDController) Execute(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	event, err := c.useCase.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, event)
}
