// jamfproapi_computer_prestages.go
// Jamf Pro Api - Computer Prestages
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-prestages-scope
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

const uriComputerPrestagesV2 = "/api/v2/computer-prestages"
const uriComputerPrestagesV3 = "/api/v3/computer-prestages"

// List

type ResponseComputerPrestagesList struct {
	TotalCount *int                       `json:"totalCount"`
	Results    []ResourceComputerPrestage `json:"results"`
}

// Responses

// ResponseDeviceScope represents the structure of the response for a specific computer prestage scope.
type ResponseDeviceScope struct {
	PrestageId  string                            `json:"prestageId"`
	Assignments []DeviceScopeSubsetAssignmentItem `json:"assignments"`
	VersionLock int                               `json:"versionLock"`
}

// AssignmentItem represents the structure of each assignment within the prestage scope.
type DeviceScopeSubsetAssignmentItem struct {
	SerialNumber   string `json:"serialNumber"`
	AssignmentDate string `json:"assignmentDate"`
	UserAssigned   string `json:"userAssigned"`
}

// ResponseComputerPrestageCreate represents the response structure for creating a building.
type ResponseComputerPrestageCreate struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// Resource

type ResourceComputerPrestage struct {
	ID                                string                                      `json:"id"`
	DisplayName                       string                                      `json:"displayName"`
	Mandatory                         bool                                        `json:"mandatory"`
	MDMRemovable                      bool                                        `json:"mdmRemovable"`
	SupportPhoneNumber                string                                      `json:"supportPhoneNumber,omitempty"`
	SupportEmailAddress               string                                      `json:"supportEmailAddress,omitempty"`
	Department                        string                                      `json:"department,omitempty"`
	DefaultPrestage                   bool                                        `json:"defaultPrestage"`
	EnrollmentSiteId                  string                                      `json:"enrollmentSiteId,omitempty"`
	KeepExistingSiteMembership        bool                                        `json:"keepExistingSiteMembership"`
	KeepExistingLocationInformation   bool                                        `json:"keepExistingLocationInformation"`
	RequireAuthentication             bool                                        `json:"requireAuthentication"`
	AuthenticationPrompt              string                                      `json:"authenticationPrompt,omitempty"`
	PreventActivationLock             bool                                        `json:"preventActivationLock"`
	EnableDeviceBasedActivationLock   bool                                        `json:"enableDeviceBasedActivationLock"`
	DeviceEnrollmentProgramInstanceId string                                      `json:"deviceEnrollmentProgramInstanceId,omitempty"`
	SkipSetupItems                    ComputerPrestageSubsetSkipSetupItems        `json:"skipSetupItems"`
	LocationInformation               ComputerPrestageSubsetLocationInformation   `json:"locationInformation"`
	PurchasingInformation             ComputerPrestageSubsetPurchasingInformation `json:"purchasingInformation"`
	AnchorCertificates                []string                                    `json:"anchorCertificates,omitempty"`
	EnrollmentCustomizationId         string                                      `json:"enrollmentCustomizationId,omitempty"`
	Language                          string                                      `json:"language,omitempty"`
	Region                            string                                      `json:"region,omitempty"`
	AutoAdvanceSetup                  bool                                        `json:"autoAdvanceSetup"`
	InstallProfilesDuringSetup        bool                                        `json:"installProfilesDuringSetup"`
	PrestageInstalledProfileIds       []string                                    `json:"prestageInstalledProfileIds,omitempty"`
	CustomPackageIds                  []string                                    `json:"customPackageIds,omitempty"`
	CustomPackageDistributionPointId  string                                      `json:"customPackageDistributionPointId,omitempty"`
	EnableRecoveryLock                bool                                        `json:"enableRecoveryLock"`
	RecoveryLockPasswordType          string                                      `json:"recoveryLockPasswordType,omitempty"`
	RecoveryLockPassword              string                                      `json:"recoveryLockPassword,omitempty"`
	RotateRecoveryLockPassword        bool                                        `json:"rotateRecoveryLockPassword"`
	ProfileUuid                       string                                      `json:"profileUuid,omitempty"`
	SiteId                            string                                      `json:"siteId,omitempty"`
	VersionLock                       int                                         `json:"versionLock,omitempty"`
	AccountSettings                   ComputerPrestageSubsetAccountSettings       `json:"accountSettings"`
}

