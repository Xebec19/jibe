// container contains all components eg, logger, config, handlers and services
// used throughout the server
package container

import (
	"context"

	"github.com/Xebec19/jibe/api/internal/db"
	"github.com/Xebec19/jibe/api/internal/layers/repositories"
	"github.com/Xebec19/jibe/api/internal/layers/services"
	"github.com/Xebec19/jibe/api/pkg/config"
	"github.com/Xebec19/jibe/api/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewContainer(ctx context.Context, cfg *config.Config, logger logger.Logger, dbpool *pgxpool.Pool, q *db.Queries) Container {
	return Container{
		Ctx:     ctx,
		Cfg:     *cfg,
		Logger:  logger,
		Dbpool:  dbpool,
		Queries: q,
	}
}

type Container struct {
	Ctx     context.Context
	Cfg     config.Config
	Logger  logger.Logger
	Dbpool  *pgxpool.Pool
	Queries *db.Queries

	// Repositories
	AuthRepository repositories.AuthRepository

	// Services
	AuthService services.AuthService
}

// initialize all repositories and save them in container
func (c *Container) SetupRepositories() {

	authRepo := repositories.NewAuthRepository(c.Ctx, &c.Logger, c.Queries)
	c.AuthRepository = authRepo
}

// initialize all services and save them in services
func (c *Container) SetupServices() {

	authSvc := services.NewAuthService(&c.Logger, c.AuthRepository)
	c.AuthService = authSvc
}
