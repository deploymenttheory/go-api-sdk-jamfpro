package models

import "encoding/xml"

// ResourceAppStoreConfigurationProfile represents the top-level structure of the plist for App Store restrictions configurations
type ResourceAppStoreConfigurationProfile struct {
	XMLName                  xml.Name                                  `xml:"plist"`
	Version                  string                                    `xml:"version,attr"`
	Payload                  AppStoreConfigurationProfileSubsetPayload `xml:"dict>array>dict"`
	PayloadDescription       string                                    `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                    `xml:"PayloadDisplayName,omitempty"`
	PayloadEnabled           string                                    `xml:"PayloadEnabled,omitempty"`
	PayloadIdentifier        string                                    `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                    `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                    `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                    `xml:"PayloadScope,omitempty"`
	PayloadType              string                                    `xml:"PayloadType,omitempty"`
	PayloadUUID              string                                    `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                                    `xml:"PayloadVersion,omitempty"`
}

// AppStoreConfigurationProfileSubsetPayload represents the content structure for configuring App Store restrictions
type AppStoreConfigurationProfileSubsetPayload struct {
	DisableSoftwareUpdateNotifications bool   `xml:"DisableSoftwareUpdateNotifications"`
	RestrictStoreDisableAppAdoption    bool   `xml:"restrict-store-disable-app-adoption"`
	RestrictStoreRequireAdminToInstall bool   `xml:"restrict-store-require-admin-to-install,omitempty"` // Deprecated
	RestrictStoreSoftwareUpdateOnly    bool   `xml:"restrict-store-softwareupdate-only"`
	PayloadIdentifier                  string `xml:"PayloadIdentifier,omitempty"`
	PayloadType                        string `xml:"PayloadType,omitempty"`
	PayloadUUID                        string `xml:"PayloadUUID,omitempty"`
	PayloadVersion                     int    `xml:"PayloadVersion,omitempty"`
}
