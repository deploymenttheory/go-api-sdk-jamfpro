// jamfproapi_cloud_azure.go
// Jamf Pro Api - Cloud Azure (Cloud IDP)
// api reference: https://developer.jamf.com/jamf-pro/reference/post_v1-cloud-azure
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
)

const uriCloudIdentityProvider = "/api/v1/cloud-azure"

type ResourceCloudIdp struct {
	CloudIdPCommon CloudIdpListItem       `json:"cloudIdPCommon"`
	Server         ResourceCloudIdpServer `json:"server"`
}

type CloudIdpListItem struct {
	DisplayName  string `json:"displayName"`
	ProviderName string `json:"providerName"`
}

type ResourceCloudIdpServer struct {
	ID                                       string                                     `json:"id"`
	TenantId                                 string                                     `json:"tenantId"`
	Enabled                                  bool                                       `json:"enabled"`
	Migrated                                 bool                                       `json:"migrated"`
	Mappings                                 CloudIdpServerSubsetCloudIdpServerMappings `json:"mappings"`
	SearchTimeout                            int                                        `json:"searchTimeout"`
	TransitiveMembershipEnabled              bool                                       `json:"transitiveMembershipEnabled"`
	TransitiveMembershipUserField            string                                     `json:"transitiveMembershipUserField"`
	TransitiveDirectoryMembershipEnabled     bool                                       `json:"transitiveDirectoryMembershipEnabled"`
	MembershipCalculationOptimizationEnabled bool                                       `json:"membershipCalculationOptimizationEnabled"`
	Code                                     string                                     `json:"code"`
}

type CloudIdpServerSubsetCloudIdpServerMappings struct {
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
type ResponseCloudIdpCreate struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// GetDefaultCloudIdentityProvider retrieves the default server configuration for the Cloud Identity Provider.
func (c *Client) GetDefaultCloudIdentityProvider() (*ResourceCloudIdpServer, error) {
	endpoint := uriCloudIdentityProvider + "/defaults/server-configuration"

	var defaultCloudIdPServer ResourceCloudIdpServer
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &defaultCloudIdPServer)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "Azure Cloud IDP", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &defaultCloudIdPServer, nil
}

// GetCloudIdentityProviderByID retrieves Cloud Identity Provider information.
func (c *Client) GetCloudIdentityProviderByID(id string) (*ResourceCloudIdp, error) {
	endpoint := fmt.Sprintf("%s/%s", uriCloudIdentityProvider, id)

	var cloudIDP ResourceCloudIdp
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &cloudIDP)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "Azure Cloud IDP", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &cloudIDP, nil
}

// CreateCloudIdentityProvider creates a new Cloud Identity Provider.
func (c *Client) CreateCloudIdentityProvider(cloudIdP *ResourceCloudIdp) (*ResponseCloudIdpCreate, error) {
	endpoint := uriCloudIdentityProvider

	var responseCreateCloudIDP ResponseCloudIdpCreate
	resp, err := c.HTTP.DoRequest("POST", endpoint, cloudIdP, &responseCreateCloudIDP)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "Azure Cloud IDP", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseCreateCloudIDP, nil
}

// UpdateCloudIdentityProviderById updates an existing Cloud Identity Provider by its ID.
func (c *Client) UpdateCloudIdentityProviderByID(id string, cloudIdPUpdate *ResourceCloudIdp) (*ResourceCloudIdp, error) {
	endpoint := fmt.Sprintf("%s%s", uriCloudIdentityProvider, id)

	var updatedCloudIDP ResourceCloudIdp
	resp, err := c.HTTP.DoRequest("PUT", endpoint, cloudIdPUpdate, &updatedCloudIDP) // or "PATCH" based on API
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "Azure Cloud IDP", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedCloudIDP, nil
}

// DeleteCloudIdentityProviderById deletes a Cloud Identity Provider by its ID.
func (c *Client) DeleteCloudIdentityProviderByID(id string) error {
	endpoint := fmt.Sprintf("%s%s", uriCloudIdentityProvider, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "Azure Cloud IDP", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	// Success, no error
	return nil
}

// GetCloudIdentityProviderDefaultServerMappings retrieves the default mappings for the Cloud Identity Provider.
func (c *Client) GetCloudIdentityProviderDefaultServerMappings() (*CloudIdpServerSubsetCloudIdpServerMappings, error) {
	endpoint := uriCloudIdentityProvider + "defaults/mappings"

	var defaultMappings CloudIdpServerSubsetCloudIdpServerMappings
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &defaultMappings)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "Azure Cloud IDP Server Mappings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &defaultMappings, nil
}
