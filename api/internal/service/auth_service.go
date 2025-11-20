package service

import (
	"github.com/Xebec19/jibe/api/internal/domain"
	"github.com/Xebec19/jibe/api/pkg/logger"
)

type AuthService interface {
	GenerateNonce(req *domain.Auth) (string, error)
}

func NewAuthService(logger *logger.Logger) AuthService {
	return authService{
		logger: logger,
	}
}

type authService struct {
	logger *logger.Logger
}

func (svc authService) GenerateNonce(req *domain.Auth) (string, error) {
	nonce, err := req.GenerateNonce(16)

	return nonce, err
}
