package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Xebec19/jibe/api/internal/lib/environment"
)

type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Version   string    `json:"version"`
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {

	config := environment.GetConfig()

	response := HealthResponse{
		Status:    "ok",
		Timestamp: time.Now(),
		Version:   config.Version,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
