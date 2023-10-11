// http_rate_handler.go

/*
Components:
Backoff Strategy: A function that calculates the delay before the next retry. It will implement exponential backoff with jitter. This strategy is more effective than a fixed delay, as it ensures that in cases of prolonged issues, the client won't keep hammering the server with a high frequency.

Response Time Monitoring: We'll introduce a mechanism to track average response times and use deviations from this average to inform our backoff strategy.

Error Classifier: A function to classify different types of errors. Only transient errors should be retried.

Rate Limit Header Parser: For future compatibility, a function that can parse common rate limit headers (like X-RateLimit-Remaining and Retry-After) and adjust behavior accordingly.

*/

package http_client

import (
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// Constants for exponential backoff with jitter
const (
	baseDelay    = 100 * time.Millisecond // Initial delay
	maxDelay     = 5 * time.Second        // Maximum delay
	jitterFactor = 0.5                    // Random jitter factor
)

// calculateBackoff calculates the next delay for retry with exponential backoff and jitter.
func calculateBackoff(retry int) time.Duration {
	delay := float64(baseDelay) * math.Pow(2, float64(retry))
	jitter := (rand.Float64() - 0.5) * jitterFactor * 2.0 // Random value between -jitterFactor and +jitterFactor
	delay *= (1.0 + jitter)

	if delay > float64(maxDelay) {
		return maxDelay
	}
	return time.Duration(delay)
}

// parseRateLimitHeaders parses common rate limit headers and adjusts behavior accordingly.
// For future compatibility.
func parseRateLimitHeaders(resp *http.Response) time.Duration {
	// Check for the Retry-After header
	if retryAfter := resp.Header.Get("Retry-After"); retryAfter != "" {
		if waitSeconds, err := strconv.Atoi(retryAfter); err == nil {
			return time.Duration(waitSeconds) * time.Second
		}
	}

	// Check for X-RateLimit-Remaining; if it's 0, use X-RateLimit-Reset to determine how long to wait
	if remaining := resp.Header.Get("X-RateLimit-Remaining"); remaining == "0" {
		if resetTimeStr := resp.Header.Get("X-RateLimit-Reset"); resetTimeStr != "" {
			if resetTimeUnix, err := strconv.ParseInt(resetTimeStr, 10, 64); err == nil {
				resetTime := time.Unix(resetTimeUnix, 0)
				return time.Until(resetTime) // Using time.Until instead of t.Sub(time.Now())
			}
		}
	}

	// No rate limiting headers found, return 0
	return 0
}
