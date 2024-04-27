/*
Device Management Profile: Passcode
Version: 	iOS 4.0+

	iPadOS 4.0+
	macOS 10.7+
	watchOS 10.0+

Ref: 			https://developer.apple.com/documentation/devicemanagement/passcode
Profile Example:

<?xml version=”1.0” encoding=”UTF-8”?>
<!DOCTYPE plist PUBLIC “-//Apple//DTD PLIST 1.0//EN” “http://www.apple.com/DTDs/PropertyList-1.0.dtd”>
<plist version=”1.0”>
<dict>

	<key>PayloadContent</key>
	<array>
	    <dict>
	        <key>allowSimple</key>
	        <true/>
	        <key>forcePIN</key>
	        <true/>
	        <key>maxFailedAttempts</key>
	        <integer>5</integer>
	        <key>maxGracePeriod</key>
	        <integer>1</integer>
	        <key>maxInactivity</key>
	        <integer>2</integer>
	        <key>maxPINAgeInDays</key>
	        <real>30</real>
	        <key>minLength</key>
	        <integer>8</integer>
	        <key>pinHistory</key>
	        <real>2</real>
	        <key>requireAlphanumeric</key>
	        <false/>
	        <key>PayloadIdentifier</key>
	        <string>com.example.mypasscodepayload</string>
	        <key>PayloadType</key>
	        <string>com.apple.mobiledevice.passwordpolicy</string>
	        <key>PayloadUUID</key>
	        <string>2a8a75e5-d17d-44d5-b062-3cb92161af9f</string>
	        <key>PayloadVersion</key>
	        <integer>1</integer>
	    </dict>
	</array>
	<key>PayloadDisplayName</key>
	<string>Passcode</string>
	<key>PayloadIdentifier</key>
	<string>com.example.myprofile</string>
	<key>PayloadType</key>
	<string>Configuration</string>
	<key>PayloadUUID</key>
	<string>e044f50d-ff67-4bcd-9f3f-d7b678091061</string>
	<key>PayloadVersion</key>
	<integer>1</integer>

</dict>
</plist>
*/
package models

// ResourcePasscodeConfigurationProfile defines the structure for managing passcode policies on devices.
type ResourcePasscodeConfigurationProfile struct {
	Version                  string                                      `plist:"version,attr"`
	PayloadContent           []PasscodeConfigurationProfileSubsetPayload `plist:"PayloadContent"`
	PayloadDescription       string                                      `plist:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                      `plist:"PayloadDisplayName,omitempty"`
	PayloadEnabled           string                                      `plist:"PayloadEnabled,omitempty"`
	PayloadIdentifier        string                                      `plist:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                      `plist:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                      `plist:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                      `plist:"PayloadScope,omitempty"`
	PayloadType              string                                      `plist:"PayloadType,omitempty"`
	PayloadUUID              string                                      `plist:"PayloadUUID,omitempty"`
	PayloadVersion           int                                         `plist:"PayloadVersion,omitempty"`
}

// PasscodeConfigurationProfileSubsetPayload represents the passcode requirements set in the policy.
type PasscodeConfigurationProfileSubsetPayload struct {
	AllowSimple         *bool        `plist:"allowSimple"`
	ForcePIN            *bool        `plist:"forcePIN"`
	MaxFailedAttempts   int          `plist:"maxFailedAttempts"`
	MaxGracePeriod      int          `plist:"maxGracePeriod"`
	MaxInactivity       int          `plist:"maxInactivity"`
	MaxPINAgeInDays     int          `plist:"maxPINAgeInDays"`
	MinLength           int          `plist:"minLength"`
	PinHistory          int          `plist:"pinHistory"`
	RequireAlphanumeric *bool        `plist:"requireAlphanumeric"`
	PayloadIdentifier   string       `plist:"PayloadIdentifier"`
	PayloadType         string       `plist:"PayloadType"`
	PayloadUUID         string       `plist:"PayloadUUID"`
	PayloadVersion      int          `plist:"PayloadVersion"`
	CustomRegex         *CustomRegex `plist:"CustomRegex,omitempty"`
}

// CustomRegex defines the regular expression for password compliance.
type CustomRegex struct {
	PasswordContentDescriptions []Description `plist:"PasswordContentDescription"`
	PasswordContentRegex        string        `plist:"PasswordContentRegex"`
}

// CustomRegex defines the regular expression for password compliance.
// Description represents a single localized description entry.
type Description struct {
	Locale      string `plist:"locale,attr"`
	Description string `plist:",chardata"`
}
