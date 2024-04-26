package models

import "encoding/xml"

// LoginWindowScriptsConfigurationProfile represents the top-level structure of the plist for configuring login window scripts
type LoginWindowScriptsConfigurationProfile struct {
	XMLName                  xml.Name                                            `xml:"plist"`
	Version                  string                                              `xml:"version,attr"`
	Dict                     LoginWindowScriptsConfigurationProfileSubsetPayload `xml:"dict"`
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

// LoginWindowScriptsConfigurationProfileSubsetPayload represents the content structure for configuring login window scripts
type LoginWindowScriptsConfigurationProfileSubsetPayload struct {
	PayloadContent     []LoginWindowScript `xml:"PayloadContent>array>dict"`
	PayloadDisplayName string              `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier  string              `xml:"PayloadIdentifier,omitempty"`
	PayloadType        string              `xml:"PayloadType,omitempty"`
	PayloadUUID        string              `xml:"PayloadUUID,omitempty"`
	PayloadVersion     int                 `xml:"PayloadVersion,omitempty"`
}

// LoginWindowScript represents a login window script
type LoginWindowScript struct {
	Loginscripts      []ScriptItem `xml:"loginscripts>array>dict,omitempty"`
	Logoutscripts     []ScriptItem `xml:"logoutscripts>array>dict,omitempty"`
	SkipLoginHook     bool         `xml:"skipLoginHook,omitempty"`
	SkipLogoutHook    bool         `xml:"skipLogoutHook,omitempty"`
	PayloadIdentifier string       `xml:"PayloadIdentifier,omitempty"`
	PayloadType       string       `xml:"PayloadType,omitempty"`
	PayloadUUID       string       `xml:"PayloadUUID,omitempty"`
	PayloadVersion    int          `xml:"PayloadVersion,omitempty"`
}

// ScriptItem represents a script item
type ScriptItem struct {
	FileData string `xml:"filedata,omitempty"`
	Filename string `xml:"filename,omitempty"`
}
