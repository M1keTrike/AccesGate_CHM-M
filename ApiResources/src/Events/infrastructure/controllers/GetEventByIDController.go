package controllers

import (
	"api_resources/src/Events/application"
	
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetEventByIDController struct {
	useCase *application.GetEventByID
}

func NewGetEventByIDController(useCase *application.GetEventByID) *GetEventByIDController {
	return &GetEventByIDController{useCase: useCase}
}

func (c *GetEventByIDController) Handle(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event, err := c.useCase.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, event)
}
