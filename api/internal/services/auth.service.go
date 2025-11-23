package services

import (
	"github.com/Xebec19/jibe/api/internal/repositories"
	"github.com/Xebec19/jibe/api/pkg/logger"
)

type AuthService interface {
	CreateNonce(addr string) (string, error)
}

func NewAuthService(logger *logger.Logger, authRepo repositories.AuthRepository) AuthService {

	return &authService{
		logger:   logger,
		authRepo: authRepo,
	}
}

type authService struct {
	authRepo repositories.AuthRepository
	logger   *logger.Logger
}

func (svc *authService) CreateNonce(addr string) (string, error) {

	return svc.authRepo.CreateNonce(addr)

}
