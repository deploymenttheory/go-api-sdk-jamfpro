/*
Device:		Management Profile: LightsOutManagementLOM
Version: 	macOS 11.0+
Ref: 			https://developer.apple.com/documentation/devicemanagement/accounts
Profile Example:
<?xml version=”1.0” encoding=”UTF-8”?>
<!DOCTYPE plist PUBLIC “-//Apple//DTD PLIST 1.0//EN” “http://www.apple.com/DTDs/PropertyList-1.0.dtd”>
<plist version=”1.0”>
<dict>

	<key>PayloadContent</key>
	<array>
	    <dict>
	        <key>EnableGuestAccount</key>
	        <true/>
	        <key>PayloadIdentifier</key>
	        <string>com.example.myaccountpayload</string>
	        <key>PayloadType</key>
	        <string>com.apple.MCX</string>
	        <key>PayloadUUID</key>
	        <string>5d4e377c-108c-44af-a46e-97a5aac1e270</string>
	        <key>PayloadVersion</key>
	        <integer>1</integer>
	    </dict>
	</array>
	<key>PayloadDisplayName</key>
	<string>Accounts</string>
	<key>PayloadIdentifier</key>
	<string>com.example.myprofile</string>
	<key>PayloadType</key>
	<string>Configuration</string>
	<key>PayloadUUID</key>
	<string>8cd28a9d-625e-4056-bbd0-43617bb8efb7</string>
	<key>PayloadVersion</key>
	<integer>1</integer>

</dict>
</plist>
*/
package models

import "encoding/xml"

// MarshalXML sets the PayloadType to a specific value during marshaling based on the location within the hierarchy.
func (r *ResourceAccountsConfigurationProfile) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// Set the top-level PayloadType for the entire plist
	r.PayloadType = "Configuration"

	// Alias the type to avoid recursion
	type alias ResourceAccountsConfigurationProfile
	return e.EncodeElement(alias(*r), start)
}

// ResourceAccountsConfigurationProfile represents the top-level structure of the plist
type ResourceAccountsConfigurationProfile struct {
	XMLName                  xml.Name                                  `xml:"plist"`
	Version                  string                                    `xml:"version,attr"`
	Payload                  AccountsConfigurationProfileSubsetPayload `xml:"dict>array>dict"`
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

// AccountsConfigurationProfileSubsetPayload represents the content structure for configuring guest accounts
type AccountsConfigurationProfileSubsetPayload struct {
	EnableGuestAccount bool   `xml:"EnableGuestAccount,omitempty"`
	PayloadIdentifier  string `xml:"PayloadIdentifier,omitempty"`
	PayloadType        string `xml:"PayloadType,omitempty"`
	PayloadUUID        string `xml:"PayloadUUID,omitempty"`
	PayloadVersion     int    `xml:"PayloadVersion,omitempty"`
}
