package controllers

import (
	"api_resources/src/Users/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllUsersController struct {
	useCase application.GetAllUsers
}

func NewGetAllUsersController(useCase application.GetAllUsers) *GetAllUsersController {
	return &GetAllUsersController{useCase: useCase}
}

func (c *GetAllUsersController) Execute(ctx *gin.Context) {
	users, err := c.useCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
