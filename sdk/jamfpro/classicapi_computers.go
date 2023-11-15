// classicapi_computers.go
// Jamf Pro Classic Api - Computers
// api reference: https://developer.jamf.com/jamf-pro/reference/computers
// Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriComputers = "/JSSResource/computers"

// Response structure for the list of computers
type ResponseComputersList struct {
	TotalCount int            `xml:"size"`
	Results    []ComputerItem `xml:"computer"`
}

type ComputerItem struct {
	ID   int    `xml:"id,omitempty" json:"id,omitempty"`
	Name string `xml:"name" json:"name"`
}

// Response structure for computer resources
type ResponseComputer struct {
	General               General                `xml:"general" json:"general"`
	Location              Location               `xml:"location" json:"location"`
	Purchasing            Purchasing             `xml:"purchasing" json:"purchasing"`
	Peripherals           Peripherals            `xml:"peripherals" json:"peripherals"`
	Hardware              Hardware               `xml:"hardware" json:"hardware"`
	Certificates          []Certificate          `xml:"certificates>certificate" json:"certificates"`
	Security              Security               `xml:"security" json:"security"`
	Software              Software               `xml:"software" json:"software"`
	ExtensionAttributes   []ExtensionAttribute   `xml:"extension_attributes>extension_attribute" json:"extension_attributes"`
	GroupsAccounts        GroupsAccounts         `xml:"groups_accounts" json:"groups_accounts"`
	ConfigurationProfiles []ConfigurationProfile `xml:"configuration_profiles>configuration_profile" json:"configuration_profiles"`
}

type General struct {
	ID                         int              `xml:"id" json:"id"`
	Name                       string           `xml:"name" json:"name"`
	MacAddress                 string           `xml:"mac_address" json:"mac_address"`
	NetworkAdapterType         string           `xml:"network_adapter_type" json:"network_adapter_type"`
	AltMacAddress              string           `xml:"alt_mac_address" json:"alt_mac_address"`
	AltNetworkAdapterType      string           `xml:"alt_network_adapter_type" json:"alt_network_adapter_type"`
	IPAddress                  string           `xml:"ip_address" json:"ip_address"`
	LastReportedIP             string           `xml:"last_reported_ip" json:"last_reported_ip"`
	SerialNumber               string           `xml:"serial_number" json:"serial_number"`
	UDID                       string           `xml:"udid" json:"udid"`
	JamfVersion                string           `xml:"jamf_version" json:"jamf_version"`
	Platform                   string           `xml:"platform" json:"platform"`
	Barcode1                   string           `xml:"barcode_1" json:"barcode_1"`
	Barcode2                   string           `xml:"barcode_2" json:"barcode_2"`
	AssetTag                   string           `xml:"asset_tag" json:"asset_tag"`
	RemoteManagement           RemoteManagement `xml:"remote_management" json:"remote_management"`
	MdmCapable                 bool             `xml:"mdm_capable" json:"mdm_capable"`
	MdmCapableUsers            MdmCapableUsers  `xml:"mdm_capable_users" json:"mdm_capable_users"`
	MdmProfileExpirationEpoch  int64            `xml:"mdm_profile_expiration_epoch" json:"mdm_profile_expiration_epoch"`
	MdmProfileExpirationUtc    string           `xml:"mdm_profile_expiration_utc" json:"mdm_profile_expiration_utc"`
	ManagementStatus           ManagementStatus `xml:"management_status" json:"management_status"`
	ReportDate                 string           `xml:"report_date" json:"report_date"`
	ReportDateEpoch            int64            `xml:"report_date_epoch" json:"report_date_epoch"`
	ReportDateUtc              string           `xml:"report_date_utc" json:"report_date_utc"`
	LastContactTime            string           `xml:"last_contact_time" json:"last_contact_time"`
	LastContactTimeEpoch       int64            `xml:"last_contact_time_epoch" json:"last_contact_time_epoch"`
	LastContactTimeUtc         string           `xml:"last_contact_time_utc" json:"last_contact_time_utc"`
	InitialEntryDate           string           `xml:"initial_entry_date" json:"initial_entry_date"`
	InitialEntryDateEpoch      int64            `xml:"initial_entry_date_epoch" json:"initial_entry_date_epoch"`
	InitialEntryDateUtc        string           `xml:"initial_entry_date_utc" json:"initial_entry_date_utc"`
	LastCloudBackupDateEpoch   int64            `xml:"last_cloud_backup_date_epoch" json:"last_cloud_backup_date_epoch"`
	LastCloudBackupDateUtc     string           `xml:"last_cloud_backup_date_utc" json:"last_cloud_backup_date_utc"`
	LastEnrolledDateEpoch      int64            `xml:"last_enrolled_date_epoch" json:"last_enrolled_date_epoch"`
	LastEnrolledDateUtc        string           `xml:"last_enrolled_date_utc" json:"last_enrolled_date_utc"`
	DistributionPoint          string           `xml:"distribution_point" json:"distribution_point"`
	Sus                        string           `xml:"sus" json:"sus"`
	Supervised                 bool             `xml:"supervised" json:"supervised"`
	Site                       Site             `xml:"site" json:"site"`
	ItunesStoreAccountIsActive bool             `xml:"itunes_store_account_is_active" json:"itunes_store_account_is_active"`
}

