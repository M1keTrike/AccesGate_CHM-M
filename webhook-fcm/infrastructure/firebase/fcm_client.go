package firebase

import (
	"context"
	"fmt"
	"log"

	"firebase.google.com/go/messaging"
)

type FCMMessage struct {
	Message struct {
		Topic        string            `json:"topic"`
		Data         map[string]string `json:"data,omitempty"`
		Notification struct {
			Title string `json:"title"`
			Body  string `json:"body"`
		} `json:"notification,omitempty"`
	} `json:"message"`
}

func SendNotification(title, body string) error {
	if App == nil {
		return fmt.Errorf("firebase no ha sido inicializado")
	}

	ctx := context.Background()
	client, err := App.Messaging(ctx)
	if err != nil {
		return fmt.Errorf("error obteniendo cliente de mensajería: %v", err)
	}

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Topic: "all", 
	}


	response, err := client.Send(ctx, message)
	if err != nil {
		return fmt.Errorf("error enviando notificación: %v", err)
	}

	log.Println("✅ Notificación enviada con éxito:", response)
	return nil
}
