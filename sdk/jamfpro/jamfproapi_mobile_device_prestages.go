// jamfproapi_mobile_device_prestages.go
// Jamf Pro Api - Mobile Device Prestages
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v2-mobile-device-prestages
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
	"net/url"

	"github.com/mitchellh/mapstructure"
)

const uriMobileDevicePrestages = "/api/v2/mobile-device-prestages"

// Structs

// List

type ResponseMobileDevicePrestagesList struct {
	TotalCount int                            `json:"totalCount"`
	Results    []ResourceMobileDevicePrestage `json:"results"`
}

// Response

type ResponseMobileDevicePrestageCreate struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// Resource

type ResourceMobileDevicePrestage struct {
	DisplayName                            string                                        `json:"displayName"`
	Mandatory                              bool                                          `json:"mandatory"`
	MdmRemovable                           bool                                          `json:"mdmRemovable"`
	SupportPhoneNumber                     string                                        `json:"supportPhoneNumber"`
	SupportEmailAddress                    string                                        `json:"supportEmailAddress"`
	Department                             string                                        `json:"department"`
	DefaultPrestage                        bool                                          `json:"defaultPrestage"`
	EnrollmentSiteID                       string                                        `json:"enrollmentSiteId"`
	KeepExistingSiteMembership             bool                                          `json:"keepExistingSiteMembership"`
	KeepExistingLocationInformation        bool                                          `json:"keepExistingLocationInformation"`
	RequireAuthentication                  bool                                          `json:"requireAuthentication"`
	AuthenticationPrompt                   string                                        `json:"authenticationPrompt"`
	PreventActivationLock                  bool                                          `json:"preventActivationLock"`
	EnableDeviceBasedActivationLock        bool                                          `json:"enableDeviceBasedActivationLock"`
	DeviceEnrollmentProgramInstanceID      string                                        `json:"deviceEnrollmentProgramInstanceId"`
	SkipSetupItems                         MobileDevicePrestageSubsetSkipSetupItems      `json:"skipSetupItems"`
	LocationInformation                    MobileDevicePrestageSubsetLocationInformation `json:"locationInformation"`
	PurchasingInformation                  MobileDevicePrestageSubsetLocationInformation `json:"purchasingInformation"`
	AnchorCertificates                     []string                                      `json:"anchorCertificates"`
	EnrollmentCustomizationID              string                                        `json:"enrollmentCustomizationId"`
	Language                               string                                        `json:"language"`
	Region                                 string                                        `json:"region"`
	AutoAdvanceSetup                       bool                                          `json:"autoAdvanceSetup"`
	AllowPairing                           bool                                          `json:"allowPairing"`
	MultiUser                              bool                                          `json:"multiUser"`
	Supervised                             bool                                          `json:"supervised"`
	MaximumSharedAccounts                  int                                           `json:"maximumSharedAccounts"`
	ConfigureDeviceBeforeSetupAssistant    bool                                          `json:"configureDeviceBeforeSetupAssistant"`
	Names                                  MobileDevicePrestageSubsetNames               `json:"names"`
	SendTimezone                           bool                                          `json:"sendTimezone"`
	Timezone                               string                                        `json:"timezone"`
	StorageQuotaSizeMegabytes              int                                           `json:"storageQuotaSizeMegabytes"`
	UseStorageQuotaSize                    bool                                          `json:"useStorageQuotaSize"`
	TemporarySessionOnly                   bool                                          `json:"temporarySessionOnly"`
	EnforceTemporarySessionTimeout         bool                                          `json:"enforceTemporarySessionTimeout"`
	TemporarySessionTimeout                int                                           `json:"temporarySessionTimeout"`
	EnforceUserSessionTimeout              bool                                          `json:"enforceUserSessionTimeout"`
	UserSessionTimeout                     int                                           `json:"userSessionTimeout"`
	ID                                     string                                        `json:"id"`
	ProfileUuid                            string                                        `json:"profileUuid"`
	SiteId                                 string                                        `json:"siteId"`
	VersionLock                            int                                           `json:"versionLock"`
	PrestageMinimumOsTargetVersionTypeIos  string                                        `json:"prestageMinimumOsTargetVersionTypeIos"`
	MinimumOsSpecificVersionIos            string                                        `json:"minimumOsSpecificVersionIos"`
	PrestageMinimumOsTargetVersionTypeIpad string                                        `json:"prestageMinimumOsTargetVersionTypeIpad"`
	MinimumOsSpecificVersionIpad           string                                        `json:"minimumOsSpecificVersionIpad"`
}

