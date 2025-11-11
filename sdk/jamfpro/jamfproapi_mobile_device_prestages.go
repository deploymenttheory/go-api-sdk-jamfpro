// jamfproapi_mobile_device_prestages.go
// Jamf Pro Api - Mobile Device Prestages
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v3-mobile-device-prestages
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
	"net/url"

	"github.com/mitchellh/mapstructure"
)

const (
	uriMobileDevicePrestagesV2 = "/api/v2/mobile-device-prestages"
	uriMobileDevicePrestagesV3 = "/api/v3/mobile-device-prestages"
)

// List

type ResponseMobileDevicePrestagesList struct {
	TotalCount int                            `json:"totalCount"`
	Results    []ResourceMobileDevicePrestage `json:"results"`
}

// Response

// ResponseMobileDeviceScope represents the structure of the response for a specific mobile device prestage scope.
type ResponseMobileDeviceScope struct {
	PrestageId  string                                  `json:"prestageId"`
	Assignments []MobileDeviceScopeSubsetAssignmentItem `json:"assignments"`
	VersionLock int                                     `json:"versionLock"`
}

// MobileDeviceScopeSubsetAssignmentItem represents the structure of each assignment within the prestage scope.
type MobileDeviceScopeSubsetAssignmentItem struct {
	SerialNumber   string `json:"serialNumber"`
	AssignmentDate string `json:"assignmentDate"`
	UserAssigned   string `json:"userAssigned"`
}

// ResponseMobileDevicePrestageCreate represents the response structure for creating a mobile device prestage.
type ResponseMobileDevicePrestageCreate struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// Resource

type ResourceMobileDevicePrestage struct {
	DisplayName                            string                                          `json:"displayName"`
	Mandatory                              bool                                            `json:"mandatory"`
	MdmRemovable                           bool                                            `json:"mdmRemovable"`
	SupportPhoneNumber                     string                                          `json:"supportPhoneNumber"`
	SupportEmailAddress                    string                                          `json:"supportEmailAddress"`
	Department                             string                                          `json:"department"`
	DefaultPrestage                        bool                                            `json:"defaultPrestage"`
	EnrollmentSiteID                       string                                          `json:"enrollmentSiteId"`
	KeepExistingSiteMembership             bool                                            `json:"keepExistingSiteMembership"`
	KeepExistingLocationInformation        bool                                            `json:"keepExistingLocationInformation"`
	RequireAuthentication                  bool                                            `json:"requireAuthentication"`
	AuthenticationPrompt                   string                                          `json:"authenticationPrompt"`
	PreventActivationLock                  bool                                            `json:"preventActivationLock"`
	EnableDeviceBasedActivationLock        bool                                            `json:"enableDeviceBasedActivationLock"`
	DeviceEnrollmentProgramInstanceID      string                                          `json:"deviceEnrollmentProgramInstanceId"`
	SkipSetupItems                         MobileDevicePrestageSubsetSkipSetupItems        `json:"skipSetupItems,omitempty"`
	LocationInformation                    MobileDevicePrestageSubsetLocationInformation   `json:"locationInformation"`
	PurchasingInformation                  MobileDevicePrestageSubsetPurchasingInformation `json:"purchasingInformation"`
	AnchorCertificates                     []string                                        `json:"anchorCertificates,omitempty"`
	EnrollmentCustomizationID              string                                          `json:"enrollmentCustomizationId,omitempty"`
	Language                               string                                          `json:"language,omitempty"`
	Region                                 string                                          `json:"region,omitempty"`
	AutoAdvanceSetup                       bool                                            `json:"autoAdvanceSetup"`
	AllowPairing                           bool                                            `json:"allowPairing"`
	MultiUser                              bool                                            `json:"multiUser"`
	Supervised                             bool                                            `json:"supervised"`
	MaximumSharedAccounts                  int                                             `json:"maximumSharedAccounts"`
	ConfigureDeviceBeforeSetupAssistant    bool                                            `json:"configureDeviceBeforeSetupAssistant"`
	Names                                  MobileDevicePrestageSubsetNames                 `json:"names"`
	SendTimezone                           bool                                            `json:"sendTimezone"`
	Timezone                               string                                          `json:"timezone"`
	StorageQuotaSizeMegabytes              int                                             `json:"storageQuotaSizeMegabytes"`
	UseStorageQuotaSize                    bool                                            `json:"useStorageQuotaSize"`
	TemporarySessionOnly                   bool                                            `json:"temporarySessionOnly,omitempty"`
	EnforceTemporarySessionTimeout         bool                                            `json:"enforceTemporarySessionTimeout,omitempty"`
	TemporarySessionTimeout                *int                                            `json:"temporarySessionTimeout,omitempty"`
	EnforceUserSessionTimeout              bool                                            `json:"enforceUserSessionTimeout,omitempty"`
	UserSessionTimeout                     *int                                            `json:"userSessionTimeout,omitempty"`
	ID                                     string                                          `json:"id"`
	ProfileUuid                            string                                          `json:"profileUuid,omitempty"`
	SiteId                                 string                                          `json:"siteId,omitempty"`
	VersionLock                            int                                             `json:"versionLock"`
	PrestageMinimumOsTargetVersionTypeIos  string                                          `json:"prestageMinimumOsTargetVersionTypeIos,omitempty"`
	MinimumOsSpecificVersionIos            string                                          `json:"minimumOsSpecificVersionIos,omitempty"`
	PrestageMinimumOsTargetVersionTypeIpad string                                          `json:"prestageMinimumOsTargetVersionTypeIpad,omitempty"`
	MinimumOsSpecificVersionIpad           string                                          `json:"minimumOsSpecificVersionIpad,omitempty"`
	RTSEnabled                             bool                                            `json:"rtsEnabled,omitempty"`
	RTSConfigProfileId                     string                                          `json:"rtsConfigProfileId,omitempty"`
	PreserveManagedApps                    bool                                            `json:"preserveManagedApps,omitempty"`
	InstallAppsDuringEnrollment            bool                                            `json:"installAppsDuringEnrollment,omitempty"`
}

