package models

import "encoding/xml"

// ResourceCertificateRevocationConfigurationProfile represents the top-level structure of the plist for Certificate Revocation configurations
type ResourceCertificateRevocationConfigurationProfile struct {
	XMLName                  xml.Name                                               `xml:"plist"`
	Version                  string                                                 `xml:"version,attr"`
	Payload                  CertificateRevocationConfigurationProfileSubsetPayload `xml:"dict>array>dict"`
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

// CertificateRevocationConfigurationProfileSubsetPayload represents the content structure for configuring Certificate Revocation settings
type CertificateRevocationConfigurationProfileSubsetPayload struct {
	PayloadContent     []CertificateRevocation `xml:"PayloadContent>dict"`
	PayloadDisplayName string                  `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier  string                  `xml:"PayloadIdentifier,omitempty"`
	PayloadType        string                  `xml:"PayloadType,omitempty"`
	PayloadUUID        string                  `xml:"PayloadUUID,omitempty"`
	PayloadVersion     int                     `xml:"PayloadVersion,omitempty"`
}

// CertificateRevocation represents the Certificate Revocation dictionary within Certificate Revocation configuration
type CertificateRevocation struct {
	EnabledForCerts    []CertificateRevocationSubjectPublicKeyInfoHashDict `xml:"EnabledForCerts"`
	PayloadDescription string                                              `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName string                                              `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier  string                                              `xml:"PayloadIdentifier,omitempty"`
	PayloadType        string                                              `xml:"PayloadType,omitempty"`
	PayloadUUID        string                                              `xml:"PayloadUUID,omitempty"`
	PayloadVersion     int                                                 `xml:"PayloadVersion,omitempty"`
}

// CertificateRevocationSubjectPublicKeyInfoHashDict represents the SubjectPublicKeyInfoHashDict dictionary within Certificate Revocation configuration
type CertificateRevocationSubjectPublicKeyInfoHashDict struct {
	Algorithm string `xml:"Algorithm"`
	Hash      string `xml:"Hash"`
}
