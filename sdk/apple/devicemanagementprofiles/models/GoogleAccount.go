package models

import "encoding/xml"

// ResourceGoogleAccountConfigurationProfile represents the top-level structure of the plist for Google account configurations
type ResourceGoogleAccountConfigurationProfile struct {
	XMLName                  xml.Name                                       `xml:"plist"`
	Version                  string                                         `xml:"version,attr"`
	Payload                  GoogleAccountConfigurationProfileSubsetPayload `xml:"dict"`
	PayloadDescription       string                                         `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                         `xml:"PayloadDisplayName,omitempty"`
	PayloadEnabled           string                                         `xml:"PayloadEnabled,omitempty"`
	PayloadIdentifier        string                                         `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                         `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                         `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                         `xml:"PayloadScope,omitempty"`
	PayloadType              string                                         `xml:"PayloadType,omitempty"`
	PayloadUUID              string                                         `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                                         `xml:"PayloadVersion,omitempty"`
}

// GoogleAccountConfigurationProfileSubsetPayload represents the content structure for configuring a Google account
type GoogleAccountConfigurationProfileSubsetPayload struct {
	AccountDescription        string                                 `xml:"AccountDescription,omitempty"`
	AccountName               string                                 `xml:"AccountName,omitempty"`
	EmailAddress              string                                 `xml:"EmailAddress,omitempty"`
	VPNUUID                   string                                 `xml:"VPNUUID,omitempty"`
	CommunicationServiceRules GoogleAccountCommunicationServiceRules `xml:"GoogleAccount.CommunicationServiceRules,omitempty"`
	PayloadIdentifier         string                                 `xml:"PayloadIdentifier,omitempty"`
	PayloadType               string                                 `xml:"PayloadType,omitempty"`
	PayloadUUID               string                                 `xml:"PayloadUUID,omitempty"`
	PayloadVersion            int                                    `xml:"PayloadVersion,omitempty"`
}

// GoogleAccountCommunicationServiceRules represents the communication service rules for a Google account
type GoogleAccountCommunicationServiceRules struct {
	DefaultServiceHandlers GoogleAccountDefaultServiceHandlers `xml:"DefaultServiceHandlers,omitempty"`
}

// GoogleAccountDefaultServiceHandlers represents the default service handlers for communication services
type GoogleAccountDefaultServiceHandlers struct {
	AudioCall string `xml:"AudioCall,omitempty"`
}
