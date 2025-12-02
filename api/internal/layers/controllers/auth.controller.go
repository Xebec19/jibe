package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Xebec19/jibe/api/internal/common/dto"
	"github.com/Xebec19/jibe/api/internal/common/schema"
	"github.com/Xebec19/jibe/api/internal/layers/domain"
	"github.com/Xebec19/jibe/api/internal/layers/services"
	"github.com/Xebec19/jibe/api/pkg/config"
	"github.com/Xebec19/jibe/api/pkg/logger"
)

type AuthController interface {
	// GenerateNonce returns a random nonce and also save it in db with expiry
	GenerateNonce(w http.ResponseWriter, r *http.Request)
	// VerifyHandler verifies the message and signature generated during SIWE
	VerifyHandler(w http.ResponseWriter, r *http.Request)
}

func NewAuthController(logger *logger.Logger, cfg *config.Config, validator schema.RequestValidator, authService services.AuthService) AuthController {
	return authController{
		logger:      *logger,
		validator:   validator,
		cfg:         cfg,
		authService: authService,
	}
}

type authController struct {
	logger      logger.Logger
	validator   schema.RequestValidator
	authService services.AuthService
	cfg         *config.Config
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

func (a authController) VerifyHandler(w http.ResponseWriter, r *http.Request) {

	var req domain.VerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		a.logger.Error("Invalid request body", err)
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	verified, err := a.authService.VerifySignature(req.Message, req.Signature)
	if err != nil {
		a.logger.Error("error: message verification failed %w", err)
		respondError(w, http.StatusBadRequest, "message verification failed")
		return
	}
	if !verified {
		a.logger.Error(fmt.Sprintf("message %s is invalid", req.Message))
		respondError(w, http.StatusBadRequest, "invalid message")
		return
	}

	a.logger.Info(fmt.Sprintf("message %s verified successfully", req.Message))

	respondJSON(w, http.StatusOK, "Message is verified", domain.VerifyResponse{
		Valid: true,
	})
}
