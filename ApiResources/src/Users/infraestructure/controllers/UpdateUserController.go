package controllers

import (
	"api_resources/src/Users/application"
	"api_resources/src/Users/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateUserController struct {
	useCase application.UpdateUser
}

func NewUpdateUserController(useCase application.UpdateUser) *UpdateUserController {
	return &UpdateUserController{useCase: useCase}
}

func (c *UpdateUserController) Execute(ctx *gin.Context) {
	var user entities.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.useCase.Execute(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
