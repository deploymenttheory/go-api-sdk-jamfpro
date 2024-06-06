// jamfproapi_health_check.go
// Jamf Pro Api - Health Check
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-health-check
// Jamf Pro API requires the structs to support an JSON data structure.

package jamfpro

import "fmt"

const uriHealthCheck = "/api/v1/health-check"

// GetHealthCheck fetches the Jamf Pro API status.
func (c *Client) GetHealthCheck() (bool, error) {
	endpoint := uriHealthCheck
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, nil)
	if err != nil {
		return false, fmt.Errorf(errMsgFailedGet, "Health Check", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	return true, nil
}
