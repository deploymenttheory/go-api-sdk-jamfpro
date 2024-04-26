package models

import "encoding/xml"

// ResourceCertificateTransparencyConfigurationProfile represents the top-level structure of the plist for Certificate Transparency configurations
type ResourceCertificateTransparencyConfigurationProfile struct {
	XMLName                  xml.Name                                                 `xml:"plist"`
	Version                  string                                                   `xml:"version,attr"`
	Dict                     CertificateTransparencyConfigurationProfileSubsetPayload `xml:"dict"`
	PayloadDescription       string                                                   `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                                   `xml:"PayloadDisplayName,omitempty"`
	PayloadEnabled           string                                                   `xml:"PayloadEnabled,omitempty"`
	PayloadIdentifier        string                                                   `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                                   `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                                   `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                                   `xml:"PayloadScope,omitempty"`
	PayloadType              string                                                   `xml:"PayloadType,omitempty"`
	PayloadUUID              string                                                   `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                                                   `xml:"PayloadVersion,omitempty"`
}

// CertificateTransparencyConfigurationProfileSubsetPayload represents the content structure for configuring Certificate Transparency settings
type CertificateTransparencyConfigurationProfileSubsetPayload struct {
	PayloadContent     []CertificateTransparency `xml:"PayloadContent>dict"`
	PayloadDisplayName string                    `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier  string                    `xml:"PayloadIdentifier,omitempty"`
	PayloadType        string                    `xml:"PayloadType,omitempty"`
	PayloadUUID        string                    `xml:"PayloadUUID,omitempty"`
	PayloadVersion     int                       `xml:"PayloadVersion,omitempty"`
}

// CertificateTransparency represents the Certificate Transparency dictionary within Certificate Transparency configuration
type CertificateTransparency struct {
	DisabledForCerts   []CertificateTransparencySubjectPublicKeyInfoHashDict `xml:"DisabledForCerts"`
	DisabledForDomains []string                                              `xml:"DisabledForDomains>string"`
	PayloadDescription string                                                `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName string                                                `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier  string                                                `xml:"PayloadIdentifier,omitempty"`
	PayloadType        string                                                `xml:"PayloadType,omitempty"`
	PayloadUUID        string                                                `xml:"PayloadUUID,omitempty"`
	PayloadVersion     int                                                   `xml:"PayloadVersion,omitempty"`
}

// CertificateTransparencySubjectPublicKeyInfoHashDict represents the SubjectPublicKeyInfoHashDict dictionary within Certificate Transparency configuration
type CertificateTransparencySubjectPublicKeyInfoHashDict struct {
	Algorithm string `xml:"Algorithm"`
	Hash      string `xml:"Hash"`
}
