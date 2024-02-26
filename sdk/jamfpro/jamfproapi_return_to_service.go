package jamfpro

import (
	"fmt"
)

const uriReturnToService = "/api/v1/return-to-service"

// GetReturnToService fetches a file list from Jamf Cloud Distribution Service
func (c *Client) GetReturnToService() ([]ResponseJCDS2List, error) {
	endpoint := uriReturnToService
	var out []ResponseJCDS2List
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "Return To Service", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return out, nil
}
