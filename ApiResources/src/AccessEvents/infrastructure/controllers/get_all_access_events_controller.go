package controllers

import (
	"api_resources/src/AccessEvents/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllAccessEventsController struct {
	useCase application.GetAllAccessEvents
}

func NewGetAllAccessEventsController(useCase application.GetAllAccessEvents) *GetAllAccessEventsController {
	return &GetAllAccessEventsController{useCase: useCase}
}



// Execute godoc
// @Summary Lista todos los eventos de acceso
// @Tags Access Events
// @Produce json
// @Success 200 {array} entities.AccessEvent
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /access-events [get]
func (c *GetAllAccessEventsController) Execute(ctx *gin.Context) {
	events, err := c.useCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, events)
}
