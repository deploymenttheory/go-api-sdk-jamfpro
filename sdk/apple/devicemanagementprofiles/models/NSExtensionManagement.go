package models

import "encoding/xml"

// ResourceNSExtensionManagementConfigurationProfile represents the top-level structure of the plist for NS Extension Management configurations
type ResourceNSExtensionManagementConfigurationProfile struct {
	XMLName                  xml.Name                                               `xml:"plist"`
	Version                  string                                                 `xml:"version,attr"`
	Dict                     NSExtensionManagementConfigurationProfileSubsetPayload `xml:"dict"`
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

// NSExtensionManagementConfigurationProfileSubsetPayload represents the content structure for configuring NS Extension Management settings
type NSExtensionManagementConfigurationProfileSubsetPayload struct {
	AllowedExtensions     []string `xml:"AllowedExtensions>string"`
	DeniedExtensionPoints []string `xml:"DeniedExtensionPoints>string"`
	DeniedExtensions      []string `xml:"DeniedExtensions>string"`
	PayloadIdentifier     string   `xml:"PayloadIdentifier,omitempty"`
	PayloadType           string   `xml:"PayloadType,omitempty"`
	PayloadUUID           string   `xml:"PayloadUUID,omitempty"`
	PayloadVersion        int      `xml:"PayloadVersion,omitempty"`
}
