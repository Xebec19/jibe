// container contains all components eg, logger, config, handlers and services
// used throughout the server
package container

import (
	"github.com/Xebec19/jibe/api/internal/container/config"
	"github.com/Xebec19/jibe/api/pkg/logger"
)

type Container interface {
}

func NewContainer(cfg *config.Config, logger logger.Logger) Container {
	return container{
		cfg:    *cfg,
		logger: logger,
	}
}

type container struct {
	cfg    config.Config
	logger logger.Logger
}
