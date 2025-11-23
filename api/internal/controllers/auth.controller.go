package controllers

import (
	"net/http"

	"github.com/Xebec19/jibe/api/internal/services"
	"github.com/Xebec19/jibe/api/pkg/logger"
)

type AuthController interface {
	GenerateNonce(w http.ResponseWriter, r *http.Request)
}

func NewAuthController(logger *logger.Logger, authService services.AuthService) AuthController {
	return authController{
		logger:      logger,
		authService: authService,
	}
}

type authController struct {
	logger      *logger.Logger
	authService services.AuthService
}

func (a authController) GenerateNonce(w http.ResponseWriter, r *http.Request) {

	// a.authService.CreateNonce()

}
