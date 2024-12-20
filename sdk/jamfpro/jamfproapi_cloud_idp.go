// jamfproapi_cloud_idp.go
// Jamf Pro Api - Cloud Identity Provider
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-idp
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// Endpoints
const uriCloudIdp = "/api/v1/cloud-idp"

// Response

// ResponseCloudIdentityProvidersList struct for cloud identity providers list
type ResponseCloudIdentityProvidersList struct {
	TotalCount int                             `json:"totalCount"`
	Results    []ResourceCloudIdentityProvider `json:"results"`
}

// Resource

// ResourceCloudIdentityProvider struct for cloud identity provider
type ResourceCloudIdentityProvider struct {
	ID           string `json:"id"`
	DisplayName  string `json:"displayName"`
	Enabled      bool   `json:"enabled"`
	ProviderName string `json:"providerName"`
}

// GetCloudIdentityProviders retrieves all cloud identity provider configurations
func (c *Client) GetCloudIdentityProviders(sort_filter string) (*ResponseCloudIdentityProvidersList, error) {
	endpoint := uriCloudIdp
	resp, err := c.DoPaginatedGet(endpoint, standardPageSize, 0, sort_filter)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "cloud identity providers", err)
	}

	var OutStruct ResponseCloudIdentityProvidersList
	OutStruct.TotalCount = resp.Size
	for _, value := range resp.Results {
		var newObj ResourceCloudIdentityProvider
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "cloud identity provider", err)
		}
		OutStruct.Results = append(OutStruct.Results, newObj)
	}

	return &OutStruct, nil
}

// ResourceCloudIdentityProviderDetails represents a single cloud identity provider configuration
type ResourceCloudIdentityProviderDetails struct {
	ID           string `json:"id"`
	DisplayName  string `json:"displayName"`
	ProviderName string `json:"providerName"`
}

// GetCloudIdentityProviderConfigurationByID retrieves a specific cloud identity provider configuration by its ID
func (c *Client) GetCloudIdentityProviderConfigurationByID(id string) (*ResourceCloudIdentityProviderDetails, error) {
	endpoint := fmt.Sprintf("%s/%s", uriCloudIdp, id)

	var cloudIdp ResourceCloudIdentityProviderDetails
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &cloudIdp)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "cloud identity provider", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &cloudIdp, nil
}