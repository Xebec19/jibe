package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/Xebec19/jibe/api/internal/common/dto"
	"github.com/Xebec19/jibe/api/internal/common/schema"
	"github.com/Xebec19/jibe/api/internal/layers/domain"
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

func (a authController) VerifyHandler(w http.ResponseWriter, r *http.Request) {

	var req domain.VerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		a.logger.Error("Invalid request body", err)
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	siweMsg, err := domain.ParseSIWEMessage(req.Message)
	if err != nil {
		a.logger.Error("Failed to parse message", err)
		respondError(w, http.StatusInternalServerError, "Failed to parse message")
		return
	}

	valid, err := domain.VerifySignature(req.Message, req.Signature, siweMsg.Address)
	if err != nil || !valid {
		respondError(w, http.StatusBadRequest, "Signature verification failed")
		return
	}

	// Verify domain
	expectedDomain := "localhost:3000"
	if siweMsg.Domain != expectedDomain {
		respondError(w, http.StatusBadRequest, "Invalid domain")
		return
	}

	// Verify expiration
	if siweMsg.ExpirationTime != "" {
		expiry, err := time.Parse(time.RFC3339, siweMsg.ExpirationTime)
		if err != nil {
			respondError(w, http.StatusBadRequest, "Expiration time parsing failed")
			return
		}
		if time.Now().After(expiry) {
			respondError(w, http.StatusBadRequest, "Message has expired")
			return
		}
	}

	// Verify not before
	if siweMsg.NotBefore != "" {
		notBefore, err := time.Parse(time.RFC3339, siweMsg.NotBefore)
		if err != nil {
			respondError(w, http.StatusBadRequest, "Invalid message")
			return
		}
		if time.Now().Before(notBefore) {
			respondError(w, http.StatusBadRequest, "Invalid message")
			return
		}
	}

	respondJSON(w, http.StatusOK, "Message is verified", domain.VerifyResponse{
		Valid:   true,
		Address: strings.ToLower(siweMsg.Address),
	})
}
