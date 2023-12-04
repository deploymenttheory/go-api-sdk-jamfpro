// classicapi_smtp_server.go
// Jamf Pro Classic Api - SMTP Server
// api reference: https://developer.jamf.com/jamf-pro/reference/smtpserver
// Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriSMTPServer = "/JSSResource/smtpserver"

// Struct for the SMTP server settings response

type ResponseSMTPServer struct {
	Enabled               bool   `xml:"enabled"`
	Host                  string `xml:"host"`
	Port                  int    `xml:"port"`
	Timeout               int    `xml:"timeout"`
	AuthorizationRequired bool   `xml:"authorization_required"`
	Username              string `xml:"username"`
	Password              string `xml:"password"`
	SSL                   bool   `xml:"ssl"`
	TLS                   bool   `xml:"tls"`
	SendFromName          string `xml:"send_from_name"`
	SendFromEmail         string `xml:"send_from_email"`
}

// GetSMTPServerInformation gets the SMTP server settings
func (c *Client) GetSMTPServerInformation() (*ResponseSMTPServer, error) {
	endpoint := uriSMTPServer

	var smtpSettings ResponseSMTPServer
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &smtpSettings)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch SMTP server settings: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &smtpSettings, nil
}

// UpdateSMTPServerInformation updates the SMTP server settings
func (c *Client) UpdateSMTPServerInformation(settings *ResponseSMTPServer) error {
	endpoint := uriSMTPServer

	// Wrap the settings with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"smtp_server"`
		*ResponseSMTPServer
	}{
		ResponseSMTPServer: settings,
	}

	// Create a dummy struct for the response
	var handleResponse struct{}

	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &handleResponse)
	if err != nil {
		return fmt.Errorf("failed to update SMTP server settings: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
