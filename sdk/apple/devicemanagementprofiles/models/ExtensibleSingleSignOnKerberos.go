package models

import "encoding/xml"

// ResourceExtensibleSingleSignOnKerberosConfigurationProfile represents the top-level structure of the plist for configuring Kerberos SSO app extension.
type ResourceExtensibleSingleSignOnKerberosConfigurationProfile struct {
	XMLName                  xml.Name                                                        `xml:"plist"`
	Version                  string                                                          `xml:"version,attr"`
	Payload                  ExtensibleSingleSignOnKerberosConfigurationProfileSubsetPayload `xml:"dict"`
	PayloadDescription       string                                                          `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                                          `xml:"PayloadDisplayName,omitempty"`
	PayloadEnabled           string                                                          `xml:"PayloadEnabled,omitempty"`
	PayloadIdentifier        string                                                          `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                                          `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                                          `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                                          `xml:"PayloadScope,omitempty"`
	PayloadType              string                                                          `xml:"PayloadType,omitempty"`
	PayloadUUID              string                                                          `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                                                          `xml:"PayloadVersion,omitempty"`
}

// ExtensibleSingleSignOnKerberosConfigurationProfileSubsetPayload represents the content structure for configuring Kerberos SSO app extension.
type ExtensibleSingleSignOnKerberosConfigurationProfileSubsetPayload struct {
	ExtensionData       ExtensibleSingleSignOnKerberosExtensionData `xml:"ExtensionData,omitempty"`
	ExtensionIdentifier string                                      `xml:"ExtensionIdentifier,omitempty"`
	Hosts               []string                                    `xml:"Hosts,omitempty"`
	Realm               string                                      `xml:"Realm,omitempty"`
	Type                string                                      `xml:"Type"`
	TeamIdentifier      string                                      `xml:"TeamIdentifier,omitempty"`
	PayloadIdentifier   string                                      `xml:"PayloadIdentifier,omitempty"`
	PayloadType         string                                      `xml:"PayloadType,omitempty"`
	PayloadUUID         string                                      `xml:"PayloadUUID,omitempty"`
	PayloadVersion      int                                         `xml:"PayloadVersion,omitempty"`
}

// ExtensibleSingleSignOnKerberosExtensionData represents the content structure for Extension Data of Kerberos SSO app extension.
type ExtensibleSingleSignOnKerberosExtensionData struct {
	AllowAutomaticLogin              bool                `xml:"allowAutomaticLogin,omitempty"`
	AllowPasswordChange              bool                `xml:"allowPasswordChange,omitempty"`
	AllowPlatformSSOAuthFallback     bool                `xml:"allowPlatformSSOAuthFallback,omitempty"`
	CacheName                        string              `xml:"cacheName,omitempty"`
	CertificateUUID                  string              `xml:"certificateUUID,omitempty"`
	CredentialBundleIdACL            []string            `xml:"credentialBundleIdACL,omitempty"`
	CredentialUseMode                string              `xml:"credentialUseMode,omitempty"`
	CustomUsernameLabel              string              `xml:"customUsernameLabel,omitempty"`
	DelayUserSetup                   bool                `xml:"delayUserSetup,omitempty"`
	DomainRealmMapping               map[string][]string `xml:"domainRealmMapping,omitempty"`
	HelpText                         string              `xml:"helpText,omitempty"`
	IncludeKerberosAppsInBundleIdACL bool                `xml:"includeKerberosAppsInBundleIdACL,omitempty"`
	IncludeManagedAppsInBundleIdACL  bool                `xml:"includeManagedAppsInBundleIdACL,omitempty"`
	IsDefaultRealm                   bool                `xml:"isDefaultRealm,omitempty"`
	MonitorCredentialsCache          bool                `xml:"monitorCredentialsCache,omitempty"`
	PrincipalName                    string              `xml:"principalName,omitempty"`
	PreferredKDCs                    []string            `xml:"preferredKDCs,omitempty"`
	PwChangeURL                      string              `xml:"pwChangeURL,omitempty"`
	PwNotificationDays               int                 `xml:"pwNotificationDays,omitempty"`
	PwExpireOverride                 int                 `xml:"pwExpireOverride,omitempty"`
	PwReqComplexity                  bool                `xml:"pwReqComplexity,omitempty"`
	PwReqHistory                     int                 `xml:"pwReqHistory,omitempty"`
	PwReqLength                      int                 `xml:"pwReqLength,omitempty"`
	PwReqMinAge                      int                 `xml:"pwReqMinAge,omitempty"`
	PwReqText                        string              `xml:"pwReqText,omitempty"`
	ReplicationTime                  int                 `xml:"replicationTime,omitempty"`
	RequireTLSForLDAP                bool                `xml:"requireTLSForLDAP,omitempty"`
	RequireUserPresence              bool                `xml:"requireUserPresence,omitempty"`
	PerformKerberosOnly              bool                `xml:"performKerberosOnly,omitempty"`
	SiteCode                         string              `xml:"siteCode,omitempty"`
	SyncLocalPassword                bool                `xml:"syncLocalPassword,omitempty"`
	UsePlatformSSOTGT                bool                `xml:"usePlatformSSOTGT,omitempty"`
	UseSiteAutoDiscovery             bool                `xml:"useSiteAutoDiscovery,omitempty"`
}
