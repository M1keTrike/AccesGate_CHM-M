package main

import (
	users "api_resources/src/Users/infraestructure"
	Nfc_cards "api_resources/src/Nfc_cards/infraestructure"
	clients "api_resources/src/clients/infraestructure"
	event_attendees "api_resources/src/EventAttendees/infraestructure"
	Events "api_resources/src/Events/infrastructure"
	
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
		AllowOrigins: []string{"http://localhost:4200"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
		}))
	
	users.Init(router)
	Nfc_cards.Init(router)	
	clients.Init(router)
	event_attendees.Init(router)
	Events.Init(router)
	router.Run(":8080")
}