// jamfproapi_api_integrations.go
// Jamf Pro Api - API Integrations
// api reference: https://developer.jamf.com/jamf-pro/reference/getallapiintegrations
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/mitchellh/mapstructure"
)

const uriApiIntegrations = "/api/v1/api-integrations"

// List

// ResponseApiIntegrations represents the structure of the response for fetching API integrations
type ResponseApiIntegrationsList struct {
	TotalCount int                      `json:"totalCount"`
	Results    []ResourceApiIntegration `json:"results"`
}

// Resource

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

// CRUD

// GetApiIntegrations fetches all API integrations
func (c *Client) GetApiIntegrations(params url.Values) (*ResponseApiIntegrationsList, error) {
	resp, err := c.DoPaginatedGet(uriApiIntegrations, params)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "api integrations", err)
	}

	var OutStruct ResponseApiIntegrationsList
	OutStruct.TotalCount = resp.Size
	for _, value := range resp.Results {
		var newObj ResourceApiIntegration
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "api integrations", err)
		}
		OutStruct.Results = append(OutStruct.Results, newObj)
	}

	return &OutStruct, nil
}

// GetApiIntegrationByID fetches an API integration by its ID
func (c *Client) GetApiIntegrationByID(id string) (*ResourceApiIntegration, error) {
	endpoint := fmt.Sprintf("%s/%s", uriApiIntegrations, id)

	var integration ResourceApiIntegration
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &integration)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "api integration", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &integration, nil
}

// GetApiIntegrationNameByID fetches an API integration by its display name and then retrieves its details using its ID
func (c *Client) GetApiIntegrationByName(name string) (*ResourceApiIntegration, error) {
	integrations, err := c.GetApiIntegrations(nil)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "api integration", err)
	}

	for _, integration := range integrations.Results {
		if integration.DisplayName == name {
			return &integration, nil
		}
	}

	return nil, fmt.Errorf(errMsgFailedGetByName, "api integration", name, errMsgNoName)
}

// CreateApiIntegration creates a new API integration
func (c *Client) CreateApiIntegration(integration *ResourceApiIntegration) (*ResourceApiIntegration, error) {
	endpoint := uriApiIntegrations

	var response ResourceApiIntegration
	resp, err := c.HTTP.DoRequest("POST", endpoint, integration, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "api integration", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateApiIntegrationByID updates an API integration by its ID
func (c *Client) UpdateApiIntegrationByID(id string, integrationUpdate *ResourceApiIntegration) (*ResourceApiIntegration, error) {
	endpoint := fmt.Sprintf(uriApiIntegrations+"/%s", id)

	var updatedIntegration ResourceApiIntegration
	resp, err := c.HTTP.DoRequest("PUT", endpoint, integrationUpdate, &updatedIntegration)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "api integration", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedIntegration, nil
}

// UpdateApiIntegrationByName updates an API integration based on its display name
func (c *Client) UpdateApiIntegrationByName(name string, integrationUpdate *ResourceApiIntegration) (*ResourceApiIntegration, error) {
	target, err := c.GetApiIntegrationByName(name)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "api integration", name, err)
	}

	target_id := strconv.Itoa(target.ID)
	resp, err := c.UpdateApiIntegrationByID(target_id, integrationUpdate)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "api integration", name, err)
	}

	return resp, nil
}

// DeleteApiIntegrationByID deletes an API integration by its ID
func (c *Client) DeleteApiIntegrationByID(id string) error {
	endpoint := fmt.Sprintf(uriApiIntegrations+"/%s", id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "api integration", id, err)
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
		return fmt.Errorf(errMsgFailedGetByName, "api integration", name, err)
	}

	target_id := strconv.Itoa(target.ID)

	err = c.DeleteApiIntegrationByID(target_id)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByName, "api integration", name, err)
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
		return nil, fmt.Errorf(errMsgFailedRefreshClientCreds, id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}
