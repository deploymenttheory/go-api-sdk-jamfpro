package jamfpro

import (
	"fmt"
)

const uriComputerPrestageV2 = "/api/v2/computer-prestages"
const uriComputerPrestageV3 = "/api/v3/computer-prestages"

type ResponseComputerPrestages struct {
	TotalCount *int               `json:"totalCount,omitempty"`
	Results    []computerPrestage `json:"results,omitempty"`
}

type computerPrestage struct {
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

func (c *Client) GetComputerPrestageIdByName(name string) (string, error) {
	var id string
	prestages, err := c.GetComputerPrestages()
	if err != nil {
		return "", err
	}

	for _, v := range prestages.Results {
		if v.DisplayName == name {
			id = v.ID
			break
		}
	}
	return id, err
}

func (c *Client) GetComputerPrestageByName(name string) (*computerPrestage, error) {
	// Fetch all computer prestages using the appropriate function
	allPrestagesResponse, err := c.GetComputerPrestages()
	if err != nil {
		return nil, err
	}

	// Iterate through the results to find the matching name
	for _, prestage := range allPrestagesResponse.Results {
		if prestage.DisplayName == name {
			return &prestage, nil // Return the matching prestage
		}
	}

	// Return an error if no match is found
	return nil, fmt.Errorf("computer prestage with name '%s' not found", name)
}

func (c *Client) GetComputerPrestages() (*ResponseComputerPrestages, error) {
	uri := uriComputerPrestageV2 // Using the constant defined earlier

	out := &ResponseComputerPrestages{}
	err := c.DoRequest("GET", uri, nil, nil, out)
	if err != nil {
		return nil, fmt.Errorf("failed to get all computer prestages: %v", err)
	}
	return out, nil
}

func (c *Client) GetComputerPrestageByID(prestageID int) (*computerPrestage, error) {
	var out *computerPrestage
	uri := fmt.Sprintf("%s/%d", uriComputerPrestageV3, prestageID)

	err := c.DoRequest("GET", uri, nil, nil, &out)
	return out, err
}

func (c *Client) CreateComputerPrestage(
	DisplayName *string,
	Mandatory *bool,
	MdmRemovable *bool,
	SupportPhoneNumber *string,
	SupportEmailAddress *string,
	Department *string,
	DefaultPrestage *bool,
	EnrollmentSiteId *string,
	KeepExistingSiteMembership *bool,
	KeepExistingLocationInformation *bool,
	RequireAuthentication *bool,
	AuthenticationPrompt *string,
	PreventActivationLock *bool,
	EnableDeviceBasedActivationLock *bool,
	DeviceEnrollmentProgramInstanceId *string,
	SkipSetupItems *map[string]bool,
	LocationInformation *ComputerPrestageDataSubsetLocationInformation,
	PurchasingInformation *ComputerPrestageDataSubsetPurchasingInformation,
	AnchorCertificates *[]string,
	EnrollmentCustomizationId *string,
	Language *string,
	Region *string,
	AutoAdvanceSetup *bool,
	InstallProfilesDuringSetup *bool,
	PrestageInstalledProfileIds *[]string,
	CustomPackageIds *[]string,
	CustomPackageDistributionPointId *string,
	EnableRecoveryLock *bool,
	RecoveryLockPasswordType *string,
	RotateRecoveryLockPassword *bool,
	AccountSettings *ComputerPrestageDataSubsetAccountSettings,
) (*computerPrestage, error) {

	in := struct {
		DisplayName                       *string                                          `json:"displayName"`
		Mandatory                         *bool                                            `json:"mandatory"`
		MdmRemovable                      *bool                                            `json:"mdmRemovable"`
		SupportPhoneNumber                *string                                          `json:"supportPhoneNumber"`
		SupportEmailAddress               *string                                          `json:"supportEmailAddress"`
		Department                        *string                                          `json:"department"`
		DefaultPrestage                   *bool                                            `json:"defaultPrestage"`
		EnrollmentSiteId                  *string                                          `json:"enrollmentSiteId"`
		KeepExistingSiteMembership        *bool                                            `json:"keepExistingSiteMembership"`
		KeepExistingLocationInformation   *bool                                            `json:"keepExistingLocationInformation"`
		RequireAuthentication             *bool                                            `json:"requireAuthentication"`
		AuthenticationPrompt              *string                                          `json:"authenticationPrompt"`
		PreventActivationLock             *bool                                            `json:"preventActivationLock"`
		EnableDeviceBasedActivationLock   *bool                                            `json:"enableDeviceBasedActivationLock"`
		DeviceEnrollmentProgramInstanceId *string                                          `json:"deviceEnrollmentProgramInstanceId"`
		SkipSetupItems                    *map[string]bool                                 `json:"skipSetupItems"`
		LocationInformation               *ComputerPrestageDataSubsetLocationInformation   `json:"locationInformation"`
		PurchasingInformation             *ComputerPrestageDataSubsetPurchasingInformation `json:"purchasingInformation"`
		AnchorCertificates                *[]string                                        `json:"anchorCertificates"`
		EnrollmentCustomizationId         *string                                          `json:"enrollmentCustomizationId"`
		Language                          *string                                          `json:"language"`
		Region                            *string                                          `json:"region"`
		AutoAdvanceSetup                  *bool                                            `json:"autoAdvanceSetup"`
		InstallProfilesDuringSetup        *bool                                            `json:"installProfilesDuringSetup"`
		PrestageInstalledProfileIds       *[]string                                        `json:"prestageInstalledProfileIds"`
		CustomPackageIds                  *[]string                                        `json:"customPackageIds"`
		CustomPackageDistributionPointId  *string                                          `json:"customPackageDistributionPointId"`
		EnableRecoveryLock                *bool                                            `json:"enableRecoveryLock"`
		RecoveryLockPasswordType          *string                                          `json:"recoveryLockPasswordType"`
		RotateRecoveryLockPassword        *bool                                            `json:"rotateRecoveryLockPassword"`
		AccountSettings                   *ComputerPrestageDataSubsetAccountSettings       `json:"accountSettings"`
	}{
		DisplayName:                       DisplayName,
		Mandatory:                         Mandatory,
		MdmRemovable:                      MdmRemovable,
		SupportPhoneNumber:                SupportPhoneNumber,
		SupportEmailAddress:               SupportEmailAddress,
		Department:                        Department,
		DefaultPrestage:                   DefaultPrestage,
		EnrollmentSiteId:                  EnrollmentSiteId,
		KeepExistingSiteMembership:        KeepExistingSiteMembership,
		KeepExistingLocationInformation:   KeepExistingLocationInformation,
		RequireAuthentication:             RequireAuthentication,
		AuthenticationPrompt:              AuthenticationPrompt,
		PreventActivationLock:             PreventActivationLock,
		EnableDeviceBasedActivationLock:   EnableDeviceBasedActivationLock,
		DeviceEnrollmentProgramInstanceId: DeviceEnrollmentProgramInstanceId,
		SkipSetupItems:                    SkipSetupItems,
		LocationInformation:               LocationInformation,
		PurchasingInformation:             PurchasingInformation,
		AnchorCertificates:                AnchorCertificates,
		EnrollmentCustomizationId:         EnrollmentCustomizationId,
		Language:                          Language,
		Region:                            Region,
		AutoAdvanceSetup:                  AutoAdvanceSetup,
		InstallProfilesDuringSetup:        InstallProfilesDuringSetup,
		PrestageInstalledProfileIds:       PrestageInstalledProfileIds,
		CustomPackageIds:                  CustomPackageIds,
		CustomPackageDistributionPointId:  CustomPackageDistributionPointId,
		EnableRecoveryLock:                EnableRecoveryLock,
		RecoveryLockPasswordType:          RecoveryLockPasswordType,
		RotateRecoveryLockPassword:        RotateRecoveryLockPassword,
		AccountSettings:                   AccountSettings,
	}

	var out *computerPrestage
	err := c.DoRequest("POST", uriComputerPrestageV3, in, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to create computer prestage: %v", err)
	}
	return out, nil
}

func (c *Client) UpdateComputerPrestage(d *computerPrestage) (*computerPrestage, error) {
	uri := fmt.Sprintf("%s/%v", uriComputerPrestageV3, d.ID)
	updatedPrestage := &computerPrestage{}

	// Perform the PUT request
	err := c.DoRequest("PUT", uri, d, nil, updatedPrestage)

	return updatedPrestage, err
}

func (c *Client) DeleteComputerPrestage(prestageID int) error {
	uri := fmt.Sprintf("%s/%v", uriComputerPrestageV3, prestageID)

	// Perform the DELETE request
	err := c.DoRequest("DELETE", uri, nil, nil, nil)

	return err
}
