package firebase

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/transport"
)

func GetAccessToken() (string, error) {
	ctx := context.Background()

	tokenFile := os.Getenv("FIREBASE_CREDENTIALS_PATH")
	if tokenFile == "" {
		return "", fmt.Errorf("🚨 ERROR: FIREBASE_CREDENTIALS_PATH no está definido en las variables de entorno")
	}

	if _, err := os.Stat(tokenFile); os.IsNotExist(err) {
		return "", fmt.Errorf("🚨 ERROR: El archivo de credenciales %s no existe", tokenFile)
	}

	creds, err := transport.Creds(ctx, option.WithCredentialsFile(tokenFile),
		option.WithScopes("https://www.googleapis.com/auth/firebase.messaging"))
	if err != nil {
		return "", fmt.Errorf("🚨 Error obteniendo credenciales: %v", err)
	}

	token, err := creds.TokenSource.Token()
	if err != nil {
		return "", fmt.Errorf("🚨 Error obteniendo token OAuth2: %v", err)
	}

	log.Println("✅ Token OAuth2 obtenido correctamente:", token.AccessToken[:20]+"...")
	return token.AccessToken, nil
}
