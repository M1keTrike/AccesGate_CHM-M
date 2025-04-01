package controllers

import (
	"api_resources/src/Devices/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllDevicesController struct {
	useCase *application.GetAllDevices
}

func NewGetAllDevicesController(useCase *application.GetAllDevices) *GetAllDevicesController {
	return &GetAllDevicesController{useCase: useCase}
}

// GetAllDevices godoc
// @Summary Obtiene todos los dispositivos
// @Tags Devices
// @Produce json
// @Success 200 {array} entities.Device
// @Security BearerAuth
// @Router /devices [get]
func (c *GetAllDevicesController) Execute(ctx *gin.Context) {
	devices, err := c.useCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, devices)
}
