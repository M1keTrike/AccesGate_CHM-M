package infraestructure

import (
	"api_resources/src/Nfc_cards/application"
	"api_resources/src/Nfc_cards/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	os := NewPostgreSQL()

	createNfcCardService := application.NewCreateNfcCard(os)
	getNfcCardService := application.NewGetNfcCardByUID(os)
	deleteNfcCardService := application.NewDeleteNfcCard(os)
	getAllNfcCardsService := application.NewGetAllNfcCards(os)

	createNfcCardController := controllers.NewCreateNfcCardController(*createNfcCardService)
	getNfcCardController := controllers.NewGetNfcCardByUIDController(*getNfcCardService)
	deleteNfcCardController := controllers.NewDeleteNfcCardController(*deleteNfcCardService)
	getAllNfcCardsController := controllers.NewGetAllNfcCardsController(*getAllNfcCardsService)

	NfcCardsRoutes(router, NfcCardsHandlers{
		create: createNfcCardController,
		get:    getNfcCardController,
		delete: deleteNfcCardController,
		getAll: getAllNfcCardsController,
	})
}
