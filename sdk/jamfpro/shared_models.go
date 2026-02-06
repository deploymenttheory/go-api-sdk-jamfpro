// shared_models.go
package jamfpro

// Shared Standalone Resources

type SharedResourceSite struct {
	ID   int    `json:"id,omitempty" xml:"id,omitempty"`
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

type SharedResourceSiteProAPI struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type SharedResourceCategory struct {
	ID   int    `json:"id,omitempty" xml:"id,omitempty"`
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

type SharedResourceSelfServiceIcon struct {
	ID       int    `json:"id,omitempty" xml:"id,omitempty"`
	URI      string `json:"uri,omitempty" xml:"uri,omitempty"`
	Data     string `json:"data,omitempty" xml:"data,omitempty"`
	Filename string `json:"filename,omitempty" xml:"filename,omitempty"`
}

type SharedResourceSelfServiceCategories struct {
	Category []SharedResourceSelfServiceCategory
}

type SharedResourceSelfServiceCategory struct {
	ID       int    `json:"id,omitempty" xml:"id,omitempty"`
	Name     string `json:"name,omitempty" xml:"name,omitempty"`
	Priority int    `json:"priority,omitempty" xml:"priority,omitempty"`
}

// Advanced Search DisplayField
type DisplayField struct {
	Name string `xml:"name"`
}

type SharedAdvancedSearchSubsetDisplayField struct {
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

// Criteria

type SharedContainerCriteria struct {
	Size      int                     `json:"size,omitempty" xml:"size,omitempty"`
	Criterion *[]SharedSubsetCriteria `json:"criterion,omitempty" xml:"criterion,omitempty"`
}

type SharedSubsetCriteria struct {
	Name         string `json:"name,omitempty" xml:"name,omitempty"`
	Priority     int    `json:"priority,omitempty" xml:"priority,omitempty"`
	AndOr        string `json:"and_or,omitempty" xml:"and_or,omitempty"`
	SearchType   string `json:"search_type,omitempty" xml:"search_type,omitempty"`
	Value        string `json:"value,omitempty" xml:"value,omitempty"`
	OpeningParen bool   `json:"opening_paren,omitempty" xml:"opening_paren,omitempty"`
	ClosingParen bool   `json:"closing_paren,omitempty" xml:"closing_paren,omitempty"`
}

// SharedSubsetCriteriaJamfProAPI represents the criteria for an Search item for jamfprom api
type SharedSubsetCriteriaJamfProAPI struct {
	Name         string `json:"name"`
	Priority     int    `json:"priority"`
	AndOr        string `json:"andOr"`
	SearchType   string `json:"searchType"`
	Value        string `json:"value"`
	OpeningParen *bool  `json:"openingParen,omitempty"`
	ClosingParen *bool  `json:"closingParen,omitempty"`
}

type SharedResourceLdapServer struct {
	ID   int    `json:"id,omitempty" xml:"id,omitempty"`
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

// ResponseError represents the structure of the error response when the force stop request fails.
type SharedResourcResponseError struct {
	HTTPStatus int                         `json:"httpStatus"`
	Errors     []SharedResourceErrorDetail `json:"errors"`
}

// ErrorDetail represents the details of an error in the response
// Used by jamfpro api MDM and managed software updates
type SharedResourceErrorDetail struct {
	Device      int    `json:"device"`
	Group       int    `json:"group"`
	Reason      string `json:"reason"`
	Code        string `json:"code"`
	Field       string `json:"field"`
	Description string `json:"description"`
	ID          string `json:"id"`
}

// ResponseError represents the structure of an error response from the API
type ResponseError struct {
	HTTPStatus int             `json:"httpStatus"`
	Errors     []ErrorInstance `json:"errors"`
}

// ErrorInstance represents a single error in the error response
type ErrorInstance struct {
	Code        string `json:"code"`
	Field       string `json:"field"`
	Description string `json:"description"`
	ID          string `json:"id"`
}

type SharedResourceInventoryListMobileDevice struct {
	MobileDeviceId                              string                                  `json:"mobileDeviceId,omitempty"`
	Udid                                        string                                  `json:"udid,omitempty"`
	AirPlayPassword                             string                                  `json:"airPlayPassword,omitempty"`
	AppAnalyticsEnabled                         bool                                    `json:"appAnalyticsEnabled,omitempty"`
	AssetTag                                    string                                  `json:"assetTag,omitempty"`
	AvailableSpaceMb                            int                                     `json:"availableSpaceMb,omitempty"`
	BatteryLevel                                int                                     `json:"batteryLevel,omitempty"`
	BatteryHealth                               string                                  `json:"batteryHealth,omitempty"`
	BluetoothLowEnergyCapable                   bool                                    `json:"bluetoothLowEnergyCapable,omitempty"`
	BluetoothMacAddress                         string                                  `json:"bluetoothMacAddress,omitempty"`
	CapacityMb                                  int                                     `json:"capacityMb,omitempty"`
	LostModeEnabledDate                         string                                  `json:"lostModeEnabledDate,omitempty"`
	DeclarativeDeviceManagementEnabled          bool                                    `json:"declarativeDeviceManagementEnabled,omitempty"`
	DeviceId                                    string                                  `json:"deviceId,omitempty"`
	DeviceLocatorServiceEnabled                 bool                                    `json:"deviceLocatorServiceEnabled,omitempty"`
	DeviceOwnershipType                         string                                  `json:"deviceOwnershipType,omitempty"`
	DevicePhoneNumber                           string                                  `json:"devicePhoneNumber,omitempty"`
	DiagnosticAndUsageReportingEnabled          bool                                    `json:"diagnosticAndUsageReportingEnabled,omitempty"`
	DisplayName                                 string                                  `json:"displayName,omitempty"`
	DoNotDisturbEnabled                         bool                                    `json:"doNotDisturbEnabled,omitempty"`
	EnrollmentSessionTokenValid                 bool                                    `json:"enrollmentSessionTokenValid,omitempty"`
	ExchangeDeviceId                            string                                  `json:"exchangeDeviceId,omitempty"`
	CloudBackupEnabled                          bool                                    `json:"cloudBackupEnabled,omitempty"`
	OsBuild                                     string                                  `json:"osBuild,omitempty"`
	OsSupplementalBuildVersion                  string                                  `json:"osSupplementalBuildVersion,omitempty"`
	OsRapidSecurityResponse                     string                                  `json:"osRapidSecurityResponse,omitempty"`
	OsVersion                                   string                                  `json:"osVersion,omitempty"`
	IpAddress                                   string                                  `json:"ipAddress,omitempty"`
	ItunesStoreAccountActive                    bool                                    `json:"itunesStoreAccountActive,omitempty"`
	JamfParentPairings                          int                                     `json:"jamfParentPairings,omitempty"`
	Languages                                   string                                  `json:"languages,omitempty"`
	LastBackupDate                              string                                  `json:"lastBackupDate,omitempty"`
	LastEnrolledDate                            string                                  `json:"lastEnrolledDate,omitempty"`
	LastCloudBackupDate                         string                                  `json:"lastCloudBackupDate,omitempty"`
	LastInventoryUpdateDate                     string                                  `json:"lastInventoryUpdateDate,omitempty"`
	Locales                                     string                                  `json:"locales,omitempty"`
	LocationServicesForSelfServiceMobileEnabled bool                                    `json:"locationServicesForSelfServiceMobileEnabled,omitempty"`
	LostModeEnabled                             bool                                    `json:"lostModeEnabled,omitempty"`
	Managed                                     bool                                    `json:"managed,omitempty"`
	ManagementId                                string                                  `json:"managementId,omitempty"`
	MdmProfileExpirationDate                    string                                  `json:"mdmProfileExpirationDate,omitempty"`
	Model                                       string                                  `json:"model,omitempty"`
	ModelIdentifier                             string                                  `json:"modelIdentifier,omitempty"`
	ModelNumber                                 string                                  `json:"modelNumber,omitempty"`
	ModemFirmwareVersion                        string                                  `json:"modemFirmwareVersion,omitempty"`
	PairedDevices                               int                                     `json:"pairedDevices,omitempty"`
	QuotaSize                                   int                                     `json:"quotaSize,omitempty"`
	ResidentUsers                               int                                     `json:"residentUsers,omitempty"`
	SerialNumber                                string                                  `json:"serialNumber,omitempty"`
	SharedIpad                                  bool                                    `json:"sharedIpad,omitempty"`
	Supervised                                  bool                                    `json:"supervised,omitempty"`
	Tethered                                    bool                                    `json:"tethered,omitempty"`
	TimeZone                                    string                                  `json:"timeZone,omitempty"`
	UsedSpacePercentage                         int                                     `json:"usedSpacePercentage,omitempty"`
	WifiMacAddress                              string                                  `json:"wifiMacAddress,omitempty"`
	Building                                    string                                  `json:"building,omitempty"`
	Department                                  string                                  `json:"department,omitempty"`
	EmailAddress                                string                                  `json:"emailAddress,omitempty"`
	FullName                                    string                                  `json:"fullName,omitempty"`
	Position                                    string                                  `json:"position,omitempty"`
	Room                                        string                                  `json:"room,omitempty"`
	UserPhoneNumber                             string                                  `json:"userPhoneNumber,omitempty"`
	Username                                    string                                  `json:"username,omitempty"`
	AppleCareId                                 string                                  `json:"appleCareId,omitempty"`
	LeaseExpirationDate                         string                                  `json:"leaseExpirationDate,omitempty"`
	LifeExpectancyYears                         int                                     `json:"lifeExpectancyYears,omitempty"`
	PoDate                                      string                                  `json:"poDate,omitempty"`
	PoNumber                                    string                                  `json:"poNumber,omitempty"`
	PurchasePrice                               string                                  `json:"purchasePrice,omitempty"`
	PurchasedOrLeased                           bool                                    `json:"purchasedOrLeased,omitempty"`
	PurchasingAccount                           string                                  `json:"purchasingAccount,omitempty"`
	PurchasingContact                           string                                  `json:"purchasingContact,omitempty"`
	Vendor                                      string                                  `json:"vendor,omitempty"`
	WarrantyExpirationDate                      string                                  `json:"warrantyExpirationDate,omitempty"`
	ActivationLockEnabled                       bool                                    `json:"activationLockEnabled,omitempty"`
	BlockEncryptionCapable                      bool                                    `json:"blockEncryptionCapable,omitempty"`
	DataProtection                              bool                                    `json:"dataProtection,omitempty"`
	FileEncryptionCapable                       bool                                    `json:"fileEncryptionCapable,omitempty"`
	HardwareEncryptionSupported                 bool                                    `json:"hardwareEncryptionSupported,omitempty"`
	JailbreakStatus                             string                                  `json:"jailbreakStatus,omitempty"`
	PasscodeCompliant                           bool                                    `json:"passcodeCompliant,omitempty"`
	PasscodeCompliantWithProfile                bool                                    `json:"passcodeCompliantWithProfile,omitempty"`
	PasscodeLockGracePeriodEnforcedSeconds      int                                     `json:"passcodeLockGracePeriodEnforcedSeconds,omitempty"`
	PasscodePresent                             bool                                    `json:"passcodePresent,omitempty"`
	PersonalDeviceProfileCurrent                bool                                    `json:"personalDeviceProfileCurrent,omitempty"`
	CarrierSettingsVersion                      string                                  `json:"carrierSettingsVersion,omitempty"`
	CellularTechnology                          string                                  `json:"cellularTechnology,omitempty"`
	CurrentCarrierNetwork                       string                                  `json:"currentCarrierNetwork,omitempty"`
	CurrentMobileCountryCode                    string                                  `json:"currentMobileCountryCode,omitempty"`
	CurrentMobileNetworkCode                    string                                  `json:"currentMobileNetworkCode,omitempty"`
	DataRoamingEnabled                          bool                                    `json:"dataRoamingEnabled,omitempty"`
	Eid                                         string                                  `json:"eid,omitempty"`
	HomeCarrierNetwork                          string                                  `json:"homeCarrierNetwork,omitempty"`
	HomeMobileCountryCode                       string                                  `json:"homeMobileCountryCode,omitempty"`
	HomeMobileNetworkCode                       string                                  `json:"homeMobileNetworkCode,omitempty"`
	Iccid                                       string                                  `json:"iccid,omitempty"`
	Imei                                        string                                  `json:"imei,omitempty"`
	Imei2                                       string                                  `json:"imei2,omitempty"`
	Meid                                        string                                  `json:"meid,omitempty"`
	PersonalHotspotEnabled                      bool                                    `json:"personalHotspotEnabled,omitempty"`
	PreferredVoiceNumber                        string                                  `json:"preferredVoiceNumber,omitempty"`
	Roaming                                     bool                                    `json:"roaming,omitempty"`
	VoiceRoamingEnabled                         string                                  `json:"voiceRoamingEnabled,omitempty"`
	LastLoggedInUsernameSelfService             string                                  `json:"lastLoggedInUsernameSelfService,omitempty"`
	LastLoggedInUsernameSelfServiceTimestamp    string                                  `json:"lastLoggedInUsernameSelfServiceTimestamp,omitempty"`
	ExtensionAttributeValueList                 []SharedResourceExtensionAttributeValue `json:"extensionAttributeValueList,omitempty"`
}

// SharedResourceExtensionAttributeValue represents a single extension attribute value returned in the inventory
type SharedResourceExtensionAttributeValue struct {
	DisplayName string `json:"displayName"`
	Value       string `json:"value,omitempty"`
}
