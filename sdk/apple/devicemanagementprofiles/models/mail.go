package models

import "encoding/xml"

// ResourceMailConfigurationProfile represents the top-level structure of the plist for configuring Mail accounts
type ResourceMailConfigurationProfile struct {
	XMLName                  xml.Name                              `xml:"plist"`
	Version                  string                                `xml:"version,attr"`
	Payload                  MailConfigurationProfileSubsetPayload `xml:"dict"`
	PayloadDescription       string                                `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier        string                                `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                `xml:"PayloadScope,omitempty"`
	PayloadType              string                                `xml:"PayloadType,omitempty"`
	PayloadUUID              string                                `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                                `xml:"PayloadVersion,omitempty"`
}

// MailConfigurationProfileSubsetPayload represents the content structure for configuring Mail accounts
type MailConfigurationProfileSubsetPayload struct {
	PayloadContent     []MailAccount `xml:"PayloadContent>array>dict"`
	PayloadDisplayName string        `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier  string        `xml:"PayloadIdentifier,omitempty"`
	PayloadType        string        `xml:"PayloadType,omitempty"`
	PayloadUUID        string        `xml:"PayloadUUID,omitempty"`
	PayloadVersion     int           `xml:"PayloadVersion,omitempty"`
}

// MailAccount represents a Mail account configuration
type MailAccount struct {
	AllowMailDrop                                  bool   `xml:"allowMailDrop,omitempty"`
	DisableMailRecentsSyncing                      bool   `xml:"disableMailRecentsSyncing,omitempty"`
	EmailAccountDescription                        string `xml:"EmailAccountDescription,omitempty"`
	EmailAccountName                               string `xml:"EmailAccountName,omitempty"`
	EmailAccountType                               string `xml:"EmailAccountType,omitempty"`
	EmailAddress                                   string `xml:"EmailAddress,omitempty"`
	IncomingMailServerAuthentication               string `xml:"IncomingMailServerAuthentication,omitempty"`
	IncomingMailServerHostName                     string `xml:"IncomingMailServerHostName,omitempty"`
	IncomingMailServerIMAPPathPrefix               string `xml:"IncomingMailServerIMAPPathPrefix,omitempty"`
	IncomingMailServerPortNumber                   int    `xml:"IncomingMailServerPortNumber,omitempty"`
	IncomingMailServerUsername                     string `xml:"IncomingMailServerUsername,omitempty"`
	IncomingMailServerUseSSL                       bool   `xml:"IncomingMailServerUseSSL,omitempty"`
	IncomingPassword                               string `xml:"IncomingPassword,omitempty"`
	OutgoingMailServerAuthentication               string `xml:"OutgoingMailServerAuthentication,omitempty"`
	OutgoingMailServerHostName                     string `xml:"OutgoingMailServerHostName,omitempty"`
	OutgoingMailServerPortNumber                   int    `xml:"OutgoingMailServerPortNumber,omitempty"`
	OutgoingMailServerUsername                     string `xml:"OutgoingMailServerUsername,omitempty"`
	OutgoingMailServerUseSSL                       bool   `xml:"OutgoingMailServerUseSSL,omitempty"`
	OutgoingPassword                               string `xml:"OutgoingPassword,omitempty"`
	OutgoingPasswordSameAsIncomingPassword         bool   `xml:"OutgoingPasswordSameAsIncomingPassword,omitempty"`
	PreventAppSheet                                bool   `xml:"PreventAppSheet,omitempty"`
	PreventMove                                    bool   `xml:"PreventMove,omitempty"`
	SMIMEEnabled                                   bool   `xml:"SMIMEEnabled,omitempty"`
	SMIMEEnableEncryptionPerMessageSwitch          bool   `xml:"SMIMEEnableEncryptionPerMessageSwitch,omitempty"`
	SMIMEEnablePerMessageSwitch                    bool   `xml:"SMIMEEnablePerMessageSwitch,omitempty"`
	SMIMEEncryptByDefault                          bool   `xml:"SMIMEEncryptByDefault,omitempty"`
	SMIMEEncryptByDefaultUserOverrideable          bool   `xml:"SMIMEEncryptByDefaultUserOverrideable,omitempty"`
	SMIMEEncryptionCertificateUUID                 string `xml:"SMIMEEncryptionCertificateUUID,omitempty"`
	SMIMEEncryptionCertificateUUIDUserOverrideable bool   `xml:"SMIMEEncryptionCertificateUUIDUserOverrideable,omitempty"`
	SMIMEEncryptionEnabled                         bool   `xml:"SMIMEEncryptionEnabled,omitempty"`
	SMIMESigningCertificateUUID                    string `xml:"SMIMESigningCertificateUUID,omitempty"`
	SMIMESigningCertificateUUIDUserOverrideable    bool   `xml:"SMIMESigningCertificateUUIDUserOverrideable,omitempty"`
	SMIMESigningEnabled                            bool   `xml:"SMIMESigningEnabled,omitempty"`
	SMIMESigningUserOverrideable                   bool   `xml:"SMIMESigningUserOverrideable,omitempty"`
	VPNUUID                                        string `xml:"VPNUUID,omitempty"`
	PayloadDisplayName                             string `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier                              string `xml:"PayloadIdentifier,omitempty"`
	PayloadType                                    string `xml:"PayloadType,omitempty"`
	PayloadUUID                                    string `xml:"PayloadUUID,omitempty"`
	PayloadVersion                                 int    `xml:"PayloadVersion,omitempty"`
}
