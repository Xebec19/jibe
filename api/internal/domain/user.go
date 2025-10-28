package domain

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateUserRequest represents the request to create a user
type CreateUserRequest struct {
	Name  string `json:"name" validate:"required,min=2,max=100"`
	Email string `json:"email" validate:"required,email"`
}

// UpdateUserRequest represents the request to update a user
type UpdateUserRequest struct {
	Name  string `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	Email string `json:"email,omitempty" validate:"omitempty,email"`
}

// Validate validates the CreateUserRequest
func (r *CreateUserRequest) Validate() error {
	if r.Name == "" {
		return ErrInvalidInput{Field: "name", Message: "name is required"}
	}
	if len(r.Name) < 2 || len(r.Name) > 100 {
		return ErrInvalidInput{Field: "name", Message: "name must be between 2 and 100 characters"}
	}
	if r.Email == "" {
		return ErrInvalidInput{Field: "email", Message: "email is required"}
	}
	// Basic email validation
	if !isValidEmail(r.Email) {
		return ErrInvalidInput{Field: "email", Message: "invalid email format"}
	}
	return nil
}

// isValidEmail performs basic email validation
func isValidEmail(email string) bool {
	// Simple validation - in production, use a proper email validation library
	if len(email) < 3 || len(email) > 254 {
		return false
	}
	atIndex := -1
	for i, c := range email {
		if c == '@' {
			if atIndex != -1 {
				return false // Multiple @ symbols
			}
			atIndex = i
		}
	}
	return atIndex > 0 && atIndex < len(email)-1
}
