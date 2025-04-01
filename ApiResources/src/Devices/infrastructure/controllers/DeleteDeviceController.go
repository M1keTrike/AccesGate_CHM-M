package controllers

import (
	"api_resources/src/Devices/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteDeviceController struct {
	useCase *application.DeleteDevice
}

func NewDeleteDeviceController(useCase *application.DeleteDevice) *DeleteDeviceController {
	return &DeleteDeviceController{useCase: useCase}
}

// DeleteDevice godoc
// @Summary Elimina un dispositivo por ID
// @Tags Devices
// @Produce json
// @Param id path int true "ID del dispositivo"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Security BearerAuth
// @Router /devices/{id} [delete]
func (c *DeleteDeviceController) Execute(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := c.useCase.Execute(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Dispositivo eliminado correctamente"})
}
