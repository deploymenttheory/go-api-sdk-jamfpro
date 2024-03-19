// jamfproapi_health_check.go
// Jamf Pro Api - Health Check
// api reference: N/A
// Jamf Pro API requires the structs to support an JSON data structure.

package jamfpro

import "fmt"

const uriHealthCheck = "/api/v1/health-check"

// GetHealthCheck fetches a file list from Jamf Cloud Distribution Service
func (c *Client) GetHealthCheck() ([]ResponseJCDS2List, error) {
	endpoint := uriHealthCheck
	var out []ResponseJCDS2List
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "Health Check", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return out, nil
}
