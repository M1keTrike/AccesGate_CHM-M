package controllers

import (
	"api_resources/src/Users/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteUserController struct {
	useCase application.DeleteUser
}

func NewDeleteUserController(useCase application.DeleteUser) *DeleteUserController {
	return &DeleteUserController{useCase: useCase}
}

func (c *DeleteUserController) Execute(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = c.useCase.Execute(id);
	if  err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
