package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Xebec19/jibe/api/internal/domain"
	"github.com/Xebec19/jibe/api/internal/repository"
	"github.com/Xebec19/jibe/api/pkg/logger"
)

// userRepository is a PostgreSQL implementation of UserRepository
type userRepository struct {
	db     *sql.DB
	logger *logger.Logger
}

// NewUserRepository creates a new PostgreSQL user repository
func NewUserRepository(db *sql.DB, logger *logger.Logger) repository.UserRepository {
	return &userRepository{
		db:     db,
		logger: logger,
	}
}

// Create creates a new user
func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	query := `
		INSERT INTO users (name, email, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	err := r.db.QueryRowContext(ctx, query, user.Name, user.Email, user.CreatedAt, user.UpdatedAt).Scan(&user.ID)
	if err != nil {
		// Check for unique constraint violation (duplicate email)
		if isDuplicateKeyError(err) {
			return domain.ErrAlreadyExists
		}
		r.logger.Error().Err(err).Msg("Failed to create user")
		return err
	}

	return nil
}

// GetByID retrieves a user by ID
func (r *userRepository) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	query := `
		SELECT id, name, email, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	user := &domain.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		r.logger.Error().Err(err).Int64("user_id", id).Msg("Failed to get user")
		return nil, err
	}

	return user, nil
}

// GetByEmail retrieves a user by email
func (r *userRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `
		SELECT id, name, email, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	user := &domain.User{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		r.logger.Error().Err(err).Str("email", email).Msg("Failed to get user by email")
		return nil, err
	}

	return user, nil
}

// List retrieves all users with pagination
func (r *userRepository) List(ctx context.Context, limit, offset int) ([]*domain.User, error) {
	query := `
		SELECT id, name, email, created_at, updated_at
		FROM users
		ORDER BY id
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to list users")
		return nil, err
	}
	defer rows.Close()

	users := make([]*domain.User, 0)
	for rows.Next() {
		user := &domain.User{}
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			r.logger.Error().Err(err).Msg("Failed to scan user")
			continue
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		r.logger.Error().Err(err).Msg("Error iterating users")
		return nil, err
	}

	return users, nil
}

// Update updates an existing user
func (r *userRepository) Update(ctx context.Context, user *domain.User) error {
	query := `
		UPDATE users
		SET name = $1, email = $2, updated_at = $3
		WHERE id = $4
	`

	user.UpdatedAt = time.Now()

	result, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.UpdatedAt, user.ID)
	if err != nil {
		if isDuplicateKeyError(err) {
			return domain.ErrAlreadyExists
		}
		r.logger.Error().Err(err).Int64("user_id", user.ID).Msg("Failed to update user")
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to get rows affected")
		return err
	}

	if rowsAffected == 0 {
		return domain.ErrNotFound
	}

	return nil
}

// Delete deletes a user by ID
func (r *userRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		r.logger.Error().Err(err).Int64("user_id", id).Msg("Failed to delete user")
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to get rows affected")
		return err
	}

	if rowsAffected == 0 {
		return domain.ErrNotFound
	}

	return nil
}

// Count returns the total number of users
func (r *userRepository) Count(ctx context.Context) (int64, error) {
	query := `SELECT COUNT(*) FROM users`

	var count int64
	err := r.db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		r.logger.Error().Err(err).Msg("Failed to count users")
		return 0, err
	}

	return count, nil
}

// isDuplicateKeyError checks if the error is a unique constraint violation
func isDuplicateKeyError(err error) bool {
	// PostgreSQL error code for unique constraint violation is 23505
	// This is a simplified check - in production, use a proper error type check
	if err == nil {
		return false
	}
	errMsg := err.Error()
	return contains(errMsg, "duplicate") || contains(errMsg, "unique") || contains(errMsg, "23505")
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 || indexSubstring(s, substr) >= 0)
}

func indexSubstring(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
