// classicapi_computers.go
// Jamf Pro Classic Api - Computers
// api reference: https://developer.jamf.com/jamf-pro/reference/computers
// Classic API requires the structs to support an XML data structure.

package jamfpro

const uriComputers = "/JSSResource/computers"

type ComputersResponse struct {
	Size    int                    `xml:"size"`
	Results []ComputerListResponse `xml:"computer"`
}

type ComputerListResponse struct {
	ID              int    `json:"id,omitempty" xml:"id,omitempty"`
	UDID            string `json:"udid,omitempty" xml:"udid,omitempty"`
	Name            string `json:"name,omitempty" xml:"name,omitempty"`
	SerialNumber    string `json:"serial_number,omitempty" xml:"serial_number,omitempty"`
	Managed         bool   `json:"managed,omitempty" xml:"managed,omitempty"`
	Model           string `json:"model,omitempty" xml:"model,omitempty"`
	Department      string `json:"department,omitempty" xml:"department,omitempty"`
	Building        string `json:"building,omitempty" xml:"building,omitempty"`
	MACAddress      string `json:"mac_address,omitempty" xml:"mac_address,omitempty"`
	ReportDateUTC   string `json:"report_date_utc,omitempty" xml:"report_date_utc,omitempty"`
	ReportDateEpoch int64  `json:"report_date_epoch,omitempty" xml:"report_date_epoch,omitempty"`
}

type Computer struct {
	General               ComputerDataSubsetGeneral               `json:"general,omitempty" xml:"general,omitempty"`
	Location              ComputerDataSubsetLocation              `json:"location,omitempty" xml:"location,omitempty"`
	Purchasing            ComputerDataSubsetPurchasing            `json:"purchasing,omitempty" xml:"purchasing,omitempty"`
	Peripherals           ComputerDataSubsetPeripherals           `json:"peripherals,omitempty" xml:"peripherals,omitempty"`
	Hardware              ComputerDataSubsetHardware              `json:"hardware,omitempty" xml:"hardware,omitempty"`
	Certificates          ComputerDataSubsetCertificates          `json:"certificates,omitempty" xml:"certificates,omitempty"`
	Security              ComputerDataSubsetSecurity              `json:"security,omitempty" xml:"security,omitempty"`
	Software              ComputerDataSubsetSoftware              `json:"software,omitempty" xml:"software,omitempty"`
	ExtensionAttributes   ComputerDataSubsetExtensionAttributes   `json:"extension_attributes,omitempty" xml:"extension_attributes,omitempty"`
	GroupAccounts         ComputerDataSubsetGroupAccounts         `json:"groups_accounts,omitempty" xml:"groups_accounts,omitempty"`
	IPhones               ComputerDataSubsetIPhones               `json:"iphones,omitempty" xml:"iphones,omitempty"`
	ConfigurationProfiles ComputerDataSubsetConfigurationProfiles `json:"configuration_profiles,omitempty" xml:"configuration_profiles,omitempty"`
}

