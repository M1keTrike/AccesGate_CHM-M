package controllers

import (
	"api_resources/src/Users/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetUserByIDController struct {
	useCase application.GetUserByID
}

func NewGetUserByIDController(useCase application.GetUserByID) *GetUserByIDController {
	return &GetUserByIDController{useCase: useCase}
}

func (c *GetUserByIDController) Execute(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.useCase.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
