package handlers

import (
	"encoding/json"
	"net/http"
)

// Response represents a standard API response
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
}

// HealthCheck handles health check requests
// @Summary      Health check
// @Description  Check if the API server is running
// @Tags         health
// @Produce      json
// @Success      200  {object}  Response
// @Router       /health [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, Response{
		Success: true,
		Data: map[string]string{
			"status": "healthy",
		},
	})
}

// Root handles root path requests
// @Summary      API root
// @Description  Get API information and version
// @Tags         root
// @Produce      json
// @Success      200  {object}  Response
// @Router       / [get]
func Root(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, Response{
		Success: true,
		Message: "Welcome to Jibe API",
		Data: map[string]string{
			"version": "1.0.0",
		},
	})
}


// respondJSON sends a JSON response
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

// respondError sends an error response
func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, Response{
		Success: false,
		Error:   message,
	})
}
