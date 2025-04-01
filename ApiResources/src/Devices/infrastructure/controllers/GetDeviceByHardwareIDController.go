package controllers

import (
	"api_resources/src/Devices/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetDeviceByHardwareIDController struct {
	useCase *application.GetDeviceByHardwareID
}

func NewGetDeviceByHardwareIDController(useCase *application.GetDeviceByHardwareID) *GetDeviceByHardwareIDController {
	return &GetDeviceByHardwareIDController{useCase: useCase}
}

// GetDeviceByHardwareID godoc
// @Summary Obtiene un dispositivo por hardware_id
// @Tags Devices
// @Produce json
// @Param hardware_id path string true "Identificador físico del dispositivo"
// @Success 200 {object} entities.Device
// @Failure 404 {object} map[string]string
// @Security BearerAuth
// @Router /devices/hardware/{hardware_id} [get]
func (c *GetDeviceByHardwareIDController) Execute(ctx *gin.Context) {
	hardwareID := ctx.Param("hardware_id")
	device, err := c.useCase.Execute(hardwareID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, device)
}
