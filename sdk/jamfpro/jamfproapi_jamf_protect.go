// jamfproapi_jamf_protect.go
// Jamf Pro Api - Jamf Protect
// api reference: https://developer.jamf.com/jamf-pro/reference/put_v1-jamf-protect
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
)

const uriJamfProtect = "/api/v1/jamf-protect"

// Structs

type ResourceJamfProtectIntegrationSettings struct {
	ID             string `json:"id"`
	APIClientID    string `json:"apiClientId"`
	APIClientName  string `json:"apiClientName"`
	RegistrationID string `json:"registrationId"`
	ProtectURL     string `json:"protectUrl"`
	LastSyncTime   string `json:"lastSyncTime"`
	SyncStatus     string `json:"syncStatus"`
	AutoInstall    bool   `json:"autoInstall"`
}

type ResourceJamfProtectIntegrationCreateSettings struct {
	ProtectURL string `json:"protectUrl"`
	ClientID   string `json:"clientId"`
	Password   string `json:"password"`
}

// CRUD

func (c *Client) GetJamfProtectIntegrationSettings() (*ResourceJamfProtectIntegrationSettings, error) {
	endpoint := uriJamfProtect
	var out ResourceJamfProtectIntegrationSettings
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "jamf protect integration settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

func (c *Client) UpdateJamfProtectIntegrationSettings(updatedSettings ResourceJamfProtectIntegrationSettings) (*ResourceJamfProtectIntegrationSettings, error) {
	// TODO - Figure out if we can update everything here or just the bool
	endpoint := uriJamfProtect
	var out ResourceJamfProtectIntegrationSettings
	resp, err := c.HTTP.DoRequest("PUT", endpoint, updatedSettings, &out)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdate, "jamf protect integration settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil

}

func (c *Client) DeleteJamfProtectIntegration() error {
	endpoint := uriJamfProtect

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDelete, "jamf protect integration", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

func (c *Client) CreateJamfProtectIntegration(createSettings ResourceJamfProtectIntegrationCreateSettings) (*ResourceJamfProtectIntegrationSettings, error) {
	endpoint := fmt.Sprintf("%s/register", uriJamfProtect)
	var out ResourceJamfProtectIntegrationSettings
	resp, err := c.HTTP.DoRequest("POST", endpoint, createSettings, &out)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdate, "jamf protect integration settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// QUERY are we bothered about the rest of the operations at this endpoint? - no i dont thinik so at this stage
