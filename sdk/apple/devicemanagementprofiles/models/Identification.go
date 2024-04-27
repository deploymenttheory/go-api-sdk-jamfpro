package models

import "encoding/xml"

// ResourceIdentificationConfigurationProfile represents the top-level structure of the plist for configuring user identification.
type ResourceIdentificationConfigurationProfile struct {
	XMLName                  xml.Name                                 `xml:"plist"`
	Version                  string                                   `xml:"version,attr"`
	Payload                  IdentificationConfigurationProfileSubset `xml:"dict"`
	PayloadDescription       string                                   `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                   `xml:"PayloadDisplayName,omitempty"`
	PayloadEnabled           string                                   `xml:"PayloadEnabled,omitempty"`
	PayloadIdentifier        string                                   `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                   `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                   `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                   `xml:"PayloadScope,omitempty"`
	PayloadType              string                                   `xml:"PayloadType,omitempty"`
	PayloadUUID              string                                   `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                                   `xml:"PayloadVersion,omitempty"`
}

// IdentificationConfigurationProfileSubset represents the content structure for configuring user identification.
type IdentificationConfigurationProfileSubset struct {
	PayloadContent     []IdentificationPayloadContent `xml:"PayloadContent,omitempty"`
	PayloadDisplayName string                         `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier  string                         `xml:"PayloadIdentifier,omitempty"`
	PayloadType        string                         `xml:"PayloadType,omitempty"`
	PayloadUUID        string                         `xml:"PayloadUUID,omitempty"`
	PayloadVersion     int                            `xml:"PayloadVersion,omitempty"`
}

// IdentificationPayloadContent represents the content structure for user identification payload.
type IdentificationPayloadContent struct {
	PayloadIdentification IdentificationPayloadIdentification `xml:"PayloadIdentification,omitempty"`
	PayloadType           string                              `xml:"PayloadType,omitempty"`
	PayloadIdentifier     string                              `xml:"PayloadIdentifier,omitempty"`
	PayloadUUID           string                              `xml:"PayloadUUID,omitempty"`
	PayloadVersion        int                                 `xml:"PayloadVersion,omitempty"`
}

// IdentificationPayloadIdentification represents the content structure for user identification details.
type IdentificationPayloadIdentification struct {
	AuthMethod    string `xml:"AuthMethod,omitempty"`
	EmailAddress  string `xml:"EmailAddress,omitempty"`
	FullName      string `xml:"FullName,omitempty"`
	Password      string `xml:"Password,omitempty"`
	Prompt        string `xml:"Prompt,omitempty"`
	PromptMessage string `xml:"PromptMessage,omitempty"`
	UserName      string `xml:"UserName,omitempty"`
}
