/*
Device:		Management Profile: LightsOutManagementLOM
Version: 	macOS 11.0+
Ref: 			https://developer.apple.com/documentation/devicemanagement/lightsoutmanagementlom
Profile Example:
<?xml version=”1.0” encoding=”UTF-8”?>
<!DOCTYPE plist PUBLIC “-//Apple//DTD PLIST 1.0//EN” “http://www.apple.com/DTDs/PropertyList-1.0.dtd”>
<plist version=”1.0”>
<dict>
    <key>PayloadContent</key>
    <array>
        <dict>
            <key>ControllerCACertificateUUIDs</key>
            <array>
                <string>D99D951C-60CC-454D-B293-AF74ADB38738</string>
            </array>
            <key>ControllerCertificateUUID</key>
            <string>6B9162FB-8B4A-4241-8C6E-2A2E80E62CA2</string>
            <key>DeviceCACertificateUUIDs</key>
            <array>
                <string>AF8A74BE-5B3E-47DC-83CD-279FC75DD63E</string>
                <string>4EC901FA-03F2-4374-A4BF-4B0BF8D249E2</string>
            </array>
            <key>PayloadDisplayName</key>
            <string>LOM Enrollment</string>
            <key>PayloadIdentifier</key>
            <string>com.apple.controller1.pkcs12.lom</string>
            <key>PayloadType</key>
            <string>com.apple.lom</string>
            <key>PayloadUUID</key>
            <string>79899C98-DA98-4A4D-8136-25E8CD33AE82</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
        <dict>
            <key>PayloadContent</key>
            <data>
            MIIDLN1R
            </data>
            <key>PayloadDisplayName</key>
            <string>Root CA Cert</string>
            <key>PayloadIdentifier</key>
            <string>com.apple.controller1.pkcs1.ca.root.lom</string>
            <key>PayloadType</key>
            <string>com.apple.security.pkcs1</string>
            <key>PayloadUUID</key>
            <string>D99D951C-60CC-454D-B293-AF74ADB38738</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
        <dict>
            <key>PayloadContent</key>
            <data>
            MIID8yw0
            </data>
            <key>PayloadDisplayName</key>
            <string>PKCS1 Root CA Cert</string>
            <key>PayloadIdentifier</key>
            <string>com.apple.device1.pkcs1.ca.root.lom</string>
            <key>PayloadType</key>
            <string>com.apple.security.pkcs1</string>
            <key>PayloadUUID</key>
            <string>AF8A74BE-5B3E-47DC-83CD-279FC75DD63E</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
        <dict>
            <key>PayloadContent</key>
            <data>
            MIIGEcCL
            </data>
            <key>PayloadDescription</key>
            <string>PKCS1Certificate</string>
            <key>PayloadDisplayName</key>
            <string>Intermediate CA Cert</string>
            <key>PayloadIdentifier</key>
            <string>com.apple.device1.pkcs1.ca.intermediate.lom</string>
            <key>PayloadOrganization</key>
            <string>Test</string>
            <key>PayloadType</key>
            <string>com.apple.security.pkcs1</string>
            <key>PayloadUUID</key>
            <string>4EC901FA-03F2-4374-A4BF-4B0BF8D249E2</string>
            <key>PayloadVersion</key>
            <integer>1</integer>
        </dict>
        <dict>
            <key>Password</key>
            <string>test</string>
            <key>PayloadCertificateFileName</key>
            <string>LOM_Test_Client.p12</string>
            <key>PayloadContent</key>
            <data>
            MIIKFAQ==
            </data>
            <key>PayloadDescription</key>
            <string>Identity from LOM_Test_Client.p12</string>
            <key>PayloadDisplayName</key>
            <string>AppleLOM_Test_controller</string>
            <key>PayloadIdentifier</key>
            <string>test.cert.LOM_Test_Client.p12.1</string>
            <key>PayloadType</key>
            <string>com.apple.security.pkcs12</string>
            <key>PayloadUUID</key>
            <string>6B9162FB-8B4A-4241-8C6E-2A2E80E62CA2</string>
        </dict>
    </array>
    <key>PayloadDisplayName</key>
    <string>Lights Out Management Controller</string>
    <key>PayloadIdentifier</key>
    <string>com.example.myprofile</string>
    <key>PayloadType</key>
    <string>Configuration</string>
    <key>PayloadUUID</key>
    <string>60f25f97-01ec-4016-a4d0-da42ae91a81d</string>
    <key>PayloadVersion</key>
    <integer>1</integer>
</dict>
</plist>
*/

package models

import "encoding/xml"

// MarshalXML sets the PayloadType to a specific value during marshaling based on the location within the hierarchy.
func (r *ResourceLightsOutManagementLOMConfigurationProfile) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// Set the top-level PayloadType for the entire plist
	r.PayloadType = "Configuration"

	// Alias the type to avoid recursion
	type alias ResourceLightsOutManagementLOMConfigurationProfile
	return e.EncodeElement(alias(*r), start)
}

// LightsOutManagementLOM represents the structure of the plist for LightsOutManagementLOM settings
type ResourceLightsOutManagementLOMConfigurationProfile struct {
	XMLName                  xml.Name                                                  `xml:"plist"`
	Version                  string                                                    `xml:"version,attr"`
	PayloadContent           []LightsOutManagementLOMConfigurationProfileSubsetPayload `xml:"PayloadContent"`
	PayloadDescription       string                                                    `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                                    `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier        string                                                    `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                                    `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                                    `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                                    `xml:"PayloadScope,omitempty"`
	PayloadType              string                                                    `xml:"PayloadType,omitempty"`
	PayloadUUID              string                                                    `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                                                    `xml:"PayloadVersion,omitempty"`
}

// LightsOutManagementLOMConfigurationProfileSubsetPayload represents the content structure for LightsOutManagementLOM payload
type LightsOutManagementLOMConfigurationProfileSubsetPayload struct {
	ControllerCACertificateUUIDs []string `xml:"ControllerCACertificateUUIDs>string,omitempty"`
	ControllerCertificateUUID    string   `xml:"ControllerCertificateUUID,omitempty"`
	DeviceCACertificateUUIDs     []string `xml:"DeviceCACertificateUUIDs>string,omitempty"`
	DeviceCertificateUUID        string   `xml:"DeviceCertificateUUID,omitempty"`
	PayloadDisplayName           string   `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier            string   `xml:"PayloadIdentifier,omitempty"`
	PayloadType                  string   `xml:"PayloadType,omitempty"`
	PayloadUUID                  string   `xml:"PayloadUUID,omitempty"`
	PayloadVersion               int      `xml:"PayloadVersion,omitempty"`
	Password                     string   `xml:"Password,omitempty"`
	PayloadCertificateFileName   string   `xml:"PayloadCertificateFileName,omitempty"`
	PayloadContent               string   `xml:"PayloadContent,omitempty"`
}
