package models

import (
	"encoding/xml"
)

// ManagedPreferences represents the payload used to configure managed preferences.
type ManagedPreferences struct {
	PayloadContent     []PayloadContent `plist:"PayloadContent"`
	PayloadDisplayName string           `plist:"PayloadDisplayName"`
	PayloadIdentifier  string           `plist:"PayloadIdentifier"`
	PayloadType        string           `plist:"PayloadType"`
	PayloadUUID        string           `plist:"PayloadUUID"`
	PayloadVersion     int              `plist:"PayloadVersion"`
}

// PayloadContent represents the content of a payload.
type PayloadContent struct {
	PayloadContent    map[string]map[string]interface{} `plist:"PayloadContent,omitempty"`
	PayloadIdentifier string                            `plist:"PayloadIdentifier"`
	PayloadType       string                            `plist:"PayloadType"`
	PayloadUUID       string                            `plist:"PayloadUUID"`
	PayloadVersion    int                               `plist:"PayloadVersion"`
}

// ProfileAvailability represents the availability of the profile.
type ProfileAvailability struct {
	DeviceChannel string `plist:"DeviceChannel"`
	MacOS         string `plist:"macOS"`
	UserChannel   string `plist:"UserChannel"`
}

// ProfileExample represents an example of the profile.
type ProfileExample struct {
	XMLName xml.Name `xml:"plist"`
	Dict    Dict     `plist:"dict"`
}

// Dict represents a dictionary in XML.
type Dict struct {
	PayloadContent     []PayloadContent `plist:"PayloadContent"`
	PayloadDisplayName string           `plist:"PayloadDisplayName"`
	PayloadIdentifier  string           `plist:"PayloadIdentifier"`
	PayloadType        string           `plist:"PayloadType"`
	PayloadUUID        string           `plist:"PayloadUUID"`
	PayloadVersion     int              `plist:"PayloadVersion"`
}

// ManagedPreferencesPreferenceDomain represents the dictionary containing app preference domains.
type ManagedPreferencesPreferenceDomain struct {
	Forced map[string]interface{} `plist:"Forced"`
}

// ManagedPreferencesPreferenceDomainSettings represents the dictionary of domain settings.
type ManagedPreferencesPreferenceDomainSettings struct {
	McxPreferenceSettings map[string]interface{} `plist:"mcx_preference_settings"`
}

// DeviceManagementProfile represents the device management profile.
type DeviceManagementProfile struct {
	ManagedPreferences                                     ManagedPreferences                         `plist:"ManagedPreferences"`
	ManagedPreferencesPrefDom                              ManagedPreferencesPreferenceDomain         `plist:"ManagedPreferences.PreferenceDomain"`
	ProfileAvailability                                    ProfileAvailability                        `plist:"ProfileAvailability"`
	AllowManualInstall                                     string                                     `plist:"AllowManualInstall"`
	RequiresSupervision                                    string                                     `plist:"RequiresSupervision"`
	RequiresUserApprovedMDM                                string                                     `plist:"RequiresUserApprovedMDM"`
	AllowedInUserEnrollment                                string                                     `plist:"AllowedInUserEnrollment"`
	AllowMultiplePayloads                                  string                                     `plist:"AllowMultiplePayloads"`
	ProfileExample                                         ProfileExample                             `plist:"ProfileExample"`
	ManagedPreferencesPrefDomSettings                      ManagedPreferencesPreferenceDomainSettings `plist:"ManagedPreferences.PreferenceDomain.Settings"`
	ManagedPreferencesPrefDomSettingsMcxPreferenceSettings map[string]interface{}                     `plist:"ManagedPreferences.PreferenceDomain.Settings.Mcx_preference_settings"`
}
