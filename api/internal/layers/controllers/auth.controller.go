package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Xebec19/jibe/api/internal/common/dto"
	"github.com/Xebec19/jibe/api/internal/common/schema"
	"github.com/Xebec19/jibe/api/internal/layers/services"
	"github.com/Xebec19/jibe/api/pkg/logger"
)

type AuthController interface {
	GenerateNonce(w http.ResponseWriter, r *http.Request)
}

func NewAuthController(logger *logger.Logger, validator schema.RequestValidator, authService services.AuthService) AuthController {
	return authController{
		logger:      *logger,
		validator:   validator,
		authService: authService,
	}
}

type authController struct {
	logger      logger.Logger
	validator   schema.RequestValidator
	authService services.AuthService
}

func (a authController) GenerateNonce(w http.ResponseWriter, r *http.Request) {

	var req dto.GenerateNonceDTO

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		a.logger.Error("request body parsing failed for creating nonce", "error", err)
		respondError(w, http.StatusBadRequest, INVALID_REQUEST_MSG)
		return
	}

	err = a.validator.Validate(req)
	if err != nil {
		a.logger.Error("invalid req body for creating nonce", "error", a.validator.FormatErrors(err))
		respondError(w, http.StatusBadRequest, INVALID_REQUEST_MSG)
		return
	}

	nonce, err := a.authService.CreateNonce(req.Eth_Addr)
	if err != nil {
		a.logger.Info("nonce creation failed", "error", err)
		respondError(w, http.StatusInternalServerError, SOMETHING_WENT_WRONG_MSG)
		return
	}

	payload := &dto.GenerateNonceResponseDTO{
		Nonce: nonce,
	}

	respondJSON(w, http.StatusCreated, RESOURCE_CREATED_MSG, payload)

}
