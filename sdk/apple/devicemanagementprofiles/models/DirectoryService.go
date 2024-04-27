package models

import "encoding/xml"

// ResourceDirectoryServiceConfigurationProfile represents the top-level structure of the plist for Directory Service configurations
type ResourceDirectoryServiceConfigurationProfile struct {
	XMLName                  xml.Name                                          `xml:"plist"`
	Version                  string                                            `xml:"version,attr"`
	Payload                  DirectoryServiceConfigurationProfileSubsetPayload `xml:"dict"`
	PayloadDescription       string                                            `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                            `xml:"PayloadDisplayName,omitempty"`
	PayloadEnabled           string                                            `xml:"PayloadEnabled,omitempty"`
	PayloadIdentifier        string                                            `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                            `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                            `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                            `xml:"PayloadScope,omitempty"`
	PayloadType              string                                            `xml:"PayloadType,omitempty"`
	PayloadUUID              string                                            `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                                            `xml:"PayloadVersion,omitempty"`
}

// DirectoryServiceConfigurationProfileSubsetPayload represents the content structure for configuring Directory Service settings
type DirectoryServiceConfigurationProfileSubsetPayload struct {
	ADAllowMultiDomainAuth        bool     `xml:"ADAllowMultiDomainAuth"`
	ADCreateMobileAccountAtLogin  bool     `xml:"ADCreateMobileAccountAtLogin"`
	ADDefaultUserShell            string   `xml:"ADDefaultUserShell,omitempty"`
	ADDomainAdminGroupList        []string `xml:"ADDomainAdminGroupList>string,omitempty"`
	ADForceHomeLocal              bool     `xml:"ADForceHomeLocal"`
	ADMapGGIDAttribute            string   `xml:"ADMapGGIDAttribute,omitempty"`
	ADMapGIDAttribute             string   `xml:"ADMapGIDAttribute,omitempty"`
	ADMapUIDAttribute             string   `xml:"ADMapUIDAttribute,omitempty"`
	ADMountStyle                  string   `xml:"ADMountStyle,omitempty"`
	ADNamespace                   string   `xml:"ADNamespace,omitempty"`
	ADOrganizationalUnit          string   `xml:"ADOrganizationalUnit,omitempty"`
	ADPacketEncrypt               string   `xml:"ADPacketEncrypt,omitempty"`
	ADPacketSign                  string   `xml:"ADPacketSign,omitempty"`
	ADPreferredDCServer           string   `xml:"ADPreferredDCServer,omitempty"`
	ADRestrictDDNS                []string `xml:"ADRestrictDDNS>string,omitempty"`
	ADTrustChangePassIntervalDays int      `xml:"ADTrustChangePassIntervalDays"`
	ADUseWindowsUNCPath           bool     `xml:"ADUseWindowsUNCPath"`
	ADWarnUserBeforeCreatingMA    bool     `xml:"ADWarnUserBeforeCreatingMA"`
	ClientID                      string   `xml:"ClientID,omitempty"`
	Description                   string   `xml:"Description,omitempty"`
	HostName                      string   `xml:"HostName"`
	Password                      string   `xml:"Password"`
	UserName                      string   `xml:"UserName"`
	PayloadIdentifier             string   `xml:"PayloadIdentifier,omitempty"`
	PayloadType                   string   `xml:"PayloadType,omitempty"`
	PayloadUUID                   string   `xml:"PayloadUUID,omitempty"`
	PayloadVersion                int      `xml:"PayloadVersion,omitempty"`
}
