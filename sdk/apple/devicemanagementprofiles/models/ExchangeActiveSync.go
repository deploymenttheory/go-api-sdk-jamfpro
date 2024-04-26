package models

import "encoding/xml"

// ResourceExchangeActiveSyncConfigurationProfile represents the top-level structure of the plist for configuring Exchange ActiveSync accounts
type ResourceExchangeActiveSyncConfigurationProfile struct {
	XMLName                  xml.Name                                            `xml:"plist"`
	Version                  string                                              `xml:"version,attr"`
	Dict                     ExchangeActiveSyncConfigurationProfileSubsetPayload `xml:"dict"`
	PayloadDescription       string                                              `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                              `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier        string                                              `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                              `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                              `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                              `xml:"PayloadScope,omitempty"`
	PayloadType              string                                              `xml:"PayloadType,omitempty"`
	PayloadUUID              string                                              `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                                              `xml:"PayloadVersion,omitempty"`
}

// ExchangeActiveSyncConfigurationProfileSubsetPayload represents the content structure for configuring Exchange ActiveSync accounts
type ExchangeActiveSyncConfigurationProfileSubsetPayload struct {
	PayloadContent     []ExchangeActiveSyncAccount `xml:"PayloadContent>array>dict"`
	PayloadDisplayName string                      `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier  string                      `xml:"PayloadIdentifier,omitempty"`
	PayloadType        string                      `xml:"PayloadType,omitempty"`
	PayloadUUID        string                      `xml:"PayloadUUID,omitempty"`
	PayloadVersion     int                         `xml:"PayloadVersion,omitempty"`
}

// ExchangeActiveSyncAccount represents an Exchange ActiveSync account configuration
type ExchangeActiveSyncAccount struct {
	AllowMailDrop                                  bool                       `xml:"allowMailDrop,omitempty"`
	Certificate                                    string                     `xml:"Certificate,omitempty"`
	CertificateName                                string                     `xml:"CertificateName,omitempty"`
	CertificatePassword                            string                     `xml:"CertificatePassword,omitempty"`
	CommunicationServiceRules                      []CommunicationServiceRule `xml:"CommunicationServiceRules>array>dict,omitempty"`
	DisableMailRecentsSyncing                      bool                       `xml:"disableMailRecentsSyncing,omitempty"`
	EmailAddress                                   string                     `xml:"EmailAddress,omitempty"`
	HeaderMagic                                    string                     `xml:"HeaderMagic,omitempty"`
	Host                                           string                     `xml:"Host,omitempty"`
	MailNumberOfPastDaysToSync                     int                        `xml:"MailNumberOfPastDaysToSync,omitempty"`
	OAuth                                          bool                       `xml:"OAuth,omitempty"`
	Password                                       string                     `xml:"Password,omitempty"`
	PayloadCertificateUUID                         string                     `xml:"PayloadCertificateUUID,omitempty"`
	PreventAppSheet                                bool                       `xml:"PreventAppSheet,omitempty"`
	PreventMove                                    bool                       `xml:"PreventMove,omitempty"`
	SMIMEEnabled                                   bool                       `xml:"SMIMEEnabled,omitempty"`
	SMIMEEncryptionEnabled                         bool                       `xml:"SMIMEEncryptionEnabled,omitempty"`
	SMIMESigningEnabled                            bool                       `xml:"SMIMESigningEnabled,omitempty"`
	SMIMESigningUserOverrideable                   bool                       `xml:"SMIMESigningUserOverrideable,omitempty"`
	SMIMEEnableEncryptionPerMessageSwitch          bool                       `xml:"SMIMEEnableEncryptionPerMessageSwitch,omitempty"`
	SMIMEEnablePerMessageSwitch                    bool                       `xml:"SMIMEEnablePerMessageSwitch,omitempty"`
	SMIMEEncryptByDefault                          bool                       `xml:"SMIMEEncryptByDefault,omitempty"`
	SMIMEEncryptByDefaultUserOverrideable          bool                       `xml:"SMIMEEncryptByDefaultUserOverrideable,omitempty"`
	SMIMEEncryptionCertificateUUID                 string                     `xml:"SMIMEEncryptionCertificateUUID,omitempty"`
	SMIMEEncryptionCertificateUUIDUserOverrideable bool                       `xml:"SMIMEEncryptionCertificateUUIDUserOverrideable,omitempty"`
	SMIMESigningCertificateUUID                    string                     `xml:"SMIMESigningCertificateUUID,omitempty"`
	SMIMESigningCertificateUUIDUserOverrideable    bool                       `xml:"SMIMESigningCertificateUUIDUserOverrideable,omitempty"`
	SSL                                            bool                       `xml:"SSL,omitempty"`
	UserName                                       string                     `xml:"UserName,omitempty"`
	EnableCalendars                                bool                       `xml:"EnableCalendars,omitempty"`
	EnableCalendarsUserOverridable                 bool                       `xml:"EnableCalendarsUserOverridable,omitempty"`
	EnableContacts                                 bool                       `xml:"EnableContacts,omitempty"`
	EnableContactsUserOverridable                  bool                       `xml:"EnableContactsUserOverridable,omitempty"`
	EnableMail                                     bool                       `xml:"EnableMail,omitempty"`
	EnableMailUserOverridable                      bool                       `xml:"EnableMailUserOverridable,omitempty"`
	EnableNotes                                    bool                       `xml:"EnableNotes,omitempty"`
	EnableNotesUserOverridable                     bool                       `xml:"EnableNotesUserOverridable,omitempty"`
	EnableReminders                                bool                       `xml:"EnableReminders,omitempty"`
	EnableRemindersUserOverridable                 bool                       `xml:"EnableRemindersUserOverridable,omitempty"`
	OAuthSignInURL                                 string                     `xml:"OAuthSignInURL,omitempty"`
	OAuthTokenRequestURL                           string                     `xml:"OAuthTokenRequestURL,omitempty"`
	OverridePreviousPassword                       bool                       `xml:"OverridePreviousPassword,omitempty"`
	VPNUUID                                        string                     `xml:"VPNUUID,omitempty"`
	PayloadIdentifier                              string                     `xml:"PayloadIdentifier,omitempty"`
	PayloadType                                    string                     `xml:"PayloadType,omitempty"`
	PayloadUUID                                    string                     `xml:"PayloadUUID,omitempty"`
	PayloadVersion                                 int                        `xml:"PayloadVersion,omitempty"`
}

// CommunicationServiceRule represents a communication service rule
type CommunicationServiceRule struct {
	Comment        string `xml:"Comment,omitempty"`
	RuleType       string `xml:"RuleType,omitempty"`
	RuleValue      string `xml:"RuleValue,omitempty"`
	TeamIdentifier string `xml:"TeamIdentifier,omitempty"`
}

// DefaultServiceHandlers represents the default service handlers for Exchange ActiveSync communication service rules
type DefaultServiceHandlers struct {
	AudioCall string `xml:"AudioCall,omitempty"`
}
