package controllers

import (
	"api_resources/src/clients/application"
	"api_resources/src/clients/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateClientController struct {
	useCase *application.CreateClientUseCase
}

func NewCreateClientController(useCase *application.CreateClientUseCase) *CreateClientController {
	return &CreateClientController{useCase: useCase}
}

func (ctrl *CreateClientController) Execute(c *gin.Context) {
	var client entities.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.useCase.Execute(&client); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, client)
}
