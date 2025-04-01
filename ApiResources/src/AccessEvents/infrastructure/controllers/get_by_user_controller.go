package controllers

import (
	"api_resources/src/AccessEvents/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetAccessEventsByUserController struct {
	useCase application.GetAccessEventsByUser
}

func NewGetAccessEventsByUserController(useCase application.GetAccessEventsByUser) *GetAccessEventsByUserController {
	return &GetAccessEventsByUserController{useCase: useCase}
}

// Execute godoc
// @Summary Obtiene eventos de acceso por usuario
// @Tags Access Events
// @Produce json
// @Param userId path int true "ID del usuario"
// @Success 200 {array} entities.AccessEvent
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /access-events/user/{userId} [get]
func (c *GetAccessEventsByUserController) Execute(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "userId inválido"})
		return
	}

	events, err := c.useCase.Execute(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, events)
}
