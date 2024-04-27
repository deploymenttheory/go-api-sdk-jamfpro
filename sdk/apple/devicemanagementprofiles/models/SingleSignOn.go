package models

import "encoding/xml"

// ResourceSingleSignOnProfile represents the top-level structure of the plist for Single Sign-On configurations
type ResourceSingleSignOnProfile struct {
	XMLName             xml.Name                         `xml:"plist"`
	Version             string                           `xml:"version,attr"`
	Payload             SingleSignOnProfileSubsetPayload `xml:"dict>array>dict"`
	PayloadDescription  string                           `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName  string                           `xml:"PayloadDisplayName,omitempty"`
	PayloadEnabled      string                           `xml:"PayloadEnabled,omitempty"`
	PayloadIdentifier   string                           `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization string                           `xml:"PayloadOrganization,omitempty"`
	PayloadScope        string                           `xml:"PayloadScope,omitempty"`
	PayloadType         string                           `xml:"PayloadType,omitempty"`
	PayloadUUID         string                           `xml:"PayloadUUID,omitempty"`
	PayloadVersion      string                           `xml:"PayloadVersion,omitempty"`
}

// SingleSignOnProfileSubsetPayload represents the content structure for configuring Single Sign-On settings
type SingleSignOnProfileSubsetPayload struct {
	PayloadContent         []SingleSignOnKerberos `xml:"PayloadContent>dict"`
	PayloadDisplayName     string                 `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier      string                 `xml:"PayloadIdentifier,omitempty"`
	PayloadType            string                 `xml:"PayloadType,omitempty"`
	PayloadUUID            string                 `xml:"PayloadUUID,omitempty"`
	PayloadVersion         int                    `xml:"PayloadVersion,omitempty"`
	PayloadCertificateUUID string                 `xml:"PayloadCertificateUUID,omitempty"`
	PrincipalName          string                 `xml:"PrincipalName,omitempty"`
	Realm                  string                 `xml:"Realm,omitempty"`
	URLPrefixMatches       []string               `xml:"URLPrefixMatches>string,omitempty"`
	AppIdentifierMatches   []string               `xml:"AppIdentifierMatches>string,omitempty"`
}

// SingleSignOnKerberos represents the Kerberos dictionary within Single Sign-On configuration
type SingleSignOnKerberos struct {
	ExtensionData        map[string]interface{} `xml:"ExtensionData>dict,omitempty"`
	ExtensionIdentifier  string                 `xml:"ExtensionIdentifier,omitempty"`
	TeamIdentifier       string                 `xml:"TeamIdentifier,omitempty"`
	Hosts                []string               `xml:"Hosts>string,omitempty"`
	Realm                string                 `xml:"Realm,omitempty"`
	Type                 string                 `xml:"Type,omitempty"`
	PayloadIdentifier    string                 `xml:"PayloadIdentifier,omitempty"`
	PayloadType          string                 `xml:"PayloadType,omitempty"`
	PayloadUUID          string                 `xml:"PayloadUUID,omitempty"`
	PayloadVersion       int                    `xml:"PayloadVersion,omitempty"`
	UseSiteAutoDiscovery bool                   `xml:"ExtensionData>useSiteAutoDiscovery,omitempty"`
}
