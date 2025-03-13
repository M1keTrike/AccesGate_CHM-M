package main

import (
	users "api_resources/src/Users/infraestructure"
	Nfc_cards "api_resources/src/Nfc_cards/infraestructure"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/gin-contrib/cors"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
		}))
	
	users.Init(router)
	Nfc_cards.Init(router)	
	router.Run(":8080")
}