// Subsets & Containers

type MobileDevicePrestageSubsetSkipSetupItems struct {
	Location              bool `json:"Location,omitempty"`
	Privacy               bool `json:"Privacy,omitempty"`
	Biometric             bool `json:"Biometric,omitempty"`
	SoftwareUpdate        bool `json:"SoftwareUpdate,omitempty"`
	Diagnostics           bool `json:"Diagnostics,omitempty"`
	IMessageAndFaceTime   bool `json:"iMessageAndFaceTime,omitempty"`
	Intelligence          bool `json:"Intelligence,omitempty"`
	TVRoom                bool `json:"TVRoom,omitempty"`
	Passcode              bool `json:"Passcode,omitempty"`
	SIMSetup              bool `json:"SIMSetup,omitempty"`
	ScreenTime            bool `json:"ScreenTime,omitempty"`
	RestoreCompleted      bool `json:"RestoreCompleted,omitempty"`
	TVProviderSignIn      bool `json:"TVProviderSignIn,omitempty"`
	Siri                  bool `json:"Siri,omitempty"`
	Restore               bool `json:"Restore,omitempty"`
	ScreenSaver           bool `json:"ScreenSaver,omitempty"`
	HomeButtonSensitivity bool `json:"HomeButtonSensitivity,omitempty"`
	CloudStorage          bool `json:"CloudStorage,omitempty"`
	ActionButton          bool `json:"ActionButton,omitempty"`
	TransferData          bool `json:"TransferData,omitempty"`
	EnableLockdownMode    bool `json:"EnableLockdownMode,omitempty"`
	Zoom                  bool `json:"Zoom,omitempty"`
	PreferredLanguage     bool `json:"PreferredLanguage,omitempty"`
	VoiceSelection        bool `json:"VoiceSelection,omitempty"`
	TVHomeScreenSync      bool `json:"TVHomeScreenSync,omitempty"`
	Safety                bool `json:"Safety,omitempty"`
	TermsOfAddress        bool `json:"TermsOfAddress,omitempty"`
	ExpressLanguage       bool `json:"ExpressLanguage,omitempty"`
	CameraButton          bool `json:"CameraButton,omitempty"`
	AppleID               bool `json:"AppleID,omitempty"`
	DisplayTone           bool `json:"DisplayTone,omitempty"`
	WatchMigration        bool `json:"WatchMigration,omitempty"`
	UpdateCompleted       bool `json:"UpdateCompleted,omitempty"`
	Appearance            bool `json:"Appearance,omitempty"`
	Android               bool `json:"Android,omitempty"`
	Payment               bool `json:"Payment,omitempty"`
	OnBoarding            bool `json:"OnBoarding,omitempty"`
	TOS                   bool `json:"TOS,omitempty"`
	Welcome               bool `json:"Welcome,omitempty"`
	SafetyAndHandling     bool `json:"SafetyAndHandling,omitempty"`
	TapToSetup            bool `json:"TapToSetup,omitempty"`
	SpokenLanguage        bool `json:"SpokenLanguage,omitempty"`
	Keyboard              bool `json:"Keyboard,omitempty"`
	Multitasking          bool `json:"Multitasking,omitempty"`
	OSShowcase            bool `json:"OSShowcase,omitempty"`
}

