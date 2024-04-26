package models

import "encoding/xml"

// ResourceCalDAVConfigurationProfile represents the top-level structure of the plist
type ResourceCalDAVConfigurationProfile struct {
	XMLName                  xml.Name                                `xml:"plist"`
	Version                  string                                  `xml:"version,attr"`
	Dict                     CalDAVConfigurationProfileSubsetPayload `xml:"dict"`
	PayloadDescription       string                                  `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                  `xml:"PayloadDisplayName,omitempty"`
	PayloadEnabled           string                                  `xml:"PayloadEnabled,omitempty"`
	PayloadIdentifier        string                                  `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                  `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                  `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                  `xml:"PayloadScope,omitempty"`
	PayloadType              string                                  `xml:"PayloadType,omitempty"`
	PayloadUUID              string                                  `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                                  `xml:"PayloadVersion,omitempty"`
}

// CalDAVConfigurationProfileSubsetPayload represents the content structure for configuring a CalDAV account
type CalDAVConfigurationProfileSubsetPayload struct {
	CalDAVAccountDescription string `xml:"CalDAVAccountDescription,omitempty"`
	CalDAVHostName           string `xml:"CalDAVHostName"`
	CalDAVPassword           string `xml:"CalDAVPassword,omitempty"`
	CalDAVPort               int    `xml:"CalDAVPort,omitempty"`
	CalDAVPrincipalURL       string `xml:"CalDAVPrincipalURL,omitempty"`
	CalDAVUsername           string `xml:"CalDAVUsername,omitempty"`
	CalDAVUseSSL             bool   `xml:"CalDAVUseSSL,omitempty"`
	VPNUUID                  string `xml:"VPNUUID,omitempty"` // Available in iOS 14 and later
	PayloadIdentifier        string `xml:"PayloadIdentifier,omitempty"`
	PayloadType              string `xml:"PayloadType,omitempty"`
	PayloadUUID              string `xml:"PayloadUUID,omitempty"`
	PayloadVersion           int    `xml:"PayloadVersion,omitempty"`
}
