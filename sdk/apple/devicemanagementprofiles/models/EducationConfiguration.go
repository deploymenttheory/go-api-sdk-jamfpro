package models

import "encoding/xml"

// ResourceEducationConfigurationProfile represents the top-level structure of the plist for configuring Education Configuration
type ResourceEducationConfigurationProfile struct {
	XMLName                  xml.Name                                   `xml:"plist"`
	Version                  string                                     `xml:"version,attr"`
	Dict                     EducationConfigurationProfileSubsetPayload `xml:"dict"`
	PayloadDescription       string                                     `xml:"PayloadDescription,omitempty"`
	PayloadDisplayName       string                                     `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier        string                                     `xml:"PayloadIdentifier,omitempty"`
	PayloadOrganization      string                                     `xml:"PayloadOrganization,omitempty"`
	PayloadRemovalDisallowed string                                     `xml:"PayloadRemovalDisallowed,omitempty"`
	PayloadScope             string                                     `xml:"PayloadScope,omitempty"`
	PayloadType              string                                     `xml:"PayloadType,omitempty"`
	PayloadUUID              string                                     `xml:"PayloadUUID,omitempty"`
	PayloadVersion           string                                     `xml:"PayloadVersion,omitempty"`
}

// EducationConfigurationProfileSubsetPayload represents the content structure for configuring Education Configuration
type EducationConfigurationProfileSubsetPayload struct {
	PayloadContent     []EducationConfiguration `xml:"PayloadContent>array>dict"`
	PayloadDisplayName string                   `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier  string                   `xml:"PayloadIdentifier,omitempty"`
	PayloadType        string                   `xml:"PayloadType,omitempty"`
	PayloadUUID        string                   `xml:"PayloadUUID,omitempty"`
	PayloadVersion     int                      `xml:"PayloadVersion,omitempty"`
}

// EducationConfiguration represents an Education Configuration payload
type EducationConfiguration struct {
	Departments                                    []DepartmentItem  `xml:"Departments,omitempty"`
	DeviceGroups                                   []DeviceGroupItem `xml:"DeviceGroups,omitempty"`
	Groups                                         []GroupItem       `xml:"Groups,omitempty"`
	LeaderPayloadCertificateAnchorUUID             []string          `xml:"LeaderPayloadCertificateAnchorUUID,omitempty"`
	MemberPayloadCertificateAnchorUUID             []string          `xml:"MemberPayloadCertificateAnchorUUID,omitempty"`
	OrganizationName                               string            `xml:"OrganizationName,omitempty"`
	OrganizationUUID                               string            `xml:"OrganizationUUID,omitempty"`
	PayloadCertificateUUID                         string            `xml:"PayloadCertificateUUID,omitempty"`
	ResourcePayloadCertificateUUID                 string            `xml:"ResourcePayloadCertificateUUID,omitempty"`
	ScreenObservationPermissionModificationAllowed bool              `xml:"ScreenObservationPermissionModificationAllowed,omitempty"`
	UserIdentifier                                 string            `xml:"UserIdentifier,omitempty"`
	Users                                          []UserItem        `xml:"Users,omitempty"`
	PayloadDisplayName                             string            `xml:"PayloadDisplayName,omitempty"`
	PayloadIdentifier                              string            `xml:"PayloadIdentifier,omitempty"`
	PayloadType                                    string            `xml:"PayloadType,omitempty"`
	PayloadUUID                                    string            `xml:"PayloadUUID,omitempty"`
	PayloadVersion                                 int               `xml:"PayloadVersion,omitempty"`
}

// DepartmentItem represents a department in the organization
type DepartmentItem struct {
	GroupBeaconIDs []int  `xml:"GroupBeaconIDs>integer"`
	Name           string `xml:"Name,omitempty"`
}

// DeviceGroupItem represents a device group in the organization
type DeviceGroupItem struct {
	Identifier    string   `xml:"Identifier,omitempty"`
	Name          string   `xml:"Name,omitempty"`
	SerialNumbers []string `xml:"SerialNumbers,omitempty"`
}

// GroupItem represents a group in the organization
type GroupItem struct {
	BeaconID               int      `xml:"BeaconID,omitempty"`
	ConfigurationSource    string   `xml:"ConfigurationSource,omitempty"`
	Description            string   `xml:"Description,omitempty"`
	DeviceGroupIdentifiers []string `xml:"DeviceGroupIdentifiers,omitempty"`
	ImageURL               string   `xml:"ImageURL,omitempty"`
	LeaderIdentifiers      []string `xml:"LeaderIdentifiers,omitempty"`
	MemberIdentifiers      []string `xml:"MemberIdentifiers,omitempty"`
	Name                   string   `xml:"Name,omitempty"`
}

// UserItem represents a user in the organization
type UserItem struct {
	AppleID            string `xml:"AppleID,omitempty"`
	FamilyName         string `xml:"FamilyName,omitempty"`
	FullScreenImageURL string `xml:"FullScreenImageURL,omitempty"`
	GivenName          string `xml:"GivenName,omitempty"`
	Identifier         string `xml:"Identifier,omitempty"`
	ImageURL           string `xml:"ImageURL,omitempty"`
	Name               string `xml:"Name,omitempty"`
	PasscodeType       string `xml:"PasscodeType,omitempty"`
	PhoneticFamilyName string `xml:"PhoneticFamilyName,omitempty"`
	PhoneticGivenName  string `xml:"PhoneticGivenName,omitempty"`
}
