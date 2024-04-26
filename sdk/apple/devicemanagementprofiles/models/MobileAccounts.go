package models

import "encoding/xml"

// ResourceMobileAccountsConfigurationProfile represents the top-level structure of the plist for Mobile Accounts configurations
type ResourceMobileAccountsConfigurationProfile struct {
	XMLName                  xml.Name                                        `xml:"plist"`
	Version                  string                                          `xml:"version,attr"`
	Dict                     MobileAccountsConfigurationProfileSubsetPayload `xml:"dict"`
	PayloadDescription       string                                          `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                          `xml:"PayloadDisplayName,omitempty"`
	PayloadEnabled           string                                          `xml:"PayloadEnabled,omitempty"`
	PayloadIdentifier        string                                          `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                          `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                          `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                          `xml:"PayloadScope,omitempty"`
	PayloadType              string                                          `xml:"PayloadType,omitempty"`
	PayloadUUID              string                                          `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                                          `xml:"PayloadVersion,omitempty"`
}

// MobileAccountsConfigurationProfileSubsetPayload represents the content structure for configuring Mobile Accounts
type MobileAccountsConfigurationProfileSubsetPayload struct {
	CachedAccountsAskForSecureTokenAuthBypass bool   `xml:"cachedaccounts.askForSecureTokenAuthBypass,omitempty"`
	CachedAccountsExpiryDeleteDisusedSeconds  int    `xml:"cachedaccounts.expiry.delete.disusedSeconds,omitempty"`
	CachedAccountsWarnOnCreateAllowNever      bool   `xml:"cachedaccounts.WarnOnCreate.allowNever,omitempty"`
	ComAppleCachedAccountsCreateAtLogin       bool   `xml:"com.apple.cachedaccounts.CreateAtLogin,omitempty"`
	ComAppleCachedAccountsWarnOnCreate        bool   `xml:"com.apple.cachedaccounts.WarnOnCreate,omitempty"`
	PayloadIdentifier                         string `xml:"PayloadIdentifier,omitempty"`
	PayloadType                               string `xml:"PayloadType,omitempty"`
	PayloadUUID                               string `xml:"PayloadUUID,omitempty"`
	PayloadVersion                            int    `xml:"PayloadVersion,omitempty"`
}
