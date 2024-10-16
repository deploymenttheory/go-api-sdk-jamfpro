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

// ResponseStatusItems represents the response structure for status items.
type ResponseStatusItems struct {
	StatusItems []StatusItem `json:"statusItems"`
}

// StatusItem represents a single status item in the status report.
type StatusItem struct {
	Key            string `json:"key"`
	Value          string `json:"value"`
	LastUpdateTime string `json:"lastUpdateTime"`
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

// GetDDMStatusItems retrieves the latest status report items for a specific device by its client management ID.
func (c *Client) GetDDMStatusItems(clientManagementId string) (*ResponseStatusItems, error) {
	endpoint := fmt.Sprintf("%s/%s/status-items", uriDeclarativeDeviceManagement, clientManagementId)

	var response ResponseStatusItems
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get DDM status items for client management ID %s: %v", clientManagementId, err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// GetDDMStatusItem retrieves the latest status report item for a specific device by its client management ID and status item key.
func (c *Client) GetDDMStatusItem(clientManagementId string, key string) (*StatusItem, error) {
	endpoint := fmt.Sprintf("%s/%s/status-items/%s", uriDeclarativeDeviceManagement, clientManagementId, key)

	var response StatusItem
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get status item for client management ID %s and key %s: %v", clientManagementId, key, err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}