type RemoteManagement struct {
	Managed            bool   `xml:"managed" json:"managed"`
	ManagementUsername string `xml:"management_username" json:"management_username"`
}

type MdmCapableUsers struct {
	MdmCapableUser string `xml:"mdm_capable_user" json:"mdm_capable_user"`
}

type ManagementStatus struct {
	EnrolledViaDep         bool `xml:"enrolled_via_dep" json:"enrolled_via_dep"`
	UserApprovedEnrollment bool `xml:"user_approved_enrollment" json:"user_approved_enrollment"`
	UserApprovedMdm        bool `xml:"user_approved_mdm" json:"user_approved_mdm"`
}

// Location details of the computer
type Location struct {
	Username     string `xml:"username" json:"username"`
	RealName     string `xml:"realname" json:"realname"`
	EmailAddress string `xml:"email_address" json:"email_address"`
	Position     string `xml:"position" json:"position"`
	Phone        string `xml:"phone" json:"phone"`
	PhoneNumber  string `xml:"phone_number" json:"phone_number"`
	Department   string `xml:"department" json:"department"`
	Building     string `xml:"building" json:"building"`
	Room         string `xml:"room" json:"room"`
}

// Purchasing information of the computer
type Purchasing struct {
	IsPurchased          bool   `xml:"is_purchased" json:"is_purchased"`
	IsLeased             bool   `xml:"is_leased" json:"is_leased"`
	PoNumber             string `xml:"po_number" json:"po_number"`
	Vendor               string `xml:"vendor" json:"vendor"`
	ApplecareID          string `xml:"applecare_id" json:"applecare_id"`
	PurchasePrice        string `xml:"purchase_price" json:"purchase_price"`
	PurchasingAccount    string `xml:"purchasing_account" json:"purchasing_account"`
	PoDate               string `xml:"po_date" json:"po_date"`
	PoDateEpoch          int64  `xml:"po_date_epoch" json:"po_date_epoch"`
	PoDateUtc            string `xml:"po_date_utc" json:"po_date_utc"`
	WarrantyExpires      string `xml:"warranty_expires" json:"warranty_expires"`
	WarrantyExpiresEpoch int64  `xml:"warranty_expires_epoch" json:"warranty_expires_epoch"`
	WarrantyExpiresUtc   string `xml:"warranty_expires_utc" json:"warranty_expires_utc"`
	LeaseExpires         string `xml:"lease_expires" json:"lease_expires"`
	LeaseExpiresEpoch    int64  `xml:"lease_expires_epoch" json:"lease_expires_epoch"`
	LeaseExpiresUtc      string `xml:"lease_expires_utc" json:"lease_expires_utc"`
	LifeExpectancy       int    `xml:"life_expectancy" json:"life_expectancy"`
	PurchasingContact    string `xml:"purchasing_contact" json:"purchasing_contact"`
	OSAppleCareID        string `xml:"os_applecare_id,omitempty" json:"os_applecare_id,omitempty"`
	OSMaintenanceExpires string `xml:"os_maintenance_expires,omitempty" json:"os_maintenance_expires,omitempty"`
}

