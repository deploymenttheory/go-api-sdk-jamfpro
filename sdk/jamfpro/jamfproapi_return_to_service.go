// jamfproapi_return_to_service.go
// Jamf Pro Api - Return to Service
// api reference: none available
// docs: https://learn.jamf.com/en-US/bundle/technical-articles/page/Return_to_Service.html
// Jamf Pro Api requires the structs to support an JSON data structure.

package jamfpro

import (
	"fmt"
)

const uriReturnToService = "/api/v1/return-to-service"

// List

// Structs to support JSON data structure
type ResponseReturnToServiceList struct {
	TotalCount int                                    `json:"totalCount"`
	Results    []ResponseReturnToServiceConfiguration `json:"results"`
}

// Response
type ResponseReturnToServiceConfiguration struct {
	ID            string `json:"id"`
	DisplayName   string `json:"displayName"`
	WifiProfileID string `json:"wifiProfileId"`
}

// Resource

// ResourceReturnToServiceConfiguration represents the structure for a Return to Service configuration
type ResourceReturnToServiceConfiguration struct {
	DisplayName                                    string `json:"displayName"`
	SsoForEnrollmentEnabled                        bool   `json:"ssoForEnrollmentEnabled"`
	SsoBypassAllowed                               bool   `json:"ssoBypassAllowed"`
	SsoEnabled                                     bool   `json:"ssoEnabled"`
	SsoForMacOsSelfServiceEnabled                  bool   `json:"ssoForMacOsSelfServiceEnabled"`
	TokenExpirationDisabled                        bool   `json:"tokenExpirationDisabled"`
	UserAttributeEnabled                           bool   `json:"userAttributeEnabled"`
	UserAttributeName                              string `json:"userAttributeName"`
	UserMapping                                    string `json:"userMapping"`
	EnrollmentSsoForAccountDrivenEnrollmentEnabled bool   `json:"enrollmentSsoForAccountDrivenEnrollmentEnabled"`
	GroupEnrollmentAccessEnabled                   bool   `json:"groupEnrollmentAccessEnabled"`
	GroupAttributeName                             string `json:"groupAttributeName"`
	GroupRdnKey                                    string `json:"groupRdnKey"`
	GroupEnrollmentAccessName                      string `json:"groupEnrollmentAccessName"`
	IdpProviderType                                string `json:"idpProviderType"`
	OtherProviderTypeName                          string `json:"otherProviderTypeName"`
	MetadataSource                                 string `json:"metadataSource"`
	SessionTimeout                                 int    `json:"sessionTimeout"`
	DeviceType                                     string `json:"deviceType"`
	Enabled                                        bool   `json:"enabled"`
}

// GetReturnToServiceConfigurations fetches a list of devices that are in the Return to Service state.
func (c *Client) GetReturnToServiceConfigurations() ([]ResponseReturnToServiceList, error) {
	endpoint := uriReturnToService
	var out []ResponseReturnToServiceList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "Return To Service Configurations", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return out, nil
}

// GetReturnToServiceConfigurationByID fetches a specific Return to Service configuration by ID.
func (c *Client) GetReturnToServiceConfigurationByID(id string) (*ResponseReturnToServiceConfiguration, error) {
	endpoint := fmt.Sprintf("%s/%s", uriReturnToService, id)
	var out ResponseReturnToServiceConfiguration
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "Return To Service Configuration", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// CreateReturnToServiceConfiguration creates a new Return to Service configuration.
func (c *Client) CreateReturnToServiceConfiguration(config ResourceReturnToServiceConfiguration) (*ResponseReturnToServiceConfiguration, error) {
	endpoint := uriReturnToService
	var out ResponseReturnToServiceConfiguration
	resp, err := c.HTTP.DoRequest("POST", endpoint, config, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "Return To Service Configuration", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// UpdateReturnToServiceConfigurationByID updates an existing Return to Service configuration by ID.
func (c *Client) UpdateReturnToServiceConfigurationByID(id string, config ResourceReturnToServiceConfiguration) (*ResponseReturnToServiceConfiguration, error) {
	endpoint := fmt.Sprintf("%s/%s", uriReturnToService, id)
	var out ResponseReturnToServiceConfiguration
	resp, err := c.HTTP.DoRequest("PUT", endpoint, config, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "Return To Service Configuration", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// DeleteReturnToServiceConfigurationByID deletes a Return to Service configuration by ID.
func (c *Client) DeleteReturnToServiceConfigurationByID(id string) error {
	endpoint := fmt.Sprintf("%s/%s", uriReturnToService, id)
	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "Return To Service Configuration", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
