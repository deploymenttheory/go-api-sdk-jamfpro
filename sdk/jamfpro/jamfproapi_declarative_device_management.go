// jamfproapi_declarative_device_management.go
// Jamf Pro Api - Declarative Device Management
// api reference: https://developer.jamf.com/jamf-pro/reference/post_v1-ddm-clientmanagementid-sync
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
)

const uriDeclarativeDeviceManagement = "/api/v1/ddm"

// ResponseError represents the structure of an error response from the API
type ResponseError struct {
	HTTPStatus int             `json:"httpStatus"`
	Errors     []ErrorInstance `json:"errors"`
}

// ErrorInstance represents a single error in the error response
type ErrorInstance struct {
	Code        string `json:"code"`
	Field       string `json:"field"`
	Description string `json:"description"`
	ID          string `json:"id"`
}

// ForceDDMSync initiates a DDM synchronization for a specific client management ID
func (c *Client) ForceDDMSync(clientManagementId string) error {
	endpoint := fmt.Sprintf("%s/%s/sync", uriDeclarativeDeviceManagement, clientManagementId)

	var errorResponse ResponseError
	resp, err := c.HTTP.DoRequest("POST", endpoint, nil, &errorResponse)
	if err != nil {
		return fmt.Errorf("failed to force DDM sync for client management ID %s: %v", clientManagementId, err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	// Check if the response contains errors
	if errorResponse.HTTPStatus != 0 {
		return fmt.Errorf("DDM sync failed for client management ID %s with status %d: %v", clientManagementId, errorResponse.HTTPStatus, errorResponse.Errors)
	}

	return nil
}
