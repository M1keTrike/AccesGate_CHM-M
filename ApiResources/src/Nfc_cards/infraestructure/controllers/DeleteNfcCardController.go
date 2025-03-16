package controllers

import (
	"api_resources/src/Nfc_cards/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteNfcCardController struct {
	useCase application.DeleteNfcCard
}

func NewDeleteNfcCardController(useCase application.DeleteNfcCard) *DeleteNfcCardController {
	return &DeleteNfcCardController{useCase: useCase}
}

func (c *DeleteNfcCardController) Execute(ctx *gin.Context) {
	uid := ctx.Param("uid")
	if err := c.useCase.Execute(uid); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
