package handlers

import (
	"encoding/json"
	"net/http"
)

// HealthResponse é a estrutura da resposta do health check
type HealthResponse struct {
	Status string `json:"status"`
}

// HealthCheck é o handler para a rota de health check
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status: "OK",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
