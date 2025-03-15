package controllers

import (
	"api_resources/src/Nfc_cards/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetNfcCardByUIDController struct {
	useCase application.GetNfcCardByUID
}

func NewGetNfcCardByUIDController(useCase application.GetNfcCardByUID) *GetNfcCardByUIDController {
	return &GetNfcCardByUIDController{useCase: useCase}
}

func (c *GetNfcCardByUIDController) Execute(ctx *gin.Context) {
	uid := ctx.Param("uid")
	card, err := c.useCase.Execute(uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, card)
}
