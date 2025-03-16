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

func SendNotification(body string) error {
	if App == nil {
		log.Println("🚨 Firebase no está inicializado correctamente")
		return fmt.Errorf("firebase no está inicializado")
	}

	ctx := context.Background()
	client, err := App.Messaging(ctx)
	if err != nil {
		log.Println("🚨 Error obteniendo cliente de mensajería:", err)
		return fmt.Errorf("error obteniendo cliente de mensajería: %v", err)
	}

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "FALLO EN EL SISTEMA",
			Body:  body,
		},
		Topic: "all",
	}

	response, err := client.Send(ctx, message)
	if err != nil {
		log.Println("🚨 Error enviando mensaje a FCM:", err)
		return fmt.Errorf("error enviando notificación: %v", err)
	}

	log.Println("✅ Notificación enviada con éxito:", response)
	return nil
}

func SubscribeToTopic(registrationToken string, topic string) error {
	if App == nil {
		return fmt.Errorf("🚨 Firebase no está inicializado")
	}

	ctx := context.Background()
	client, err := App.Messaging(ctx)
	if err != nil {
		return fmt.Errorf("🚨 Error obteniendo cliente de mensajería: %v", err)
	}

	response, err := client.SubscribeToTopic(ctx, []string{registrationToken}, topic)
	if err != nil {
		return fmt.Errorf("🚨 Error suscribiendo token al tema %s: %v", topic, err)
	}

	log.Printf("✅ %d token(s) suscrito(s) al tema '%s'\n", response.SuccessCount, topic)
	return nil
}
