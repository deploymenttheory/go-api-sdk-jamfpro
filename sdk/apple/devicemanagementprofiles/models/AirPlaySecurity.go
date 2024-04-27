package models

import "encoding/xml"

// ResourceAirPlaySecurityConfigurationProfile represents the top-level structure of the plist for AirPlay Security configurations
type ResourceAirPlaySecurityConfigurationProfile struct {
	XMLName                  xml.Name                                         `xml:"plist"`
	Version                  string                                           `xml:"version,attr"`
	Payload                  AirPlaySecurityConfigurationProfileSubsetPayload `xml:"dict>array>dict"`
	PayloadDescription       string                                           `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                           `xml:"PayloadDisplayName,omitempty"`
	PayloadEnabled           string                                           `xml:"PayloadEnabled,omitempty"`
	PayloadIdentifier        string                                           `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                           `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                           `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                           `xml:"PayloadScope,omitempty"`
	PayloadType              string                                           `xml:"PayloadType,omitempty"`
	PayloadUUID              string                                           `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                                           `xml:"PayloadVersion,omitempty"`
}

// AirPlaySecurityConfigurationProfileSubsetPayload represents the content structure for configuring AirPlay security settings
type AirPlaySecurityConfigurationProfileSubsetPayload struct {
	AccessType        string `xml:"AccessType,omitempty"`
	Password          string `xml:"Password,omitempty"`
	SecurityType      string `xml:"SecurityType,omitempty"`
	PayloadIdentifier string `xml:"PayloadIdentifier,omitempty"`
	PayloadType       string `xml:"PayloadType,omitempty"`
	PayloadUUID       string `xml:"PayloadUUID,omitempty"`
	PayloadVersion    int    `xml:"PayloadVersion,omitempty"`
}
