package main

import (

	users "api_resources/src/Users/infraestructure"
	Nfc_cards "api_resources/src/Nfc_cards/infraestructure"
	clients "api_resources/src/clients/infraestructure"
	event_attendees "api_resources/src/EventAttendees/infraestructure"
	Events "api_resources/src/Events/infrastructure"
	nfc_assignments "api_resources/src/NfcCardAssignments/infrastructure"
	

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"



	_ "api_resources/docs" // 👈 Swagger docs

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title AccessGate API
// @version 1.0
// @description API REST con arquitectura hexagonal para gestión de usuarios, tarjetas NFC y clientes.
// @termsOfService https://accessgate.com/terms/

// @contact.name Equipo de Desarrollo AccessGate
// @contact.email contacto@accessgate.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8084
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	router := gin.Default()

	// CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Rutas principales
	users.Init(router)
	Nfc_cards.Init(router)
	clients.Init(router)
	event_attendees.Init(router)
	Events.Init(router)
	nfc_assignments.Init(router)
	//router.Run(":8080")


	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Inicia servidor
	router.Run(":8084")
}
