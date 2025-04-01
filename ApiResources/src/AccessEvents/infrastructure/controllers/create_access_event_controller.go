package controllers

import (
	"api_resources/src/AccessEvents/application"
	"api_resources/src/AccessEvents/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateAccessEventController struct {
	useCase application.CreateAccessEvent
}

func NewCreateAccessEventController(useCase application.CreateAccessEvent) *CreateAccessEventController {
	return &CreateAccessEventController{useCase: useCase}
}

// Execute godoc
// @Summary Crea un nuevo evento de acceso
// @Tags Access Events
// @Accept json
// @Produce json
// @Param event body entities.AccessEvent true "Evento de acceso a registrar"
// @Success 201 {object} entities.AccessEvent
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /access-events [post]
func (c *CreateAccessEventController) Execute(ctx *gin.Context) {
	var event entities.AccessEvent
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.useCase.Execute(&event); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, event)
}
