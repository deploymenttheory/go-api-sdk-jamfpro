package models

import "encoding/xml"

// ResourceActiveDirectoryCertificateProfile represents the top-level structure of the plist for Active Directory Certificate configurations
type ResourceActiveDirectoryCertificateProfile struct {
	XMLName                  xml.Name                                       `xml:"plist"`
	Version                  string                                         `xml:"version,attr"`
	Payload                  ActiveDirectoryCertificateProfileSubsetPayload `xml:"dict"`
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

// ActiveDirectoryCertificateProfileSubsetPayload represents the content structure for configuring Active Directory Certificate settings
type ActiveDirectoryCertificateProfileSubsetPayload struct {
	PayloadContent     []ActiveDirectoryCertificate `xml:"PayloadContent>dict"`
	PayloadDisplayName string                       `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier  string                       `xml:"PayloadIdentifier,omitempty"`
	PayloadType        string                       `xml:"PayloadType,omitempty"`
	PayloadUUID        string                       `xml:"PayloadUUID,omitempty"`
	PayloadVersion     int                          `xml:"PayloadVersion,omitempty"`
}

// ActiveDirectoryCertificate represents the Active Directory Certificate dictionary within Active Directory Certificate configuration
type ActiveDirectoryCertificate struct {
	CertServer                      string `xml:"CertServer,omitempty"`
	CertTemplate                    string `xml:"CertTemplate,omitempty"`
	CertificateAcquisitionMechanism string `xml:"CertificateAcquisitionMechanism,omitempty"`
	CertificateAuthority            string `xml:"CertificateAuthority,omitempty"`
	CertificateRenewalTimeInterval  int    `xml:"CertificateRenewalTimeInterval,omitempty"`
	Description                     string `xml:"Description,omitempty"`
	EnableAutoRenewal               bool   `xml:"EnableAutoRenewal,omitempty"`
	KeyIsExtractable                bool   `xml:"KeyIsExtractable,omitempty"`
	Keysize                         int    `xml:"Keysize,omitempty"`
	PromptForCredentials            bool   `xml:"PromptForCredentials,omitempty"`
	PayloadIdentifier               string `xml:"PayloadIdentifier,omitempty"`
	PayloadType                     string `xml:"PayloadType,omitempty"`
	PayloadUUID                     string `xml:"PayloadUUID,omitempty"`
	PayloadVersion                  int    `xml:"PayloadVersion,omitempty"`
}
