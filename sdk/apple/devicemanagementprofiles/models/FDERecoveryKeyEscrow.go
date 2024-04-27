package models

import "encoding/xml"

// FDERecoveryKeyEscrowConfigurationProfile represents the top-level structure of the plist for configuring FileVault recovery key escrow
type FDERecoveryKeyEscrowConfigurationProfile struct {
	XMLName                  xml.Name                                              `xml:"plist"`
	Version                  string                                                `xml:"version,attr"`
	Payload                  FDERecoveryKeyEscrowConfigurationProfileSubsetPayload `xml:"dict"`
	PayloadDescription       string                                                `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                                `xml:"PayloadDisplayName,omitempty"`
	PayloadEnabled           string                                                `xml:"PayloadEnabled,omitempty"`
	PayloadIdentifier        string                                                `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                                `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                                `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                                `xml:"PayloadScope,omitempty"`
	PayloadType              string                                                `xml:"PayloadType,omitempty"`
	PayloadUUID              string                                                `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                                                `xml:"PayloadVersion,omitempty"`
}

// FDERecoveryKeyEscrowConfigurationProfileSubsetPayload represents the content structure for configuring FileVault recovery key escrow
type FDERecoveryKeyEscrowConfigurationProfileSubsetPayload struct {
	PayloadContent     []FDERecoveryKeyEscrow `xml:"PayloadContent>dict"`
	PayloadDisplayName string                 `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier  string                 `xml:"PayloadIdentifier,omitempty"`
	PayloadType        string                 `xml:"PayloadType,omitempty"`
	PayloadUUID        string                 `xml:"PayloadUUID,omitempty"`
	PayloadVersion     int                    `xml:"PayloadVersion,omitempty"`
}

// FDERecoveryKeyEscrow represents the FileVault recovery key escrow dictionary within the configuration
type FDERecoveryKeyEscrow struct {
	Defer                                  bool   `xml:"Defer"`
	DeferDontAskAtUserLogout               bool   `xml:"DeferDontAskAtUserLogout"`
	DeferForceAtUserLoginMaxBypassAttempts int    `xml:"DeferForceAtUserLoginMaxBypassAttempts"`
	Enable                                 string `xml:"Enable"`
	EncryptCertPayloadUUID                 string `xml:"EncryptCertPayloadUUID"`
	Location                               string `xml:"Location"`
	ShowRecoveryKey                        bool   `xml:"ShowRecoveryKey"`
	UseKeychain                            bool   `xml:"UseKeychain"`
	UseRecoveryKey                         bool   `xml:"UseRecoveryKey"`
	UserEntersMissingInfo                  bool   `xml:"UserEntersMissingInfo"`
	PayloadIdentifier                      string `xml:"PayloadIdentifier,omitempty"`
	PayloadType                            string `xml:"PayloadType,omitempty"`
	PayloadUUID                            string `xml:"PayloadUUID,omitempty"`
	PayloadVersion                         int    `xml:"PayloadVersion,omitempty"`
}
