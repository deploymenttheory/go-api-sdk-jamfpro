package models

import "encoding/xml"

// LoginWindowLoginItemsConfigurationProfile represents the top-level structure of the plist for configuring login behavior
type LoginWindowLoginItemsConfigurationProfile struct {
	XMLName                  xml.Name                                               `xml:"plist"`
	Version                  string                                                 `xml:"version,attr"`
	Payload                  LoginWindowLoginItemsConfigurationProfileSubsetPayload `xml:"dict>array>dict"`
	PayloadDescription       string                                                 `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                                 `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier        string                                                 `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                                 `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                                 `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                                 `xml:"PayloadScope,omitempty"`
	PayloadType              string                                                 `xml:"PayloadType,omitempty"`
	PayloadUUID              string                                                 `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                                                 `xml:"PayloadVersion,omitempty"`
}

// LoginWindowLoginItemsConfigurationProfileSubsetPayload represents the content structure for configuring login behavior
type LoginWindowLoginItemsConfigurationProfileSubsetPayload struct {
	PayloadContent     []LoginWindowLoginItem `xml:"PayloadContent>array>dict"`
	PayloadDisplayName string                 `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier  string                 `xml:"PayloadIdentifier,omitempty"`
	PayloadType        string                 `xml:"PayloadType,omitempty"`
	PayloadUUID        string                 `xml:"PayloadUUID,omitempty"`
	PayloadVersion     int                    `xml:"PayloadVersion,omitempty"`
}

// LoginWindowLoginItem represents the login item configuration
type LoginWindowLoginItem struct {
	DisableLoginItemsSuppression bool   `xml:"DisableLoginItemsSuppression"`
	PayloadIdentifier            string `xml:"PayloadIdentifier,omitempty"`
	PayloadType                  string `xml:"PayloadType,omitempty"`
	PayloadUUID                  string `xml:"PayloadUUID,omitempty"`
	PayloadVersion               int    `xml:"PayloadVersion,omitempty"`
}
