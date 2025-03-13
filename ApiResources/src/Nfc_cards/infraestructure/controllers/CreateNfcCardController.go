package controllers

import (
	"api_resources/src/Nfc_cards/application"
	"api_resources/src/Nfc_cards/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateNfcCardController struct {
	useCase application.CreateNfcCard
}

func NewCreateNfcCardController(useCase application.CreateNfcCard) *CreateNfcCardController {
	return &CreateNfcCardController{useCase: useCase}
}

func (c *CreateNfcCardController) Execute(ctx *gin.Context) {
	var card entities.NfcCard
	if err := ctx.ShouldBindJSON(&card); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.useCase.Execute(&card); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}
