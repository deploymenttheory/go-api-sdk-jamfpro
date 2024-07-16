// classicapi_mobile_devices.go
// Jamf Pro Classic Api - Mobile Devices
// API reference: https://developer.jamf.com/jamf-pro/reference/mobiledevices
// Jamf Pro Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriMobileDevices = "/JSSResource/mobiledevices"

// List

// ResponseMobileDevicesList represents the structure for a list of mobile devices.
type ResponseMobileDeviceList struct {
	MobileDevices []MobileDeviceListItem `xml:"mobile_device"`
}

type MobileDeviceListItem struct {
	ID              int    `xml:"id"`
	Name            string `xml:"name"`
	DeviceName      string `xml:"device_name"`
	UDID            string `xml:"udid"`
	SerialNumber    string `xml:"serial_number"`
	PhoneNumber     string `xml:"phone_number"`
	WifiMacAddress  string `xml:"wifi_mac_address"`
	Managed         bool   `xml:"managed"`
	Supervised      bool   `xml:"supervised"`
	Model           string `xml:"model"`
	ModelIdentifier string `xml:"model_identifier"`
	ModelDisplay    string `xml:"model_display"`
	Username        string `xml:"username"`
}

// Resource

// ResourceMobileDevice represents the structure for a of a mobile device.
type ResourceMobileDevice struct {
	General               MobileDeviceSubsetGeneral                `xml:"general"`
	Location              MobileDeviceSubsetLocation               `xml:"location"`
	Purchasing            MobileDeviceSubsetPurchasing             `xml:"purchasing"`
	Applications          []MobileDeviceSubsetApplication          `xml:"applications>application"`
	SecurityObject        MobileDeviceSubsetSecurity               `xml:"security_object"`
	Network               MobileDeviceSubsetNetwork                `xml:"network"`
	Certificates          []MobileDeviceSubsetCertificate          `xml:"certificates>certificate"`
	ConfigurationProfiles []MobileDeviceSubsetConfigurationProfile `xml:"configuration_profiles>configuration_profile"`
	ProvisioningProfiles  []MobileDeviceSubsetProvisioningProfile  `xml:"provisioning_profiles>mobile_device_provisioning_profile"`
	MobileDeviceGroups    []MobileDeviceSubsetGroup                `xml:"mobile_device_groups>mobile_device_group"`
	ExtensionAttributes   []MobileDeviceSubsetExtensionAttribute   `xml:"extension_attributes>extension_attribute"`
}

// Subsets & Containers

type MobileDeviceSubsetGeneral struct {
	ID                                 int    `xml:"id"`
	DisplayName                        string `xml:"display_name"`
	DeviceName                         string `xml:"device_name"`
	Name                               string `xml:"name"`
	AssetTag                           string `xml:"asset_tag"`
	LastInventoryUpdate                string `xml:"last_inventory_update"`
	LastInventoryUpdateEpoch           int64  `xml:"last_inventory_update_epoch"`
	LastInventoryUpdateUTC             string `xml:"last_inventory_update_utc"`
	Capacity                           int    `xml:"capacity"`
	CapacityMB                         int    `xml:"capacity_mb"`
	Available                          int    `xml:"available"`
	AvailableMB                        int    `xml:"available_mb"`
	PercentageUsed                     int    `xml:"percentage_used"`
	OSType                             string `xml:"os_type"`
	OSVersion                          string `xml:"os_version"`
	OSBuild                            string `xml:"os_build"`
	SerialNumber                       string `xml:"serial_number"`
	UDID                               string `xml:"udid"`
	InitialEntryDateEpoch              int64  `xml:"initial_entry_date_epoch"`
	InitialEntryDateUTC                string `xml:"initial_entry_date_utc"`
	PhoneNumber                        string `xml:"phone_number"`
	IPAddress                          string `xml:"ip_address"`
	WifiMacAddress                     string `xml:"wifi_mac_address"`
	BluetoothMacAddress                string `xml:"bluetooth_mac_address"`
	ModemFirmware                      string `xml:"modem_firmware"`
	Model                              string `xml:"model"`
	ModelIdentifier                    string `xml:"model_identifier"`
	ModelNumber                        string `xml:"model_number"`
	ModelDisplay                       string `xml:"model_display"`
	DeviceOwnershipLevel               string `xml:"device_ownership_level"`
	LastEnrollmentEpoch                int64  `xml:"last_enrollment_epoch"`
	LastEnrollmentUTC                  string `xml:"last_enrollment_utc"`
	Managed                            bool   `xml:"managed"`
	Supervised                         bool   `xml:"supervised"`
	ExchangeActiveSyncDeviceIdentifier string `xml:"exchange_activesync_device_identifier"`
	Shared                             string `xml:"shared"`
	Tethered                           string `xml:"tethered"`
	BatteryLevel                       int    `xml:"battery_level"`
	BLECapable                         bool   `xml:"ble_capable"`
	DeviceLocatorServiceEnabled        bool   `xml:"device_locator_service_enabled"`
	DoNotDisturbEnabled                bool   `xml:"do_not_disturb_enabled"`
	CloudBackupEnabled                 bool   `xml:"cloud_backup_enabled"`
	LastCloudBackupDateEpoch           int64  `xml:"last_cloud_backup_date_epoch"`
	LastCloudBackupDateUTC             string `xml:"last_cloud_backup_date_utc"`
	LocationServicesEnabled            bool   `xml:"location_services_enabled"`
	ItunesStoreAccountIsActive         bool   `xml:"itunes_store_account_is_active"`
	LastBackupTimeEpoch                int64  `xml:"last_backup_time_epoch"`
	LastBackupTimeUTC                  string `xml:"last_backup_time_utc"`
}

