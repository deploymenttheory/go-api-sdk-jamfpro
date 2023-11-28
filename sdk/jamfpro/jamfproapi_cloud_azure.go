// jamfproapi_cloud_azure.go
// Jamf Pro Api - Cloud Azure (Cloud IDP)
// api reference: https://developer.jamf.com/jamf-pro/reference/post_v1-cloud-azure
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
)

const uriCloudIdentityProvider = "/api/v1/cloud-azure/"

type ResponseCloudIDP struct {
	CloudIdPCommon CloudIdPCommon `json:"cloudIdPCommon"`
	Server         CloudIdPServer `json:"server"`
}

type CloudIdPCommon struct {
	DisplayName  string `json:"displayName"`
	ProviderName string `json:"providerName"`
}

type CloudIdPServer struct {
	ID                                       string                 `json:"id"`
	TenantId                                 string                 `json:"tenantId"`
	Enabled                                  bool                   `json:"enabled"`
	Migrated                                 bool                   `json:"migrated"`
	Mappings                                 CloudIdPServerMappings `json:"mappings"`
	SearchTimeout                            int                    `json:"searchTimeout"`
	TransitiveMembershipEnabled              bool                   `json:"transitiveMembershipEnabled"`
	TransitiveMembershipUserField            string                 `json:"transitiveMembershipUserField"`
	TransitiveDirectoryMembershipEnabled     bool                   `json:"transitiveDirectoryMembershipEnabled"`
	MembershipCalculationOptimizationEnabled bool                   `json:"membershipCalculationOptimizationEnabled"`
	Code                                     string                 `json:"code"`
}

type CloudIdPServerMappings struct {
	UserId     string `json:"userId"`
	UserName   string `json:"userName"`
	RealName   string `json:"realName"`
	Email      string `json:"email"`
	Department string `json:"department"`
	Building   string `json:"building"`
	Room       string `json:"room"`
	Phone      string `json:"phone"`
	Position   string `json:"position"`
	GroupId    string `json:"groupId"`
	GroupName  string `json:"groupName"`
}

// ResponseCloudIDPCreate represents the response received after creating a Cloud Identity Provider.
type ResponseCloudIDPCreate struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// GetDefaultCloudIdentityProvider retrieves the default server configuration for the Cloud Identity Provider.
func (c *Client) GetDefaultCloudIdentityProvider() (*CloudIdPServer, error) {
	endpoint := uriCloudIdentityProvider + "defaults/server-configuration"

	var defaultCloudIdPServer CloudIdPServer
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &defaultCloudIdPServer)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch default Azure Cloud Identity Provider server configuration: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &defaultCloudIdPServer, nil
}

// GetCloudIdentityProviderByID retrieves Cloud Identity Provider information.
func (c *Client) GetCloudIdentityProviderByID(id string) (*ResponseCloudIDP, error) {
	endpoint := fmt.Sprintf("%s%s", uriCloudIdentityProvider, id)

	var cloudIDP ResponseCloudIDP
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &cloudIDP)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Azure Cloud Identity Provider: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &cloudIDP, nil
}

// GetCloudIdentityProviderDefaultServerMappings retrieves the default mappings for the Cloud Identity Provider.
func (c *Client) GetCloudIdentityProviderDefaultServerMappings() (*CloudIdPServerMappings, error) {
	endpoint := uriCloudIdentityProvider + "defaults/mappings"

	var defaultMappings CloudIdPServerMappings
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &defaultMappings)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Azure Cloud IDP default mappings: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &defaultMappings, nil
}

// CreateCloudIdentityProvider creates a new Cloud Identity Provider.
func (c *Client) CreateCloudIdentityProvider(cloudIdPData *ResponseCloudIDP) (*ResponseCloudIDPCreate, error) {
	endpoint := uriCloudIdentityProvider

	var responseCreateCloudIDP ResponseCloudIDPCreate
	resp, err := c.HTTP.DoRequest("POST", endpoint, cloudIdPData, &responseCreateCloudIDP)
	if err != nil {
		return nil, fmt.Errorf("failed to create Azure Cloud Identity Provider: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseCreateCloudIDP, nil
}

// UpdateCloudIdentityProviderById updates an existing Cloud Identity Provider by its ID.
func (c *Client) UpdateCloudIdentityProviderById(id string, cloudIdPData *ResponseCloudIDP) (*ResponseCloudIDP, error) {
	endpoint := fmt.Sprintf("%s%s", uriCloudIdentityProvider, id)

	var updatedCloudIDP ResponseCloudIDP
	resp, err := c.HTTP.DoRequest("PUT", endpoint, cloudIdPData, &updatedCloudIDP) // or "PATCH" based on API
	if err != nil {
		return nil, fmt.Errorf("failed to update Azure Cloud Identity Provider with ID %s: %v", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedCloudIDP, nil
}

// DeleteCloudIdentityProviderById deletes a Cloud Identity Provider by its ID.
func (c *Client) DeleteCloudIdentityProviderById(id string) error {
	endpoint := fmt.Sprintf("%s%s", uriCloudIdentityProvider, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete Azure Cloud Identity Provider with ID %s: %v", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	// Success, no error
	return nil
}
