package models

import "encoding/xml"

// ResourceCertificatePKCS1ConfigurationProfile represents the top-level structure of the plist for Certificate PKCS1 configurations
type ResourceCertificatePKCS1ConfigurationProfile struct {
	XMLName                  xml.Name                                          `xml:"plist"`
	Version                  string                                            `xml:"version,attr"`
	Payload                  CertificatePKCS1ConfigurationProfileSubsetPayload `xml:"dict"`
	PayloadDescription       string                                            `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                            `xml:"PayloadDisplayName,omitempty"`
	PayloadEnabled           string                                            `xml:"PayloadEnabled,omitempty"`
	PayloadIdentifier        string                                            `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                            `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                            `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                            `xml:"PayloadScope,omitempty"`
	PayloadType              string                                            `xml:"PayloadType,omitempty"`
	PayloadUUID              string                                            `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                                            `xml:"PayloadVersion,omitempty"`
}

// CertificatePKCS1ConfigurationProfileSubsetPayload represents the content structure for configuring Certificate PKCS1 settings
type CertificatePKCS1ConfigurationProfileSubsetPayload struct {
	PayloadContent     []CertificatePKCS1 `xml:"PayloadContent>dict"`
	PayloadDisplayName string             `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier  string             `xml:"PayloadIdentifier,omitempty"`
	PayloadType        string             `xml:"PayloadType,omitempty"`
	PayloadUUID        string             `xml:"PayloadUUID,omitempty"`
	PayloadVersion     int                `xml:"PayloadVersion,omitempty"`
}

// CertificatePKCS1 represents the Certificate PKCS1 dictionary within Certificate PKCS1 configuration
type CertificatePKCS1 struct {
	PayloadCertificateFileName string `xml:"PayloadCertificateFileName,omitempty"`
	PayloadContent             string `xml:"PayloadContent,omitempty"`
	PayloadIdentifier          string `xml:"PayloadIdentifier,omitempty"`
	PayloadType                string `xml:"PayloadType,omitempty"`
	PayloadUUID                string `xml:"PayloadUUID,omitempty"`
	PayloadVersion             int    `xml:"PayloadVersion,omitempty"`
	Password                   string `xml:"Password,omitempty"`
}
