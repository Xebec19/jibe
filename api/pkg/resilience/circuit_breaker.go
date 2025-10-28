package resilience

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/Xebec19/jibe/api/pkg/logger"
)

// Circuit breaker states
const (
	StateClosed   = "closed"
	StateOpen     = "open"
	StateHalfOpen = "half_open"
)

var (
	ErrCircuitOpen     = errors.New("circuit breaker is open")
	ErrTooManyRequests = errors.New("too many requests")
)

// CircuitBreaker implements the circuit breaker pattern
// It prevents cascading failures by stopping requests to failing services
type CircuitBreaker struct {
	maxFailures  int
	timeout      time.Duration
	resetTimeout time.Duration
	logger       *logger.Logger

	mu            sync.RWMutex
	state         string
	failures      int
	lastFailTime  time.Time
	lastStateTime time.Time
}

// NewCircuitBreaker creates a new circuit breaker
func NewCircuitBreaker(maxFailures int, timeoutSeconds int, resetTimeoutSeconds int, log *logger.Logger) *CircuitBreaker {
	return &CircuitBreaker{
		maxFailures:   maxFailures,
		timeout:       time.Duration(timeoutSeconds) * time.Second,
		resetTimeout:  time.Duration(resetTimeoutSeconds) * time.Second,
		logger:        log,
		state:         StateClosed,
		lastStateTime: time.Now(),
	}
}

// Execute runs a function with circuit breaker protection
func (cb *CircuitBreaker) Execute(ctx context.Context, fn func() error) error {
	if err := cb.beforeRequest(); err != nil {
		return err
	}

	// Execute the function with timeout
	done := make(chan error, 1)
	go func() {
		done <- fn()
	}()

	select {
	case err := <-done:
		cb.afterRequest(err)
		return err
	case <-time.After(cb.timeout):
		err := errors.New("operation timeout")
		cb.afterRequest(err)
		return err
	case <-ctx.Done():
		err := ctx.Err()
		cb.afterRequest(err)
		return err
	}
}

// beforeRequest checks if the request can proceed
func (cb *CircuitBreaker) beforeRequest() error {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	now := time.Now()

	switch cb.state {
	case StateOpen:
		// Check if it's time to try half-open
		if now.Sub(cb.lastFailTime) > cb.resetTimeout {
			cb.state = StateHalfOpen
			cb.lastStateTime = now
			cb.logger.Info().Msg("Circuit breaker moved to half-open state")
			return nil
		}
		return ErrCircuitOpen

	case StateHalfOpen:
		// Only allow one request in half-open state
		return nil

	case StateClosed:
		return nil

	default:
		return nil
	}
}

// afterRequest updates the circuit breaker state based on the result
func (cb *CircuitBreaker) afterRequest(err error) {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	now := time.Now()

	if err != nil {
		cb.failures++
		cb.lastFailTime = now

		if cb.state == StateHalfOpen {
			// Failed in half-open state, go back to open
			cb.state = StateOpen
			cb.lastStateTime = now
			cb.logger.Warn().
				Err(err).
				Msg("Circuit breaker reopened after failed half-open attempt")
			return
		}

		if cb.failures >= cb.maxFailures {
			cb.state = StateOpen
			cb.lastStateTime = now
			cb.logger.Warn().
				Int("failures", cb.failures).
				Msg("Circuit breaker opened due to failures")
		}
	} else {
		// Success
		if cb.state == StateHalfOpen {
			cb.state = StateClosed
			cb.failures = 0
			cb.lastStateTime = now
			cb.logger.Info().Msg("Circuit breaker closed after successful half-open attempt")
		} else if cb.state == StateClosed {
			// Reset failure count on success
			cb.failures = 0
		}
	}
}

// GetState returns the current state of the circuit breaker
func (cb *CircuitBreaker) GetState() string {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	return cb.state
}

// GetFailures returns the current failure count
func (cb *CircuitBreaker) GetFailures() int {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	return cb.failures
}

// Reset manually resets the circuit breaker
func (cb *CircuitBreaker) Reset() {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	cb.state = StateClosed
	cb.failures = 0
	cb.lastStateTime = time.Now()
	cb.logger.Info().Msg("Circuit breaker manually reset")
}
