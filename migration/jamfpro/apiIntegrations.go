// apiIntegrations.go
// Jamf Pro Api
// Jamf Pro Api requires the structs to support JSON.

package jamfpro

import (
	"fmt"
)

const uriAPIIntegrations = "/api/v1/api-integrations"

type ResponseAPIIntegration struct {
	TotalCount int              `json:"totalCount"`
	Results    []APIIntegration `json:"results"`
}

type APIIntegration struct {
	ID                         int      `json:"id"`
	AuthorizationScopes        []string `json:"authorizationScopes"`
	DisplayName                string   `json:"displayName"`
	Enabled                    bool     `json:"enabled"`
	AccessTokenLifetimeSeconds int      `json:"accessTokenLifetimeSeconds"`
	AppType                    string   `json:"appType"`
	ClientId                   string   `json:"clientId"`
}

func (c *Client) GetApiIntegrationIdByName(name string) (int, error) {
	var id int
	integrations, err := c.GetApiIntegrations()
	if err != nil {
		return 0, err
	}

	for _, v := range integrations.Results {
		if v.DisplayName == name {
			id = v.ID
			break
		}
	}
	return id, err
}

func (c *Client) GetApiIntegrationByName(name string) (*APIIntegration, error) {
	allIntegrationsResponse, err := c.GetApiIntegrations()
	if err != nil {
		return nil, err
	}

	for _, integration := range allIntegrationsResponse.Results {
		if integration.DisplayName == name {
			return &integration, nil
		}
	}

	return nil, fmt.Errorf("API integration with name '%s' not found", name)
}

func (c *Client) GetApiIntegrations() (*ResponseAPIIntegration, error) {
	uri := fmt.Sprintf("%s?page=0&page-size=100&sort=id%%3Aasc", uriAPIIntegrations)

	var out ResponseAPIIntegration
	err := c.DoRequest("GET", uri, nil, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to get API integrations: %v", err)
	}

	return &out, nil
}

func (c *Client) GetApiIntegrationByID(integrationID int) (*APIIntegration, error) {
	uri := fmt.Sprintf("%s/%d", uriAPIIntegrations, integrationID)

	var out APIIntegration
	err := c.DoRequest("GET", uri, nil, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to get API integration by ID: %v", err)
	}

	return &out, nil
}

func (c *Client) CreateApiIntegration(
	displayName *string,
	authorizationScopes *[]string,
	enabled *bool,
	accessTokenLifetimeSeconds *int) (*APIIntegration, error) {

	in := struct {
		DisplayName                *string   `json:"displayName"`
		Enabled                    *bool     `json:"enabled"`
		AuthorizationScopes        *[]string `json:"authorizationScopes"`
		AccessTokenLifetimeSeconds *int      `json:"accessTokenLifetimeSeconds"`
	}{
		DisplayName:                displayName,
		Enabled:                    enabled,
		AuthorizationScopes:        authorizationScopes,
		AccessTokenLifetimeSeconds: accessTokenLifetimeSeconds,
	}

	var out *APIIntegration

	err := c.DoRequest("POST", uriAPIIntegrations, in, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to create API integration: %v", err)
	}
	return out, nil
}

func (c *Client) UpdateApiIntegration(d *APIIntegration) (*APIIntegration, error) {
	uri := fmt.Sprintf("%s/%d", uriAPIIntegrations, d.ID)
	updatedIntegration := &APIIntegration{}

	// Perform the PUT request
	err := c.DoRequest("PUT", uri, d, nil, updatedIntegration)
	if err != nil {
		return nil, fmt.Errorf("failed to update API integration: %v", err)
	}
	return updatedIntegration, nil
}

func (c *Client) DeleteApiIntegration(integrationID int) error {
	uri := fmt.Sprintf("%s/%d", uriAPIIntegrations, integrationID)

	// Perform the DELETE request
	err := c.DoRequest("DELETE", uri, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete API integration: %v", err)
	}
	return nil
}
