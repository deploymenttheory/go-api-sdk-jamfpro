package models

import "encoding/xml"

// ServiceManagementManagedLoginItemsConfigurationProfile represents the top-level structure of the plist for configuring managed login items
type ServiceManagementManagedLoginItemsConfigurationProfile struct {
	XMLName                  xml.Name                                                            `xml:"plist"`
	Version                  string                                                              `xml:"version,attr"`
	Dict                     ServiceManagementManagedLoginItemsConfigurationProfileSubsetPayload `xml:"dict"`
	PayloadDescription       string                                                              `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                                              `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier        string                                                              `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                                              `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                                              `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                                              `xml:"PayloadScope,omitempty"`
	PayloadType              string                                                              `xml:"PayloadType,omitempty"`
	PayloadUUID              string                                                              `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                                                              `xml:"PayloadVersion,omitempty"`
}

// ServiceManagementManagedLoginItemsConfigurationProfileSubsetPayload represents the content structure for configuring managed login items
type ServiceManagementManagedLoginItemsConfigurationProfileSubsetPayload struct {
	PayloadContent     []ServiceManagementManagedLoginItemsRule `xml:"PayloadContent>array>dict"`
	PayloadDisplayName string                                   `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier  string                                   `xml:"PayloadIdentifier,omitempty"`
	PayloadType        string                                   `xml:"PayloadType,omitempty"`
	PayloadUUID        string                                   `xml:"PayloadUUID,omitempty"`
	PayloadVersion     int                                      `xml:"PayloadVersion,omitempty"`
}

// ServiceManagementManagedLoginItemsRule represents a managed login item rule
type ServiceManagementManagedLoginItemsRule struct {
	Rules             []RuleItem `xml:"Rules>array>dict,omitempty"`
	PayloadIdentifier string     `xml:"PayloadIdentifier,omitempty"`
	PayloadType       string     `xml:"PayloadType,omitempty"`
	PayloadUUID       string     `xml:"PayloadUUID,omitempty"`
	PayloadVersion    int        `xml:"PayloadVersion,omitempty"`
}

// RuleItem represents a rule item
type RuleItem struct {
	Comment        string `xml:"Comment,omitempty"`
	RuleType       string `xml:"RuleType,omitempty"`
	RuleValue      string `xml:"RuleValue,omitempty"`
	TeamIdentifier string `xml:"TeamIdentifier,omitempty"`
}