type ComputerDataSubsetGeneral struct {
	ID                    int    `xml:"id,omitempty"`
	Name                  string `xml:"name,omitempty"`
	NetworkAdapterType    string `xml:"network_adapter_type,omitempty"`
	MacAddress            string `xml:"mac_address,omitempty"`
	AltNetworkAdapterType string `xml:"alt_network_adapter_type,omitempty"`
	AltMacAddress         string `xml:"alt_mac_address,omitempty"`
	IPAddress             string `xml:"ip_address,omitempty"`
	LastReportedIP        string `xml:"last_reported_ip,omitempty"`
	SerialNumber          string `xml:"serial_number,omitempty"`
	UDID                  string `xml:"udid,omitempty"`
	JamfVersion           string `xml:"jamf_version,omitempty"`
	Platform              string `xml:"platform,omitempty"`
	Barcode1              string `xml:"barcode_1,omitempty"`
	Barcode2              string `xml:"barcode_2,omitempty"`
	AssetTag              string `xml:"asset_tag,omitempty"`
	RemoteManagement      struct {
		Managed                  bool   `xml:"managed,omitempty"`
		ManagementUsername       string `xml:"management_username,omitempty"`
		ManagementPasswordSha256 string `xml:"management_password_sha256,omitempty"`
	} `xml:"remote_management,omitempty"`
	Supervised      string `xml:"supervised,omitempty"`
	MdmCapable      string `xml:"mdm_capable,omitempty"`
	MdmCapableUsers []struct {
		MdmCapableUser string `xml:"mdm_capable_user,omitempty"`
	} `xml:"mdm_capable_users,omitempty"`
	ManagementStatus struct {
		EnrolledViaDep  bool `xml:"enrolled_via_dep,omitempty"`
		UserApprovedMdm bool `xml:"user_approved_mdm,omitempty"`
	} `xml:"management_status,omitempty"`
	ReportDate                string `xml:"report_date,omitempty"`
	ReportDateEpoch           string `xml:"report_date_epoch,omitempty"`
	ReportDateUtc             string `xml:"report_date_utc,omitempty"`
	LastContactTime           string `xml:"last_contact_time,omitempty"`
	LastContactTimeEpoch      string `xml:"last_contact_time_epoch,omitempty"`
	LastContactTimeUtc        string `xml:"last_contact_time_utc,omitempty"`
	InitialEntryDate          string `xml:"initial_entry_date,omitempty"`
	InitialEntryDateEpoch     string `xml:"initial_entry_date_epoch,omitempty"`
	InitialEntryDateUtc       string `xml:"initial_entry_date_utc,omitempty"`
	LastCloudBackupDateEpoch  string `xml:"last_cloud_backup_date_epoch,omitempty"`
	LastCloudBackupDateUtc    string `xml:"last_cloud_backup_date_utc,omitempty"`
	LastEnrolledDateEpoch     string `xml:"last_enrolled_date_epoch,omitempty"`
	LastEnrolledDateUtc       string `xml:"last_enrolled_date_utc,omitempty"`
	MdmProfileExpirationEpoch string `xml:"mdm_profile_expiration_epoch,omitempty"`
	MdmProfileExpirationUtc   string `xml:"mdm_profile_expiration_utc,omitempty"`
	DistributionPoint         string `xml:"distribution_point,omitempty"`
	Sus                       string `xml:"sus,omitempty"`
	Site                      struct {
		ID   int    `xml:"id,omitempty"`
		Name string `xml:"name,omitempty"`
	} `xml:"site,omitempty"`
	ItunesStoreAccountIsActive string `xml:"itunes_store_account_is_active"`
}

type ComputerDataSubsetLocation struct {
	Username     string `xml:"username,omitempty"`
	RealName     string `xml:"real_name,omitempty"`
	EmailAddress string `xml:"email_address,omitempty"`
	Position     string `xml:"position,omitempty"`
	Phone        string `xml:"phone,omitempty"`
	PhoneNumber  string `xml:"phone_number,omitempty"`
	Department   string `xml:"department,omitempty"`
	Building     string `xml:"building,omitempty"`
	Room         string `xml:"room,omitempty"`
}

type ComputerDataSubsetPurchasing struct {
	IsPurchased          bool   `xml:"is_purchased,omitempty"`
	IsLeased             bool   `xml:"is_leased,omitempty"`
	PoNumber             string `xml:"po_number,omitempty"`
	Vendor               string `xml:"vendor,omitempty"`
	ApplecareID          string `xml:"applecare_id,omitempty"`
	PurchasePrice        string `xml:"purchase_price,omitempty"`
	PurchasingAccount    string `xml:"purchasing_account,omitempty"`
	PoDate               string `xml:"po_date,omitempty"`
	PoDateEpoch          int64  `xml:"po_date_epoch,omitempty"`
	PoDateUtc            string `xml:"po_date_utc,omitempty"`
	WarrantyExpires      string `xml:"warranty_expires,omitempty"`
	WarrantyExpiresEpoch int64  `xml:"warranty_expires_epoch,omitempty"`
	WarrantyExpiresUtc   string `xml:"warranty_expires_utc,omitempty"`
	LeaseExpires         string `xml:"lease_expires,omitempty"`
	LeaseExpiresEpoch    int64  `xml:"lease_expires_epoch,omitempty"`
	LeaseExpiresUtc      string `xml:"lease_expires_utc,omitempty"`
	LifeExpectancy       int    `xml:"life_expectancy,omitempty"`
	PurchasingContact    string `xml:"purchasing_contact,omitempty"`
	OsApplecareID        string `xml:"os_applecare_id,omitempty"`
	OsMaintenanceExpires string `xml:"os_maintenance_expires,omitempty"`
	Attachments          string `xml:"attachments,omitempty"`
}

