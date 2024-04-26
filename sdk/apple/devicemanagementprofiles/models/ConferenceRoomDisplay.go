package models

import "encoding/xml"

// ResourceConferenceRoomDisplayConfigurationProfile represents the top-level structure of the plist for Conference Room Display mode configurations
type ResourceConferenceRoomDisplayConfigurationProfile struct {
	XMLName                  xml.Name                                               `xml:"plist"`
	Version                  string                                                 `xml:"version,attr"`
	Dict                     ConferenceRoomDisplayConfigurationProfileSubsetPayload `xml:"dict"`
	PayloadDescription       string                                                 `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                                 `xml:"PayloadDisplayName,omitempty"`
	PayloadEnabled           string                                                 `xml:"PayloadEnabled,omitempty"`
	PayloadIdentifier        string                                                 `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                                 `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                                 `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                                 `xml:"PayloadScope,omitempty"`
	PayloadType              string                                                 `xml:"PayloadType,omitempty"`
	PayloadUUID              string                                                 `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                                                 `xml:"PayloadVersion,omitempty"`
}

// ConferenceRoomDisplayConfigurationProfileSubset represents the content structure for configuring Conference Room Display settings
type ConferenceRoomDisplayConfigurationProfileSubsetPayload struct {
	Message           string `xml:"Message"`
	PayloadIdentifier string `xml:"PayloadIdentifier,omitempty"`
	PayloadType       string `xml:"PayloadType,omitempty"`
	PayloadUUID       string `xml:"PayloadUUID,omitempty"`
	PayloadVersion    int    `xml:"PayloadVersion,omitempty"`
}
