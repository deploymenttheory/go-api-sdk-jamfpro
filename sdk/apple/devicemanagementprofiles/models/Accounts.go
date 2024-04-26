package models

import "encoding/xml"

// ResourceAccountsConfigurationProfile represents the top-level structure of the plist
type ResourceAccountsConfigurationProfile struct {
	XMLName                  xml.Name                                  `xml:"plist"`
	Version                  string                                    `xml:"version,attr"`
	Dict                     AccountsConfigurationProfileSubsetPayload `xml:"dict"`
	PayloadDescription       string                                    `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                    `xml:"PayloadDisplayName,omitempty"`
	PayloadEnabled           string                                    `xml:"PayloadEnabled,omitempty"`
	PayloadIdentifier        string                                    `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                    `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                    `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                    `xml:"PayloadScope,omitempty"`
	PayloadType              string                                    `xml:"PayloadType,omitempty"`
	PayloadUUID              string                                    `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                                    `xml:"PayloadVersion,omitempty"`
}

// AccountsConfigurationProfileSubsetPayload represents the content structure for configuring guest accounts
type AccountsConfigurationProfileSubsetPayload struct {
	EnableGuestAccount bool   `xml:"EnableGuestAccount,omitempty"`
	PayloadIdentifier  string `xml:"PayloadIdentifier,omitempty"`
	PayloadType        string `xml:"PayloadType,omitempty"`
	PayloadUUID        string `xml:"PayloadUUID,omitempty"`
	PayloadVersion     int    `xml:"PayloadVersion,omitempty"`
}
