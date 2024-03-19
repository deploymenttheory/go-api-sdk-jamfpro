// jamfproapi_return_to_service.go
// Jamf Pro Api - Return to Service
// api reference: none available
// docs: https://learn.jamf.com/en-US/bundle/technical-articles/page/Return_to_Service.html
// Jamf Pro Api requires the structs to support an JSON data structure.

package jamfpro

import (
	"fmt"
)

const uriReturnToService = "/api/v1/return-to-service"

// GetReturnToService fetches a list of devices that are in the Return to Service state.
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
