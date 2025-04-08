package controllers

import (
	"api_resources/src/Users/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetUsersByCreatedByController struct {
	useCase *application.GetUsersByCreatedByUseCase
}

func NewGetUsersByCreatedByController(useCase *application.GetUsersByCreatedByUseCase) *GetUsersByCreatedByController {
	return &GetUsersByCreatedByController{
		useCase: useCase,
	}
}

func (c *GetUsersByCreatedByController) Execute(ctx *gin.Context) {
	createdByStr := ctx.Param("created_by")
	createdBy, err := strconv.Atoi(createdByStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid created_by parameter"})
		return
	}

	users, err := c.useCase.Execute(createdBy)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