// Peripherals associated with the computer
type Peripherals struct {
	Size        int          `xml:"size" json:"size"`
	Peripherals []Peripheral `xml:"peripheral" json:"peripherals"`
}

type Peripheral struct {
	ID          int                    `xml:"id" json:"id"`
	BarCode1    string                 `xml:"bar_code_1" json:"bar_code_1"`
	BarCode2    string                 `xml:"bar_code_2" json:"bar_code_2"`
	Type        string                 `xml:"type" json:"type"`
	Fields      PeripheralFields       `xml:"fields" json:"fields"`
	Purchasing  PeripheralPurchasing   `xml:"purchasing" json:"purchasing"`
	Attachments []PeripheralAttachment `xml:"attachments>attachment" json:"attachments"`
}

type PeripheralFields struct {
	Field []Field `xml:"field" json:"field"`
}

type Field struct {
	Name  string `xml:"name" json:"name"`
	Value string `xml:"value" json:"value"`
}

// PeripheralPurchasing contains purchasing information for peripherals
type PeripheralPurchasing struct {
	IsPurchased          bool   `xml:"is_purchased" json:"is_purchased"`
	IsLeased             bool   `xml:"is_leased" json:"is_leased"`
	PoNumber             string `xml:"po_number" json:"po_number"`
	Vendor               string `xml:"vendor" json:"vendor"`
	ApplecareID          string `xml:"applecare_id" json:"applecare_id"`
	PurchasePrice        string `xml:"purchase_price" json:"purchase_price"`
	PurchasingAccount    string `xml:"purchasing_account" json:"purchasing_account"`
	PoDate               string `xml:"po_date" json:"po_date"`
	PoDateEpoch          int64  `xml:"po_date_epoch" json:"po_date_epoch"`
	PoDateUtc            string `xml:"po_date_utc" json:"po_date_utc"`
	WarrantyExpires      string `xml:"warranty_expires" json:"warranty_expires"`
	WarrantyExpiresEpoch int64  `xml:"warranty_expires_epoch" json:"warranty_expires_epoch"`
	WarrantyExpiresUtc   string `xml:"warranty_expires_utc" json:"warranty_expires_utc"`
	LeaseExpires         string `xml:"lease_expires" json:"lease_expires"`
	LeaseExpiresEpoch    int64  `xml:"lease_expires_epoch" json:"lease_expires_epoch"`
	LeaseExpiresUtc      string `xml:"lease_expires_utc" json:"lease_expires_utc"`
	LifeExpectancy       int    `xml:"life_expectancy" json:"life_expectancy"`
	PurchasingContact    string `xml:"purchasing_contact" json:"purchasing_contact"`
}

// Attachment represents an attachment to a peripheral
type PeripheralAttachment struct {
	Size     int    `xml:"size" json:"size"`
	ID       int    `xml:"id" json:"id"`
	Filename string `xml:"filename" json:"filename"`
	URI      string `xml:"uri" json:"uri"`
}

