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

type ResponseComputerPrestagesList struct {
	TotalCount *int                       `json:"totalCount"`
	Results    []ResourceComputerPrestage `json:"results"`
}

type ResourceComputerPrestage struct {
	DisplayName                       string                                      `json:"displayName"`
	Mandatory                         bool                                        `json:"mandatory"`
	MDMRemovable                      bool                                        `json:"mdmRemovable"`
	SupportPhoneNumber                string                                      `json:"supportPhoneNumber"`
	SupportEmailAddress               string                                      `json:"supportEmailAddress"`
	Department                        string                                      `json:"department"`
	DefaultPrestage                   bool                                        `json:"defaultPrestage"`
	EnrollmentSiteId                  string                                      `json:"enrollmentSiteId"`
	KeepExistingSiteMembership        bool                                        `json:"keepExistingSiteMembership"`
	KeepExistingLocationInformation   bool                                        `json:"keepExistingLocationInformation"`
	RequireAuthentication             bool                                        `json:"requireAuthentication"`
	AuthenticationPrompt              string                                      `json:"authenticationPrompt"`
	PreventActivationLock             bool                                        `json:"preventActivationLock"`
	EnableDeviceBasedActivationLock   bool                                        `json:"enableDeviceBasedActivationLock"`
	DeviceEnrollmentProgramInstanceId string                                      `json:"deviceEnrollmentProgramInstanceId"`
	SkipSetupItems                    map[string]bool                             `json:"skipSetupItems"`
	LocationInformation               ComputerPrestageSubsetLocationInformation   `json:"locationInformation"`
	PurchasingInformation             ComputerPrestageSubsetPurchasingInformation `json:"purchasingInformation"`
	AnchorCertificates                []string                                    `json:"anchorCertificates"`
	EnrollmentCustomizationId         string                                      `json:"enrollmentCustomizationId"`
	Language                          string                                      `json:"language"`
	Region                            string                                      `json:"region"`
	AutoAdvanceSetup                  bool                                        `json:"autoAdvanceSetup"`
	InstallProfilesDuringSetup        bool                                        `json:"installProfilesDuringSetup"`
	PrestageInstalledProfileIds       []string                                    `json:"prestageInstalledProfileIds"`
	CustomPackageIds                  []string                                    `json:"customPackageIds"`
	CustomPackageDistributionPointId  string                                      `json:"customPackageDistributionPointId"`
	EnableRecoveryLock                bool                                        `json:"enableRecoveryLock"`
	RecoveryLockPasswordType          string                                      `json:"recoveryLockPasswordType"`
	RecoveryLockPassword              string                                      `json:"recoveryLockPassword"`
	RotateRecoveryLockPassword        bool                                        `json:"rotateRecoveryLockPassword"`
	ID                                string                                      `json:"id"`
	ProfileUuid                       string                                      `json:"profileUuid"`
	SiteId                            string                                      `json:"siteId"`
	VersionLock                       int                                         `json:"versionLock"`
	AccountSettings                   ComputerPrestageSubsetAccountSettings       `json:"accountSettings"`
}

type ComputerPrestageSubsetLocationInformation struct {
	Username     string `json:"username"`
	Realname     string `json:"realname"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Room         string `json:"room"`
	Position     string `json:"position"`
	DepartmentId string `json:"departmentId"`
	BuildingId   string `json:"buildingId"`
	ID           string `json:"id"`
	VersionLock  int    `json:"versionLock"`
}

type ComputerPrestageSubsetPurchasingInformation struct {
	ID                string `json:"id"`
	Leased            bool   `json:"leased"`
	Purchased         bool   `json:"purchased"`
	AppleCareId       string `json:"appleCareId"`
	PONumber          string `json:"poNumber"`
	Vendor            string `json:"vendor"`
	PurchasePrice     string `json:"purchasePrice"`
	LifeExpectancy    int    `json:"lifeExpectancy"`
	PurchasingAccount string `json:"purchasingAccount"`
	PurchasingContact string `json:"purchasingContact"`
	LeaseDate         string `json:"leaseDate"`
	PODate            string `json:"poDate"`
	WarrantyDate      string `json:"warrantyDate"`
	VersionLock       int    `json:"versionLock"`
}

type ComputerPrestageSubsetAccountSettings struct {
	ID                                      string `json:"id"`
	PayloadConfigured                       bool   `json:"payloadConfigured"`
	LocalAdminAccountEnabled                bool   `json:"localAdminAccountEnabled"`
	AdminUsername                           string `json:"adminUsername"`
	AdminPassword                           string `json:"adminPassword"`
	HiddenAdminAccount                      bool   `json:"hiddenAdminAccount"`
	LocalUserManaged                        bool   `json:"localUserManaged"`
	UserAccountType                         string `json:"userAccountType"`
	VersionLock                             int    `json:"versionLock"`
	PrefillPrimaryAccountInfoFeatureEnabled bool   `json:"prefillPrimaryAccountInfoFeatureEnabled"`
	PrefillType                             string `json:"prefillType"`
	PrefillAccountFullName                  string `json:"prefillAccountFullName"`
	PrefillAccountUserName                  string `json:"prefillAccountUserName"`
	PreventPrefillInfoFromModification      bool   `json:"preventPrefillInfoFromModification"`
}

// ResponseComputerPrestageCreate represents the response structure for creating a building.
type ResponseComputerPrestageCreate struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

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

	return nil, fmt.Errorf(errMsgFailedGetByName, "computer prestage", name, err)
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
