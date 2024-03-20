package jamfpro

import (
	"fmt"
	"net/http"
	"time"
)

// PingJamfResource checks if a resource is available by making a GET request to the resource's URL, using the provided resourceId.
// It retries the request with exponential backoff until a 200 OK response is received or the max wait time is reached.
// resourceURL: The base URL to check for resource availability.
// resourceId: The unique identifier of the resource to be checked.
// maxWaitTime: Maximum time to wait for the resource to become available.
func (c *Client) PingJamfResource(resourceURL string, resourceId string, maxWaitTime time.Duration) error {
	var backoffTime = 1 * time.Second
	const backoffFactor = 2
	const maxBackoffTime = 30 * time.Second

	// Construct the full URL with the resource ID
	fullURL := fmt.Sprintf("%s/id/%s", resourceURL, resourceId)

	startTime := time.Now()

	for {
		// Make a request to check resource availability
		resp, err := c.HTTP.DoRequest("GET", fullURL, nil, nil)
		if err != nil {
			return fmt.Errorf("error making request to check resource availability at %s: %v", fullURL, err)
		}

		if resp.StatusCode == http.StatusOK {
			// Resource is available
			return nil
		}

		if time.Since(startTime) > maxWaitTime {
			// Max wait time exceeded
			return fmt.Errorf("timed out waiting for resource to become available at %s", fullURL)
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
