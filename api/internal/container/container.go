// container contains all components eg, logger, config, handlers and services
// used throughout the server
package container

import (
	"github.com/Xebec19/jibe/api/internal/container/config"
	"github.com/Xebec19/jibe/api/pkg/logger"
)

func NewContainer(cfg *config.Config, logger logger.Logger) Container {
	return Container{
		Cfg:    *cfg,
		Logger: logger,
	}
}

type Container struct {
	Cfg    config.Config
	Logger logger.Logger
}
