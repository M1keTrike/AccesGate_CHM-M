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

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)

		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("🚨 Error cargando el archivo .env")
	}

	fmt.Println("🔍 FIREBASE_CREDENTIALS_PATH:", os.Getenv("FIREBASE_CREDENTIALS_PATH"))
	fmt.Println("🔍 FCM_PROJECT_ID:", os.Getenv("FCM_PROJECT_ID"))
	fmt.Println("🔍 SERVER_PORT:", os.Getenv("SERVER_PORT"))

	err = firebase.InitializeFirebase()
	if err != nil {
		log.Fatalf("🚨 Error al inicializar Firebase: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/webhook", router.WebhookHandler)
	mux.HandleFunc("/token", router.TokenHandler)
	mux.HandleFunc("/subscribe", router.SubscribeToTopicHandler)

	handlerWithCORS := corsMiddleware(mux)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("🚀 Servidor corriendo en el puerto:", port)
	log.Fatal(http.ListenAndServe(":"+port, handlerWithCORS))
}
