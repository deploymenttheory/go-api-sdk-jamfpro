package jamfpro

import (
	"fmt"
	"net/http"
	"time"
)

// WaitForResourceAvailable checks if a resource is available by making a GET request to the resource's URL.
// It retries the request with exponential backoff until a 200 OK response is received or the max wait time is reached.
// resourceURL: The URL to check for resource availability.
// maxWaitTime: Maximum time to wait for the resource to become available.
// client: The HTTP client used to make the request. Ensure it's properly authenticated to access the resource.
func (c *Client) WaitForResourceAvailable(resourceURL string, maxWaitTime time.Duration) error {
	var backoffTime = 1 * time.Second
	const backoffFactor = 2
	const maxBackoffTime = 30 * time.Second

	startTime := time.Now()

	for {
		// Make a request to check resource availability
		resp, err := c.HTTP.DoRequest("GET", resourceURL, nil, nil)
		if err != nil {
			return fmt.Errorf("error making request to check resource availability: %v", err)
		}

		if resp.StatusCode == http.StatusOK {
			// Resource is available
			return nil
		}

		if time.Since(startTime) > maxWaitTime {
			// Max wait time exceeded
			return fmt.Errorf("timed out waiting for resource to become available")
		}

		// Wait for backoff time before retrying
		time.Sleep(backoffTime)

		// Increase backoff time for the next iteration, capped at maxBackoffTime
		backoffTime *= time.Duration(backoffFactor)
		if backoffTime > maxBackoffTime {
			backoffTime = maxBackoffTime
		}
	}
}
