package main

import (
	event_attendees "api_resources/src/EventAttendees/infraestructure"

	nfc_assignments "api_resources/src/NfcCardAssignments/infrastructure"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	access_events "api_resources/src/AccessEvents/infrastructure"
	devices "api_resources/src/Devices/infrastructure"
	events "api_resources/src/Events/infrastructure"
	nfc_cards "api_resources/src/Nfc_cards/infraestructure"
	users "api_resources/src/Users/infraestructure"
	clients "api_resources/src/clients/infraestructure"

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
	devices.Init(router)

	clients.Init(router)
	nfc_cards.Init(router)
	events.Init(router)
	event_attendees.Init(router)
	access_events.Init(router)

	nfc_assignments.Init(router)

	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Inicia servidor
	router.Run(":8084")
}
