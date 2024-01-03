// Refactor Complete

// classicapi_gsx_connection.go
// Jamf Pro Classic Api - GSX Connection
// api reference: https://developer.jamf.com/jamf-pro/reference/gsxconnection
// Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriGSXConnection = "/JSSResource/gsxconnection"

// Resource

type ResourceGSXConnection struct {
	Enabled       bool   `xml:"enabled"`
	Username      string `xml:"username"`
	AccountNumber int    `xml:"account_number"`
	Region        string `xml:"region"`
	URI           string `xml:"uri"`
}

// CRUD

// GetGSXConnectionInformation gets the GSX connection settings
func (c *Client) GetGSXConnectionInformation() (*ResourceGSXConnection, error) {
	endpoint := uriGSXConnection

	var gsxConnectionSettings ResourceGSXConnection
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &gsxConnectionSettings)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch GSX Connection settings: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &gsxConnectionSettings, nil
}

// UpdateGSXConnectionInformation updates the GSX connection settings
func (c *Client) UpdateGSXConnectionInformation(settings *ResourceGSXConnection) error {
	endpoint := uriGSXConnection

	requestBody := struct {
		XMLName xml.Name `xml:"gsx_connection"`
		*ResourceGSXConnection
	}{
		ResourceGSXConnection: settings,
	}

	var handleResponse struct{}

	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &handleResponse)
	if err != nil {
		return fmt.Errorf("failed to update GSX Connection settings: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
