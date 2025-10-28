package container

import (
	"github.com/Xebec19/jibe/api/internal/config"
	"github.com/Xebec19/jibe/api/internal/repository"
	"github.com/Xebec19/jibe/api/internal/repository/memory"
	"github.com/Xebec19/jibe/api/internal/service"
	"github.com/Xebec19/jibe/api/pkg/logger"
	"github.com/Xebec19/jibe/api/pkg/resilience"
)

// Container holds all application dependencies
// This implements the Dependency Injection pattern
type Container struct {
	Config *config.Config
	Logger *logger.Logger

	// Repositories
	UserRepository repository.UserRepository

	// Services
	UserService service.UserService

	// Resilience
	CircuitBreaker *resilience.CircuitBreaker
	RetryPolicy    *resilience.RetryPolicy
}

// New creates a new dependency injection container
func New(cfg *config.Config, log *logger.Logger) *Container {
	c := &Container{
		Config: cfg,
		Logger: log,
	}

	// Initialize resilience components
	c.initResilience()

	// Initialize repositories
	c.initRepositories()

	// Initialize services
	c.initServices()

	return c
}

// initResilience initializes resilience patterns
func (c *Container) initResilience() {
	// Circuit breaker configuration
	c.CircuitBreaker = resilience.NewCircuitBreaker(
		5,   // Max failures before opening
		10,  // Timeout in seconds
		30,  // Reset timeout in seconds
		c.Logger,
	)

	// Retry policy configuration
	c.RetryPolicy = resilience.NewRetryPolicy(
		3,     // Max retries
		100,   // Initial backoff in ms
		2.0,   // Backoff multiplier
		5000,  // Max backoff in ms
		c.Logger,
	)
}

// initRepositories initializes all repositories
func (c *Container) initRepositories() {
	// For now, use in-memory repository
	// In production, this would be a database implementation
	c.UserRepository = memory.NewUserRepository()

	c.Logger.Info().Msg("Repositories initialized")
}

// initServices initializes all services
func (c *Container) initServices() {
	c.UserService = service.NewUserService(c.UserRepository, c.Logger)

	c.Logger.Info().Msg("Services initialized")
}

// Shutdown performs cleanup when the application stops
func (c *Container) Shutdown() {
	c.Logger.Info().Msg("Shutting down container")
	// Add any cleanup logic here (e.g., closing database connections)
}
