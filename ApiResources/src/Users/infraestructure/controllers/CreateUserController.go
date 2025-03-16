package controllers

import (
	"api_resources/src/Users/application"
	"api_resources/src/Users/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserController struct {
	useCase application.CreateUser
}

func NewCreateUserController(useCase application.CreateUser) *CreateUserController {
	return &CreateUserController{useCase: useCase}
}

func (c *CreateUserController) Execute(ctx *gin.Context) {
	var user entities.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.useCase.Execute(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}
