package models

import "encoding/xml"

// LoginWindowConfigurationProfile represents the top-level structure of the plist for configuring login window behavior
type LoginWindowConfigurationProfile struct {
	XMLName                  xml.Name                                     `xml:"plist"`
	Version                  string                                       `xml:"version,attr"`
	Dict                     LoginWindowConfigurationProfileSubsetPayload `xml:"dict"`
	PayloadDescription       string                                       `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                       `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier        string                                       `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                       `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                       `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                       `xml:"PayloadScope,omitempty"`
	PayloadType              string                                       `xml:"PayloadType,omitempty"`
	PayloadUUID              string                                       `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                                       `xml:"PayloadVersion,omitempty"`
}

// LoginWindowConfigurationProfileSubsetPayload represents the content structure for configuring login window behavior
type LoginWindowConfigurationProfileSubsetPayload struct {
	PayloadContent     []LoginWindowSetting `xml:"PayloadContent>array>dict"`
	PayloadDisplayName string               `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier  string               `xml:"PayloadIdentifier,omitempty"`
	PayloadType        string               `xml:"PayloadType,omitempty"`
	PayloadUUID        string               `xml:"PayloadUUID,omitempty"`
	PayloadVersion     int                  `xml:"PayloadVersion,omitempty"`
}

// LoginWindowSetting represents a login window setting
type LoginWindowSetting struct {
	AdminHostInfo                 string   `xml:"AdminHostInfo,omitempty"`
	AllowList                     []string `xml:"AllowList>array>string,omitempty"`
	AutologinPassword             string   `xml:"AutologinPassword,omitempty"`
	AutologinUsername             string   `xml:"AutologinUsername,omitempty"`
	DenyList                      []string `xml:"DenyList>array>string,omitempty"`
	DisableConsoleAccess          bool     `xml:"DisableConsoleAccess,omitempty"`
	DisableFDEAutoLogin           bool     `xml:"DisableFDEAutoLogin,omitempty"`
	DisableScreenLockImmediate    bool     `xml:"DisableScreenLockImmediate,omitempty"`
	HideAdminUsers                bool     `xml:"HideAdminUsers,omitempty"`
	HideLocalUsers                bool     `xml:"HideLocalUsers,omitempty"`
	HideMobileAccounts            bool     `xml:"HideMobileAccounts,omitempty"`
	IncludeNetworkUser            bool     `xml:"IncludeNetworkUser,omitempty"`
	LoginwindowText               string   `xml:"LoginwindowText,omitempty"`
	LogOutDisabledWhileLoggedIn   bool     `xml:"LogOutDisabledWhileLoggedIn,omitempty"`
	PowerOffDisabledWhileLoggedIn bool     `xml:"PowerOffDisabledWhileLoggedIn,omitempty"`
	RestartDisabled               bool     `xml:"RestartDisabled,omitempty"`
	RestartDisabledWhileLoggedIn  bool     `xml:"RestartDisabledWhileLoggedIn,omitempty"`
	SHOWFULLNAME                  bool     `xml:"SHOWFULLNAME,omitempty"`
	SHOWOTHERUSERS_MANAGED        bool     `xml:"SHOWOTHERUSERS_MANAGED,omitempty"`
	ShowInputMenu                 bool     `xml:"showInputMenu,omitempty"`
	ShutDownDisabled              bool     `xml:"ShutDownDisabled,omitempty"`
	SleepDisabled                 bool     `xml:"SleepDisabled,omitempty"`
	ShutDownDisabledWhileLoggedIn bool     `xml:"ShutDownDisabledWhileLoggedIn,omitempty"`
	PayloadIdentifier             string   `xml:"PayloadIdentifier,omitempty"`
	PayloadType                   string   `xml:"PayloadType,omitempty"`
	PayloadUUID                   string   `xml:"PayloadUUID,omitempty"`
	PayloadVersion                int      `xml:"PayloadVersion,omitempty"`
}
