package infraestructure

import (
	"api_resources/src/Nfc_cards/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

type NfcCardsHandlers struct {
	create *controllers.CreateNfcCardController
	get    *controllers.GetNfcCardByUIDController
	delete *controllers.DeleteNfcCardController
	getAll *controllers.GetAllNfcCardsController
}

func NfcCardsRoutes(router *gin.Engine, handlers NfcCardsHandlers) {
	router.POST("/nfc_cards", handlers.create.Execute)
	router.GET("/nfc_cards/:uid", handlers.get.Execute)
	router.DELETE("/nfc_cards/:uid", handlers.delete.Execute)
	router.GET("/nfc_cards", handlers.getAll.Execute)
}