// Subsets & Containers

type ComputerPrestageSubsetSkipSetupItems struct {
	Biometric         bool `json:"Biometric"`
	TermsOfAddress    bool `json:"TermsOfAddress"`
	FileVault         bool `json:"FileVault"`
	ICloudDiagnostics bool `json:"iCloudDiagnostics"`
	Diagnostics       bool `json:"Diagnostics"`
	Accessibility     bool `json:"Accessibility"`
	AppleID           bool `json:"AppleID"`
	ScreenTime        bool `json:"ScreenTime"`
	Siri              bool `json:"Siri"`
	DisplayTone       bool `json:"DisplayTone"`
	Restore           bool `json:"Restore"`
	Appearance        bool `json:"Appearance"`
	Privacy           bool `json:"Privacy"`
	Payment           bool `json:"Payment"`
	Registration      bool `json:"Registration"`
	TOS               bool `json:"TOS"`
	ICloudStorage     bool `json:"iCloudStorage"`
	Location          bool `json:"Location"`
}

type ComputerPrestageSubsetLocationInformation struct {
	Username     string `json:"username,omitempty"`
	Realname     string `json:"realname,omitempty"`
	Phone        string `json:"phone,omitempty"`
	Email        string `json:"email,omitempty"`
	Room         string `json:"room,omitempty"`
	Position     string `json:"position,omitempty"`
	DepartmentId string `json:"departmentId,omitempty"`
	BuildingId   string `json:"buildingId,omitempty"`
	ID           string `json:"id,omitempty"`
	VersionLock  int    `json:"versionLock,omitempty"`
}

type ComputerPrestageSubsetPurchasingInformation struct {
	ID                string `json:"id,omitempty"`
	Leased            bool   `json:"leased"`
	Purchased         bool   `json:"purchased"`
	AppleCareId       string `json:"appleCareId,omitempty"`
	PONumber          string `json:"poNumber,omitempty"`
	Vendor            string `json:"vendor,omitempty"`
	PurchasePrice     string `json:"purchasePrice,omitempty"`
	LifeExpectancy    int    `json:"lifeExpectancy,omitempty"`
	PurchasingAccount string `json:"purchasingAccount,omitempty"`
	PurchasingContact string `json:"purchasingContact,omitempty"`
	LeaseDate         string `json:"leaseDate,omitempty"`
	PODate            string `json:"poDate,omitempty"`
	WarrantyDate      string `json:"warrantyDate,omitempty"`
	VersionLock       int    `json:"versionLock,omitempty"`
}

type ComputerPrestageSubsetAccountSettings struct {
	ID                                      string `json:"id,omitempty"`
	PayloadConfigured                       bool   `json:"payloadConfigured"`
	LocalAdminAccountEnabled                bool   `json:"localAdminAccountEnabled"`
	AdminUsername                           string `json:"adminUsername,omitempty"`
	AdminPassword                           string `json:"adminPassword,omitempty"`
	HiddenAdminAccount                      bool   `json:"hiddenAdminAccount"`
	LocalUserManaged                        bool   `json:"localUserManaged"`
	UserAccountType                         string `json:"userAccountType,omitempty"`
	VersionLock                             int    `json:"versionLock,omitempty"`
	PrefillPrimaryAccountInfoFeatureEnabled bool   `json:"prefillPrimaryAccountInfoFeatureEnabled"`
	PrefillType                             string `json:"prefillType,omitempty"`
	PrefillAccountFullName                  string `json:"prefillAccountFullName,omitempty"`
	PrefillAccountUserName                  string `json:"prefillAccountUserName,omitempty"`
	PreventPrefillInfoFromModification      bool   `json:"preventPrefillInfoFromModification"`
}

// CRUD