type MobileDevicePrestageSubsetLocationInformation struct {
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

type MobileDevicePrestageSubsetPurchasingInformation struct {
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

type MobileDevicePrestageSubsetNames struct {
	AssignNamesUsing       string                                `json:"assignNamesUsing"`
	PrestageDeviceNames    []MobileDevicePrestageSubsetNamesName `json:"prestageDeviceNames,omitempty"`
	DeviceNamePrefix       string                                `json:"deviceNamePrefix,omitempty"`
	DeviceNameSuffix       string                                `json:"deviceNameSuffix,omitempty"`
	SingleDeviceName       string                                `json:"singleDeviceName,omitempty"`
	ManageNames            bool                                  `json:"manageNames,omitempty"`
	DeviceNamingConfigured bool                                  `json:"deviceNamingConfigured,omitempty"`
}

type MobileDevicePrestageSubsetNamesName struct {
	ID         string `json:"id,omitempty"`
	DeviceName string `json:"deviceName,omitempty"`
	Used       bool   `json:"used,omitempty"`
}

// CRUD

// GetMobileDevicePrestages retrieves a list of all mobile prestages
func (c *Client) GetMobileDevicePrestages(params url.Values) (*ResponseMobileDevicePrestagesList, error) {
	endpoint := uriMobileDevicePrestagesV3
	resp, err := c.DoPaginatedGet(endpoint, params)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "mobile device prestages", err)
	}

	var out ResponseMobileDevicePrestagesList
	out.TotalCount = resp.Size

	for _, value := range resp.Results {
		var newObj ResourceMobileDevicePrestage
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "mobile device prestage", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil
}

// GetMobileDevicePrestageByID retrieves a single mobile prestage from the supplied ID
func (c *Client) GetMobileDevicePrestageByID(id string) (*ResourceMobileDevicePrestage, error) {
	endpoint := fmt.Sprintf("%s/%s", uriMobileDevicePrestagesV3, id)
	var out ResourceMobileDevicePrestage

	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "mobile device prestage", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// GetMobileDevicePrestageByName retrieves a specific mobile device prestage by its name.
func (c *Client) GetMobileDevicePrestageByName(name string) (*ResourceMobileDevicePrestage, error) {
	prestages, err := c.GetMobileDevicePrestages(nil)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "mobile device prestages", err)
	}

	for _, value := range prestages.Results {
		if value.DisplayName == name {
			return &value, nil
		}
	}

	return nil, fmt.Errorf(errMsgFailedGetByName, "mobile device prestage", name, errMsgNoName)
}

// CreateMobileDevicePrestage creates a new mobile prestage and returns the id
func (c *Client) CreateMobileDevicePrestage(newPrestage ResourceMobileDevicePrestage) (*ResponseMobileDevicePrestageCreate, error) {
	endpoint := uriMobileDevicePrestagesV3
	var out ResponseMobileDevicePrestageCreate
	resp, err := c.HTTP.DoRequest("POST", endpoint, newPrestage, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "mobile device prestage", err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// UpdateMobileDevicePrestageByID updates a mobile device prestage by its ID.
func (c *Client) UpdateMobileDevicePrestageByID(id string, prestageUpdate *ResourceMobileDevicePrestage) (*ResourceMobileDevicePrestage, error) {
	endpoint := fmt.Sprintf("%s/%s", uriMobileDevicePrestagesV3, id)

	var updatedPrestage ResourceMobileDevicePrestage
	resp, err := c.HTTP.DoRequest("PUT", endpoint, prestageUpdate, &updatedPrestage)
	if err != nil {
		return nil, fmt.Errorf("failed to update mobile device prestage with ID %s: %v", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedPrestage, nil
}

// UpdateMobileDevicePrestageByNameByID updates a mobile prestage based on its display name.
func (c *Client) UpdateMobileDevicePrestageByName(name string, prestageUpdate *ResourceMobileDevicePrestage) (*ResourceMobileDevicePrestage, error) {
	target, err := c.GetMobileDevicePrestageByName(name)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "mobile device prestage", name, err)
	}

	target_id := target.ID
	resp, err := c.UpdateMobileDevicePrestageByID(target_id, prestageUpdate)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "mobile device prestage", name, err)
	}

	return resp, nil
}

// DeleteMobileDevicePrestageByID a mobile prestage at the given id
func (c *Client) DeleteMobileDevicePrestageByID(id string) error {
	endpoint := fmt.Sprintf("%s/%s", uriMobileDevicePrestagesV3, id)
	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "mobile device prestage", id, err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteMobileDevicePrestageByNameByID deletes a mobile device prestage by its name.
func (c *Client) DeleteMobileDevicePrestageByName(name string) error {
	target, err := c.GetMobileDevicePrestageByName(name)
	if err != nil {
		return fmt.Errorf(errMsgFailedPaginatedGet, "mobile device prestages", err)
	}

	target_id := target.ID

	err = c.DeleteMobileDevicePrestageByID(target_id)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByName, "mobile device prestage", name, err)
	}

	return nil
}

// GetDeviceScopeForMobileDevicePrestage retrieves the device scope for a specific mobile device prestage by its ID.
func (c *Client) GetDeviceScopeForMobileDevicePrestageByID(id string) (*ResponseMobileDeviceScope, error) {
	endpoint := fmt.Sprintf("%s/%s/scope", uriMobileDevicePrestagesV2, id)

	var deviceScope ResponseMobileDeviceScope
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &deviceScope)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "mobile device prestage scope", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &deviceScope, nil
}