type MobileDeviceSubsetLocation struct {
	Username     string `xml:"username"`
	RealName     string `xml:"realname"`
	EmailAddress string `xml:"email_address"`
	Position     string `xml:"position"`
	Phone        string `xml:"phone"`
	PhoneNumber  string `xml:"phone_number"`
	Department   string `xml:"department"`
	Building     string `xml:"building"`
	Room         int    `xml:"room"`
}

type MobileDeviceSubsetPurchasing struct {
	IsPurchased          bool   `xml:"is_purchased"`
	IsLeased             bool   `xml:"is_leased"`
	PONumber             string `xml:"po_number"`
	Vendor               string `xml:"vendor"`
	ApplecareID          string `xml:"applecare_id"`
	PurchasePrice        string `xml:"purchase_price"`
	PurchasingAccount    string `xml:"purchasing_account"`
	PODate               string `xml:"po_date"`
	PODateEpoch          int64  `xml:"po_date_epoch"`
	PODateUTC            string `xml:"po_date_utc"`
	WarrantyExpires      string `xml:"warranty_expires"`
	WarrantyExpiresEpoch int64  `xml:"warranty_expires_epoch"`
	WarrantyExpiresUTC   string `xml:"warranty_expires_utc"`
	LeaseExpires         string `xml:"lease_expires"`
	LeaseExpiresEpoch    int64  `xml:"lease_expires_epoch"`
	LeaseExpiresUTC      string `xml:"lease_expires_utc"`
	LifeExpectancy       int    `xml:"life_expectancy"`
	PurchasingContact    string `xml:"purchasing_contact"`
}

type MobileDeviceSubsetApplication struct {
	ApplicationName    string `xml:"application_name"`
	ApplicationVersion string `xml:"application_version"`
	Identifier         string `xml:"identifier"`
}

type MobileDeviceSubsetSecurity struct {
	DataProtection                  bool    `xml:"data_protection"`
	BlockLevelEncryptionCapable     bool    `xml:"block_level_encryption_capable"`
	FileLevelEncryptionCapable      bool    `xml:"file_level_encryption_capable"`
	PasscodePresent                 bool    `xml:"passcode_present"`
	PasscodeCompliant               bool    `xml:"passcode_compliant"`
	PasscodeCompliantWithProfile    bool    `xml:"passcode_compliant_with_profile"`
	PasscodeLockGracePeriodEnforced string  `xml:"passcode_lock_grace_period_enforced"`
	HardwareEncryption              string  `xml:"hardware_encryption"`
	ActivationLockEnabled           bool    `xml:"activation_lock_enabled"`
	JailbreakDetected               string  `xml:"jailbreak_detected"`
	LostModeEnabled                 bool    `xml:"lost_mode_enabled"`
	LostModeEnforced                bool    `xml:"lost_mode_enforced"`
	LostModeEnableIssuedEpoch       int64   `xml:"lost_mode_enable_issued_epoch"`
	LostModeEnableIssuedUTC         string  `xml:"lost_mode_enable_issued_utc"`
	LostModeMessage                 string  `xml:"lost_mode_message"`
	LostModePhone                   string  `xml:"lost_mode_phone"`
	LostModeFootnote                string  `xml:"lost_mode_footnote"`
	LostLocationEpoch               int64   `xml:"lost_location_epoch"`
	LostLocationUTC                 string  `xml:"lost_location_utc"`
	LostLocationLatitude            float64 `xml:"lost_location_latitude"`
	LostLocationLongitude           float64 `xml:"lost_location_longitude"`
	LostLocationAltitude            float64 `xml:"lost_location_altitude"`
	LostLocationSpeed               float64 `xml:"lost_location_speed"`
	LostLocationCourse              float64 `xml:"lost_location_course"`
	LostLocationHorizontalAccuracy  float64 `xml:"lost_location_horizontal_accuracy"`
	LostLocationVerticalAccuracy    float64 `xml:"lost_location_vertical_accuracy"`
}

