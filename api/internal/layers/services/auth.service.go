package services

import (
	"fmt"
	"time"

	"github.com/Xebec19/jibe/api/internal/layers/domain"
	"github.com/Xebec19/jibe/api/internal/layers/repositories"
	"github.com/Xebec19/jibe/api/pkg/config"
	"github.com/Xebec19/jibe/api/pkg/jwt"
	"github.com/Xebec19/jibe/api/pkg/logger"
)

const (
	ACCESS_TOKEN_DURATION time.Duration = 30 * time.Minute
)

type AuthService interface {
	// CreateNonce create a random nonce, saves it in db and return it
	CreateNonce(addr string) (string, error)

	// VerifySignature verifies a siwe message with the given signature and returns if it
	// is valid and public address of the account
	VerifySignature(message, signature string) (bool, string, error)

	// SignJWTToken signs a jwt token
	SignJWTToken(addr string) (string, error)
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

func (svc *authService) VerifySignature(message, signature string) (bool, string, error) {

	siweMsg, err := domain.ParseSIWEMessage(message)
	if err != nil {
		return false, "", fmt.Errorf("message parsing failed %w", err)
	}

	valid, err := domain.VerifySignature(message, signature, siweMsg.Address)
	if err != nil || !valid {
		return false, "", fmt.Errorf("signature verification failed %w", err)
	}

	// Verify domain
	expectedDomain := svc.cfg.Domain
	if siweMsg.Domain != expectedDomain {
		return false, "", fmt.Errorf("invalid domain")
	}

	// Verify nonce
	isValid, err := svc.authRepo.CheckNonce(siweMsg.Nonce, siweMsg.Address)

	if err != nil {
		return false, "", fmt.Errorf("nonce verification failed %w", err)
	}

	if !isValid {
		return false, "", fmt.Errorf("invalid nonce")
	}

	// Verify expiration
	if siweMsg.ExpirationTime != "" {
		expiry, err := time.Parse(time.RFC3339, siweMsg.ExpirationTime)
		if err != nil {
			return false, "", fmt.Errorf("message expiration time parsing failed %d", err)
		}
		if time.Now().After(expiry) {
			return false, "", fmt.Errorf("message has expired %w", err)
		}
	}

	// Verify not before
	if siweMsg.NotBefore != "" {
		notBefore, err := time.Parse(time.RFC3339, siweMsg.NotBefore)
		if err != nil {
			return false, "", fmt.Errorf("invalid message %w", err)
		}
		if time.Now().Before(notBefore) {
			return false, "", fmt.Errorf("invalid message %w", err)
		}
	}

	return true, siweMsg.Address, nil
}

func (svc *authService) SignJWTToken(addr string) (string, error) {

	jti, err := svc.authRepo.CreateAccessToken(addr, time.Now().Add(ACCESS_TOKEN_DURATION))
	if err != nil {
		return "", fmt.Errorf("access token creation failed %w", err)
	}

	claims := jwt.TokenJWTClaims{
		Iss: svc.cfg.Domain,
		Sub: addr,
		Aud: svc.cfg.Domain,
		Exp: time.Now().Add(ACCESS_TOKEN_DURATION),
		Iat: time.Now(),
		Nbf: time.Now(),
		Jti: jti,
	}

	token, err := jwt.Token(claims, []byte(svc.cfg.JwtSecret))

	return token, err
}