// Hardware details of the computer
type Hardware struct {
	Make                        string           `xml:"make" json:"make"`
	Model                       string           `xml:"model" json:"model"`
	ModelIdentifier             string           `xml:"model_identifier" json:"model_identifier"`
	OsName                      string           `xml:"os_name" json:"os_name"`
	OsVersion                   string           `xml:"os_version" json:"os_version"`
	OsBuild                     string           `xml:"os_build" json:"os_build"`
	MasterPasswordSet           bool             `xml:"master_password_set" json:"master_password_set"`
	ActiveDirectoryStatus       string           `xml:"active_directory_status" json:"active_directory_status"`
	ServicePack                 string           `xml:"service_pack" json:"service_pack"`
	ProcessorType               string           `xml:"processor_type" json:"processor_type"`
	ProcessorArchitecture       string           `xml:"processor_architecture" json:"processor_architecture"`
	ProcessorSpeed              int              `xml:"processor_speed" json:"processor_speed"`
	ProcessorSpeedMhz           int              `xml:"processor_speed_mhz" json:"processor_speed_mhz"`
	NumberProcessors            int              `xml:"number_processors" json:"number_processors"`
	NumberCores                 int              `xml:"number_cores" json:"number_cores"`
	TotalRam                    int              `xml:"total_ram" json:"total_ram"`
	TotalRamMb                  int              `xml:"total_ram_mb" json:"total_ram_mb"`
	BootRom                     string           `xml:"boot_rom" json:"boot_rom"`
	BusSpeed                    int              `xml:"bus_speed" json:"bus_speed"`
	BusSpeedMhz                 int              `xml:"bus_speed_mhz" json:"bus_speed_mhz"`
	BatteryCapacity             int              `xml:"battery_capacity" json:"battery_capacity"`
	CacheSize                   int              `xml:"cache_size" json:"cache_size"`
	CacheSizeKb                 int              `xml:"cache_size_kb" json:"cache_size_kb"`
	AvailableRamSlots           int              `xml:"available_ram_slots" json:"available_ram_slots"`
	OpticalDrive                string           `xml:"optical_drive" json:"optical_drive"`
	NicSpeed                    string           `xml:"nic_speed" json:"nic_speed"`
	SmcVersion                  string           `xml:"smc_version" json:"smc_version"`
	BleCapable                  bool             `xml:"ble_capable" json:"ble_capable"`
	SipStatus                   string           `xml:"sip_status" json:"sip_status"`
	GatekeeperStatus            string           `xml:"gatekeeper_status" json:"gatekeeper_status"`
	XprotectVersion             string           `xml:"xprotect_version" json:"xprotect_version"`
	InstitutionalRecoveryKey    string           `xml:"institutional_recovery_key" json:"institutional_recovery_key"`
	DiskEncryptionConfiguration string           `xml:"disk_encryption_configuration" json:"disk_encryption_configuration"`
	SoftwareUpdateDeviceID      string           `xml:"software_update_device_id,omitempty" json:"software_update_device_id,omitempty"`
	IsAppleSilicon              bool             `xml:"is_apple_silicon,omitempty" json:"is_apple_silicon,omitempty"`
	SupportsIosAppInstalls      bool             `xml:"supports_ios_app_installs,omitempty" json:"supports_ios_app_installs,omitempty"`
	Filevault2Users             []Filevault2User `xml:"filevault2_users>user" json:"filevault_2_users"`
	Storage                     []StorageDevice  `xml:"storage>device" json:"storage"`
	MappedPrinters              []Printer        `xml:"mapped_printers>printer" json:"mapped_printers"`
}

// Filevault2User represents a user associated with FileVault 2
type Filevault2User struct {
	User string `xml:"user" json:"user"`
}

// StorageDevice represents a storage device of the computer
type StorageDevice struct {
	Disk            string      `xml:"disk" json:"disk"`
	Model           string      `xml:"model" json:"model"`
	Revision        string      `xml:"revision" json:"revision"`
	SerialNumber    string      `xml:"serial_number" json:"serial_number"`
	Size            int         `xml:"size" json:"size"`
	DriveCapacityMb int         `xml:"drive_capacity_mb" json:"drive_capacity_mb"`
	ConnectionType  string      `xml:"connection_type" json:"connection_type"`
	SmartStatus     string      `xml:"smart_status" json:"smart_status"`
	Partitions      []Partition `xml:"partition" json:"partitions"`
}

// Partition represents a partition on a storage device
type Partition struct {
	Name                 string `xml:"name" json:"name"`
	Size                 int    `xml:"size" json:"size"`
	Type                 string `xml:"type" json:"type"`
	PartitionCapacityMb  int    `xml:"partition_capacity_mb" json:"partition_capacity_mb"`
	PercentageFull       int    `xml:"percentage_full" json:"percentage_full"`
	FilevaultStatus      string `xml:"filevault_status" json:"filevault_status"`
	FilevaultPercent     int    `xml:"filevault_percent" json:"filevault_percent"`
	Filevault2Status     string `xml:"filevault2_status" json:"filevault2_status"`
	Filevault2Percent    int    `xml:"filevault2_percent" json:"filevault2_percent"`
	BootDriveAvailableMb int    `xml:"boot_drive_available_mb" json:"boot_drive_available_mb"`
	LvgUUID              string `xml:"lvgUUID" json:"lvgUUID"`
	LvUUID               string `xml:"lvUUID" json:"lvUUID"`
	PvUUID               string `xml:"pvUUID" json:"pvUUID"`
}

