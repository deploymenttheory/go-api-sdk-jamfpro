package models

import "encoding/xml"

// ResourceAirPlayConfigurationProfile represents the top-level structure of the plist for AirPlay configurations
type ResourceAirPlayConfigurationProfile struct {
	XMLName                  xml.Name                                 `xml:"plist"`
	Version                  string                                   `xml:"version,attr"`
	Payload                  AirPlayConfigurationProfileSubsetPayload `xml:"dict>array>dict"`
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

// AirPlayConfigurationProfileSubsetPayload represents the content structure for configuring AirPlay settings
type AirPlayConfigurationProfileSubsetPayload struct {
	AllowList         []AirPlayAllowListItem `xml:"AllowList>AirPlayAllowListItem,omitempty"`
	Passwords         []AirPlayPasswordsItem `xml:"Passwords>AirPlayPasswordsItem,omitempty"`
	PayloadIdentifier string                 `xml:"PayloadIdentifier,omitempty"`
	PayloadType       string                 `xml:"PayloadType,omitempty"`
	PayloadUUID       string                 `xml:"PayloadUUID,omitempty"`
	PayloadVersion    int                    `xml:"PayloadVersion,omitempty"`
}

// AirPlayAllowListItem represents an item in the list of allowed AirPlay destinations
type AirPlayAllowListItem struct {
	DeviceID string `xml:"DeviceID,omitempty"`
}

// AirPlayPasswordsItem represents an item defining a password for an AirPlay destination
type AirPlayPasswordsItem struct {
	DeviceID   string `xml:"DeviceID,omitempty"`
	DeviceName string `xml:"DeviceName,omitempty"`
	Password   string `xml:"Password,omitempty"`
}
