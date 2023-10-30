// jamfproapi_api_integrations.go
// Jamf Pro Api - API Integrations
// api reference: https://developer.jamf.com/jamf-pro/reference/getallapiintegrations
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import "fmt"

const uriApiIntegrations = "/api/v1/api-integrations"

// ResponseApiIntegrations represents the structure of the response for fetching API integrations
type ResponseApiIntegrations struct {
	TotalCount int              `json:"totalCount"`
	Results    []ApiIntegration `json:"results"`
}

// Integration represents the details of an individual API integration
type ApiIntegration struct {
	ID                         int      `json:"id,omitempty"`
	AuthorizationScopes        []string `json:"authorizationScopes,omitempty"`
	DisplayName                string   `json:"displayName,omitempty"`
	Enabled                    bool     `json:"enabled,omitempty"`
	AccessTokenLifetimeSeconds int      `json:"accessTokenLifetimeSeconds,omitempty"`
	AppType                    string   `json:"appType,omitempty"`
	ClientID                   string   `json:"clientId,omitempty"`
}

// Integration represents the details of Api client credentials
type ClientCredentials struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

// GetApiIntegrations fetches all API integrations
func (c *Client) GetApiIntegrations() (*ResponseApiIntegrations, error) {
	var integrationsList ResponseApiIntegrations
	resp, err := c.HTTP.DoRequest("GET", uriApiIntegrations, nil, &integrationsList)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Jamf API integrations: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &integrationsList, nil
}

// GetApiIntegrationByID fetches an API integration by its ID
func (c *Client) GetApiIntegrationByID(id int) (*ApiIntegration, error) {
	// Construct the URL with the provided ID
	endpoint := fmt.Sprintf("%s/%d", uriApiIntegrations, id)

	var integration ApiIntegration
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
func (c *Client) GetApiIntegrationNameByID(name string) (*ApiIntegration, error) {
	integrationsList, err := c.GetApiIntegrations()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all Jamf API integrations: %v", err)
	}

	// Search for the integration with the given name
	for _, integration := range integrationsList.Results {
		if integration.DisplayName == name {
			return c.GetApiIntegrationByID(integration.ID)
		}
	}

	return nil, fmt.Errorf("no Jamf API integration found with the name %s", name)
}

// CreateApiIntegration creates a new API integration
func (c *Client) CreateApiIntegration(integration *ApiIntegration) (*ApiIntegration, error) {
	endpoint := uriApiIntegrations

	var response ApiIntegration
	resp, err := c.HTTP.DoRequest("POST", endpoint, integration, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to create Jamf API integration: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// CreateClientCredentialsByApiRoleID creates new client credentials for an API integration by its ID
func (c *Client) CreateClientCredentialsByApiRoleID(id string) (*ClientCredentials, error) {
	// Construct the URL with the provided ID
	endpoint := fmt.Sprintf(uriApiIntegrations+"/%s/client-credentials", id)

	var response ClientCredentials
	resp, err := c.HTTP.DoRequest("POST", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to create client credentials for Jamf API integration with ID %s: %v", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateApiIntegrationByID updates an API integration by its ID
func (c *Client) UpdateApiIntegrationByID(id string, integrationUpdate *ApiIntegration) (*ApiIntegration, error) {
	// Construct the URL with the provided ID
	endpoint := fmt.Sprintf(uriApiIntegrations+"/%s", id)

	var updatedIntegration ApiIntegration
	resp, err := c.HTTP.DoRequest("PUT", endpoint, integrationUpdate, &updatedIntegration)
	if err != nil {
		return nil, fmt.Errorf("failed to update Jamf API integration with ID %s: %v", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedIntegration, nil
}

// UpdateApiIntegrationByName updates an API integration based on its display name
func (c *Client) UpdateApiIntegrationByName(name string, updatedIntegration *ApiIntegration) (*ApiIntegration, error) {
	integrationsList, err := c.GetApiIntegrations()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all Jamf API integrations: %v", err)
	}

	// Search for the integration with the given name
	for _, integration := range integrationsList.Results {
		if integration.DisplayName == name {
			// Update the integration with the provided ID
			return c.UpdateApiIntegrationByID(fmt.Sprintf("%d", integration.ID), updatedIntegration)
		}
	}

	return nil, fmt.Errorf("no Jamf API integration found with the name %s", name)
}

// UpdateClientCredentialsByApiIntegrationID updates client credentials for an API integration by its ID
func (c *Client) UpdateClientCredentialsByApiIntegrationID(id string) (*ClientCredentials, error) {
	// Construct the URL with the provided ID
	endpoint := fmt.Sprintf(uriApiIntegrations+"/%s/client-credentials", id)

	var updatedCredentials ClientCredentials
	resp, err := c.HTTP.DoRequest("POST", endpoint, nil, &updatedCredentials)
	if err != nil {
		return nil, fmt.Errorf("failed to update client credentials for Jamf API integration with ID %s: %v", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedCredentials, nil
}

// DeleteApiIntegrationByID deletes an API integration by its ID
func (c *Client) DeleteApiIntegrationByID(id string) error {
	// Construct the URL with the provided ID
	endpoint := fmt.Sprintf(uriApiIntegrations+"/%s", id)

	// Perform the DELETE request
	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete Jamf API integration with ID %s: %v", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteApiIntegrationByName deletes an API integration by its display name
func (c *Client) DeleteApiIntegrationByName(name string) error {
	integrationsList, err := c.GetApiIntegrations()
	if err != nil {
		return fmt.Errorf("failed to fetch all Jamf API integrations: %v", err)
	}

	// Search for the integration with the given name
	for _, integration := range integrationsList.Results {
		if integration.DisplayName == name {
			return c.DeleteApiIntegrationByID(fmt.Sprintf("%d", integration.ID))
		}
	}

	return fmt.Errorf("no Jamf API integration found with the name %s", name)
}
