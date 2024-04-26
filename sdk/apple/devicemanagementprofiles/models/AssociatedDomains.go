package models

import "encoding/xml"

// ResourceAssociatedDomainsConfigurationProfile represents the top-level structure of the plist for Associated Domains configurations
type ResourceAssociatedDomainsConfigurationProfile struct {
	XMLName                  xml.Name                                           `xml:"plist"`
	Version                  string                                             `xml:"version,attr"`
	Dict                     AssociatedDomainsConfigurationProfileSubsetPayload `xml:"dict"`
	PayloadDescription       string                                             `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                             `xml:"PayloadDisplayName,omitempty"`
	PayloadEnabled           string                                             `xml:"PayloadEnabled,omitempty"`
	PayloadIdentifier        string                                             `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                             `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                             `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                             `xml:"PayloadScope,omitempty"`
	PayloadType              string                                             `xml:"PayloadType,omitempty"`
	PayloadUUID              string                                             `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                                             `xml:"PayloadVersion,omitempty"`
}

// AssociatedDomainsConfigurationProfileSubsetPayload represents the content structure for configuring associated domains
type AssociatedDomainsConfigurationProfileSubsetPayload struct {
	Configuration     []AssociatedDomainsConfigurationItem `xml:"Configuration>AssociatedDomainsConfigurationItem"`
	PayloadIdentifier string                               `xml:"PayloadIdentifier,omitempty"`
	PayloadType       string                               `xml:"PayloadType,omitempty"`
	PayloadUUID       string                               `xml:"PayloadUUID,omitempty"`
	PayloadVersion    int                                  `xml:"PayloadVersion,omitempty"`
}

// AssociatedDomainsConfigurationItem represents a mapping of apps to their associated domains
type AssociatedDomainsConfigurationItem struct {
	ApplicationIdentifier string   `xml:"ApplicationIdentifier"`
	AssociatedDomains     []string `xml:"AssociatedDomains>string"`
	EnableDirectDownloads bool     `xml:"EnableDirectDownloads,omitempty"`
}
