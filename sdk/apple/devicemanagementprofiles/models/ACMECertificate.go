package models

import "encoding/xml"

// ResourceACMECertificateProfile represents the top-level structure of the plist for ACME Certificate configurations
type ResourceACMECertificateProfile struct {
	XMLName                  xml.Name                            `xml:"plist"`
	Version                  string                              `xml:"version,attr"`
	Dict                     ACMECertificateProfileSubsetPayload `xml:"dict"`
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

// ACMECertificateProfileSubsetPayload represents the content structure for configuring ACME Certificate settings
type ACMECertificateProfileSubsetPayload struct {
	PayloadContent     []ACMECertificate `xml:"PayloadContent>dict"`
	PayloadDisplayName string            `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier  string            `xml:"PayloadIdentifier,omitempty"`
	PayloadType        string            `xml:"PayloadType,omitempty"`
	PayloadUUID        string            `xml:"PayloadUUID,omitempty"`
	PayloadVersion     int               `xml:"PayloadVersion,omitempty"`
}

// ACMECertificate represents the ACME Certificate dictionary within ACME Certificate configuration
type ACMECertificate struct {
	ClientIdentifier  string                        `xml:"ClientIdentifier,omitempty"`
	ExtendedKeyUsage  []string                      `xml:"ExtendedKeyUsage>string,omitempty"`
	HardwareBound     bool                          `xml:"HardwareBound,omitempty"`
	KeySize           int                           `xml:"KeySize,omitempty"`
	KeyType           string                        `xml:"KeyType,omitempty"`
	UsageFlags        int                           `xml:"UsageFlags,omitempty"`
	PayloadIdentifier string                        `xml:"PayloadIdentifier,omitempty"`
	PayloadType       string                        `xml:"PayloadType,omitempty"`
	PayloadUUID       string                        `xml:"PayloadUUID,omitempty"`
	PayloadVersion    int                           `xml:"PayloadVersion,omitempty"`
	Subject           [][][2]string                 `xml:"Subject>array>array>string,omitempty"`
	SubjectAltName    ACMECertificateSubjectAltName `xml:"SubjectAltName,omitempty"`
	DirectoryURL      string                        `xml:"DirectoryURL,omitempty"`
}

// ACMECertificateSubjectAltName represents the Subject Alt Name details within ACME Certificate configuration
type ACMECertificateSubjectAltName struct {
	DNSName                   string `xml:"dNSName,omitempty"`
	NTPrincipalName           string `xml:"ntPrincipalName,omitempty"`
	RFC822Name                string `xml:"rfc822Name,omitempty"`
	UniformResourceIdentifier string `xml:"uniformResourceIdentifier,omitempty"`
}
