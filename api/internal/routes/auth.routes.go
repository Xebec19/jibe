package routes

import (
	"github.com/Xebec19/jibe/api/internal/container"
	"github.com/Xebec19/jibe/api/internal/controllers"
	"github.com/gorilla/mux"
)

func registerAuthRoutes(r *mux.Router, c container.Container) {

	authController := controllers.NewAuthController(&c.Logger, c.AuthService)

	authApi := r.PathPrefix("/v1/auth").Subrouter()

	authApi.HandleFunc("/generate-nonce", authController.GenerateNonce).Methods("POST")
}