// Printer represents a printer mapped to the computer
type Printer struct {
	Name     string `xml:"name" json:"name"`
	URI      string `xml:"uri" json:"uri"`
	Type     string `xml:"type" json:"type"`
	Location string `xml:"location" json:"location"`
}

// Certificate represents a certificate installed on the computer
type Certificate struct {
	CommonName   string `xml:"common_name" json:"common_name"`
	Identity     bool   `xml:"identity" json:"identity"`
	ExpiresUtc   string `xml:"expires_utc" json:"expires_utc"`
	ExpiresEpoch int64  `xml:"expires_epoch" json:"expires_epoch"`
	Name         string `xml:"name" json:"name"`
}

// Security contains security-related information of the computer
type Security struct {
	ActivationLock      bool   `xml:"activation_lock" json:"activation_lock"`
	RecoveryLockEnabled bool   `xml:"recovery_lock_enabled" json:"recovery_lock_enabled"`
	SecureBootLevel     string `xml:"secure_boot_level" json:"secure_boot_level"`
	ExternalBootLevel   string `xml:"external_boot_level" json:"external_boot_level"`
	FirewallEnabled     bool   `xml:"firewall_enabled" json:"firewall_enabled"`
}

// Software holds information about the software installed on the computer
type Software struct {
	UnixExecutables          []string         `xml:"unix_executables>string" json:"unix_executables"`
	LicensedSoftware         []string         `xml:"licensed_software>string" json:"licensed_software"`
	InstalledByCasper        []string         `xml:"installed_by_casper>package" json:"installed_by_casper"`
	InstalledByInstallerSwu  []string         `xml:"installed_by_installer_swu>package" json:"installed_by_installer_swu"`
	CachedByCasper           []string         `xml:"cached_by_casper>package" json:"cached_by_casper"`
	AvailableSoftwareUpdates []string         `xml:"available_software_updates>name" json:"available_software_updates"`
	AvailableUpdates         []SoftwareUpdate `xml:"available_updates>update" json:"available_updates"`
	RunningServices          []string         `xml:"running_services>name" json:"running_services"`
	Applications             []Application    `xml:"applications>application" json:"applications"`
	Fonts                    []Font           `xml:"fonts>font" json:"fonts"`
	Plugins                  []Plugin         `xml:"plugins>plugin" json:"plugins"`
}

// SoftwareUpdate represents an available software update
type SoftwareUpdate struct {
	Name        string `xml:"name" json:"name"`
	PackageName string `xml:"package_name" json:"package_name"`
	Version     string `xml:"version" json:"version"`
}

// Application represents an application installed on the computer
type Application struct {
	Name    string `xml:"name" json:"name"`
	Path    string `xml:"path" json:"path"`
	Version string `xml:"version" json:"version"`
}

// Font represents a font installed on the computer
type Font struct {
	Name    string `xml:"name" json:"name"`
	Path    string `xml:"path" json:"path"`
	Version string `xml:"version" json:"version"`
}

// Plugin represents a plugin installed on the computer
type Plugin struct {
	Name    string `xml:"name" json:"name"`
	Path    string `xml:"path" json:"path"`
	Version string `xml:"version" json:"version"`
}

// ExtensionAttribute represents a custom extension attribute of the computer
type ExtensionAttribute struct {
	ID    int    `xml:"id" json:"id"`
	Name  string `xml:"name" json:"name"`
	Type  string `xml:"type" json:"type"`
	Value string `xml:"value" json:"value"`
}

// GroupsAccounts contains information about group memberships and local accounts
type GroupsAccounts struct {
	ComputerGroupMemberships []GroupMembership `xml:"computer_group_memberships>group" json:"computer_group_memberships"`
	LocalAccounts            []LocalAccount    `xml:"local_accounts>user" json:"local_accounts"`
}

