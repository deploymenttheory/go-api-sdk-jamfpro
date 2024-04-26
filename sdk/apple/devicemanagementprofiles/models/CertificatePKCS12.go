package models

import "encoding/xml"

// ResourceCertificatePKCS12ConfigurationProfile represents the top-level structure of the plist for Certificate PKCS12 configurations
type ResourceCertificatePKCS12ConfigurationProfile struct {
	XMLName                  xml.Name                                            `xml:"plist"`
	Version                  string                                              `xml:"version,attr"`
	Dict                     CertificatePKCS12ConfigurationProfileSSubsetPayload `xml:"dict"`
	PayloadDescription       string                                              `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                              `xml:"PayloadDisplayName,omitempty"`
	PayloadEnabled           string                                              `xml:"PayloadEnabled,omitempty"`
	PayloadIdentifier        string                                              `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                              `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                              `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                              `xml:"PayloadScope,omitempty"`
	PayloadType              string                                              `xml:"PayloadType,omitempty"`
	PayloadUUID              string                                              `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                                              `xml:"PayloadVersion,omitempty"`
}

// CertificatePKCS12ConfigurationProfileSSubsetPayload represents the content structure for configuring Certificate PKCS12 settings
type CertificatePKCS12ConfigurationProfileSSubsetPayload struct {
	PayloadContent     []CertificatePKCS12 `xml:"PayloadContent>dict"`
	PayloadDisplayName string              `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier  string              `xml:"PayloadIdentifier,omitempty"`
	PayloadType        string              `xml:"PayloadType,omitempty"`
	PayloadUUID        string              `xml:"PayloadUUID,omitempty"`
	PayloadVersion     int                 `xml:"PayloadVersion,omitempty"`
}

// CertificatePKCS12 represents the Certificate PKCS12 dictionary within Certificate PKCS12 configuration
type CertificatePKCS12 struct {
	AllowAllAppsAccess         bool   `xml:"AllowAllAppsAccess,omitempty"`
	KeyIsExtractable           bool   `xml:"KeyIsExtractable,omitempty"`
	Password                   string `xml:"Password,omitempty"`
	PayloadCertificateFileName string `xml:"PayloadCertificateFileName,omitempty"`
	PayloadContent             string `xml:"PayloadContent,omitempty"`
	PayloadIdentifier          string `xml:"PayloadIdentifier,omitempty"`
	PayloadType                string `xml:"PayloadType,omitempty"`
	PayloadUUID                string `xml:"PayloadUUID,omitempty"`
	PayloadVersion             int    `xml:"PayloadVersion,omitempty"`
}
