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

// Execute godoc
// @Summary      Crea una nueva tarjeta NFC
// @Description  Registra una nueva tarjeta NFC con UID y estado.
// @Tags         NFC Cards
// @Accept       json
// @Produce      json
// @Param        card  body      entities.NfcCard  true  "Datos de la tarjeta NFC"
// @Success      201   {string}  string            "Created"
// @Failure      400   {object}  map[string]string "Error de validación de entrada"
// @Failure      500   {object}  map[string]string "Error interno del servidor"
// @Router       /nfc_cards [post]
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
