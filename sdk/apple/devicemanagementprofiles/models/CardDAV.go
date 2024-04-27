package models

import "encoding/xml"

// ResourceCardDAVConfigurationProfile represents the top-level structure of the plist
type ResourceCardDAVConfigurationProfile struct {
	XMLName                  xml.Name                                 `xml:"plist"`
	Version                  string                                   `xml:"version,attr"`
	Payload                  CardDAVConfigurationProfileSubsetPayload `xml:"dict"`
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

// CardDAVConfigurationProfileSubsetPayload represents the content structure for configuring a CardDAV account
type CardDAVConfigurationProfileSubsetPayload struct {
	CardDAVAccountDescription string                           `xml:"CardDAVAccountDescription,omitempty"`
	CardDAVHostName           string                           `xml:"CardDAVHostName,omitempty"`
	CardDAVPassword           string                           `xml:"CardDAVPassword,omitempty"`
	CardDAVPort               int                              `xml:"CardDAVPort,omitempty"`
	CardDAVPrincipalURL       string                           `xml:"CardDAVPrincipalURL,omitempty"`
	CardDAVUsername           string                           `xml:"CardDAVUsername,omitempty"`
	CardDAVUseSSL             bool                             `xml:"CardDAVUseSSL,omitempty"`
	CommunicationServiceRules CardDAVCommunicationServiceRules `xml:"CardDAV.CommunicationServiceRules,omitempty"`
	VPNUUID                   string                           `xml:"VPNUUID,omitempty"`
	PayloadIdentifier         string                           `xml:"PayloadIdentifier,omitempty"`
	PayloadType               string                           `xml:"PayloadType,omitempty"`
	PayloadUUID               string                           `xml:"PayloadUUID,omitempty"`
	PayloadVersion            int                              `xml:"PayloadVersion,omitempty"`
}

// CardDAVCommunicationServiceRules represents the set of rules and default handlers for communication services associated with a CardDAV account
type CardDAVCommunicationServiceRules struct {
	DefaultServiceHandlers CardDAVDefaultServiceHandlers `xml:"DefaultServiceHandlers,omitempty"`
}

// CardDAVDefaultServiceHandlers represents the default service handlers for communication services
type CardDAVDefaultServiceHandlers struct {
	AudioCall string `xml:"AudioCall,omitempty"`
}
