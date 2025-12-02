package services

import (
	"fmt"
	"time"

	"github.com/Xebec19/jibe/api/internal/layers/domain"
	"github.com/Xebec19/jibe/api/internal/layers/repositories"
	"github.com/Xebec19/jibe/api/pkg/config"
	"github.com/Xebec19/jibe/api/pkg/logger"
)

type AuthService interface {
	CreateNonce(addr string) (string, error)

	VerifySignature(message, signature string) (bool, error)
}

func NewAuthService(logger logger.Logger, cfg *config.Config, authRepo repositories.AuthRepository) AuthService {

	return &authService{
		logger:   logger,
		cfg:      cfg,
		authRepo: authRepo,
	}
}

type authService struct {
	logger   logger.Logger
	cfg      *config.Config
	authRepo repositories.AuthRepository
}

func (svc *authService) CreateNonce(addr string) (string, error) {

	return svc.authRepo.CreateNonce(addr)

}

func (svc *authService) VerifySignature(message, signature string) (bool, error) {

	siweMsg, err := domain.ParseSIWEMessage(message)
	if err != nil {
		return false, fmt.Errorf("message parsing failed %w", err)
	}

	valid, err := domain.VerifySignature(message, signature, siweMsg.Address)
	if err != nil || !valid {
		return false, fmt.Errorf("signature verification failed %w", err)
	}

	// Verify domain
	expectedDomain := svc.cfg.Domain
	if siweMsg.Domain != expectedDomain {
		return false, fmt.Errorf("invalid domain")
	}

	// Verify nonce
	isValid, err := svc.authRepo.CheckNonce(siweMsg.Nonce, siweMsg.Address)

	if err != nil {
		return false, fmt.Errorf("nonce verification failed %w", err)
	}

	if !isValid {
		return false, fmt.Errorf("invalid nonce")
	}

	// Verify expiration
	if siweMsg.ExpirationTime != "" {
		expiry, err := time.Parse(time.RFC3339, siweMsg.ExpirationTime)
		if err != nil {
			return false, fmt.Errorf("message expiration time parsing failed %d", err)
		}
		if time.Now().After(expiry) {
			return false, fmt.Errorf("message has expired %w", err)
		}
	}

	// Verify not before
	if siweMsg.NotBefore != "" {
		notBefore, err := time.Parse(time.RFC3339, siweMsg.NotBefore)
		if err != nil {
			return false, fmt.Errorf("invalid message %w", err)
		}
		if time.Now().Before(notBefore) {
			return false, fmt.Errorf("invalid message %w", err)
		}
	}

	return true, nil
}
