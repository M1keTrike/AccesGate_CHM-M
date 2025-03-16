package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/M1keTrike/Accessgate/application"
	"github.com/M1keTrike/Accessgate/domain"
)

var webhookService = application.NewWebhookService()

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var event domain.Event

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		log.Println("🚨 Error decodificando JSON:", err)
		http.Error(w, "Error en el formato JSON", http.StatusBadRequest)
		return
	}

	log.Println("📩 Webhook recibido:", event.Event)

	err = webhookService.ProcessEvent(event)
	if err != nil {
		log.Println("🚨 Error procesando el evento:", err) 
		http.Error(w, "Error procesando el evento", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("✅ Webhook procesado correctamente"))
}
