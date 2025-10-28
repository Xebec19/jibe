package resilience

import (
	"context"
	"errors"
	"time"

	"github.com/Xebec19/jibe/api/pkg/logger"
)

var (
	ErrMaxRetriesExceeded = errors.New("max retries exceeded")
)

// RetryPolicy implements exponential backoff retry pattern
type RetryPolicy struct {
	maxRetries        int
	initialBackoff    time.Duration
	backoffMultiplier float64
	maxBackoff        time.Duration
	logger            *logger.Logger
}

// NewRetryPolicy creates a new retry policy with exponential backoff
func NewRetryPolicy(maxRetries int, initialBackoffMs int, backoffMultiplier float64, maxBackoffMs int, log *logger.Logger) *RetryPolicy {
	return &RetryPolicy{
		maxRetries:        maxRetries,
		initialBackoff:    time.Duration(initialBackoffMs) * time.Millisecond,
		backoffMultiplier: backoffMultiplier,
		maxBackoff:        time.Duration(maxBackoffMs) * time.Millisecond,
		logger:            log,
	}
}

// Execute runs a function with retry logic and exponential backoff
func (r *RetryPolicy) Execute(ctx context.Context, fn func() error) error {
	var lastErr error
	backoff := r.initialBackoff

	for attempt := 0; attempt <= r.maxRetries; attempt++ {
		// Execute the function
		err := fn()
		if err == nil {
			// Success
			if attempt > 0 {
				r.logger.Info().
					Int("attempts", attempt+1).
					Msg("Operation succeeded after retries")
			}
			return nil
		}

		lastErr = err

		// Check if we should retry
		if !r.shouldRetry(err) {
			r.logger.Debug().
				Err(err).
				Msg("Error is not retryable")
			return err
		}

		// Check if we've exhausted retries
		if attempt >= r.maxRetries {
			r.logger.Warn().
				Err(err).
				Int("attempts", attempt+1).
				Msg("Max retries exceeded")
			break
		}

		// Log retry attempt
		r.logger.Debug().
			Err(err).
			Int("attempt", attempt+1).
			Dur("backoff", backoff).
			Msg("Retrying operation")

		// Wait with exponential backoff
		select {
		case <-time.After(backoff):
			// Calculate next backoff
			backoff = time.Duration(float64(backoff) * r.backoffMultiplier)
			if backoff > r.maxBackoff {
				backoff = r.maxBackoff
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return lastErr
}

// ExecuteWithJitter runs a function with retry logic and exponential backoff with jitter
// Jitter helps prevent thundering herd problem
func (r *RetryPolicy) ExecuteWithJitter(ctx context.Context, fn func() error, jitterFactor float64) error {
	var lastErr error
	backoff := r.initialBackoff

	for attempt := 0; attempt <= r.maxRetries; attempt++ {
		err := fn()
		if err == nil {
			if attempt > 0 {
				r.logger.Info().
					Int("attempts", attempt+1).
					Msg("Operation succeeded after retries")
			}
			return nil
		}

		lastErr = err

		if !r.shouldRetry(err) {
			r.logger.Debug().
				Err(err).
				Msg("Error is not retryable")
			return err
		}

		if attempt >= r.maxRetries {
			r.logger.Warn().
				Err(err).
				Int("attempts", attempt+1).
				Msg("Max retries exceeded")
			break
		}

		// Apply jitter (random factor to prevent thundering herd)
		jitteredBackoff := time.Duration(float64(backoff) * (1.0 - jitterFactor/2.0 + jitterFactor*float64(time.Now().UnixNano()%100)/100.0))

		r.logger.Debug().
			Err(err).
			Int("attempt", attempt+1).
			Dur("backoff", jitteredBackoff).
			Msg("Retrying operation with jitter")

		select {
		case <-time.After(jitteredBackoff):
			backoff = time.Duration(float64(backoff) * r.backoffMultiplier)
			if backoff > r.maxBackoff {
				backoff = r.maxBackoff
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return lastErr
}

// shouldRetry determines if an error is retryable
func (r *RetryPolicy) shouldRetry(err error) bool {
	// Add logic to determine if error is retryable
	// For now, retry all errors except context cancellation
	if errors.Is(err, context.Canceled) {
		return false
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return false
	}
	// Don't retry circuit breaker open errors
	if errors.Is(err, ErrCircuitOpen) {
		return false
	}
	return true
}

// SetMaxRetries updates the maximum number of retries
func (r *RetryPolicy) SetMaxRetries(maxRetries int) {
	r.maxRetries = maxRetries
}

// GetMaxRetries returns the maximum number of retries
func (r *RetryPolicy) GetMaxRetries() int {
	return r.maxRetries
}
