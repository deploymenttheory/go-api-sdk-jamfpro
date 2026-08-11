// jamfproapi_mobile_devices_inventory.go
// Jamf Pro Api - Mobile Device Inventory
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v2-mobile-devices-detail
// Jamf Pro API requires the structs to support a JSON data structure.

/*
Shared Resources in this Endpoint:
- SharedResourceInventoryListMobileDevice
*/

package jamfpro

import (
	"fmt"
	"net/url"

	"github.com/mitchellh/mapstructure"
)

const uriMobileDevicesInventory = "/api/v2/mobile-devices/detail"

// mobileDeviceInventorySections contains all available sections for mobile device inventory API requests
var mobileDeviceInventorySections = []string{
	"GENERAL", "HARDWARE", "USER_AND_LOCATION", "PURCHASING", "SECURITY",
	"APPLICATIONS", "EBOOKS", "NETWORK", "SERVICE_SUBSCRIPTIONS", "CERTIFICATES",
	"PROFILES", "USER_PROFILES", "PROVISIONING_PROFILES", "SHARED_USERS",
	"GROUPS", "EXTENSION_ATTRIBUTES",
}

// Response Structures

type ResourceMobileDeviceInventory struct {
	MobileDeviceId                     string                                  `json:"mobileDeviceId,omitempty"`
	DeviceType                         string                                  `json:"deviceType,omitempty"`
	Hardware                           MobileDeviceHardware                    `json:"hardware,omitempty"`
	General                            MobileDeviceGeneral                     `json:"general,omitempty"`
	UserAndLocation                    MobileDeviceUserAndLocation             `json:"userAndLocation,omitempty"`
	Purchasing                         MobileDevicePurchasing                  `json:"purchasing,omitempty"`
	Security                           MobileDeviceSecurity                    `json:"security,omitempty"`
	Network                            MobileDeviceNetwork                     `json:"network,omitempty"`
	Applications                       []MobileDeviceApplication               `json:"applications,omitempty"`
	Ebooks                             []MobileDeviceEbook                     `json:"ebooks,omitempty"`
	ServiceSubscriptions               []MobileDeviceServiceSubscription       `json:"serviceSubscriptions,omitempty"`
	Certificates                       []MobileDeviceCertificate               `json:"certificates,omitempty"`
	Profiles                           []MobileDeviceProfile                   `json:"profiles,omitempty"`
	UserProfiles                       []MobileDeviceProfile                   `json:"userProfiles,omitempty"`
	ProvisioningProfiles               []MobileDeviceProvisioningProfile       `json:"provisioningProfiles,omitempty"`
	SharedUsers                        []MobileDeviceSharedUser                `json:"sharedUsers,omitempty"`
	Groups                             []MobileDeviceGroup                     `json:"groups,omitempty"`
	ExtensionAttributes                []SharedResourceExtensionAttributeValue `json:"extensionAttributes,omitempty"`
}

type MobileDeviceHardware struct {
	CapacityMb                int    `json:"capacityMb,omitempty"`
	AvailableSpaceMb          int    `json:"availableSpaceMb,omitempty"`
	UsedSpacePercentage       int    `json:"usedSpacePercentage,omitempty"`
	BatteryLevel              int    `json:"batteryLevel,omitempty"`
	BatteryHealth             string `json:"batteryHealth,omitempty"`
	SerialNumber              string `json:"serialNumber,omitempty"`
	WifiMacAddress            string `json:"wifiMacAddress,omitempty"`
	BluetoothMacAddress       string `json:"bluetoothMacAddress,omitempty"`
	BluetoothLowEnergyCapable bool   `json:"bluetoothLowEnergyCapable,omitempty"`
	Model                     string `json:"model,omitempty"`
	ModelIdentifier           string `json:"modelIdentifier,omitempty"`
	ModelNumber               string `json:"modelNumber,omitempty"`
}

