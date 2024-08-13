// jamfproapi_csa.go
// Jamf Pro Api - CSA Tokens
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-csa-token
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import "fmt"

const uriCSATokenExchange = "/api/v1/csa/token"

// Response

// ResponseCSATokenExchangeDetails represents the response structure for the CSA token exchange details.
type ResponseCSATokenExchangeDetails struct {
	TenantID          string   `json:"tenantId"`
	Subject           string   `json:"subject"`
	RefreshExpiration int      `json:"refreshExpiration"`
	Scopes            []string `json:"scopes"`
}

// ResponseCSATenantID represents the response structure for the CSA tenant ID.
type ResponseCSATenantID struct {
	TenantID string `json:"tenantId"`
}

// CRUD

// GetCSATokenExchangeDetails retrieves details regarding the CSA token exchange.
func (c *Client) GetCSATokenExchangeDetails() (*ResponseCSATokenExchangeDetails, error) {
	endpoint := uriCSATokenExchange
	var details ResponseCSATokenExchangeDetails

	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &details)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "csa token exchange details", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &details, nil
}

// GetCSATenantID retrieves the CSA tenant ID.
func (c *Client) GetCSATenantID() (*ResponseCSATenantID, error) {
	endpoint := fmt.Sprintf("%s/tenant-id", uriCSATokenExchange)
	var tenantID ResponseCSATenantID

	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &tenantID)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "CSA tenant ID", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &tenantID, nil
}

// DeleteCSATokenExchange deletes the CSA token exchange, disabling Jamf Pro's ability to authenticate with cloud-hosted services.
func (c *Client) DeleteCSATokenExchange() (*SharedResourcResponseError, error) {
	endpoint := uriCSATokenExchange

	var responseError SharedResourcResponseError
	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, &responseError)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedDelete, "CSA token exchange", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != 204 {
		return &responseError, fmt.Errorf("failed to delete CSA token exchange: %v", responseError)
	}

	return nil, nil
}
