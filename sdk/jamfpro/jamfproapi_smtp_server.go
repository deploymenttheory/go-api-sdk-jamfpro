// jamfproapi_smtp_server.go
package jamfpro

import (
	"fmt"
)

const uriSMTPServer = "/api/v2/smtp-server"

// Resource
type ResourceSMTPServer struct {
	Enabled               bool                                     `json:"enabled"`
	AuthenticationType    string                                   `json:"authenticationType"`
	ConnectionSettings    ResourceSMTPServerConnectionSettings     `json:"connectionSettings"`
	SenderSettings        ResourceSMTPServerSenderSettings         `json:"senderSettings"`
	BasicAuthCredentials  *ResourceSMTPServerBasicAuthCredentials  `json:"basicAuthCredentials,omitempty"`
	GraphApiCredentials   *ResourceSMTPServerGraphApiCredentials   `json:"graphApiCredentials,omitempty"`
	GoogleMailCredentials *ResourceSMTPServerGoogleMailCredentials `json:"googleMailCredentials,omitempty"`
}

// Connection Settings struct
type ResourceSMTPServerConnectionSettings struct {
	Host              string `json:"host"`
	Port              int    `json:"port"`
	EncryptionType    string `json:"encryptionType"`
	ConnectionTimeout int    `json:"connectionTimeout"`
}

// Sender Settings struct
type ResourceSMTPServerSenderSettings struct {
	DisplayName  string `json:"displayName"`
	EmailAddress string `json:"emailAddress"`
}

// Basic Auth Credentials struct
type ResourceSMTPServerBasicAuthCredentials struct {
	Username string `json:"username"`
}

// Graph API Credentials struct
type ResourceSMTPServerGraphApiCredentials struct {
	TenantId string `json:"tenantId"`
	ClientId string `json:"clientId"`
}

// Authentication struct
type ResourceSMTPServerAuthentication struct {
	EmailAddress string `json:"emailAddress"`
	Status       string `json:"status"`
}

// Google Mail Credentials struct
type ResourceSMTPServerGoogleMailCredentials struct {
	ClientId        string                             `json:"clientId"`
	Authentications []ResourceSMTPServerAuthentication `json:"authentications,omitempty"`
}

// GetSMTPServerInformation gets the SMTP server settings
func (c *Client) GetSMTPServerInformation() (*ResourceSMTPServer, error) {
	endpoint := uriSMTPServer

	var smtpSettings ResourceSMTPServer
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &smtpSettings)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "smtp server", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &smtpSettings, nil
}

// UpdateSMTPServerInformation updates the SMTP server settings
func (c *Client) UpdateSMTPServerInformation(settings *ResourceSMTPServer) (*ResourceSMTPServer, error) {
	endpoint := uriSMTPServer

	var updatedSettings ResourceSMTPServer
	resp, err := c.HTTP.DoRequest("PUT", endpoint, settings, &updatedSettings)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdate, "smtp server", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedSettings, nil
}
