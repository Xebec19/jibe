package routes

import (
	"github.com/Xebec19/jibe/api/internal/layers/container"
	"github.com/Xebec19/jibe/api/internal/layers/controllers"
	"github.com/gorilla/mux"
)

func registerHealthRoutes(r *mux.Router, c container.Container) {

	healthController := controllers.NewHealthController(&c.Logger)

	api := r.PathPrefix("/health").Subrouter()

	api.HandleFunc("/", healthController.GetHealthCheckpoint).Methods("GET")
}
