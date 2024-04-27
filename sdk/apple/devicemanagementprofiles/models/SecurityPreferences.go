/*
Device Management Profile: Security Preferences
Version: 	macOS 10.10+

Ref: 			https://developer.apple.com/documentation/devicemanagement/securitypreferences

Profile Example:
<?xml version=”1.0” encoding=”UTF-8”?>
<!DOCTYPE plist PUBLIC “-//Apple//DTD PLIST 1.0//EN” “http://www.apple.com/DTDs/PropertyList-1.0.dtd”>
<plist version=”1.0”>
<dict>

	<key>PayloadContent</key>
	<array>
	    <dict>
	        <key>dontAllowFireWallUI</key>
	        <true/>
	        <key>PayloadIdentifier</key>
	        <string>com.example.mysecuritypreferencespayload</string>
	        <key>PayloadType</key>
	        <string>com.apple.preference.security</string>
	        <key>PayloadUUID</key>
	        <string>d99bb019-a61d-447f-8fed-8f223cc56be3</string>
	        <key>PayloadVersion</key>
	        <integer>1</integer>
	    </dict>
	</array>
	<key>PayloadDisplayName</key>
	<string>Security Preferences</string>
	<key>PayloadIdentifier</key>
	<string>com.example.myprofile</string>
	<key>PayloadType</key>
	<string>Configuration</string>
	<key>PayloadUUID</key>
	<string>b44b6a04-6527-4333-87e5-46422e8a5844</string>
	<key>PayloadVersion</key>
	<integer>1</integer>

</dict>
</plist>
*/
package models

// ResourceSecurityPreferencesConfigurationProfile defines the structure for managing passcode policies on devices.
type ResourceSecurityPreferencesConfigurationProfile struct {
	Version                  string                                                 `plist:"version,attr"`
	PayloadContent           []SecurityPreferencesConfigurationProfileSubsetPayload `plist:"PayloadContent"`
	PayloadDescription       string                                                 `plist:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                                 `plist:"PayloadDisplayName,omitempty"`
	PayloadEnabled           string                                                 `plist:"PayloadEnabled,omitempty"`
	PayloadIdentifier        string                                                 `plist:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                                 `plist:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                                 `plist:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                                 `plist:"PayloadScope,omitempty"`
	PayloadType              string                                                 `plist:"PayloadType,omitempty"`
	PayloadUUID              string                                                 `plist:"PayloadUUID,omitempty"`
	PayloadVersion           int                                                    `plist:"PayloadVersion,omitempty"`
}

// SecurityPreferencesConfigurationProfileSubsetPayload represents the specific security preferences set in the policy.
type SecurityPreferencesConfigurationProfileSubsetPayload struct {
	DontAllowFireWallUI      *bool  `plist:"dontAllowFireWallUI"`
	DontAllowLockMessageUI   *bool  `plist:"dontAllowLockMessageUI"`
	DontAllowPasswordResetUI *bool  `plist:"dontAllowPasswordResetUI"`
	PayloadIdentifier        string `plist:"PayloadIdentifier"`
	PayloadType              string `plist:"PayloadType"`
	PayloadUUID              string `plist:"PayloadUUID"`
	PayloadVersion           int    `plist:"PayloadVersion"`
}
