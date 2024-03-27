// jamfproapi_smtp_server.go
// Jamf Pro Api - SMTP Server
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-smtp-server
// Jamf Pro API requires the structs to support an JSON data structure.

package jamfpro

import (
	"fmt"
)

const uriSMTPServer = "/api/v1/smtp-server" // Constant for the URL

// Updated struct for the SMTP server settings based on the new JSON data model
type ResourceSMTPServer struct {
	Enabled                bool   `json:"enabled"`
	Server                 string `json:"server,omitempty"`
	Port                   int    `json:"port,omitempty"`
	EncryptionType         string `json:"encryptionType,omitempty"`
	ConnectionTimeout      int    `json:"connectionTimeout,omitempty"`
	SenderDisplayName      string `json:"senderDisplayName,omitempty"`
	SenderEmailAddress     string `json:"senderEmailAddress,omitempty"`
	RequiresAuthentication bool   `json:"requiresAuthentication"`
	Username               string `json:"username,omitempty"`
}

// GetSMTPServerInformation gets the SMTP server settings
func (c *Client) GetSMTPServerInformation() (*ResourceSMTPServer, error) {
	endpoint := uriSMTPServer

	var smtpSettings ResourceSMTPServer
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &smtpSettings)
	if err != nil {
		return nil, fmt.Errorf("failed to get smtp server information: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &smtpSettings, nil
}

// UpdateSMTPServerInformation updates the SMTP server settings
func (c *Client) UpdateSMTPServerInformation(settings *ResourceSMTPServer) error {
	endpoint := uriSMTPServer

	// No need to wrap settings for JSON
	resp, err := c.HTTP.DoRequest("PUT", endpoint, settings, nil)
	if err != nil {
		return fmt.Errorf("failed to update smtp server information: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