// GroupMembership represents a membership of a computer in a group
type GroupMembership struct {
	Group string `xml:"group" json:"group"`
}

// LocalAccount represents a local account on the computer
type LocalAccount struct {
	Name             string `xml:"name" json:"name"`
	RealName         string `xml:"realname" json:"realname"`
	UID              string `xml:"uid" json:"uid"`
	Home             string `xml:"home" json:"home"`
	HomeSize         string `xml:"home_size" json:"home_size"`
	HomeSizeMb       int    `xml:"home_size_mb" json:"home_size_mb"`
	Administrator    bool   `xml:"administrator" json:"administrator"`
	FilevaultEnabled bool   `xml:"filevault_enabled" json:"filevault_enabled"`
}

// ConfigurationProfile represents a configuration profile applied to the computer
type ConfigurationProfile struct {
	ID          int    `xml:"id" json:"id"`
	Name        string `xml:"name" json:"name"`
	UUID        string `xml:"uuid" json:"uuid"`
	IsRemovable bool   `xml:"is_removable" json:"is_removable"`
}

// GetComputers retrieves all computers
func (c *Client) GetComputers() (*ResponseComputersList, error) {
	endpoint := uriComputers

	var computersList ResponseComputersList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &computersList)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch computers: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &computersList, nil
}

// GetComputerByID retrieves the computer details by its ID.
func (c *Client) GetComputerByID(id int) (*ResponseComputer, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriComputers, id)

	var computer ResponseComputer
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &computer)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch computer by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &computer, nil
}

// GetComputerByName retrieves the computer by its name
func (c *Client) GetComputerByName(name string) (*ResponseComputer, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriComputers, name)

	var computer ResponseComputer
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &computer)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch computer by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &computer, nil
}

// CreateComputer creates a new computer.
func (c *Client) CreateComputer(computer ResponseComputer) (*ResponseComputer, error) {
	endpoint := uriComputers

	// Check if site is not provided in the General subset and set default values
	if computer.General.Site.ID == 0 && computer.General.Site.Name == "" {
		computer.General.Site = Site{
			ID:   -1,
			Name: "None",
		}
	}

	// The requestBody struct should mirror the Computer struct, including all nested structs
	requestBody := struct {
		XMLName xml.Name `xml:"computer"`
		ResponseComputer
	}{
		ResponseComputer: computer,
	}

	var response ResponseComputer
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to create computer: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateComputerByID updates the details of a computer by its ID.
func (c *Client) UpdateComputerByID(id int, computer ResponseComputer) (*ResponseComputer, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriComputers, id)

	// Check if site is not provided in the General subset and set default values
	if computer.General.Site.ID == 0 && computer.General.Site.Name == "" {
		computer.General.Site = Site{
			ID:   -1,
			Name: "None",
		}
	}

	// The requestBody struct should mirror the Computer struct, including all nested structs
	requestBody := struct {
		XMLName xml.Name `xml:"computer"`
		ResponseComputer
	}{
		ResponseComputer: computer,
	}

	var response ResponseComputer
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to update computer by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateComputerByName updates the details of a computer by its name.
func (c *Client) UpdateComputerByName(name string, computer ResponseComputer) (*ResponseComputer, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriComputers, name)

	// Check if site is not provided in the General subset and set default values
	if computer.General.Site.ID == 0 && computer.General.Site.Name == "" {
		computer.General.Site = Site{
			ID:   -1,
			Name: "None",
		}
	}

	// The requestBody struct should mirror the Computer struct, including all nested structs
	requestBody := struct {
		XMLName xml.Name `xml:"computer"`
		ResponseComputer
	}{
		ResponseComputer: computer,
	}

	var response ResponseComputer
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to update computer by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// DeleteComputerByID deletes an existing Computer by its ID
func (c *Client) DeleteComputerByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriComputers, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete computer by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteComputerByName deletes an existing computer by its name
func (c *Client) DeleteComputerByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriComputers, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete computer by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
