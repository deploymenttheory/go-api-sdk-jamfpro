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

import "encoding/xml"

// ResourcePasscodeConfigurationProfile defines the structure for managing passcode policies on devices.
type ResourcePasscodeConfigurationProfile struct {
	XMLName                  xml.Name                                  `xml:"plist"`
	Version                  string                                    `xml:"version,attr"`
	Payload                  PasscodeConfigurationProfileSubsetPayload `xml:"dict>array>dict"`
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

// PasscodeConfigurationProfileSubsetPayload represents the passcode requirements set in the policy.
type PasscodeConfigurationProfileSubsetPayload struct {
	AllowSimple         bool         `xml:"allowSimple"`
	ForcePIN            bool         `xml:"forcePIN"`
	MaxFailedAttempts   int          `xml:"maxFailedAttempts"`
	MaxGracePeriod      int          `xml:"maxGracePeriod"`
	MaxInactivity       int          `xml:"maxInactivity"`
	MaxPINAgeInDays     float64      `xml:"maxPINAgeInDays"`
	MinLength           int          `xml:"minLength"`
	PinHistory          float64      `xml:"pinHistory"`
	RequireAlphanumeric bool         `xml:"requireAlphanumeric"`
	PayloadIdentifier   string       `xml:"PayloadIdentifier"`
	PayloadType         string       `xml:"PayloadType"`
	PayloadUUID         string       `xml:"PayloadUUID"`
	PayloadVersion      int          `xml:"PayloadVersion"`
	CustomRegex         *CustomRegex `xml:"CustomRegex,omitempty"`
}

// CustomRegex defines the regular expression for password compliance.
type CustomRegex struct {
	PasswordContentDescription map[string]string `xml:"PasswordContentDescription"`
	PasswordContentRegex       string            `xml:"PasswordContentRegex"`
}
