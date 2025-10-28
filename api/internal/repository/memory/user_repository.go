package memory

import (
	"context"
	"sync"
	"time"

	"github.com/Xebec19/jibe/api/internal/domain"
	"github.com/Xebec19/jibe/api/internal/repository"
)

// userRepository is an in-memory implementation of UserRepository
// This is useful for development and testing
type userRepository struct {
	mu      sync.RWMutex
	users   map[int64]*domain.User
	nextID  int64
	byEmail map[string]*domain.User
}

// NewUserRepository creates a new in-memory user repository
func NewUserRepository() repository.UserRepository {
	return &userRepository{
		users:   make(map[int64]*domain.User),
		byEmail: make(map[string]*domain.User),
		nextID:  1,
	}
}

// Create creates a new user
func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check if email already exists
	if _, exists := r.byEmail[user.Email]; exists {
		return domain.ErrAlreadyExists
	}

	user.ID = r.nextID
	r.nextID++
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	r.users[user.ID] = user
	r.byEmail[user.Email] = user

	return nil
}

// GetByID retrieves a user by ID
func (r *userRepository) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, domain.ErrNotFound
	}

	// Return a copy to prevent external modifications
	userCopy := *user
	return &userCopy, nil
}

// GetByEmail retrieves a user by email
func (r *userRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.byEmail[email]
	if !exists {
		return nil, domain.ErrNotFound
	}

	// Return a copy to prevent external modifications
	userCopy := *user
	return &userCopy, nil
}

// List retrieves all users with pagination
func (r *userRepository) List(ctx context.Context, limit, offset int) ([]*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users := make([]*domain.User, 0, len(r.users))
	for _, user := range r.users {
		userCopy := *user
		users = append(users, &userCopy)
	}

	// Simple pagination
	start := offset
	if start > len(users) {
		return []*domain.User{}, nil
	}

	end := start + limit
	if end > len(users) {
		end = len(users)
	}

	return users[start:end], nil
}

// Update updates an existing user
func (r *userRepository) Update(ctx context.Context, user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	existing, exists := r.users[user.ID]
	if !exists {
		return domain.ErrNotFound
	}

	// If email changed, check for conflicts and update index
	if existing.Email != user.Email {
		if _, exists := r.byEmail[user.Email]; exists {
			return domain.ErrAlreadyExists
		}
		delete(r.byEmail, existing.Email)
		r.byEmail[user.Email] = user
	}

	user.UpdatedAt = time.Now()
	r.users[user.ID] = user

	return nil
}

// Delete deletes a user by ID
func (r *userRepository) Delete(ctx context.Context, id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, exists := r.users[id]
	if !exists {
		return domain.ErrNotFound
	}

	delete(r.users, id)
	delete(r.byEmail, user.Email)

	return nil
}

// Count returns the total number of users
func (r *userRepository) Count(ctx context.Context) (int64, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return int64(len(r.users)), nil
}
