package server

import (
	"github.com/Xebec19/jibe/api/internal/auth"
	"github.com/Xebec19/jibe/api/internal/middleware"
	"github.com/gorilla/mux"
)

func createRoutes() *mux.Router {
	r := mux.NewRouter()

	// Add middleware to all routes
	r.Use(middleware.Recovery)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.SecurityHeaders)
	r.Use(middleware.CORS)

	// Health check endpoints
	r.HandleFunc("/health", healthCheckHandler).Methods("GET")
	r.HandleFunc("/ready", healthCheckHandler).Methods("GET") // You might want to make this more sophisticated

	// API routes
	api := r.PathPrefix("/api").Subrouter()
	auth.CreateRoutes(api)

	return r
}
