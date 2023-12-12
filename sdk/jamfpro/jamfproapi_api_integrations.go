// jamfproapi_api_integrations.go
// Jamf Pro Api - API Integrations
// api reference: https://developer.jamf.com/jamf-pro/reference/getallapiintegrations
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

const uriApiIntegrations = "/api/v1/api-integrations"

// ResponseApiIntegrations represents the structure of the response for fetching API integrations
type ResponseApiIntegrationsList struct {
	Size    int                      `json:"totalCount"`
	Results []ResourceApiIntegration `json:"results"`
}

// Integration represents the details of an individual API integration
type ResourceApiIntegration struct {
	ID                         int      `json:"id,omitempty"`
	AuthorizationScopes        []string `json:"authorizationScopes,omitempty"`
	DisplayName                string   `json:"displayName,omitempty"`
	Enabled                    bool     `json:"enabled,omitempty"`
	AccessTokenLifetimeSeconds int      `json:"accessTokenLifetimeSeconds,omitempty"`
	AppType                    string   `json:"appType,omitempty"`
	ClientID                   string   `json:"clientId,omitempty"`
}

// Integration represents the details of Api client credentials
type ResourceClientCredentials struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

// GetApiIntegrations fetches all API integrations
func (c *Client) GetApiIntegrations() (*ResponseApiIntegrationsList, error) {
	endpoint := uriApiIntegrations
	resp, err := c.DoPaginatedGet(endpoint, standardPageSize, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Jamf API integrations: %v", err)
	}

	var OutStruct ResponseApiIntegrationsList
	OutStruct.Size = resp.Size
	for _, value := range resp.Results {
		var newObj ResourceApiIntegration
		mapstructure.Decode(value, &newObj)
		OutStruct.Results = append(OutStruct.Results, newObj)
	}

	return &OutStruct, nil
}

// GetApiIntegrationByID fetches an API integration by its ID
func (c *Client) GetApiIntegrationByID(id int) (*ResourceApiIntegration, error) {
	endpoint := fmt.Sprintf("%s/%d", uriApiIntegrations, id)

	var integration ResourceApiIntegration
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &integration)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Jamf API integration ID %d: %v", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &integration, nil
}

// GetApiIntegrationNameByID fetches an API integration by its display name and then retrieves its details using its ID
func (c *Client) GetApiIntegrationByName(name string) (*ResourceApiIntegration, error) {
	integrations, err := c.GetApiIntegrations()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all Jamf API integrations: %v", err)
	}

	for _, integration := range integrations.Results {
		if integration.DisplayName == name {
			return &integration, nil
		}
	}

	return nil, fmt.Errorf("no Jamf API integration found with the name %s", name)
}

// CreateApiIntegration creates a new API integration
func (c *Client) CreateApiIntegration(integration *ResourceApiIntegration) (*ResourceApiIntegration, error) {
	endpoint := uriApiIntegrations

	var response ResourceApiIntegration
	resp, err := c.HTTP.DoRequest("POST", endpoint, integration, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to create Jamf API integration: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateApiIntegrationByID updates an API integration by its ID
func (c *Client) UpdateApiIntegrationByID(id int, integrationUpdate *ResourceApiIntegration) (*ResourceApiIntegration, error) {
	// Construct the URL with the provided ID
	endpoint := fmt.Sprintf(uriApiIntegrations+"/%d", id)

	var updatedIntegration ResourceApiIntegration
	resp, err := c.HTTP.DoRequest("PUT", endpoint, integrationUpdate, &updatedIntegration)
	if err != nil {
		return nil, fmt.Errorf("failed to update Jamf API integration with ID %d: %v", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedIntegration, nil
}

// UpdateApiIntegrationByName updates an API integration based on its display name
func (c *Client) UpdateApiIntegrationByName(name string, updatedIntegration *ResourceApiIntegration) (*ResourceApiIntegration, error) {
	target, err := c.GetApiIntegrationByName(name)
	if err != nil {
		return nil, fmt.Errorf("failed to find integration, %v", err)
	}

	target_id := target.ID
	resp, err := c.UpdateApiIntegrationByID(target_id, updatedIntegration)
	if err != nil {
		return nil, fmt.Errorf("failed to update api integration, %v", err)
	}

	return resp, nil
}

// DeleteApiIntegrationByID deletes an API integration by its ID
func (c *Client) DeleteApiIntegrationByID(id int) error {
	// Construct the URL with the provided ID
	endpoint := fmt.Sprintf(uriApiIntegrations+"/%d", id)

	// Perform the DELETE request
	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete Jamf API integration with ID %d: %v", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteApiIntegrationByName deletes an API integration by its display name
func (c *Client) DeleteApiIntegrationByName(name string) error {
	target, err := c.GetApiIntegrationByName(name)
	if err != nil {
		return fmt.Errorf("failed to find api integration, %v", err)
	}

	target_id := target.ID

	err = c.DeleteApiIntegrationByID(target_id)
	if err != nil {
		return fmt.Errorf("failed to delete api integration, %v", err)
	}

	return nil
}

// Client Credentials

// RefreshClientCredentialsByApiRoleID creates new client credentials for an API integration by its ID
func (c *Client) RefreshClientCredentialsByApiRoleID(id string) (*ResourceClientCredentials, error) {
	endpoint := fmt.Sprintf(uriApiIntegrations+"/%s/client-credentials", id)

	var response ResourceClientCredentials
	resp, err := c.HTTP.DoRequest("POST", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to create client credentials for Jamf API integration with ID %s: %v", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}