// Don't have example data for this to construct resulting struct. Please cut a PR to populate if needed
type ComputerDataSubsetPeripherals struct {
}

type ComputerDataSubsetHardware struct {
	Make                        string `xml:"make,omitempty"`
	Model                       string `xml:"model,omitempty"`
	ModelIdentifier             string `xml:"model_identifier,omitempty"`
	OsName                      string `xml:"os_name,omitempty"`
	OsVersion                   string `xml:"os_version,omitempty"`
	OsBuild                     string `xml:"os_build,omitempty"`
	SoftwareUpdateDeviceID      string `xml:"software_update_device_id,omitempty"`
	ActiveDirectoryStatus       string `xml:"active_directory_status,omitempty"`
	ServicePack                 string `xml:"service_pack,omitempty"`
	ProcessorType               string `xml:"processor_type,omitempty"`
	IsAppleSilicon              bool   `xml:"is_apple_silicon,omitempty"`
	ProcessorArchitecture       string `xml:"processor_architecture,omitempty"`
	ProcessorSpeed              int    `xml:"processor_speed,omitempty"`
	ProcessorSpeedMhz           int    `xml:"processor_speed_mhz,omitempty"`
	NumberProcessors            int    `xml:"number_processors,omitempty"`
	NumberCores                 int    `xml:"number_cores,omitempty"`
	TotalRAM                    int    `xml:"total_ram,omitempty"`
	TotalRAMMb                  int    `xml:"total_ram_mb,omitempty"`
	BootRom                     string `xml:"boot_rom,omitempty"`
	BusSpeed                    int    `xml:"bus_speed,omitempty"`
	BusSpeedMhz                 int    `xml:"bus_speed_mhz,omitempty"`
	BatteryCapacity             int    `xml:"battery_capacity,omitempty"`
	CacheSize                   int    `xml:"cache_size,omitempty"`
	CacheSizeKb                 int    `xml:"cache_size_kb,omitempty"`
	AvailableRAMSlots           int    `xml:"available_ram_slots,omitempty"`
	OpticalDrive                string `xml:"optical_drive,omitempty"`
	NicSpeed                    string `xml:"nic_speed,omitempty"`
	SmcVersion                  string `xml:"smc_version,omitempty"`
	BleCapable                  bool   `xml:"ble_capable,omitempty"`
	SupportsIosAppInstalls      bool   `xml:"supports_ios_app_installs,omitempty"`
	SipStatus                   string `xml:"sip_status,omitempty"`
	GatekeeperStatus            string `xml:"gatekeeper_status,omitempty"`
	XprotectVersion             string `xml:"xprotect_version,omitempty"`
	InstitutionalRecoveryKey    string `xml:"institutional_recovery_key,omitempty"`
	DiskEncryptionConfiguration string `xml:"disk_encryption_configuration,omitempty"`
	Filevault2Users             []struct {
		User string `xml:"user,omitempty"`
	} `xml:"filevault2_users,omitempty"`
	Storage struct {
		Devices []struct {
			Disk            string `xml:"disk,omitempty"`
			Model           string `xml:"model,omitempty"`
			Revision        string `xml:"revision,omitempty"`
			SerialNumber    string `xml:"serial_number,omitempty"`
			Size            int    `xml:"size,omitempty"`
			DriveCapacityMb int    `xml:"drive_capacity_mb,omitempty"`
			ConnectionType  string `xml:"connection_type,omitempty"`
			SmartStatus     string `xml:"smart_status,omitempty"`
			Partitions      struct {
				Partitions []struct {
					Name                 string `xml:"name,omitempty"`
					Size                 int    `xml:"size,omitempty"`
					Type                 string `xml:"type,omitempty"`
					PartitionCapacityMb  int    `xml:"partition_capacity_mb,omitempty"`
					PercentageFull       int    `xml:"percentage_full,omitempty"`
					AvailableMb          int    `xml:"available_mb,omitempty"`
					FilevaultStatus      string `xml:"filevault_status,omitempty"`
					FilevaultPercent     int    `xml:"filevault_percent,omitempty"`
					Filevault2Status     string `xml:"filevault2_status,omitempty"`
					Filevault2Percent    int    `xml:"filevault2_percent,omitempty"`
					BootDriveAvailableMb int    `xml:"boot_drive_available_mb,omitempty"`
					LvgUUID              string `xml:"lvgUUID,omitempty"`
					LvUUID               string `xml:"lvUUID,omitempty"`
					PvUUID               string `xml:"pvUUID,omitempty"`
				} `xml:"partition,omitempty"`
			} `xml:"partitions,omitempty"`
		} `xml:"device,omitempty"`
	} `xml:"storage,omitempty"`
	MappedPrinters string `xml:"mapped_printers,omitempty"`
}

