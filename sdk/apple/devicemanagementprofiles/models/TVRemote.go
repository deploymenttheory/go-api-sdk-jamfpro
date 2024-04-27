package models

import "encoding/xml"

// ResourceTVRemoteConfigurationProfile represents the top-level structure of the plist for TV Remote configurations
type ResourceTVRemoteConfigurationProfile struct {
	XMLName                  xml.Name                                  `xml:"plist"`
	Version                  string                                    `xml:"version,attr"`
	Payload                  TVRemoteConfigurationProfileSubsetPayload `xml:"dict"`
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

// TVRemoteConfigurationProfileSubsetPayload represents the content structure for configuring TV Remote settings
type TVRemoteConfigurationProfileSubsetPayload struct {
	AllowedRemotes    []TVRemoteAllowedRemotesItem `xml:"AllowedRemotes>TVRemoteAllowedRemotesItem"`
	AllowedTVs        []TVRemoteAllowedTVsItem     `xml:"AllowedTVs>TVRemoteAllowedTVsItem"`
	PayloadIdentifier string                       `xml:"PayloadIdentifier,omitempty"`
	PayloadType       string                       `xml:"PayloadType,omitempty"`
	PayloadUUID       string                       `xml:"PayloadUUID,omitempty"`
	PayloadVersion    int                          `xml:"PayloadVersion,omitempty"`
}

// TVRemoteAllowedRemotesItem represents a device that Apple TV can connect to
type TVRemoteAllowedRemotesItem struct {
	RemoteDeviceID string `xml:"RemoteDeviceID"`
}

// TVRemoteAllowedTVsItem represents an Apple TV device that the remote can connect to
type TVRemoteAllowedTVsItem struct {
	TVDeviceID   string `xml:"TVDeviceID"`
	TVDeviceName string `xml:"TVDeviceName,omitempty"`
}
