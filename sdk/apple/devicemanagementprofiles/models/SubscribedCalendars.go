package models

import "encoding/xml"

// ResourceSubscribedCalendarsConfigurationProfile represents the top-level structure of the plist for subscribed calendar configurations
type ResourceSubscribedCalendarsConfigurationProfile struct {
	XMLName                  xml.Name                                             `xml:"plist"`
	Version                  string                                               `xml:"version,attr"`
	Dict                     SubscribedCalendarsConfigurationProfileSubsetPayload `xml:"dict"`
	PayloadDescription       string                                               `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                               `xml:"PayloadDisplayName,omitempty"`
	PayloadEnabled           string                                               `xml:"PayloadEnabled,omitempty"`
	PayloadIdentifier        string                                               `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                               `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                               `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                               `xml:"PayloadScope,omitempty"`
	PayloadType              string                                               `xml:"PayloadType,omitempty"`
	PayloadUUID              string                                               `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                                               `xml:"PayloadVersion,omitempty"`
}

// SubscribedCalendarsConfigurationProfileSubsetPayload represents the content structure for configuring subscribed calendars
type SubscribedCalendarsConfigurationProfileSubsetPayload struct {
	SubCalAccountDescription string `xml:"SubCalAccountDescription,omitempty"`
	SubCalAccountHostName    string `xml:"SubCalAccountHostName,omitempty"`
	SubCalAccountPassword    string `xml:"SubCalAccountPassword,omitempty"`
	SubCalAccountUsername    string `xml:"SubCalAccountUsername,omitempty"`
	SubCalAccountUseSSL      bool   `xml:"SubCalAccountUseSSL,omitempty"`
	VPNUUID                  string `xml:"VPNUUID,omitempty"`
	PayloadIdentifier        string `xml:"PayloadIdentifier,omitempty"`
	PayloadType              string `xml:"PayloadType,omitempty"`
	PayloadUUID              string `xml:"PayloadUUID,omitempty"`
	PayloadVersion           int    `xml:"PayloadVersion,omitempty"`
}
