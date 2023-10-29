// jamfproapi_computer_prestages.go

package jamfpro

import (
	"fmt"
)

const uriComputerPrestagesV2 = "/api/v2/computer-prestages/scope"
const uriComputerPrestagesV3 = "/api/v3/computer-prestages"

// ResponseComputerPrestages represents the structure of the response for fetching computer pre-stages
type ResponseComputerPrestages struct {
	TotalCount *int               `json:"totalCount,omitempty"`
	Results    []ComputerPrestage `json:"results,omitempty"`
}

type ComputerPrestage struct {
	DisplayName                       string                                             `json:"displayName"`
	Mandatory                         bool                                               `json:"mandatory"`
	MdmRemovable                      bool                                               `json:"mdmRemovable"`
	SupportPhoneNumber                string                                             `json:"supportPhoneNumber"`
	SupportEmailAddress               string                                             `json:"supportEmailAddress"`
	Department                        string                                             `json:"department"`
	DefaultPrestage                   bool                                               `json:"defaultPrestage"`
	EnrollmentSiteId                  string                                             `json:"enrollmentSiteId"`
	KeepExistingSiteMembership        bool                                               `json:"keepExistingSiteMembership"`
	KeepExistingLocationInformation   bool                                               `json:"keepExistingLocationInformation"`
	RequireAuthentication             bool                                               `json:"requireAuthentication"`
	AuthenticationPrompt              string                                             `json:"authenticationPrompt"`
	PreventActivationLock             bool                                               `json:"preventActivationLock"`
	EnableDeviceBasedActivationLock   bool                                               `json:"enableDeviceBasedActivationLock"`
	DeviceEnrollmentProgramInstanceId string                                             `json:"deviceEnrollmentProgramInstanceId"`
	SkipSetupItems                    *[]ComputerPrestageDataSubsetSkipSetupItems        `json:"skipSetupItems"`
	LocationInformation               *[]ComputerPrestageDataSubsetLocationInformation   `json:"locationInformation"`
	PurchasingInformation             *[]ComputerPrestageDataSubsetPurchasingInformation `json:"purchasingInformation"`
	AnchorCertificates                []string                                           `json:"anchorCertificates"`
	EnrollmentCustomizationId         string                                             `json:"enrollmentCustomizationId"`
	Language                          string                                             `json:"language"`
	Region                            string                                             `json:"region"`
	AutoAdvanceSetup                  bool                                               `json:"autoAdvanceSetup"`
	InstallProfilesDuringSetup        *bool                                              `json:"installProfilesDuringSetup"`
	PrestageInstalledProfileIds       []string                                           `json:"prestageInstalledProfileIds"`
	CustomPackageIds                  []string                                           `json:"customPackageIds"`
	CustomPackageDistributionPointId  string                                             `json:"customPackageDistributionPointId"`
	EnableRecoveryLock                *bool                                              `json:"enableRecoveryLock"`
	RecoveryLockPasswordType          *string                                            `json:"recoveryLockPasswordType"`
	RotateRecoveryLockPassword        *bool                                              `json:"rotateRecoveryLockPassword"`
	ID                                string                                             `json:"id"`
	ProfileUuid                       string                                             `json:"profileUuid"`
	SiteId                            string                                             `json:"siteId"`
	VersionLock                       int                                                `json:"versionLock"`
	AccountSettings                   *[]ComputerPrestageDataSubsetAccountSettings       `json:"accountSettings"`
}

type ComputerPrestageDataSubsetSkipSetupItems struct {
	Location bool `json:"Location"`
	Privacy  bool `json:"Privacy"`
}

type ComputerPrestageDataSubsetLocationInformation struct {
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

type ComputerPrestageDataSubsetPurchasingInformation struct {
	ID                string `json:"id"`
	Leased            bool   `json:"leased"`
	Purchased         bool   `json:"purchased"`
	AppleCareId       string `json:"appleCareId"`
	PoNumber          string `json:"poNumber"`
	Vendor            string `json:"vendor"`
	PurchasePrice     string `json:"purchasePrice"`
	LifeExpectancy    int    `json:"lifeExpectancy"`
	PurchasingAccount string `json:"purchasingAccount"`
	PurchasingContact string `json:"purchasingContact"`
	LeaseDate         string `json:"leaseDate"`
	PoDate            string `json:"poDate"`
	WarrantyDate      string `json:"warrantyDate"`
	VersionLock       int    `json:"versionLock"`
}

type ComputerPrestageDataSubsetAccountSettings struct {
	ID                                      *string `json:"id"`
	PayloadConfigured                       *bool   `json:"payloadConfigured"`
	LocalAdminAccountEnabled                *bool   `json:"localAdminAccountEnabled"`
	AdminUsername                           *string `json:"adminUsername"`
	AdminPassword                           *string `json:"adminPassword"`
	HiddenAdminAccount                      *bool   `json:"hiddenAdminAccount"`
	LocalUserManaged                        *bool   `json:"localUserManaged"`
	UserAccountType                         *string `json:"userAccountType"`
	VersionLock                             *int    `json:"versionLock"`
	PrefillPrimaryAccountInfoFeatureEnabled *bool   `json:"prefillPrimaryAccountInfoFeatureEnabled"`
	PrefillType                             *string `json:"prefillType"`
	PrefillAccountFullName                  *string `json:"prefillAccountFullName"`
	PrefillAccountUserName                  *string `json:"prefillAccountUserName"`
	PreventPrefillInfoFromModification      *bool   `json:"preventPrefillInfoFromModification"`
}

// GetComputerPrestages fetches all computer pre-stages
func (c *Client) GetComputerPrestages() (*ResponseComputerPrestages, error) {
	var preStagesList ResponseComputerPrestages
	resp, err := c.HTTP.DoRequest("GET", uriComputerPrestagesV2, nil, &preStagesList)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Jamf computer pre-stages: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &preStagesList, nil
}

// GetComputerPrestageByID fetches a computer pre-stage by its ID
func (c *Client) GetComputerPrestageByID(id string) (*ComputerPrestage, error) {
	endpoint := fmt.Sprintf(uriComputerPrestagesV3+"/%s", id)

	var preStage ComputerPrestage
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &preStage)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Jamf computer pre-stage ID %s: %v", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &preStage, nil
}

// GetComputerPrestageByNameByID fetches a Jamf computer pre-stage by its display name and then retrieves its details using its ID
func (c *Client) GetComputerPrestageByNameByID(name string) (*ComputerPrestage, error) {
	preStagesList, err := c.GetComputerPrestages()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all Jamf computer pre-stages: %v", err)
	}

	// Search for the pre-stage with the given name
	for _, preStage := range preStagesList.Results {
		fmt.Printf("Comparing desired name '%s' with pre-stage name '%s'\n", name, preStage.DisplayName) // Debug log
		if preStage.DisplayName == name {
			return c.GetComputerPrestageByID(preStage.ID)
		}
	}

	return nil, fmt.Errorf("no Jamf computer pre-stage found with the name %s", name)
}
