package models

import "encoding/xml"

// ResourceExchangeWebServicesConfigurationProfile represents the top-level structure of the plist for configuring Exchange Web Services accounts
type ResourceExchangeWebServicesConfigurationProfile struct {
	XMLName                  xml.Name                                             `xml:"plist"`
	Version                  string                                               `xml:"version,attr"`
	Payload                  ExchangeWebServicesConfigurationProfileSubsetPayload `xml:"dict>array>dict"`
	PayloadDescription       string                                               `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                               `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier        string                                               `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                               `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                               `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                               `xml:"PayloadScope,omitempty"`
	PayloadType              string                                               `xml:"PayloadType,omitempty"`
	PayloadUUID              string                                               `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                                               `xml:"PayloadVersion,omitempty"`
}

// ExchangeWebServicesConfigurationProfileSubsetPayload represents the content structure for configuring Exchange Web Services accounts
type ExchangeWebServicesConfigurationProfileSubsetPayload struct {
	PayloadContent     []ExchangeWebServicesAccount `xml:"PayloadContent>array>dict"`
	PayloadDisplayName string                       `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier  string                       `xml:"PayloadIdentifier,omitempty"`
	PayloadType        string                       `xml:"PayloadType,omitempty"`
	PayloadUUID        string                       `xml:"PayloadUUID,omitempty"`
	PayloadVersion     int                          `xml:"PayloadVersion,omitempty"`
}

// ExchangeWebServicesAccount represents an Exchange Web Services account configuration
type ExchangeWebServicesAccount struct {
	AllowMailDrop                 bool   `xml:"allowMailDrop,omitempty"`
	AuthenticationCertificateUUID string `xml:"AuthenticationCertificateUUID,omitempty"`
	EmailAddress                  string `xml:"EmailAddress,omitempty"`
	ExternalHost                  string `xml:"ExternalHost,omitempty"`
	ExternalPath                  string `xml:"ExternalPath,omitempty"`
	ExternalPort                  int    `xml:"ExternalPort,omitempty"`
	ExternalSSL                   bool   `xml:"ExternalSSL,omitempty"`
	Host                          string `xml:"Host,omitempty"`
	OAuth                         bool   `xml:"OAuth,omitempty"`
	OAuthSignInURL                string `xml:"OAuthSignInURL,omitempty"`
	Password                      string `xml:"Password,omitempty"`
	Path                          string `xml:"Path,omitempty"`
	PayloadCertificateUUID        string `xml:"PayloadCertificateUUID,omitempty"`
	Port                          int    `xml:"Port,omitempty"`
	SSL                           bool   `xml:"SSL,omitempty"`
	UserName                      string `xml:"UserName,omitempty"`
	PayloadDisplayName            string `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier             string `xml:"PayloadIdentifier,omitempty"`
	PayloadType                   string `xml:"PayloadType,omitempty"`
	PayloadUUID                   string `xml:"PayloadUUID,omitempty"`
	PayloadVersion                int    `xml:"PayloadVersion,omitempty"`
}
