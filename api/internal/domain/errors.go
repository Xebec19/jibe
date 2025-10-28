package domain

import (
	"errors"
	"fmt"
)

// Common domain errors
var (
	ErrNotFound          = errors.New("resource not found")
	ErrAlreadyExists     = errors.New("resource already exists")
	ErrUnauthorized      = errors.New("unauthorized")
	ErrForbidden         = errors.New("forbidden")
	ErrInternalServer    = errors.New("internal server error")
	ErrServiceUnavailable = errors.New("service unavailable")
	ErrTooManyRequests   = errors.New("too many requests")
)

// ErrInvalidInput represents a validation error
type ErrInvalidInput struct {
	Field   string
	Message string
}

func (e ErrInvalidInput) Error() string {
	return fmt.Sprintf("invalid input for field '%s': %s", e.Field, e.Message)
}

// ErrValidation represents multiple validation errors
type ErrValidation struct {
	Errors []ErrInvalidInput
}

func (e ErrValidation) Error() string {
	if len(e.Errors) == 0 {
		return "validation error"
	}
	return fmt.Sprintf("validation failed: %s", e.Errors[0].Error())
}
