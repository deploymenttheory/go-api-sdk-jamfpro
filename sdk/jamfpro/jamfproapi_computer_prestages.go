// jamfproapi_computer_prestages.go

package jamfpro

import (
	"fmt"
)

const uriComputerPrestagesV3 = "/api/v3/computer-prestages"

// ResponseComputerPrestages represents the structure of the response for fetching computer pre-stages
type ResponseComputerPrestages struct {
	TotalCount int                `json:"totalCount"`
	Results    []ComputerPrestage `json:"results"`
}

// ComputerPrestage represents the top-level structure for a computer prestage response.
type ComputerPrestage struct {
	DisplayName                       string                `json:"displayName"`
	Mandatory                         bool                  `json:"mandatory"`
	MdmRemovable                      bool                  `json:"mdmRemovable"`
	SupportPhoneNumber                string                `json:"supportPhoneNumber"`
	SupportEmailAddress               string                `json:"supportEmailAddress"`
	Department                        string                `json:"department"`
	DefaultPrestage                   bool                  `json:"defaultPrestage"`
	EnrollmentSiteId                  string                `json:"enrollmentSiteId"`
	KeepExistingSiteMembership        bool                  `json:"keepExistingSiteMembership"`
	KeepExistingLocationInformation   bool                  `json:"keepExistingLocationInformation"`
	RequireAuthentication             bool                  `json:"requireAuthentication"`
	AuthenticationPrompt              string                `json:"authenticationPrompt"`
	PreventActivationLock             bool                  `json:"preventActivationLock"`
	EnableDeviceBasedActivationLock   bool                  `json:"enableDeviceBasedActivationLock"`
	DeviceEnrollmentProgramInstanceId string                `json:"deviceEnrollmentProgramInstanceId"`
	SkipSetupItems                    SkipSetupItems        `json:"skipSetupItems"`
	LocationInformation               LocationInformation   `json:"locationInformation"`
	PurchasingInformation             PurchasingInformation `json:"purchasingInformation"`
	AnchorCertificates                []string              `json:"anchorCertificates"`
	EnrollmentCustomizationId         string                `json:"enrollmentCustomizationId"`
	Language                          string                `json:"language"`
	Region                            string                `json:"region"`
	AutoAdvanceSetup                  bool                  `json:"autoAdvanceSetup"`
	InstallProfilesDuringSetup        bool                  `json:"installProfilesDuringSetup"`
	PrestageInstalledProfileIds       []string              `json:"prestageInstalledProfileIds"`
	CustomPackageIds                  []string              `json:"customPackageIds"`
	CustomPackageDistributionPointId  string                `json:"customPackageDistributionPointId"`
	EnableRecoveryLock                bool                  `json:"enableRecoveryLock"`
	RecoveryLockPasswordType          string                `json:"recoveryLockPasswordType"`
	RotateRecoveryLockPassword        bool                  `json:"rotateRecoveryLockPassword"`
	ID                                string                `json:"id"`
	ProfileUuid                       string                `json:"profileUuid"`
	SiteId                            string                `json:"siteId"`
	VersionLock                       int                   `json:"versionLock"`
	AccountSettings                   AccountSettings       `json:"accountSettings"`
}

// SkipSetupItems represents the structure for skipping setup items.
type SkipSetupItems struct {
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

// LocationInformation represents the structure for location information.
type LocationInformation struct {
	Username     string `json:"username"`
	RealName     string `json:"realname"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Room         string `json:"room"`
	Position     string `json:"position"`
	DepartmentId string `json:"departmentId"`
	BuildingId   string `json:"buildingId"`
	ID           string `json:"id"`
	VersionLock  int    `json:"versionLock"`
}

// PurchasingInformation represents the structure for purchasing information.
type PurchasingInformation struct {
	ID                string `json:"id"`
	Leased            bool   `json:"leased"`
	Purchased         bool   `json:"purchased"`
	AppleCareId       string `json:"appleCareId"`
	PoNumber          string `json:"poNumber"`
	Vendor            string `json:"vendor"`
	PurchasePrice     string `json:"purchasePrice"`
	LifeExpectancy    int    `json:"lifeExpectancy"`
	WarrantyExpires   string `json:"warrantyExpires"`
	LeaseExpires      string `json:"leaseExpires"`
	PurchaseDate      string `json:"purchaseDate"`
	PoDate            string `json:"poDate"`
	VersionLock       int    `json:"versionLock"`
	IsPurchased       bool   `json:"isPurchased"`
	IsLeased          bool   `json:"isLeased"`
	PurchasingAccount string `json:"purchasingAccount"`
	WarrantyDaysLeft  int    `json:"warrantyDaysLeft"`
	LeaseDaysLeft     int    `json:"leaseDaysLeft"`
	Attachments       []struct {
		ID           string `json:"id"`
		Filename     string `json:"filename"`
		Filetype     string `json:"filetype"`
		Filesize     int    `json:"filesize"`
		VersionLock  int    `json:"versionLock"`
		DownloadLink string `json:"downloadLink"`
	} `json:"attachments"`
}

// AccountSettings represents the structure for account settings.
type AccountSettings struct {
	ID                                      string `json:"id"`
	PayloadConfigured                       bool   `json:"payloadConfigured"`
	LocalAdminAccountEnabled                bool   `json:"localAdminAccountEnabled"`
	AdminUsername                           string `json:"adminUsername"`
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

// GetComputerPrestages fetches all computer pre-stages
func (c *Client) GetComputerPrestages() (*ResponseComputerPrestages, error) {
	var preStagesList ResponseComputerPrestages
	resp, err := c.HTTP.DoRequest("GET", uriComputerPrestagesV3, nil, &preStagesList)
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
