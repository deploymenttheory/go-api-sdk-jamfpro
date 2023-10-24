// http_concurrency_management.go
// Package http_client provides utilities to manage HTTP client interactions, including concurrency control.
// The Concurrency Manager ensures no more than a certain number of concurrent requests (e.g., 5 for Jamf Pro) are sent at the same time. This is managed using a semaphore
package http_client

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

// ConcurrencyManager controls the number of concurrent HTTP requests.
type ConcurrencyManager struct {
	sem       chan struct{}
	logger    Logger
	debugMode bool
}

type requestIDKey struct{}

// NewConcurrencyManager initializes a new ConcurrencyManager with the given limit.
func NewConcurrencyManager(limit int, logger Logger, debugMode bool) *ConcurrencyManager {
	if logger == nil {
		logger = &defaultLogger{} // Assuming this is the default logger implementation
	}
	return &ConcurrencyManager{
		sem:       make(chan struct{}, limit),
		logger:    logger,
		debugMode: debugMode,
	}
}

// Acquire blocks until a token is available to proceed, or until the context is done.
func (c *ConcurrencyManager) Acquire(ctx context.Context) (uuid.UUID, error) {
	requestID := uuid.New()

	select {
	case c.sem <- struct{}{}:
		if c.debugMode {
			utilizedTokens := len(c.sem)
			availableTokens := cap(c.sem) - len(c.sem)
			c.logger.Debug(fmt.Sprintf("[ConcurrencyTokenID: %s] Acquired concurrency token. Details [Utilized tokens: %d. Available tokens: %d.]", requestID, utilizedTokens, availableTokens))
		}
		return requestID, nil
	case <-ctx.Done():
		c.logger.Warn(fmt.Sprintf("[ConcurrencyTokenID: %s] Failed to acquire concurrency token, context done", requestID))
		return requestID, ctx.Err()
	}
}

// Release releases a token, allowing another request to proceed.
func (c *ConcurrencyManager) Release(requestID uuid.UUID) {
	<-c.sem
	if c.debugMode {
		utilizedTokens := len(c.sem)
		availableTokens := cap(c.sem) - len(c.sem)
		c.logger.Debug(fmt.Sprintf("[ConcurrencyTokenID: %s] Released concurrency token. Details [Utilized tokens: %d. Available tokens: %d.]", requestID, utilizedTokens, availableTokens))
	}
}
