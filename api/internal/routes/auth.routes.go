package routes

import (
	"github.com/Xebec19/jibe/api/internal/layers/container"
	"github.com/Xebec19/jibe/api/internal/layers/controllers"
	"github.com/gorilla/mux"
)

func registerAuthRoutes(r *mux.Router, c container.Container) {

	authController := controllers.NewAuthController(&c.Logger, c.AuthService)

	authApi := r.PathPrefix("/v1/auth").Subrouter()

	authApi.HandleFunc("/generate-nonce", authController.GenerateNonce).Methods("POST")
}
