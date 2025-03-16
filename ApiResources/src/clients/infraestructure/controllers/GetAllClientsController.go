package controllers

import (
	"api_resources/src/clients/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllClientsController struct {
	useCase *application.GetAllClientsUseCase
}

func NewGetAllClientsController(useCase *application.GetAllClientsUseCase) *GetAllClientsController {
	return &GetAllClientsController{useCase: useCase}
}

func (ctrl *GetAllClientsController) Execute(c *gin.Context) {
	clients, err := ctrl.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, clients)
}