type MobileDeviceGeneral struct {
	Udid                                string `json:"udid,omitempty"`
	DisplayName                         string `json:"displayName,omitempty"`
	AssetTag                            string `json:"assetTag,omitempty"`
	SoftwareUpdateDeviceId              string `json:"softwareUpdateDeviceId,omitempty"`
	IpAddress                           string `json:"ipAddress,omitempty"`
	Managed                             bool   `json:"managed,omitempty"`
	Supervised                          bool   `json:"supervised,omitempty"`
	DeviceOwnershipType                 string `json:"deviceOwnershipType,omitempty"`
	EnrollmentMethod                    string `json:"enrollmentMethod,omitempty"`
	EnrollmentSessionTokenValid         bool   `json:"enrollmentSessionTokenValid,omitempty"`
	LastEnrolledDate                    string `json:"lastEnrolledDate,omitempty"`
	MdmProfileExpirationDate            string `json:"mdmProfileExpirationDate,omitempty"`
	TimeZone                            string `json:"timeZone,omitempty"`
	DeclarativeDeviceManagementEnabled  bool   `json:"declarativeDeviceManagementEnabled,omitempty"`
	OsVersion                           string `json:"osVersion,omitempty"`
	OsBuild                             string `json:"osBuild,omitempty"`
	OsSupplementalBuildVersion          string `json:"osSupplementalBuildVersion,omitempty"`
	OsRapidSecurityResponse             string `json:"osRapidSecurityResponse,omitempty"`
	LastInventoryUpdateDate             string `json:"lastInventoryUpdateDate,omitempty"`
	LastCloudBackupDate                 string `json:"lastCloudBackupDate,omitempty"`
	LastBackupDate                      string `json:"lastBackupDate,omitempty"`
	CloudBackupEnabled                  bool   `json:"cloudBackupEnabled,omitempty"`
	DeviceLocatorServiceEnabled         bool   `json:"deviceLocatorServiceEnabled,omitempty"`
	DoNotDisturbEnabled                 bool   `json:"doNotDisturbEnabled,omitempty"`
	LostModeEnabled                     bool   `json:"lostModeEnabled,omitempty"`
	LostModeEnabledDate                 string `json:"lostModeEnabledDate,omitempty"`
	ItunesStoreAccountActive            bool   `json:"itunesStoreAccountActive,omitempty"`
	Languages                           string `json:"languages,omitempty"`
	Locales                             string `json:"locales,omitempty"`
	SharedIpad                          bool   `json:"sharedIpad,omitempty"`
	QuotaSize                           int    `json:"quotaSize,omitempty"`
	ResidentUsers                       int    `json:"residentUsers,omitempty"`
	DeviceId                            string `json:"deviceId,omitempty"`
	ManagementId                        string `json:"managementId,omitempty"`
	ExchangeDeviceId                    string `json:"exchangeDeviceId,omitempty"`
	Tethered                            bool   `json:"tethered,omitempty"`
}

type MobileDeviceUserAndLocation struct {
	Username        string `json:"username,omitempty"`
	FullName        string `json:"fullName,omitempty"`
	EmailAddress    string `json:"emailAddress,omitempty"`
	PhoneNumber     string `json:"phoneNumber,omitempty"`
	Position        string `json:"position,omitempty"`
	Department      string `json:"department,omitempty"`
	Building        string `json:"building,omitempty"`
	Room            string `json:"room,omitempty"`
}

type MobileDevicePurchasing struct {
	PurchasedOrLeased       bool   `json:"purchasedOrLeased,omitempty"`
	PoNumber                string `json:"poNumber,omitempty"`
	PoDate                  string `json:"poDate,omitempty"`
	Vendor                  string `json:"vendor,omitempty"`
	PurchasePrice           string `json:"purchasePrice,omitempty"`
	PurchasingAccount       string `json:"purchasingAccount,omitempty"`
	PurchasingContact       string `json:"purchasingContact,omitempty"`
	WarrantyExpirationDate  string `json:"warrantyExpirationDate,omitempty"`
	AppleCareId             string `json:"appleCareId,omitempty"`
	LeaseExpirationDate     string `json:"leaseExpirationDate,omitempty"`
	LifeExpectancyYears     int    `json:"lifeExpectancyYears,omitempty"`
}

type MobileDeviceSecurity struct {
	ActivationLockEnabled                  bool   `json:"activationLockEnabled,omitempty"`
	DataProtection                         bool   `json:"dataProtection,omitempty"`
	BlockEncryptionCapable                 bool   `json:"blockEncryptionCapable,omitempty"`
	FileEncryptionCapable                  bool   `json:"fileEncryptionCapable,omitempty"`
	HardwareEncryptionSupported            bool   `json:"hardwareEncryptionSupported,omitempty"`
	PasscodePresent                        bool   `json:"passcodePresent,omitempty"`
	PasscodeCompliant                      bool   `json:"passcodeCompliant,omitempty"`
	PasscodeCompliantWithProfile           bool   `json:"passcodeCompliantWithProfile,omitempty"`
	PasscodeLockGracePeriodEnforcedSeconds int    `json:"passcodeLockGracePeriodEnforcedSeconds,omitempty"`
	PersonalDeviceProfileCurrent           bool   `json:"personalDeviceProfileCurrent,omitempty"`
	JailbreakStatus                        string `json:"jailbreakStatus,omitempty"`
}

type MobileDeviceNetwork struct {
	CellularTechnology       string `json:"cellularTechnology,omitempty"`
	VoiceRoamingEnabled      bool   `json:"voiceRoamingEnabled,omitempty"`
	Imei                     string `json:"imei,omitempty"`
	Imei2                    string `json:"imei2,omitempty"`
	Iccid                    string `json:"iccid,omitempty"`
	Eid                      string `json:"eid,omitempty"`
	Meid                     string `json:"meid,omitempty"`
	CurrentCarrierNetwork    string `json:"currentCarrierNetwork,omitempty"`
	HomeCarrierNetwork       string `json:"homeCarrierNetwork,omitempty"`
	CurrentMobileCountryCode string `json:"currentMobileCountryCode,omitempty"`
	CurrentMobileNetworkCode string `json:"currentMobileNetworkCode,omitempty"`
	HomeMobileCountryCode    string `json:"homeMobileCountryCode,omitempty"`
	HomeMobileNetworkCode    string `json:"homeMobileNetworkCode,omitempty"`
	CarrierSettingsVersion   string `json:"carrierSettingsVersion,omitempty"`
	DataRoamingEnabled       bool   `json:"dataRoamingEnabled,omitempty"`
	Roaming                  bool   `json:"roaming,omitempty"`
	PersonalHotspotEnabled   bool   `json:"personalHotspotEnabled,omitempty"`
	DevicePhoneNumber        string `json:"devicePhoneNumber,omitempty"`
	ModemFirmwareVersion     string `json:"modemFirmwareVersion,omitempty"`
}

