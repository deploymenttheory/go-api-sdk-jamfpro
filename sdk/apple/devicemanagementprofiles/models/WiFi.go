package models

import "encoding/xml"

// ResourceWiFiConfigurationProfile represents the top-level structure of the plist
type ResourceWiFiConfigurationProfile struct {
	XMLName                  xml.Name                              `xml:"plist"`
	Version                  string                                `xml:"version,attr"`
	Payload                  WiFiConfigurationProfileSubsetPayload `xml:"dict"`
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

// WiFiConfigurationProfileSubsetPayload represents the nested dictionary structure within the plist
type WiFiConfigurationProfileSubsetPayload struct {
	XMLName                       xml.Name                                             `xml:"dict"`
	SSID_STR                      string                                               `xml:"SSID_STR,omitempty"`
	HIDDEN_NETWORK                string                                               `xml:"HIDDEN_NETWORK,omitempty"`
	AutoJoin                      string                                               `xml:"AutoJoin,omitempty"`
	SetupModes                    []string                                             `xml:"SetupModes,omitempty"`
	EncryptionType                string                                               `xml:"EncryptionType,omitempty"`
	IsHotspot                     string                                               `xml:"IsHotspot,omitempty"`
	DomainName                    string                                               `xml:"DomainName,omitempty"`
	ServiceProviderRoamingEnabled string                                               `xml:"ServiceProviderRoamingEnabled,omitempty"`
	RoamingConsortiumOIs          []string                                             `xml:"RoamingConsortiumOIs,omitempty"`
	NAIRealmNames                 []string                                             `xml:"NAIRealmNames,omitempty"`
	MCCAndMNCs                    []string                                             `xml:"MCCAndMNCs,omitempty"`
	DisplayedOperatorName         string                                               `xml:"DisplayedOperatorName,omitempty"`
	ProxyType                     string                                               `xml:"ProxyType,omitempty"`
	CaptiveBypass                 string                                               `xml:"CaptiveBypass,omitempty"`
	QoSMarkingPolicy              WiFiConfigurationProfileSubsetQoSMarkingPolicy       `xml:"QoSMarkingPolicy,omitempty"`
	ProxyServer                   string                                               `xml:"ProxyServer,omitempty"`
	ProxyServerPort               string                                               `xml:"ProxyServerPort,omitempty"`
	ProxyUsername                 string                                               `xml:"ProxyUsername,omitempty"`
	ProxyPassword                 string                                               `xml:"ProxyPassword,omitempty"`
	ProxyPACURL                   string                                               `xml:"ProxyPACURL,omitempty"`
	ProxyPACFallbackAllowed       string                                               `xml:"ProxyPACFallbackAllowed,omitempty"`
	EAPClientConfiguration        WiFiConfigurationProfileSubsetEAPClientConfiguration `xml:"EAPClientConfiguration,omitempty"`
	HESSID                        string                                               `xml:"HESSID,omitempty"`
	EnableIPv6                    string                                               `xml:"EnableIPv6,omitempty"`
	TLSCertificateRequired        string                                               `xml:"TLSCertificateRequired,omitempty"`
	TLSCertificateIsRequired      string                                               `xml:"TLSCertificateIsRequired,omitempty"`
	TLSMaximumVersion             string                                               `xml:"TLSMaximumVersion,omitempty"`
	TLSMinimumVersion             string                                               `xml:"TLSMinimumVersion,omitempty"`
	TLSTrustedCertificates        []string                                             `xml:"TLSTrustedCertificates,omitempty"`
	TLSTrustedServerNames         []string                                             `xml:"TLSTrustedServerNames,omitempty"`
	TTLSInnerAuthentication       string                                               `xml:"TTLSInnerAuthentication,omitempty"`
	UserName                      string                                               `xml:"UserName,omitempty"`
	UserPassword                  string                                               `xml:"UserPassword,omitempty"`
}

// WiFiConfigurationProfileSubsetEAPClientConfiguration represents the nested dictionary structure for EAP client configuration
type WiFiConfigurationProfileSubsetEAPClientConfiguration struct {
	XMLName                        xml.Name `xml:"EAPClientConfiguration"`
	AcceptEAPTypes                 []int    `xml:"AcceptEAPTypes>integer,omitempty"`
	EAPFASTProvisionPAC            string   `xml:"EAPFASTProvisionPAC,omitempty"`
	EAPFASTProvisionPACAnonymously string   `xml:"EAPFASTProvisionPACAnonymously,omitempty"`
	EAPFASTUsePAC                  string   `xml:"EAPFASTUsePAC,omitempty"`
	EAPSIMNumberOfRANDs            int      `xml:"EAPSIMNumberOfRANDs,omitempty"`
	OneTimeUserPassword            string   `xml:"OneTimeUserPassword,omitempty"`
	OuterIdentity                  string   `xml:"OuterIdentity,omitempty"`
	SystemModeCredentialsSource    string   `xml:"SystemModeCredentialsSource,omitempty"`
	TLSMaximumVersion              string   `xml:"TLSMaximumVersion,omitempty"`
	TLSMinimumVersion              string   `xml:"TLSMinimumVersion,omitempty"`
	TTLSInnerAuthentication        string   `xml:"TTLSInnerAuthentication,omitempty"`
	UserName                       string   `xml:"UserName,omitempty"`
	UserPassword                   string   `xml:"UserPassword,omitempty"`
}

// WiFiConfigurationProfileSubsetQoSMarkingPolicy represents the nested dictionary structure for QoS marking policy
type WiFiConfigurationProfileSubsetQoSMarkingPolicy struct {
	XMLName                             xml.Name `xml:"QoSMarkingPolicy"`
	QoSMarkingAppleAudioVideoCalls      string   `xml:"QoSMarkingAppleAudioVideoCalls,omitempty"`
	QoSMarkingEnabled                   string   `xml:"QoSMarkingEnabled,omitempty"`
	QoSMarkingWhitelistedAppIdentifiers []string `xml:"QoSMarkingWhitelistedAppIdentifiers>string,omitempty"`
}
