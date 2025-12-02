package routes

import (
	"github.com/Xebec19/jibe/api/internal/layers/container"
	"github.com/Xebec19/jibe/api/internal/layers/controllers"
	"github.com/Xebec19/jibe/api/internal/middleware"
	"github.com/gorilla/mux"
)

func registerAuthRoutes(r *mux.Router, c container.Container) {

	authController := controllers.NewAuthController(&c.Logger, &c.Cfg, c.Validator, c.AuthService)

	authApi := r.PathPrefix("/v1/auth").Subrouter()

	authApi.Use(middleware.BodySizeLimit(c.Cfg.MaxBodySizeAllowed))

	authApi.HandleFunc("/generate-nonce", authController.GenerateNonce).Methods("POST")

	authApi.HandleFunc("/verify", authController.VerifyHandler).Methods("POST")

}