type ComputerDataSubsetCertificates struct {
	Certificates []struct {
		CommonName   string `xml:"common_name,omitempty"`
		Identity     bool   `xml:"identity,omitempty"`
		ExpiresUtc   string `xml:"expires_utc,omitempty"`
		ExpiresEpoch int64  `xml:"expires_epoch,omitempty"`
		Name         string `xml:"name,omitempty"`
	} `xml:"certificate,omitempty"`
}

type ComputerDataSubsetSecurity struct {
	ActivationLock      bool   `xml:"activation_lock,omitempty"`
	RecoveryLockEnabled bool   `xml:"recovery_lock_enabled,omitempty"`
	SecureBootLevel     string `xml:"secure_boot_level,omitempty"`
	ExternalBootLevel   string `xml:"external_boot_level,omitempty"`
	FirewallEnabled     bool   `xml:"firewall_enabled,omitempty"`
}

type ComputerDataSubsetSoftware struct {
	// UnixExecutables   string `xml:"unix_executables,omitempty"`  - unknown format, accepting PRs
	// LicensedSoftware  string `xml:"licensed_software,omitempty"` - unknown format, accepting PRs
	InstalledByCasper []struct {
		Package string `xml:"package,omitempty"`
	} `xml:"installed_by_casper,omitempty"`
	InstalledByInstallerSwu []struct {
		Package []string `xml:"package,omitempty"`
	} `xml:"installed_by_installer_swu,omitempty"`
	CachedByCasper []struct {
		Package string `xml:"package,omitempty"`
	} `xml:"cached_by_casper,omitempty"`
	AvailableSoftwareUpdates []struct {
		Name string `xml:"name,omitempty"`
	} `xml:"available_software_updates,omitempty"`
	AvailableUpdates []struct {
		Update struct {
			Text        string `xml:",chardata,omitempty"`
			Name        string `xml:"name,omitempty"`
			PackageName string `xml:"package_name,omitempty"`
			Version     string `xml:"version,omitempty"`
		} `xml:"update,omitempty"`
	} `xml:"available_updates,omitempty"`
	RunningServices struct {
		Names []string `xml:"name,omitempty"`
	} `xml:"running_services,omitempty"`
	Applications struct {
		Size         int `xml:"size,omitempty"`
		Applications []struct {
			Name     string `xml:"name,omitempty"`
			Path     string `xml:"path,omitempty"`
			Version  string `xml:"version,omitempty"`
			BundleID string `xml:"bundle_id,omitempty"`
		} `xml:"application,omitempty"`
	} `xml:"applications,omitempty"`
}