// GetComputerPrestagesV3 retrieves all computer prestage information with optional sorting.
func (c *Client) GetComputerPrestages(sort_filter string) (*ResponseComputerPrestagesList, error) {
	resp, err := c.DoPaginatedGet(
		uriComputerPrestagesV3,
		standardPageSize,
		startingPageNumber,
		sort_filter,
	)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "computer prestages", err)
	}

	var out ResponseComputerPrestagesList
	out.TotalCount = &resp.Size
	for _, value := range resp.Results {
		var newObj ResourceComputerPrestage
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "computer prestages", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil
}

// GetComputerPrestageByID retrieves a specific computer prestage by its ID.
func (c *Client) GetComputerPrestageByID(id string) (*ResourceComputerPrestage, error) {
	endpoint := fmt.Sprintf("%s/%s", uriComputerPrestagesV3, id)

	var prestage ResourceComputerPrestage
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &prestage)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "computer prestage", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &prestage, nil
}

// GetComputerPrestageByName retrieves a specific computer prestage by its name.
func (c *Client) GetComputerPrestageByName(name string) (*ResourceComputerPrestage, error) {
	prestages, err := c.GetComputerPrestages("")
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "computer prestages", err)
	}

	for _, value := range prestages.Results {
		if value.DisplayName == name {
			return &value, nil
		}
	}

	return nil, fmt.Errorf(errMsgFailedGetByName, "computer prestage", name, errMsgNoName)
}

// CreateComputerPrestage creates a new computer prestage with the given details.
func (c *Client) CreateComputerPrestage(prestage *ResourceComputerPrestage) (*ResponseComputerPrestageCreate, error) {
	endpoint := uriComputerPrestagesV3

	var creationResponse ResponseComputerPrestageCreate
	resp, err := c.HTTP.DoRequest("POST", endpoint, prestage, &creationResponse)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "computer prestage", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &creationResponse, nil
}

// UpdateComputerPrestageByID updates a computer prestage by its ID.
func (c *Client) UpdateComputerPrestageByID(id string, prestageUpdate *ResourceComputerPrestage) (*ResourceComputerPrestage, error) {
	endpoint := fmt.Sprintf("%s/%s", uriComputerPrestagesV3, id)

	var updatedPrestage ResourceComputerPrestage
	resp, err := c.HTTP.DoRequest("PUT", endpoint, prestageUpdate, &updatedPrestage)
	if err != nil {
		return nil, fmt.Errorf("failed to update computer prestage with ID %s: %v", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedPrestage, nil
}

// UpdateComputerPrestageByNameByID updates a computer prestage based on its display name.
func (c *Client) UpdateComputerPrestageByName(name string, prestageUpdate *ResourceComputerPrestage) (*ResourceComputerPrestage, error) {
	target, err := c.GetComputerPrestageByName(name)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "computer prestage", name, err)
	}

	target_id := target.ID
	resp, err := c.UpdateComputerPrestageByID(target_id, prestageUpdate)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "computer prestage", name, err)
	}

	return resp, nil
}

// DeleteComputerPrestageByID deletes a computer prestage by its ID
func (c *Client) DeleteComputerPrestageByID(id string) error {
	endpoint := fmt.Sprintf("%s/%s", uriComputerPrestagesV3, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "computer prestage", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteComputerPrestageByNameByID deletes a computer prestage by its name.
func (c *Client) DeleteComputerPrestageByName(name string) error {
	target, err := c.GetComputerPrestageByName(name)
	if err != nil {
		return fmt.Errorf(errMsgFailedPaginatedGet, "computer prestages", err)
	}

	target_id := target.ID

	err = c.DeleteComputerPrestageByID(target_id)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByName, "computer prestage", name, err)
	}

	return nil
}

// GetDeviceScopeForComputerPrestage retrieves the device scope for a specific computer prestage by its ID.
func (c *Client) GetDeviceScopeForComputerPrestageByID(id string) (*ResponseDeviceScope, error) {
	endpoint := fmt.Sprintf("%s/%s/scope", uriComputerPrestagesV2, id)

	var deviceScope ResponseDeviceScope
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &deviceScope)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "computer prestage scope", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &deviceScope, nil
}
