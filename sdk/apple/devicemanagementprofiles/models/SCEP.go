package models

import "encoding/xml"

// ResourceSCEPConfigurationProfile represents the top-level structure of the plist for SCEP configurations
type ResourceSCEPConfigurationProfile struct {
	XMLName                  xml.Name                              `xml:"plist"`
	Version                  string                                `xml:"version,attr"`
	Dict                     SCEPConfigurationProfileSubsetPayload `xml:"dict"`
	PayloadDescription       string                                `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                `xml:"PayloadDisplayName,omitempty"`
	PayloadEnabled           string                                `xml:"PayloadEnabled,omitempty"`
	PayloadIdentifier        string                                `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                `xml:"PayloadScope,omitempty"`
	PayloadType              string                                `xml:"PayloadType,omitempty"`
	PayloadUUID              string                                `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                                `xml:"PayloadVersion,omitempty"`
}

// SCEPConfigurationProfileSubsetPayload represents the content structure for configuring SCEP settings
type SCEPConfigurationProfileSubsetPayload struct {
	PayloadContent     []SCEP `xml:"PayloadContent>dict"`
	PayloadDisplayName string `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier  string `xml:"PayloadIdentifier,omitempty"`
	PayloadType        string `xml:"PayloadType,omitempty"`
	PayloadUUID        string `xml:"PayloadUUID,omitempty"`
	PayloadVersion     int    `xml:"PayloadVersion,omitempty"`
}

// SCEP represents the SCEP dictionary within SCEP configuration
type SCEP struct {
	AllowAllAppsAccess bool               `xml:"AllowAllAppsAccess"`
	CAFingerprint      string             `xml:"CAFingerprint"`
	Challenge          string             `xml:"Challenge"`
	KeyType            string             `xml:"KeyType"`
	KeyUsage           int                `xml:"KeyUsage"`
	KeyIsExtractable   bool               `xml:"KeyIsExtractable"`
	Keysize            int                `xml:"Keysize"`
	Name               string             `xml:"Name"`
	Retries            int                `xml:"Retries"`
	RetryDelay         int                `xml:"RetryDelay"`
	Subject            [][][]string       `xml:"Subject>array>array"`
	SubjectAltName     SCEPSubjectAltName `xml:"SubjectAltName"`
	URL                string             `xml:"URL"`
	PayloadDescription string             `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName string             `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier  string             `xml:"PayloadIdentifier,omitempty"`
	PayloadType        string             `xml:"PayloadType,omitempty"`
	PayloadUUID        string             `xml:"PayloadUUID,omitempty"`
	PayloadVersion     int                `xml:"PayloadVersion,omitempty"`
}

// SCEPSubjectAltName represents the SubjectAltName dictionary within SCEP configuration
type SCEPSubjectAltName struct {
	DNSName                   string `xml:"dNSName"`
	NTPrincipalName           string `xml:"ntPrincipalName"`
	RFC822Name                string `xml:"rfc822Name"`
	UniformResourceIdentifier string `xml:"uniformResourceIdentifier"`
}
