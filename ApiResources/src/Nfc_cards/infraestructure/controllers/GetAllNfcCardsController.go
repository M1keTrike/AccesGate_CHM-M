package controllers

import (
	"api_resources/src/Nfc_cards/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllNfcCardsController struct {
	useCase application.GetAllNfcCards
}

func NewGetAllNfcCardsController(useCase application.GetAllNfcCards) *GetAllNfcCardsController {
	return &GetAllNfcCardsController{useCase: useCase}
}

func (c *GetAllNfcCardsController) Execute(ctx *gin.Context) {
	cards, err := c.useCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, cards)
}
