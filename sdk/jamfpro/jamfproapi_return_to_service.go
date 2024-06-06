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
	Results    []ResourceReturnToServiceConfiguration `json:"results"`
}

// Resource
type ResourceReturnToServiceConfiguration struct {
	ID            string `json:"id"`
	DisplayName   string `json:"displayName"`
	WifiProfileID string `json:"wifiProfileId"`
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
func (c *Client) GetReturnToServiceConfigurationByID(id string) (*ResourceReturnToServiceConfiguration, error) {
	endpoint := fmt.Sprintf("%s/%s", uriReturnToService, id)
	var out ResourceReturnToServiceConfiguration
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
func (c *Client) CreateReturnToServiceConfiguration(config ResourceReturnToServiceConfiguration) (*ResourceReturnToServiceConfiguration, error) {
	endpoint := uriReturnToService
	var out ResourceReturnToServiceConfiguration
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
func (c *Client) UpdateReturnToServiceConfigurationByID(id string, config ResourceReturnToServiceConfiguration) (*ResourceReturnToServiceConfiguration, error) {
	endpoint := fmt.Sprintf("%s/%s", uriReturnToService, id)
	var out ResourceReturnToServiceConfiguration
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
