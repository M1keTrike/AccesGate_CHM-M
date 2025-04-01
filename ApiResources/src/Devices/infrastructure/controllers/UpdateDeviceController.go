package controllers

import (
	"api_resources/src/Devices/application"
	"api_resources/src/Devices/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateDeviceController struct {
	useCase *application.UpdateDevice
}

func NewUpdateDeviceController(useCase *application.UpdateDevice) *UpdateDeviceController {
	return &UpdateDeviceController{useCase: useCase}
}

// UpdateDevice godoc
// @Summary Actualiza un dispositivo existente
// @Tags Devices
// @Accept json
// @Produce json
// @Param id path int true "ID del dispositivo"
// @Param device body entities.Device true "Dispositivo actualizado"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Security BearerAuth
// @Router /devices/{id} [put]
func (c *UpdateDeviceController) Execute(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var device entities.Device
	if err := ctx.ShouldBindJSON(&device); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	device.ID = id

	if err := c.useCase.Execute(&device); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar dispositivo"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Dispositivo actualizado correctamente"})
}
