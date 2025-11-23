package controllers

import (
	"net/http"

	"github.com/Xebec19/jibe/api/pkg/logger"
)

type HealthController interface {
	GetHealthCheckpoint(w http.ResponseWriter, r *http.Request)
}

func NewHealthController(logger *logger.Logger) HealthController {
	return healthController{
		logger: logger,
	}
}

type healthController struct {
	logger *logger.Logger
}

func (h healthController) GetHealthCheckpoint(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, "alive", nil)
}
