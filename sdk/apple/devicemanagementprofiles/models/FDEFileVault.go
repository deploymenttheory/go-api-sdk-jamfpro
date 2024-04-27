package models

import "encoding/xml"

// FDEFileVaultConfigurationProfile represents the top-level structure of the plist for configuring FileVault
type FDEFileVaultConfigurationProfile struct {
	XMLName                  xml.Name                                      `xml:"plist"`
	Version                  string                                        `xml:"version,attr"`
	Payload                  FDEFileVaultConfigurationProfileSubsetPayload `xml:"dict>array>dict"`
	PayloadDescription       string                                        `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                        `xml:"PayloadDisplayName,omitempty"`
	PayloadEnabled           string                                        `xml:"PayloadEnabled,omitempty"`
	PayloadIdentifier        string                                        `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                        `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                        `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                        `xml:"PayloadScope,omitempty"`
	PayloadType              string                                        `xml:"PayloadType,omitempty"`
	PayloadUUID              string                                        `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                                        `xml:"PayloadVersion,omitempty"`
}

// FDEFileVaultConfigurationProfileSubsetPayload represents the content structure for configuring FileVault
type FDEFileVaultConfigurationProfileSubsetPayload struct {
	PayloadContent     []FDEFileVault `xml:"PayloadContent>dict"`
	PayloadDisplayName string         `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier  string         `xml:"PayloadIdentifier,omitempty"`
	PayloadType        string         `xml:"PayloadType,omitempty"`
	PayloadUUID        string         `xml:"PayloadUUID,omitempty"`
	PayloadVersion     int            `xml:"PayloadVersion,omitempty"`
}

// FDEFileVault represents the FileVault dictionary within the FileVault configuration
type FDEFileVault struct {
	Defer                                  bool   `xml:"Defer"`
	DeferDontAskAtUserLogout               bool   `xml:"DeferDontAskAtUserLogout"`
	DeferForceAtUserLoginMaxBypassAttempts int    `xml:"DeferForceAtUserLoginMaxBypassAttempts"`
	Enable                                 string `xml:"Enable"`
	ForceEnableInSetupAssistant            bool   `xml:"ForceEnableInSetupAssistant"`
	OutputPath                             string `xml:"OutputPath,omitempty"`
	Password                               string `xml:"Password,omitempty"`
	PayloadCertificateUUID                 string `xml:"PayloadCertificateUUID,omitempty"`
	ShowRecoveryKey                        bool   `xml:"ShowRecoveryKey"`
	UseKeychain                            bool   `xml:"UseKeychain"`
	UseRecoveryKey                         bool   `xml:"UseRecoveryKey"`
	UserEntersMissingInfo                  bool   `xml:"UserEntersMissingInfo"`
	Username                               string `xml:"Username,omitempty"`
	PayloadDescription                     string `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName                     string `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier                      string `xml:"PayloadIdentifier,omitempty"`
	PayloadType                            string `xml:"PayloadType,omitempty"`
	PayloadUUID                            string `xml:"PayloadUUID,omitempty"`
	PayloadVersion                         int    `xml:"PayloadVersion,omitempty"`
}
