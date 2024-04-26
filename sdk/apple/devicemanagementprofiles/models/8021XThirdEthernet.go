package models

import "encoding/xml"

// Resource8021XThirdEthernetConfigurationProfile represents the top-level structure of the plist for 802.1X Third Ethernet configurations
type Resource8021XThirdEthernetConfigurationProfile struct {
	XMLName                  xml.Name                                             `xml:"plist"`
	Version                  string                                               `xml:"version,attr"`
	Dict                     X8021XThirdEthernetConfigurationProfileSubsetPayload `xml:"dict"`
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

// X8021XThirdEthernetConfigurationProfileSubsetPayload represents the content structure for configuring 802.1X Third Ethernet settings
type X8021XThirdEthernetConfigurationProfileSubsetPayload struct {
	PayloadContent     []X8021XThirdEthernet `xml:"PayloadContent>dict"`
	PayloadDisplayName string                `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier  string                `xml:"PayloadIdentifier,omitempty"`
	PayloadType        string                `xml:"PayloadType,omitempty"`
	PayloadUUID        string                `xml:"PayloadUUID,omitempty"`
	PayloadVersion     int                   `xml:"PayloadVersion,omitempty"`
}

// X8021XThirdEthernet represents the 802.1X Third Ethernet dictionary within 802.1X Third Ethernet configuration
type X8021XThirdEthernet struct {
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
