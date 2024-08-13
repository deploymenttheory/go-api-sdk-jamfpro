// jamfproapi_slasa.go
// Jamf Pro Api - SLASA
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-slasa
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import "fmt"

// Struct for handling the response

// ResponseSLASAStatus represents the response structure for the SLASA acceptance status.
type ResponseSLASAStatus struct {
	SLASAAcceptanceStatus string `json:"slasaAcceptanceStatus"`
}

// GetSLASAStatus retrieves the status of SLASA (whether it has been accepted or not).
func (c *Client) GetSLASAStatus() (*ResponseSLASAStatus, error) {
	endpoint := fmt.Sprintf("%s/slasa", uriManagedSoftwareUpdates)

	var slasaStatus ResponseSLASAStatus
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &slasaStatus)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve SLASA status: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &slasaStatus, nil
}

// AcceptSLASA accepts the SLASA (Software License Agreement Service Acceptance).
func (c *Client) AcceptSLASA() error {
	endpoint := fmt.Sprintf("%s/slasa", uriManagedSoftwareUpdates)

	resp, err := c.HTTP.DoRequest("POST", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to accept SLASA: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
