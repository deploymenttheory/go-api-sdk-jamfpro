package models

import "encoding/xml"

// Resource8021XGlobalEthernetConfigurationProfile represents the top-level structure of the plist for 802.1X Global Ethernet configurations
type Resource8021XGlobalEthernetConfigurationProfile struct {
	XMLName                  xml.Name                                              `xml:"plist"`
	Version                  string                                                `xml:"version,attr"`
	Dict                     X8021XGlobalEthernetConfigurationProfileSubsetPayload `xml:"dict"`
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

// X8021XGlobalEthernetConfigurationProfileSubsetPayload represents the content structure for configuring 802.1X Global Ethernet settings
type X8021XGlobalEthernetConfigurationProfileSubsetPayload struct {
	PayloadContent     []X8021XGlobalEthernet `xml:"PayloadContent>dict"`
	PayloadDisplayName string                 `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier  string                 `xml:"PayloadIdentifier,omitempty"`
	PayloadType        string                 `xml:"PayloadType,omitempty"`
	PayloadUUID        string                 `xml:"PayloadUUID,omitempty"`
	PayloadVersion     int                    `xml:"PayloadVersion,omitempty"`
}

// X8021XGlobalEthernet represents the 802.1X Global Ethernet dictionary within 802.1X Global Ethernet configuration
type X8021XGlobalEthernet struct {
	AuthenticationMethod   string                  `xml:"AuthenticationMethod,omitempty"`
	AutoJoin               bool                    `xml:"AutoJoin"`
	CaptiveBypass          bool                    `xml:"CaptiveBypass"`
	EAPClientConfiguration XEAPClientConfiguration `xml:"EAPClientConfiguration,omitempty"`
	EncryptionType         string                  `xml:"EncryptionType,omitempty"`
	HiddenNetwork          bool                    `xml:"HIDDEN_NETWORK"`
	Interface              string                  `xml:"Interface"`
	Password               string                  `xml:"Password,omitempty"`
	ProxyType              string                  `xml:"ProxyType,omitempty"`
	SetupModes             []string                `xml:"SetupModes>string"`
	PayloadDescription     string                  `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName     string                  `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier      string                  `xml:"PayloadIdentifier,omitempty"`
	PayloadType            string                  `xml:"PayloadType,omitempty"`
	PayloadUUID            string                  `xml:"PayloadUUID,omitempty"`
	PayloadVersion         int                     `xml:"PayloadVersion,omitempty"`
}
