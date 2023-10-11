// cloudAzure.go
// Jamf Pro Api
// Jamf Pro API requires the structs to support JSON.

package jamfpro

import (
	"fmt"
)

const uriAPICloudAzure = "/api/v1/cloud-azure"

type ResponseDefaultServerConfiguration struct {
	ID                                       string                  `json:"id"`
	TenantId                                 string                  `json:"tenantId"`
	Enabled                                  bool                    `json:"enabled"`
	Migrated                                 bool                    `json:"migrated"`
	Mappings                                 ResponseDefaultMappings `json:"mappings"`
	SearchTimeout                            int                     `json:"searchTimeout"`
	TransitiveMembershipEnabled              bool                    `json:"transitiveMembershipEnabled"`
	TransitiveMembershipUserField            string                  `json:"transitiveMembershipUserField"`
	TransitiveDirectoryMembershipEnabled     bool                    `json:"transitiveDirectoryMembershipEnabled"`
	MembershipCalculationOptimizationEnabled bool                    `json:"membershipCalculationOptimizationEnabled"`
}

type ResponseDefaultMappings struct {
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

type ResponseCloudIdentityProvider struct {
	CloudIdPCommon CloudIdPCommon                          `json:"cloudIdPCommon"`
	Server         AzureCloudIdentityProviderConfiguration `json:"server"`
}

type CloudIdPCommon struct {
	ID           string `json:"id"`
	DisplayName  string `json:"displayName"`
	ProviderName string `json:"providerName"`
}

type AzureCloudIdentityProviderConfiguration struct {
	ID                                       string                  `json:"id"`
	TenantId                                 string                  `json:"tenantId"`
	Enabled                                  bool                    `json:"enabled"`
	Migrated                                 bool                    `json:"migrated"`
	Mappings                                 ResponseDefaultMappings `json:"mappings"`
	SearchTimeout                            int                     `json:"searchTimeout"`
	TransitiveMembershipEnabled              bool                    `json:"transitiveMembershipEnabled"`
	TransitiveMembershipUserField            string                  `json:"transitiveMembershipUserField"`
	TransitiveDirectoryMembershipEnabled     bool                    `json:"transitiveDirectoryMembershipEnabled"`
	MembershipCalculationOptimizationEnabled bool                    `json:"membershipCalculationOptimizationEnabled"`
}

type RequestAzureCloudIdentityProvider struct {
	CloudIdPCommon RequestCloudIdPCommon `json:"cloudIdPCommon"`
	Server         RequestServer         `json:"server"`
}

type RequestCloudIdPCommon struct {
	DisplayName  string `json:"displayName"`
	ProviderName string `json:"providerName"` // AZURE
}

type RequestServer struct {
	TenantId                                 string                 `json:"tenantId"`
	Enabled                                  bool                   `json:"enabled"`
	Mappings                                 RequestDefaultMappings `json:"mappings"`
	SearchTimeout                            int                    `json:"searchTimeout"`
	TransitiveMembershipEnabled              bool                   `json:"transitiveMembershipEnabled"`
	TransitiveMembershipUserField            string                 `json:"transitiveMembershipUserField"`
	TransitiveDirectoryMembershipEnabled     bool                   `json:"transitiveDirectoryMembershipEnabled"`
	MembershipCalculationOptimizationEnabled bool                   `json:"membershipCalculationOptimizationEnabled"`
	Code                                     string                 `json:"code"`
}

type RequestDefaultMappings struct {
	SearchTimeout int    `json:"searchTimeout"`
	UserId        string `json:"userId,omitempty"`
	UserName      string `json:"userName,omitempty"`
	RealName      string `json:"realName,omitempty"`
	Email         string `json:"email,omitempty"`
	Department    string `json:"department,omitempty"`
	Building      string `json:"building,omitempty"`
	Room          string `json:"room,omitempty"`
	Phone         string `json:"phone,omitempty"`
	Position      string `json:"position,omitempty"`
	GroupId       string `json:"groupId,omitempty"`
	GroupName     string `json:"groupName,omitempty"`
}

// GetDefaultMappings retrieves the default mappings for Azure cloud.
func (c *Client) GetDefaultMappings() (*ResponseDefaultMappings, error) {
	url := fmt.Sprintf("%s/defaults/mappings", uriAPICloudAzure)

	var defaultMappings ResponseDefaultMappings
	if err := c.DoRequest("GET", url, nil, nil, &defaultMappings); err != nil {
		return nil, fmt.Errorf("failed to get default mappings: %v", err)
	}

	return &defaultMappings, nil
}

// GetDefaultServerConfiguration retrieves the default server configuration for Azure cloud.
func (c *Client) GetDefaultServerConfiguration() (*ResponseDefaultServerConfiguration, error) {
	url := fmt.Sprintf("%s/defaults/server-configuration", uriAPICloudAzure)

	var defaultServerConfig ResponseDefaultServerConfiguration
	if err := c.DoRequest("GET", url, nil, nil, &defaultServerConfig); err != nil {
		return nil, fmt.Errorf("failed to get default server configuration: %v", err)
	}

	return &defaultServerConfig, nil
}

// GetAzureCloudIdentityProviderByID retrieves the Azure Cloud Identity Provider configuration by its ID.
func (c *Client) GetAzureCloudIdentityProviderByID(id string) (*ResponseCloudIdentityProvider, error) {
	url := fmt.Sprintf("%s/%s", uriAPICloudAzure, id)

	var cloudIdentityProviderConfig ResponseCloudIdentityProvider
	if err := c.DoRequest("GET", url, nil, nil, &cloudIdentityProviderConfig); err != nil {
		return nil, fmt.Errorf("failed to get Azure Cloud Identity Provider configuration by ID: %v", err)
	}

	return &cloudIdentityProviderConfig, nil
}

// CreateAzureCloudIdentityProviderConfiguration creates a new Azure Cloud Identity Provider configuration.
func (c *Client) CreateAzureCloudIdentityProviderConfiguration(config *RequestAzureCloudIdentityProvider) (*ResponseCloudIdentityProvider, error) {
	url := uriAPICloudAzure

	var responseConfig ResponseCloudIdentityProvider
	if err := c.DoRequest("POST", url, config, nil, &responseConfig); err != nil {
		return nil, fmt.Errorf("failed to create Azure Cloud Identity Provider configuration: %v", err)
	}

	return &responseConfig, nil
}

// UpdateAzureCloudIdentityProviderConfiguration updates an existing Azure Cloud Identity Provider configuration.
func (c *Client) UpdateAzureCloudIdentityProviderConfiguration(id string, config *RequestAzureCloudIdentityProvider) (*ResponseCloudIdentityProvider, error) {
	url := fmt.Sprintf("%s/%s", uriAPICloudAzure, id)

	var updatedConfig ResponseCloudIdentityProvider
	if err := c.DoRequest("PUT", url, config, nil, &updatedConfig); err != nil {
		return nil, fmt.Errorf("failed to update Azure Cloud Identity Provider configuration: %v", err)
	}

	return &updatedConfig, nil
}

// DeleteCloudIdentityProviderConfiguration deletes the Azure Cloud Identity Provider configuration with the given ID.
func (c *Client) DeleteCloudIdentityProviderConfiguration(id string) error {
	url := fmt.Sprintf("%s/%s", uriAPICloudAzure, id)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete Azure Cloud Identity Provider configuration: %v", err)
	}

	return nil
}
