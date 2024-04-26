package models

import "encoding/xml"

// ResourceCertificateRootProfile represents the top-level structure of the plist for Certificate Root configurations
type ResourceCertificateRootProfile struct {
	XMLName                  xml.Name                            `xml:"plist"`
	Version                  string                              `xml:"version,attr"`
	Dict                     CertificateRootProfileSubsetPayload `xml:"dict"`
	PayloadDescription       string                              `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                              `xml:"PayloadDisplayName,omitempty"`
	PayloadEnabled           string                              `xml:"PayloadEnabled,omitempty"`
	PayloadIdentifier        string                              `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                              `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                              `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                              `xml:"PayloadScope,omitempty"`
	PayloadType              string                              `xml:"PayloadType,omitempty"`
	PayloadUUID              string                              `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                              `xml:"PayloadVersion,omitempty"`
}

// CertificateRootProfileSubsetPayload represents the content structure for configuring Certificate Root settings
type CertificateRootProfileSubsetPayload struct {
	PayloadContent     []CertificateRoot `xml:"PayloadContent>dict"`
	PayloadDisplayName string            `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier  string            `xml:"PayloadIdentifier,omitempty"`
	PayloadType        string            `xml:"PayloadType,omitempty"`
	PayloadUUID        string            `xml:"PayloadUUID,omitempty"`
	PayloadVersion     int               `xml:"PayloadVersion,omitempty"`
}

// CertificateRoot represents the Certificate Root dictionary within Certificate Root configuration
type CertificateRoot struct {
	PayloadCertificateFileName string `xml:"PayloadCertificateFileName,omitempty"`
	PayloadContent             string `xml:"PayloadContent,omitempty"`
	PayloadDescription         string `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName         string `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier          string `xml:"PayloadIdentifier,omitempty"`
	PayloadType                string `xml:"PayloadType,omitempty"`
	PayloadUUID                string `xml:"PayloadUUID,omitempty"`
	PayloadVersion             int    `xml:"PayloadVersion,omitempty"`
}
