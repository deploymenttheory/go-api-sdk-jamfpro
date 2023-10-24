package http_client_test

import (
	"context"
	"testing"
	"time"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"

	"github.com/google/uuid"
)

func TestAcquireAndRelease(t *testing.T) {
	manager := http_client.NewConcurrencyManager(2, nil, false)

	acquiredCount := 0
	releasedCount := 0
	finishedCount := 0

	for i := 0; i < 5; i++ {
		go func() {
			_, err := manager.Acquire(context.Background())
			if err == nil {
				acquiredCount++
			}

			// Simulate some work
			time.Sleep(100 * time.Millisecond)

			manager.Release(uuid.New()) // Note: You'll need to adjust how you handle requestID here.
			releasedCount++
			finishedCount++
		}()
	}

	// Give enough time for all goroutines to either acquire the token or get blocked
	time.Sleep(500 * time.Millisecond)

	if acquiredCount != 2 {
		t.Errorf("Expected 2 goroutines to acquire tokens, but got %d", acquiredCount)
	}

	// Release a token to check if a blocked goroutine can now acquire it
	manager.Release(uuid.New())

	// Give time for another goroutine to acquire the token
	time.Sleep(100 * time.Millisecond)

	if releasedCount <= 2 {
		t.Errorf("Expected more than 2 goroutines to have released tokens, but got %d", releasedCount)
	}

	// Ensure all goroutines have finished
	time.Sleep(1 * time.Second)
	if finishedCount != 5 {
		t.Errorf("Expected all 5 goroutines to finish, but got %d", finishedCount)
	}
}