type MobileDeviceSubsetNetwork struct {
	HomeCarrierNetwork       string `xml:"home_carrier_network"`
	CellularTechnology       string `xml:"cellular_technology"`
	VoiceRoamingEnabled      string `xml:"voice_roaming_enabled"`
	IMEI                     string `xml:"imei"`
	ICCID                    string `xml:"iccid"`
	CurrentCarrierNetwork    string `xml:"current_carrier_network"`
	CarrierSettingsVersion   int    `xml:"carrier_settings_version"`
	CurrentMobileCountryCode int    `xml:"current_mobile_country_code"`
	CurrentMobileNetworkCode int    `xml:"current_mobile_network_code"`
	HomeMobileCountryCode    int    `xml:"home_mobile_country_code"`
	HomeMobileNetworkCode    int    `xml:"home_mobile_network_code"`
	DataRoamingEnabled       bool   `xml:"data_roaming_enabled"`
	PhoneNumber              string `xml:"phone_number"`
}

type MobileDeviceSubsetCertificate struct {
	CommonName string `xml:"common_name"`
	Identity   bool   `xml:"identity"`
}

type MobileDeviceSubsetConfigurationProfile struct {
	DisplayName string `xml:"display_name"`
	Version     int    `xml:"version"`
	Identifier  string `xml:"identifier"`
	UUID        string `xml:"uuid"`
}

type MobileDeviceSubsetProvisioningProfile struct {
	DisplayName         string `xml:"display_name"`
	ExpirationDate      string `xml:"expiration_date"`
	ExpirationDateEpoch int64  `xml:"expiration_date_epoch"`
	ExpirationDateUTC   string `xml:"expiration_date_utc"`
	UUID                string `xml:"uuid"`
}

type MobileDeviceSubsetGroup struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type MobileDeviceSubsetExtensionAttribute struct {
	ID    int    `xml:"id"`
	Name  string `xml:"name"`
	Type  string `xml:"type"`
	Value string `xml:"value"`
}

// CRUD

// GetMobileDevices retrieves a list of all mobile devices.
func (c *Client) GetMobileDevices() (*ResponseMobileDeviceList, error) {
	endpoint := uriMobileDevices

	var mobileDevices ResponseMobileDeviceList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &mobileDevices)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "mobile devices", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &mobileDevices, nil
}

// GetMobileDeviceByID retrieves a specific mobile device by its ID.
func (c *Client) GetMobileDeviceByID(id int) (*ResourceMobileDevice, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriMobileDevices, id)

	var device ResourceMobileDevice
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &device)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "mobile device", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &device, nil
}

// GetMobileDeviceByName retrieves a specific mobile device by its name.
func (c *Client) GetMobileDeviceByName(name string) (*ResourceMobileDevice, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMobileDevices, name)

	var device ResourceMobileDevice
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &device)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "mobile device", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &device, nil
}

// GetMobileDeviceByIDAndDataSubset retrieves a specific subset of data for a mobile device by its ID.
func (c *Client) GetMobileDeviceByIDAndDataSubset(id int, subset string) (*ResourceMobileDevice, error) {
	endpoint := fmt.Sprintf("%s/id/%d/subset/%s", uriMobileDevices, id, subset)

	var deviceSubset ResourceMobileDevice
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &deviceSubset)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "mobile device with data subset", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &deviceSubset, nil
}

// GetMobileDeviceByNameAndDataSubset retrieves a specific subset of data for a mobile device by its name.
func (c *Client) GetMobileDeviceByNameAndDataSubset(name, subset string) (*ResourceMobileDevice, error) {
	endpoint := fmt.Sprintf("%s/name/%s/subset/%s", uriMobileDevices, name, subset)

	var deviceSubset ResourceMobileDevice
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &deviceSubset)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "mobile device with data subset", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &deviceSubset, nil
}

// CreateMobileDevice creates a new mobile device device.
func (c *Client) CreateMobileDevice(attribute *ResourceMobileDevice) (*ResourceMobileDevice, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriMobileDevices)

	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device"`
		*ResourceMobileDevice
	}{
		ResourceMobileDevice: attribute,
	}

	var responseAttribute ResourceMobileDevice
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &responseAttribute)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "mobile device", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseAttribute, nil
}

// UpdateMobileDeviceByID updates a mobile device by its ID.
func (c *Client) UpdateMobileDeviceByID(id int, attribute *ResourceMobileDevice) (*ResourceMobileDevice, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriMobileDevices, id)

	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device"`
		*ResourceMobileDevice
	}{
		ResourceMobileDevice: attribute,
	}

	var responseAttribute ResourceMobileDevice
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responseAttribute)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "mobile device", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseAttribute, nil
}

// UpdateMobileDeviceByName updates a mobile device by its name.
func (c *Client) UpdateMobileDeviceByName(name string, attribute *ResourceMobileDevice) (*ResourceMobileDevice, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMobileDevices, name)

	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device"`
		*ResourceMobileDevice
	}{
		ResourceMobileDevice: attribute,
	}

	var responseAttribute ResourceMobileDevice
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responseAttribute)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "mobile device", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseAttribute, nil
}

// DeleteMobileDeviceByID deletes a mobile device by its ID.
func (c *Client) DeleteMobileDeviceByID(id string) error {
	endpoint := fmt.Sprintf("%s/id/%s", uriMobileDevices, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "mobile device", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteMobileDeviceByName deletes a mobile device by its name.
func (c *Client) DeleteMobileDeviceByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriMobileDevices, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByName, "mobile device", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
