package models

import "encoding/xml"

// FDEFileVaultOptionsConfigurationProfile represents the top-level structure of the plist for configuring FileVault options
type FDEFileVaultOptionsConfigurationProfile struct {
	XMLName                  xml.Name                                             `xml:"plist"`
	Version                  string                                               `xml:"version,attr"`
	Dict                     FDEFileVaultOptionsConfigurationProfileSubsetPayload `xml:"dict"`
	PayloadDescription       string                                               `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                               `xml:"PayloadDisplayName,omitempty"`
	PayloadEnabled           string                                               `xml:"PayloadEnabled,omitempty"`
	PayloadIdentifier        string                                               `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                               `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                               `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                               `xml:"PayloadScope,omitempty"`
	PayloadType              string                                               `xml:"PayloadType,omitempty"`
	PayloadUUID              string                                               `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                                               `xml:"PayloadVersion,omitempty"`
}

// FDEFileVaultOptionsConfigurationProfileSubsetPayload represents the content structure for configuring FileVault options
type FDEFileVaultOptionsConfigurationProfileSubsetPayload struct {
	PayloadContent     []FDEFileVaultOptions `xml:"PayloadContent>dict"`
	PayloadDisplayName string                `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier  string                `xml:"PayloadIdentifier,omitempty"`
	PayloadType        string                `xml:"PayloadType,omitempty"`
	PayloadUUID        string                `xml:"PayloadUUID,omitempty"`
	PayloadVersion     int                   `xml:"PayloadVersion,omitempty"`
}

// FDEFileVaultOptions represents the FileVault options dictionary within the FileVault options configuration
type FDEFileVaultOptions struct {
	DestroyFVKeyOnStandby bool   `xml:"DestroyFVKeyOnStandby"`
	DontAllowFDEDisable   bool   `xml:"dontAllowFDEDisable"`
	DontAllowFDEEnable    bool   `xml:"dontAllowFDEEnable"`
	PayloadDescription    string `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName    string `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier     string `xml:"PayloadIdentifier,omitempty"`
	PayloadType           string `xml:"PayloadType,omitempty"`
	PayloadUUID           string `xml:"PayloadUUID,omitempty"`
	PayloadVersion        int    `xml:"PayloadVersion,omitempty"`
}
