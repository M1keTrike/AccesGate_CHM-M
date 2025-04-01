package controllers

import (
	"api_resources/src/AccessEvents/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetAccessEventsByDeviceController struct {
	useCase application.GetAccessEventsByDevice
}

func NewGetAccessEventsByDeviceController(useCase application.GetAccessEventsByDevice) *GetAccessEventsByDeviceController {
	return &GetAccessEventsByDeviceController{useCase: useCase}
}

// GetAccessEventsByDevice godoc
// @Summary Obtiene eventos de acceso por dispositivo
// @Tags Access Events
// @Produce json
// @Param deviceId path int true "ID del dispositivo"
// @Success 200 {array} entities.AccessEvent
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /access-events/device/{deviceId} [get]
func (c *GetAccessEventsByDeviceController) Execute(ctx *gin.Context) {
	deviceID, err := strconv.Atoi(ctx.Param("deviceId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "deviceId inválido"})
		return
	}

	events, err := c.useCase.Execute(deviceID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, events)
}