// Subsets

type MobileDevicePrestageSubsetSkipSetupItems struct {
	Location              bool `json:"Location"`
	Privacy               bool `json:"Privacy"`
	Biometric             bool `json:"Biometric"`
	SoftwareUpdate        bool `json:"SoftwareUpdate"`
	Diagnostics           bool `json:"Diagnostics"`
	IMessageAndFaceTime   bool `json:"iMessageAndFaceTime"`
	Intelligence          bool `json:"Intelligence"`
	TVRoom                bool `json:"TVRoom"`
	Passcode              bool `json:"Passcode"`
	SIMSetup              bool `json:"SIMSetup"`
	ScreenTime            bool `json:"ScreenTime"`
	RestoreCompleted      bool `json:"RestoreCompleted"`
	TVProviderSignIn      bool `json:"TVProviderSignIn"`
	Siri                  bool `json:"Siri"`
	Restore               bool `json:"Restore"`
	ScreenSaver           bool `json:"ScreenSaver"`
	HomeButtonSensitivity bool `json:"HomeButtonSensitivity"`
	CloudStorage          bool `json:"CloudStorage"`
	ActionButton          bool `json:"ActionButton"`
	TransferData          bool `json:"TransferData"`
	EnableLockdownMode    bool `json:"EnableLockdownMode"`
	Zoom                  bool `json:"Zoom"`
	PreferredLanguage     bool `json:"PreferredLanguage"`
	VoiceSelection        bool `json:"VoiceSelection"`
	TVHomeScreenSync      bool `json:"TVHomeScreenSync"`
	Safety                bool `json:"Safety"`
	TermsOfAddress        bool `json:"TermsOfAddress"`
	ExpressLanguage       bool `json:"ExpressLanguage"`
	CameraButton          bool `json:"CameraButton"`
	AppleID               bool `json:"AppleID"`
	DisplayTone           bool `json:"DisplayTone"`
	WatchMigration        bool `json:"WatchMigration"`
	UpdateCompleted       bool `json:"UpdateCompleted"`
	Appearance            bool `json:"Appearance"`
	Android               bool `json:"Android"`
	Payment               bool `json:"Payment"`
	OnBoarding            bool `json:"OnBoarding"`
	TOS                   bool `json:"TOS"`
	Welcome               bool `json:"Welcome"`
	TapToSetup            bool `json:"TapToSetup"`
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
	PrestageDeviceNames    []MobileDevicePrestageSubsetNamesName `json:"prestageDeviceNames"`
	DeviceNamePrefix       string                                `json:"deviceNamePrefix"`
	DeviceNameSuffix       string                                `json:"deviceNameSuffix"`
	SingleDeviceName       string                                `json:"singleDeviceName"`
	ManageNames            bool                                  `json:"manageNames"`
	DeviceNamingConfigured bool                                  `json:"deviceNamingConfigured"`
}

type MobileDevicePrestageSubsetNamesName struct {
	ID         string `json:"id"`
	DeviceName string `json:"deviceName"`
	Used       bool   `json:"used"`
}

// CRUD

// GetMobileDevicePrestages retrieves a list of all mobile prestages
func (c *Client) GetMobileDevicePrestages(params url.Values) (*ResponseMobileDevicePrestagesList, error) {
	endpoint := uriMobileDevicePrestages
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
	endpoint := fmt.Sprintf("%s/%s", uriMobileDevicePrestages, id)
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

// CreateMobileDevicePrestage creates a new mobile prestage and returns the id
func (c *Client) CreateMobileDevicePrestage(newPrestage ResourceMobileDevicePrestage) (*ResponseMobileDevicePrestageCreate, error) {
	endpoint := uriMobileDevicePrestages
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

// DeleteMobileDevicePrestageByID a mobile prestage at the given id
func (c *Client) DeleteMobileDevicePrestageByID(id string) error {
	endpoint := fmt.Sprintf("%s/%s", uriMobileDevicePrestages, id)
	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "mobile device prestage", id, err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return nil
}

// QUERY Which other endpoints required here? I think something to do with the scopes & syncs but strugging to make sense of it.
