package models

import "encoding/xml"

// ResourceCertificatePreferenceConfgurationProfile represents the top-level structure of the plist for Certificate Preference configurations
type ResourceCertificatePreferenceConfgurationProfile struct {
	XMLName                  xml.Name                                                `xml:"plist"`
	Version                  string                                                  `xml:"version,attr"`
	Payload                  CertificatePreferenceConfigurationProfileSSubsetPayload `xml:"dict"`
	PayloadDescription       string                                                  `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                                  `xml:"PayloadDisplayName,omitempty"`
	PayloadEnabled           string                                                  `xml:"PayloadEnabled,omitempty"`
	PayloadIdentifier        string                                                  `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                                  `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                                  `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                                  `xml:"PayloadScope,omitempty"`
	PayloadType              string                                                  `xml:"PayloadType,omitempty"`
	PayloadUUID              string                                                  `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                                                  `xml:"PayloadVersion,omitempty"`
}

// CertificatePreferenceConfigurationProfileSSubsetPayload represents the content structure for configuring Certificate Preference settings
type CertificatePreferenceConfigurationProfileSSubsetPayload struct {
	PayloadContent     []CertificatePreference `xml:"PayloadContent>dict"`
	PayloadDisplayName string                  `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier  string                  `xml:"PayloadIdentifier,omitempty"`
	PayloadType        string                  `xml:"PayloadType,omitempty"`
	PayloadUUID        string                  `xml:"PayloadUUID,omitempty"`
	PayloadVersion     int                     `xml:"PayloadVersion,omitempty"`
}

// CertificatePreference represents the Certificate Preference dictionary within Certificate Preference configuration
type CertificatePreference struct {
	Name                   string `xml:"Name,omitempty"`
	PayloadCertificateUUID string `xml:"PayloadCertificateUUID,omitempty"`
	PayloadContent         string `xml:"PayloadContent,omitempty"`
	PayloadIdentifier      string `xml:"PayloadIdentifier,omitempty"`
	PayloadType            string `xml:"PayloadType,omitempty"`
	PayloadUUID            string `xml:"PayloadUUID,omitempty"`
	PayloadVersion         int    `xml:"PayloadVersion,omitempty"`
}
