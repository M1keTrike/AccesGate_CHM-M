package firebase

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var App *firebase.App

func InitializeFirebase() error {
	credentialsPath := os.Getenv("FIREBASE_CREDENTIALS_PATH")
	if credentialsPath == "" {
		return fmt.Errorf("la variable de entorno FIREBASE_CREDENTIALS_PATH no está configurada")
	}

	opt := option.WithCredentialsFile(credentialsPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return fmt.Errorf("error inicializando Firebase: %v", err)
	}

	App = app
	log.Println("✅ Firebase inicializado correctamente")
	return nil
}
