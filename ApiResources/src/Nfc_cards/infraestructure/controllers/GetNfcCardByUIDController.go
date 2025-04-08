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

// Execute godoc
// @Summary      Obtiene una tarjeta NFC por UID
// @Description  Retorna los datos de una tarjeta NFC específica según su UID.
// @Tags         NFC Cards
// @Produce      json
// @Param        uid  path      string           true  "UID de la tarjeta NFC"
// @Success      200  {object}  entities.NfcCard
// @Failure      500  {object}  map[string]string "Error interno del servidor"
// @Router       /nfc_cards/{uid} [get]
func (c *GetNfcCardByUIDController) Execute(ctx *gin.Context) {
	uid := ctx.Param("uid")
	card, err := c.useCase.Execute(uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, card)
}
