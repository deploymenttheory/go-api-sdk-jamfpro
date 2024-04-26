package models

import "encoding/xml"

// LightsOutManagementLOM represents the structure of the plist for LightsOutManagementLOM settings
type LightsOutManagementLOM struct {
	XMLName                  xml.Name                                                  `xml:"plist"`
	Version                  string                                                    `xml:"version,attr"`
	PayloadContent           []LightsOutManagementLOMConfigurationProfileSubsetPayload `xml:"PayloadContent"`
	PayloadDescription       string                                                    `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                                    `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier        string                                                    `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                                    `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                                    `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                                    `xml:"PayloadScope,omitempty"`
	PayloadType              string                                                    `xml:"PayloadType,omitempty"`
	PayloadUUID              string                                                    `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                                                    `xml:"PayloadVersion,omitempty"`
}

// LightsOutManagementLOMConfigurationProfileSubsetPayload represents the content structure for LightsOutManagementLOM payload
type LightsOutManagementLOMConfigurationProfileSubsetPayload struct {
	ControllerCACertificateUUIDs []string `xml:"ControllerCACertificateUUIDs>string,omitempty"`
	ControllerCertificateUUID    string   `xml:"ControllerCertificateUUID,omitempty"`
	DeviceCACertificateUUIDs     []string `xml:"DeviceCACertificateUUIDs>string,omitempty"`
	DeviceCertificateUUID        string   `xml:"DeviceCertificateUUID,omitempty"`
	PayloadDisplayName           string   `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier            string   `xml:"PayloadIdentifier,omitempty"`
	PayloadType                  string   `xml:"PayloadType,omitempty"`
	PayloadUUID                  string   `xml:"PayloadUUID,omitempty"`
	PayloadVersion               int      `xml:"PayloadVersion,omitempty"`
	Password                     string   `xml:"Password,omitempty"`
	PayloadCertificateFileName   string   `xml:"PayloadCertificateFileName,omitempty"`
	PayloadContent               string   `xml:"PayloadContent,omitempty"`
}
