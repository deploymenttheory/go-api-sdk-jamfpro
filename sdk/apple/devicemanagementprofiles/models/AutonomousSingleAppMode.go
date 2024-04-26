package models

import "encoding/xml"

// ResourceAutonomousSingleAppModeConfigurationProfile represents the top-level structure of the plist for Autonomous Single App Mode configurations
type ResourceAutonomousSingleAppModeConfigurationProfile struct {
	XMLName                  xml.Name                                                 `xml:"plist"`
	Version                  string                                                   `xml:"version,attr"`
	Dict                     AutonomousSingleAppModeConfigurationProfileSubsetPayload `xml:"dict"`
	PayloadDescription       string                                                   `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                                   `xml:"PayloadDisplayName,omitempty"`
	PayloadEnabled           string                                                   `xml:"PayloadEnabled,omitempty"`
	PayloadIdentifier        string                                                   `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                                   `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                                   `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                                   `xml:"PayloadScope,omitempty"`
	PayloadType              string                                                   `xml:"PayloadType,omitempty"`
	PayloadUUID              string                                                   `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                                                   `xml:"PayloadVersion,omitempty"`
}

// AutonomousSingleAppModeConfigurationProfileSubsetPayload represents the content structure for configuring Autonomous Single App Mode settings
type AutonomousSingleAppModeConfigurationProfileSubsetPayload struct {
	AllowedApplications []AutonomousSingleAppModeAllowedApplicationsItem `xml:"AllowedApplications>AutonomousSingleAppModeAllowedApplicationsItem"`
	PayloadIdentifier   string                                           `xml:"PayloadIdentifier,omitempty"`
	PayloadType         string                                           `xml:"PayloadType,omitempty"`
	PayloadUUID         string                                           `xml:"PayloadUUID,omitempty"`
	PayloadVersion      int                                              `xml:"PayloadVersion,omitempty"`
}

// AutonomousSingleAppModeAllowedApplicationsItem represents a dictionary that specifies an app granted access to Accessibility APIs
type AutonomousSingleAppModeAllowedApplicationsItem struct {
	BundleIdentifier string `xml:"BundleIdentifier"`
	TeamIdentifier   string `xml:"TeamIdentifier"`
}
