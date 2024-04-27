package models

import "encoding/xml"

// ResourceCertificatePEMProfile represents the top-level structure of the plist for Certificate PEM configurations
type ResourceCertificatePEMProfile struct {
	XMLName                  xml.Name                           `xml:"plist"`
	Version                  string                             `xml:"version,attr"`
	Payload                  CertificatePEMProfileSubsetPayload `xml:"dict>array>dict"`
	PayloadDescription       string                             `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                             `xml:"PayloadDisplayName,omitempty"`
	PayloadEnabled           string                             `xml:"PayloadEnabled,omitempty"`
	PayloadIdentifier        string                             `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                             `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                             `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                             `xml:"PayloadScope,omitempty"`
	PayloadType              string                             `xml:"PayloadType,omitempty"`
	PayloadUUID              string                             `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                             `xml:"PayloadVersion,omitempty"`
}

// CertificatePEMProfileSubsetPayload represents the content structure for configuring Certificate PEM settings
type CertificatePEMProfileSubsetPayload struct {
	PayloadContent     []CertificatePEM `xml:"PayloadContent>dict"`
	PayloadDisplayName string           `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier  string           `xml:"PayloadIdentifier,omitempty"`
	PayloadType        string           `xml:"PayloadType,omitempty"`
	PayloadUUID        string           `xml:"PayloadUUID,omitempty"`
	PayloadVersion     int              `xml:"PayloadVersion,omitempty"`
}

// CertificatePEM represents the Certificate PEM dictionary within Certificate PEM configuration
type CertificatePEM struct {
	PayloadCertificateFileName string `xml:"PayloadCertificateFileName,omitempty"`
	PayloadContent             string `xml:"PayloadContent,omitempty"`
	PayloadIdentifier          string `xml:"PayloadIdentifier,omitempty"`
	PayloadType                string `xml:"PayloadType,omitempty"`
	PayloadUUID                string `xml:"PayloadUUID,omitempty"`
	PayloadVersion             int    `xml:"PayloadVersion,omitempty"`
}
