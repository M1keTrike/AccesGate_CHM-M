package controllers

import (
	"api_resources/src/Users/application"
	"api_resources/src/Users/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserAdminController struct {
	useCase *application.CreateUserAdminUseCase
}

func NewCreateUserAdminController(useCase *application.CreateUserAdminUseCase) *CreateUserAdminController {
	return &CreateUserAdminController{useCase: useCase}
}

func (c *CreateUserAdminController) Execute(ctx *gin.Context) {
	var user entities.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if err := c.useCase.Execute(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create admin user"})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}