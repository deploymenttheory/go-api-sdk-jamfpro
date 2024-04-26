package models

import "encoding/xml"

// ResourceAppLockConfigurationProfile represents the top-level structure of the plist for App Lock configurations
type ResourceAppLockConfigurationProfile struct {
	XMLName                  xml.Name                                 `xml:"plist"`
	Version                  string                                   `xml:"version,attr"`
	Dict                     AppLockConfigurationProfileSubsetPayload `xml:"dict"`
	PayloadDescription       string                                   `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                   `xml:"PayloadDisplayName,omitempty"`
	PayloadEnabled           string                                   `xml:"PayloadEnabled,omitempty"`
	PayloadIdentifier        string                                   `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                   `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                   `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                   `xml:"PayloadScope,omitempty"`
	PayloadType              string                                   `xml:"PayloadType,omitempty"`
	PayloadUUID              string                                   `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                                   `xml:"PayloadVersion,omitempty"`
}

// AppLockConfigurationProfileSubsetPayload represents the content structure for configuring App Lock settings
type AppLockConfigurationProfileSubsetPayload struct {
	App               AppLockApp `xml:"App"`
	PayloadIdentifier string     `xml:"PayloadIdentifier,omitempty"`
	PayloadType       string     `xml:"PayloadType,omitempty"`
	PayloadUUID       string     `xml:"PayloadUUID,omitempty"`
	PayloadVersion    int        `xml:"PayloadVersion,omitempty"`
}

// AppLockApp represents the app-specific information for App Lock
type AppLockApp struct {
	Identifier         string                       `xml:"Identifier"`
	Options            AppLockAppOptions            `xml:"Options,omitempty"`
	UserEnabledOptions AppLockAppUserEnabledOptions `xml:"UserEnabledOptions,omitempty"`
}

// AppLockAppOptions represents the dictionary of options to set for the app
type AppLockAppOptions struct {
	DisableAutoLock        bool `xml:"DisableAutoLock,omitempty"`
	DisableDeviceRotation  bool `xml:"DisableDeviceRotation,omitempty"`
	DisableRingerSwitch    bool `xml:"DisableRingerSwitch,omitempty"`
	DisableSleepWakeButton bool `xml:"DisableSleepWakeButton,omitempty"`
	DisableTouch           bool `xml:"DisableTouch,omitempty"`
	DisableVolumeButtons   bool `xml:"DisableVolumeButtons,omitempty"`
	EnableAssistiveTouch   bool `xml:"EnableAssistiveTouch,omitempty"`
	EnableInvertColors     bool `xml:"EnableInvertColors,omitempty"`
	EnableMonoAudio        bool `xml:"EnableMonoAudio,omitempty"`
	EnableSpeakSelection   bool `xml:"EnableSpeakSelection,omitempty"`
	EnableVoiceOver        bool `xml:"EnableVoiceOver,omitempty"`
	EnableZoom             bool `xml:"EnableZoom,omitempty"`
	EnableVoiceControl     bool `xml:"EnableVoiceControl,omitempty"`
}

// AppLockAppUserEnabledOptions represents the dictionary of user-editable options for the app
type AppLockAppUserEnabledOptions struct {
	AssistiveTouch bool `xml:"AssistiveTouch,omitempty"`
	InvertColors   bool `xml:"InvertColors,omitempty"`
	VoiceOver      bool `xml:"VoiceOver,omitempty"`
	Zoom           bool `xml:"Zoom,omitempty"`
	VoiceControl   bool `xml:"VoiceControl,omitempty"`
}