type MobileDeviceApplication struct {
	Name                     string `json:"name,omitempty"`
	Version                  string `json:"version,omitempty"`
	ShortVersion             string `json:"shortVersion,omitempty"`
	BundleSize               string `json:"bundleSize,omitempty"`
	DynamicSize              string `json:"dynamicSize,omitempty"`
	External                 bool   `json:"external,omitempty"`
	Identifier               string `json:"identifier,omitempty"`
}

type MobileDeviceEbook struct {
	Name       string `json:"name,omitempty"`
	Author     string `json:"author,omitempty"`
	Version    string `json:"version,omitempty"`
	Kind       string `json:"kind,omitempty"`
}

type MobileDeviceServiceSubscription struct {
	Name              string `json:"name,omitempty"`
	CarrierName       string `json:"carrierName,omitempty"`
	PhoneNumber       string `json:"phoneNumber,omitempty"`
	Iccid             string `json:"iccid,omitempty"`
	Eid               string `json:"eid,omitempty"`
	RoamingEnabled    bool   `json:"roamingEnabled,omitempty"`
}

type MobileDeviceCertificate struct {
	CommonName string `json:"commonName,omitempty"`
	Identity   bool   `json:"identity,omitempty"`
}

type MobileDeviceProfile struct {
	DisplayName string `json:"displayName,omitempty"`
	Version     string `json:"version,omitempty"`
	Uuid        string `json:"uuid,omitempty"`
	Identifier  string `json:"identifier,omitempty"`
}

type MobileDeviceProvisioningProfile struct {
	DisplayName    string `json:"displayName,omitempty"`
	ExpirationDate string `json:"expirationDate,omitempty"`
	Uuid           string `json:"uuid,omitempty"`
}

type MobileDeviceSharedUser struct {
	Username         string `json:"username,omitempty"`
	DataToSync       int    `json:"dataToSync,omitempty"`
	DataQuota        int    `json:"dataQuota,omitempty"`
	DataUsed         int    `json:"dataUsed,omitempty"`
}

type MobileDeviceGroup struct {
	GroupId string `json:"groupId,omitempty"`
	Name    string `json:"name,omitempty"`
}

type ResponseMobileDeviceInventoryList struct {
	TotalCount int                              `json:"totalCount"`
	Results    []ResourceMobileDeviceInventory  `json:"results"`
}

// CRUD

func (c *Client) GetMobileDevicesInventory(params url.Values) (*ResponseMobileDeviceInventoryList, error) {
	resp, err := c.DoPaginatedGet(uriMobileDevicesInventory, params)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "mobile-devices-inventories", err)
	}

	var out ResponseMobileDeviceInventoryList
	out.TotalCount = resp.Size

	for _, value := range resp.Results {
		var newObj ResourceMobileDeviceInventory
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "mobile-device-inventory", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil
}

func (c *Client) GetMobileDeviceInventoryByID(id string) (*ResourceMobileDeviceInventory, error) {
	params := url.Values{}
	params.Set("filter", fmt.Sprintf("mobileDeviceId==\"%s\"", id))
	for _, section := range mobileDeviceInventorySections {
		params.Add("section", section)
	}

	inventories, err := c.GetMobileDevicesInventory(params)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "mobile device inventory", err)
	}

	if len(inventories.Results) == 0 {
		return nil, fmt.Errorf("mobile device inventory with ID '%s' not found", id)
	}

	return &inventories.Results[0], nil
}

func (c *Client) GetMobileDeviceInventoryBySerialNumber(serialNumber string) (*ResourceMobileDeviceInventory, error) {
	params := url.Values{}
	params.Set("filter", fmt.Sprintf("serialNumber==\"%s\"", serialNumber))
	for _, section := range mobileDeviceInventorySections {
		params.Add("section", section)
	}

	inventories, err := c.GetMobileDevicesInventory(params)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "mobile device inventory", err)
	}

	if len(inventories.Results) == 0 {
		return nil, fmt.Errorf("failed to find mobile device inventory by serial number '%s'", serialNumber)
	}

	return &inventories.Results[0], nil
}

func (c *Client) GetMobileDeviceInventoryByName(name string) (*ResourceMobileDeviceInventory, error) {
	params := url.Values{}
	for _, section := range mobileDeviceInventorySections {
		params.Add("section", section)
	}

	inventories, err := c.GetMobileDevicesInventory(params)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "mobile device inventory", err)
	}

	for _, inventory := range inventories.Results {
		if inventory.General.DisplayName == name {
			return &inventory, nil
		}
	}

	return nil, fmt.Errorf("mobile device inventory with name '%s' not found", name)
}
