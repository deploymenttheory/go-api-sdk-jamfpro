package models

import "encoding/xml"

// ResourceIdentityPreferenceConfigurationProfile represents the top-level structure of the plist for configuring user identity preference.
type ResourceIdentityPreferenceConfigurationProfile struct {
	XMLName                  xml.Name                                     `xml:"plist"`
	Version                  string                                       `xml:"version,attr"`
	Dict                     IdentityPreferenceConfigurationProfileSubset `xml:"dict"`
	PayloadDescription       string                                       `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                       `xml:"PayloadDisplayName,omitempty"`
	PayloadEnabled           string                                       `xml:"PayloadEnabled,omitempty"`
	PayloadIdentifier        string                                       `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                       `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                       `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                       `xml:"PayloadScope,omitempty"`
	PayloadType              string                                       `xml:"PayloadType,omitempty"`
	PayloadUUID              string                                       `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                                       `xml:"PayloadVersion,omitempty"`
}

// IdentityPreferenceConfigurationProfileSubset represents the content structure for configuring user identity preference.
type IdentityPreferenceConfigurationProfileSubset struct {
	PayloadContent     []IdentityPreferencePayloadContent `xml:"PayloadContent,omitempty"`
	PayloadDescription string                             `xml:"PayloadDescription,omitempty"`
	PayloadScope       string                             `xml:"PayloadScope,omitempty"`
	PayloadType        string                             `xml:"PayloadType,omitempty"`
	PayloadUUID        string                             `xml:"PayloadUUID,omitempty"`
	PayloadVersion     int                                `xml:"PayloadVersion,omitempty"`
}

// IdentityPreferencePayloadContent represents the content structure for user identity preference payload.
type IdentityPreferencePayloadContent struct {
	Name                   string `xml:"Name,omitempty"`
	PayloadCertificateUUID string `xml:"PayloadCertificateUUID,omitempty"`
	PayloadIdentifier      string `xml:"PayloadIdentifier,omitempty"`
	PayloadType            string `xml:"PayloadType,omitempty"`
	PayloadUUID            string `xml:"PayloadUUID,omitempty"`
	PayloadVersion         int    `xml:"PayloadVersion,omitempty"`
}
