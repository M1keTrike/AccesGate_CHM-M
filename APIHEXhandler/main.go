package main

import (
	"log"
	"os"

	aglomerationDeps "github.com/M1keTrike/EventDriven/messages_aglomeration/dependencies"
	"github.com/M1keTrike/EventDriven/messages_nfc/dependencies"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	r := gin.Default()

	messageNFCDependencies := dependencies.NewMessageDependencies()
	messageNFCDependencies.Execute(r)

	messageAglomerationDependencies := aglomerationDeps.NewMessageAgloemrationDependencies()
	messageAglomerationDependencies.Execute(r)
	r.Run(":" + PORT)

}
