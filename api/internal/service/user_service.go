package service

import (
	"context"

	"github.com/Xebec19/jibe/api/internal/domain"
	"github.com/Xebec19/jibe/api/internal/repository"
	"github.com/Xebec19/jibe/api/pkg/logger"
)

// UserService defines the interface for user business logic
type UserService interface {
	// CreateUser creates a new user
	CreateUser(ctx context.Context, req *domain.CreateUserRequest) (*domain.User, error)

	// GetUser retrieves a user by ID
	GetUser(ctx context.Context, id int64) (*domain.User, error)

	// GetUserByEmail retrieves a user by email
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)

	// ListUsers retrieves all users with pagination
	ListUsers(ctx context.Context, limit, offset int) ([]*domain.User, error)

	// UpdateUser updates an existing user
	UpdateUser(ctx context.Context, id int64, req *domain.UpdateUserRequest) (*domain.User, error)

	// DeleteUser deletes a user by ID
	DeleteUser(ctx context.Context, id int64) error
}

// userService implements UserService
type userService struct {
	repo   repository.UserRepository
	logger *logger.Logger
}

// NewUserService creates a new user service
func NewUserService(repo repository.UserRepository, log *logger.Logger) UserService {
	return &userService{
		repo:   repo,
		logger: log,
	}
}

// CreateUser creates a new user
func (s *userService) CreateUser(ctx context.Context, req *domain.CreateUserRequest) (*domain.User, error) {
	// Validate request
	if err := req.Validate(); err != nil {
		s.logger.Warn().
			Err(err).
			Str("email", req.Email).
			Msg("Invalid user creation request")
		return nil, err
	}

	// Check if user with email already exists
	existingUser, err := s.repo.GetByEmail(ctx, req.Email)
	if err == nil && existingUser != nil {
		s.logger.Warn().
			Str("email", req.Email).
			Msg("User with email already exists")
		return nil, domain.ErrAlreadyExists
	}

	// Create user
	user := &domain.User{
		Name:  req.Name,
		Email: req.Email,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		s.logger.Error().
			Err(err).
			Str("email", req.Email).
			Msg("Failed to create user")
		return nil, err
	}

	s.logger.Info().
		Int64("user_id", user.ID).
		Str("email", user.Email).
		Msg("User created successfully")

	return user, nil
}

// GetUser retrieves a user by ID
func (s *userService) GetUser(ctx context.Context, id int64) (*domain.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.logger.Debug().
			Err(err).
			Int64("user_id", id).
			Msg("User not found")
		return nil, err
	}

	return user, nil
}

// GetUserByEmail retrieves a user by email
func (s *userService) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		s.logger.Debug().
			Err(err).
			Str("email", email).
			Msg("User not found")
		return nil, err
	}

	return user, nil
}

// ListUsers retrieves all users with pagination
func (s *userService) ListUsers(ctx context.Context, limit, offset int) ([]*domain.User, error) {
	// Set default and maximum limits
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	users, err := s.repo.List(ctx, limit, offset)
	if err != nil {
		s.logger.Error().
			Err(err).
			Int("limit", limit).
			Int("offset", offset).
			Msg("Failed to list users")
		return nil, err
	}

	s.logger.Debug().
		Int("count", len(users)).
		Int("limit", limit).
		Int("offset", offset).
		Msg("Users listed successfully")

	return users, nil
}

// UpdateUser updates an existing user
func (s *userService) UpdateUser(ctx context.Context, id int64, req *domain.UpdateUserRequest) (*domain.User, error) {
	// Get existing user
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.logger.Debug().
			Err(err).
			Int64("user_id", id).
			Msg("User not found for update")
		return nil, err
	}

	// Update fields if provided
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}

	// Save updates
	if err := s.repo.Update(ctx, user); err != nil {
		s.logger.Error().
			Err(err).
			Int64("user_id", id).
			Msg("Failed to update user")
		return nil, err
	}

	s.logger.Info().
		Int64("user_id", user.ID).
		Msg("User updated successfully")

	return user, nil
}

// DeleteUser deletes a user by ID
func (s *userService) DeleteUser(ctx context.Context, id int64) error {
	// Check if user exists
	if _, err := s.repo.GetByID(ctx, id); err != nil {
		s.logger.Debug().
			Err(err).
			Int64("user_id", id).
			Msg("User not found for deletion")
		return err
	}

	// Delete user
	if err := s.repo.Delete(ctx, id); err != nil {
		s.logger.Error().
			Err(err).
			Int64("user_id", id).
			Msg("Failed to delete user")
		return err
	}

	s.logger.Info().
		Int64("user_id", id).
		Msg("User deleted successfully")

	return nil
}