type ComputerDataSubsetExtensionAttributes struct {
	ExtensionAttributes []struct {
		ID         int    `xml:"id,omitempty"`
		Name       string `xml:"name,omitempty"`
		Type       string `xml:"type,omitempty"`
		MultiValue bool   `xml:"multi_value,omitempty"`
		Value      string `xml:"value,omitempty"`
	} `xml:"extension_attribute,omitempty"`
}

type ComputerDataSubsetGroupAccounts struct {
	ComputerGroupMemberships struct {
		Groups []string `xml:"group,omitempty"`
	} `xml:"computer_group_memberships,omitempty"`
	LocalAccounts struct {
		Users []struct {
			Name             string `xml:"name,omitempty"`
			Realname         string `xml:"realname,omitempty"`
			UID              int    `xml:"uid,omitempty"`
			Home             string `xml:"home,omitempty"`
			HomeSize         string `xml:"home_size,omitempty"`
			HomeSizeMb       int    `xml:"home_size_mb,omitempty"`
			Administrator    bool   `xml:"administrator,omitempty"`
			FilevaultEnabled bool   `xml:"filevault_enabled,omitempty"`
		} `xml:"user,omitempty"`
	} `xml:"local_accounts,omitempty"`
	UserInventories struct {
		DisableAutomaticLogin bool `xml:"disable_automatic_login,omitempty"`
		Users                 []struct {
			Username                     string `xml:"username,omitempty"`
			PasswordHistoryDepth         string `xml:"password_history_depth,omitempty"`
			PasswordMinLength            string `xml:"password_min_length,omitempty"`
			PasswordMaxAge               string `xml:"password_max_age,omitempty"`
			PasswordMinComplexCharacters string `xml:"password_min_complex_characters,omitempty"`
			PasswordRequireAlphanumeric  string `xml:"password_require_alphanumeric,omitempty"`
		} `xml:"user,omitempty"`
	} `xml:"user_inventories,omitempty"`
}

// Don't have example data for this to construct resulting struct. Please cut a PR to populate if needed

type ComputerDataSubsetIPhones struct {
}

type ComputerDataSubsetConfigurationProfiles struct {
	Size                 string `xml:"size,omitempty"`
	ConfigurationProfile []struct {
		ID          int    `xml:"id,omitempty"`
		Name        string `xml:"name,omitempty"`
		UUID        string `xml:"uuid,omitempty"`
		IsRemovable bool   `xml:"is_removable,omitempty"`
	} `xml:"configuration_profile,omitempty"`
}
type ComputerScope struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
	UDID string `xml:"udid,omitempty"`
}

type ComputerDataSubsetName string

const (
	ComputerDataSubsetNameGeneral               ComputerDataSubsetName = "General"
	ComputerDataSubsetNameLocation              ComputerDataSubsetName = "Location"
	ComputerDataSubsetNamePurchasing            ComputerDataSubsetName = "Purchasing"
	ComputerDataSubsetNamePeripherals           ComputerDataSubsetName = "Peripherals"
	ComputerDataSubsetNameHardware              ComputerDataSubsetName = "Hardware"
	ComputerDataSubsetNameCertificates          ComputerDataSubsetName = "Certificates"
	ComputerDataSubsetNameSecurity              ComputerDataSubsetName = "Security"
	ComputerDataSubsetNameSoftware              ComputerDataSubsetName = "Software"
	ComputerDataSubsetNameExtensionAttributes   ComputerDataSubsetName = "ExtensionAttributes"
	ComputerDataSubsetNameGroupAccounts         ComputerDataSubsetName = "GroupsAccounts"
	ComputerDataSubsetNameIPhones               ComputerDataSubsetName = "iphones"
	ComputerDataSubsetNameConfigurationProfiles ComputerDataSubsetName = "ConfigurationProfiles"
)

//TODO
