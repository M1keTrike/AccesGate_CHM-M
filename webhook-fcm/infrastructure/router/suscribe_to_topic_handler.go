package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/M1keTrike/Accessgate/application"
	"github.com/M1keTrike/Accessgate/domain"
)

func SubscribeToTopicHandler(w http.ResponseWriter, r *http.Request) {
	var req domain.SubscribeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "🚨 Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	if req.Token == "" {
		http.Error(w, "🚨 Token de registro requerido", http.StatusBadRequest)
		return
	}

	service := application.NewSuscribeToTokenService()
	err = service.Execute(req.Token)
	if err != nil {
		http.Error(w, fmt.Sprintf("🚨 Error en la suscripción: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "✅ Suscripción exitosa al tema 'all'"}`))
}
