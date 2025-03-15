package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/M1keTrike/Accessgate/infrastructure/firebase"
	"github.com/M1keTrike/Accessgate/infrastructure/router"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	err = firebase.InitializeFirebase()
	if err != nil {
		log.Fatalf("Error al inicializar Firebase: %v", err)
	}

	http.HandleFunc("/webhook", router.WebhookHandler)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8085"
	}

	fmt.Println("🚀 Servidor corriendo en el puerto:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
