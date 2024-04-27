package models

import "encoding/xml"

// ResourceLDAPConfigurationProfile represents the top-level structure of the plist for LDAP account configurations
type ResourceLDAPConfigurationProfile struct {
	XMLName                  xml.Name                              `xml:"plist"`
	Version                  string                                `xml:"version,attr"`
	Payload                  LDAPConfigurationProfileSubsetPayload `xml:"dict"`
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

// LDAPConfigurationProfileSubsetPayload represents the content structure for configuring an LDAP account
type LDAPConfigurationProfileSubsetPayload struct {
	LDAPAccountDescription string                   `xml:"LDAPAccountDescription,omitempty"`
	LDAPAccountHostName    string                   `xml:"LDAPAccountHostName,omitempty"`
	LDAPAccountPassword    string                   `xml:"LDAPAccountPassword,omitempty"`
	LDAPAccountUserName    string                   `xml:"LDAPAccountUserName,omitempty"`
	LDAPAccountUseSSL      bool                     `xml:"LDAPAccountUseSSL,omitempty"`
	LDAPSearchSettings     []LDAPSearchSettingsItem `xml:"LDAPSearchSettings>LDAPSearchSettingsItem,omitempty"`
	VPNUUID                string                   `xml:"VPNUUID,omitempty"`
	PayloadIdentifier      string                   `xml:"PayloadIdentifier,omitempty"`
	PayloadType            string                   `xml:"PayloadType,omitempty"`
	PayloadUUID            string                   `xml:"PayloadUUID,omitempty"`
	PayloadVersion         int                      `xml:"PayloadVersion,omitempty"`
}

// LDAPSearchSettingsItem represents an individual LDAP search setting
type LDAPSearchSettingsItem struct {
	LDAPSearchSettingDescription string `xml:"LDAPSearchSettingDescription,omitempty"`
	LDAPSearchSettingScope       string `xml:"LDAPSearchSettingScope,omitempty"`
	LDAPSearchSettingSearchBase  string `xml:"LDAPSearchSettingSearchBase,omitempty"`
}
