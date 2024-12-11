// jamfproapi_startup_status.go
// Jamf Pro Api - Startup Status
// api reference: https://developer.jamf.com/jamf-pro/reference/get_startup-status
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
)

const uriStartupStatus = "/api/startup-status"

// Resource structure
type ResourceStartupStatus struct {
	Step                    string `json:"step"`
	StepCode                string `json:"stepCode"`
	StepParam               string `json:"stepParam"`
	Percentage              int    `json:"percentage"`
	Warning                 string `json:"warning"`
	WarningCode             string `json:"warningCode"`
	WarningParam            string `json:"warningParam"`
	Error                   string `json:"error"`
	ErrorCode               string `json:"errorCode"`
	SetupAssistantNecessary bool   `json:"setupAssistantNecessary"`
}

// CRUD

// GetStartupStatus retrieves the status of the Jamf Pro server using API.
func (c *Client) GetStartupStatus() (*ResourceStartupStatus, error) {
	endpoint := uriStartupStatus

	var jamfProStatus ResourceStartupStatus
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &jamfProStatus)
	if err != nil {
		return nil, fmt.Errorf("failed to get Jamf Pro status informations: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &jamfProStatus, nil
}
