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

// Execute godoc
// @Summary      Elimina una tarjeta NFC
// @Description  Elimina una tarjeta NFC usando su UID.
// @Tags         NFC Cards
// @Produce      json
// @Param        uid  path      string            true  "UID de la tarjeta NFC"
// @Success      200  {string}  string            "OK"
// @Failure      500  {object}  map[string]string "Error interno del servidor"
// @Router       /nfc_cards/{uid} [delete]
func (c *DeleteNfcCardController) Execute(ctx *gin.Context) {
	uid := ctx.Param("uid")
	if err := c.useCase.Execute(uid); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
