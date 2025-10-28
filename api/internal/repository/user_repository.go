package repository

import (
	"context"

	"github.com/Xebec19/jibe/api/internal/domain"
)

// UserRepository defines the interface for user data operations
type UserRepository interface {
	// Create creates a new user
	Create(ctx context.Context, user *domain.User) error

	// GetByID retrieves a user by ID
	GetByID(ctx context.Context, id int64) (*domain.User, error)

	// GetByEmail retrieves a user by email
	GetByEmail(ctx context.Context, email string) (*domain.User, error)

	// List retrieves all users with pagination
	List(ctx context.Context, limit, offset int) ([]*domain.User, error)

	// Update updates an existing user
	Update(ctx context.Context, user *domain.User) error

	// Delete deletes a user by ID
	Delete(ctx context.Context, id int64) error

	// Count returns the total number of users
	Count(ctx context.Context) (int64, error)
}
