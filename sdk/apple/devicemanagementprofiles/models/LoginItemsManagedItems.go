package models

import "encoding/xml"

// LoginItemsManagedItemsConfigurationProfile represents the top-level structure of the plist for configuring login items
type LoginItemsManagedItemsConfigurationProfile struct {
	XMLName                  xml.Name                                                `xml:"plist"`
	Version                  string                                                  `xml:"version,attr"`
	Payload                  LoginItemsManagedItemsConfigurationProfileSubsetPayload `xml:"dict>array>dict"`
	PayloadDescription       string                                                  `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                                  `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier        string                                                  `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                                  `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                                  `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                                  `xml:"PayloadScope,omitempty"`
	PayloadType              string                                                  `xml:"PayloadType,omitempty"`
	PayloadUUID              string                                                  `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                                                  `xml:"PayloadVersion,omitempty"`
}

// LoginItemsManagedItemsConfigurationProfileSubsetPayload represents the content structure for configuring login items
type LoginItemsManagedItemsConfigurationProfileSubsetPayload struct {
	PayloadContent     []LoginItemsManagedItem `xml:"PayloadContent>dict"`
	PayloadDisplayName string                  `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier  string                  `xml:"PayloadIdentifier,omitempty"`
	PayloadType        string                  `xml:"PayloadType,omitempty"`
	PayloadUUID        string                  `xml:"PayloadUUID,omitempty"`
	PayloadVersion     int                     `xml:"PayloadVersion,omitempty"`
}

// LoginItemsManagedItem represents a login item within the configuration
type LoginItemsManagedItem struct {
	AutoLaunchedApplicationDictionaryManaged []LoginItem `xml:"AutoLaunchedApplicationDictionary-managed>array>dict"`
	PayloadIdentifier                        string      `xml:"PayloadIdentifier,omitempty"`
	PayloadType                              string      `xml:"PayloadType,omitempty"`
	PayloadUUID                              string      `xml:"PayloadUUID,omitempty"`
	PayloadVersion                           int         `xml:"PayloadVersion,omitempty"`
}

// LoginItem represents the details about a login item
type LoginItem struct {
	Path string `xml:"Path"`
	Hide bool   `xml:"Hide"`
}
