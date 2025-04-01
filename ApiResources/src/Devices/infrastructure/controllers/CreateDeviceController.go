package controllers

import (
	"api_resources/src/Devices/application"
	"api_resources/src/Devices/domain/entities"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateDeviceController struct {
	useCase *application.CreateDevice
}

func NewCreateDeviceController(useCase *application.CreateDevice) *CreateDeviceController {
	return &CreateDeviceController{useCase: useCase}
}

// CreateDevice godoc
// @Summary Crea un nuevo dispositivo
// @Tags Devices
// @Accept json
// @Produce json
// @Param device body entities.Device true "Nuevo dispositivo"
// @Success 201 {object} entities.Device
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /devices [post]
func (c *CreateDeviceController) Execute(ctx *gin.Context) {
	var device entities.Device
	if err := ctx.ShouldBindJSON(&device); err != nil {
		log.Print(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	if err := c.useCase.Execute(&device); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, device)
}
