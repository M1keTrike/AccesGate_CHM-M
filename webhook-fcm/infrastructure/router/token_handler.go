package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/M1keTrike/Accessgate/infrastructure/firebase"
)

func TokenHandler(w http.ResponseWriter, r *http.Request) {
	token, err := firebase.GetAccessToken()
	if err != nil {
		log.Println("🚨 Error obteniendo OAuth2 Token:", err)
		http.Error(w, "Error obteniendo token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
