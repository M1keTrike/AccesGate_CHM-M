package controllers

import (
	"api_resources/src/Devices/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetDeviceByIDController struct {
	useCase *application.GetDeviceByID
}

func NewGetDeviceByIDController(useCase *application.GetDeviceByID) *GetDeviceByIDController {
	return &GetDeviceByIDController{useCase: useCase}
}

// GetDeviceByID godoc
// @Summary Obtiene un dispositivo por ID
// @Tags Devices
// @Produce json
// @Param id path int true "ID del dispositivo"
// @Success 200 {object} entities.Device
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Security BearerAuth
// @Router /devices/{id} [get]
func (c *GetDeviceByIDController) Execute(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	device, err := c.useCase.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, device)
}
