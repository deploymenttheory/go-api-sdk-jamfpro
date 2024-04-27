package models

import "encoding/xml"

// ResourceExtensibleSingleSignOnConfigurationProfile represents the top-level structure of the plist for configuring SSO app extension.
type ResourceExtensibleSingleSignOnConfigurationProfile struct {
	XMLName                  xml.Name                                                `xml:"plist"`
	Version                  string                                                  `xml:"version,attr"`
	Payload                  ExtensibleSingleSignOnConfigurationProfileSubsetPayload `xml:"dict>array>dict"`
	PayloadDescription       string                                                  `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                                  `xml:"PayloadDisplayName,omitempty"`
	PayloadEnabled           string                                                  `xml:"PayloadEnabled,omitempty"`
	PayloadIdentifier        string                                                  `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                                  `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                                  `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                                  `xml:"PayloadScope,omitempty"`
	PayloadType              string                                                  `xml:"PayloadType,omitempty"`
	PayloadUUID              string                                                  `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                                                  `xml:"PayloadVersion,omitempty"`
}

// ExtensibleSingleSignOnConfigurationProfileSubsetPayload represents the content structure for configuring SSO app extension.
type ExtensibleSingleSignOnConfigurationProfileSubsetPayload struct {
	AuthenticationMethod    string                                         `xml:"AuthenticationMethod,omitempty"`
	DeniedBundleIdentifiers []string                                       `xml:"DeniedBundleIdentifiers,omitempty"`
	ExtensionData           map[string]interface{}                         `xml:"ExtensionData,omitempty"`
	ExtensionIdentifier     string                                         `xml:"ExtensionIdentifier"`
	Hosts                   []string                                       `xml:"Hosts,omitempty"`
	PlatformSSO             PayloadSubsetExtensibleSingleSignOnPlatformSSO `xml:"PlatformSSO,omitempty"`
	Realm                   string                                         `xml:"Realm,omitempty"`
	RegistrationToken       string                                         `xml:"RegistrationToken,omitempty"`
	ScreenLockedBehavior    string                                         `xml:"ScreenLockedBehavior,omitempty"`
	TeamIdentifier          string                                         `xml:"TeamIdentifier,omitempty"`
	Type                    string                                         `xml:"Type"`
	URLs                    []string                                       `xml:"URLs,omitempty"`
	PayloadIdentifier       string                                         `xml:"PayloadIdentifier,omitempty"`
	PayloadType             string                                         `xml:"PayloadType,omitempty"`
	PayloadUUID             string                                         `xml:"PayloadUUID,omitempty"`
	PayloadVersion          int                                            `xml:"PayloadVersion,omitempty"`
}

// PayloadSubsetExtensibleSingleSignOnPlatformSSO represents the structure for configuring Platform SSO.
type PayloadSubsetExtensibleSingleSignOnPlatformSSO struct {
	AccountDisplayName       string                                   `xml:"AccountDisplayName,omitempty"`
	AdditionalGroups         []string                                 `xml:"AdditionalGroups,omitempty"`
	AdministratorGroups      []string                                 `xml:"AdministratorGroups,omitempty"`
	AuthenticationMethod     string                                   `xml:"AuthenticationMethod,omitempty"`
	AuthorizationGroups      map[string]string                        `xml:"AuthorizationGroups,omitempty"`
	EnableAuthorization      bool                                     `xml:"EnableAuthorization,omitempty"`
	EnableCreateUserAtLogin  bool                                     `xml:"EnableCreateUserAtLogin,omitempty"`
	LoginFrequency           int                                      `xml:"LoginFrequency,omitempty"`
	NewUserAuthorizationMode string                                   `xml:"NewUserAuthorizationMode,omitempty"`
	TokenToUserMapping       ExtensibleSingleSignOnTokenToUserMapping `xml:"TokenToUserMapping,omitempty"`
	UserAuthorizationMode    string                                   `xml:"UserAuthorizationMode,omitempty"`
	UseSharedDeviceKeys      bool                                     `xml:"UseSharedDeviceKeys,omitempty"`
}

// ExtensibleSingleSignOnTokenToUserMapping represents the attribute mapping for creating new users or for authorization.
type ExtensibleSingleSignOnTokenToUserMapping struct {
	AccountName string `xml:"AccountName,omitempty"`
	FullName    string `xml:"FullName,omitempty"`
}
