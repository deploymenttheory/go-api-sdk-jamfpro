package jamfpro

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const uriComputerInventory = "/api/v1/computers-inventory"
const uriComputerInventoryDetail = "/api/v1/computers-inventory-detail"

type ComputerInventoriesResponse struct {
	TotalCount int                 `json:"totalCount,omitempty"`
	Results    []ComputerInventory `json:"results,omitempty"`
}

type ComputerInventory struct {
	ID                    string                                              `json:"id,omitempty"`
	Udid                  string                                              `json:"udid,omitempty"`
	General               *ComputerInventoryDataSubsetGeneral                 `json:"general,omitempty"`
	DiskEncryption        *ComputerInventoryDataSubsetDiskEncryption          `json:"diskEncryption,omitempty"`
	Purchasing            *ComputerInventoryDataSubsetPurchasing              `json:"purchasing,omitempty"`
	Applications          *[]ComputerInventoryDataSubsetApplications          `json:"applications,omitempty"`
	Storage               *ComputerInventoryDataSubsetStorage                 `json:"storage,omitempty"`
	UserAndLocation       *ComputerInventoryDataSubsetUserAndLocation         `json:"userAndLocation,omitempty"`
	ConfigurationProfiles *[]ComputerInventoryDataSubsetConfigurationProfiles `json:"configurationProfiles,omitempty"`
	Printers              *[]ComputerInventoryDataSubsetPrinters              `json:"printers,omitempty"`
	Services              *[]ComputerInventoryDataSubsetServices              `json:"services,omitempty"`
	Hardware              *ComputerInventoryDataSubsetHardware                `json:"hardware,omitempty"`
	LocalUserAccounts     *[]ComputerInventoryDataSubsetLocalUserAccounts     `json:"localUserAccounts,omitempty"`
	Certificates          *[]ComputerInventoryDataSubsetCertificates          `json:"certificates,omitempty"`
	Attachments           *[]ComputerInventoryDataSubsetAttachments           `json:"attachments,omitempty"`
	Plugins               *[]ComputerInventoryDataSubsetPlugins               `json:"plugins,omitempty"`
	PackageReceipts       *ComputerInventoryDataSubsetPackageReceipts         `json:"packageReceipts,omitempty"`
	Fonts                 *[]ComputerInventoryDataSubsetFonts                 `json:"fonts,omitempty"`
	Security              *ComputerInventoryDataSubsetSecurity                `json:"security,omitempty"`
	OperatingSystem       *ComputerInventoryDataSubsetOperatingSystem         `json:"operatingSystem,omitempty"`
	LicensedSoftware      *[]ComputerInventoryDataSubsetLicensedSoftware      `json:"licensedSoftware,omitempty"`
	Ibeacons              *[]ComputerInventoryDataSubsetIbeacons              `json:"ibeacons,omitempty"`
	SoftwareUpdates       *[]ComputerInventoryDataSubsetSoftwareUpdates       `json:"softwareUpdates,omitempty"`
	ExtensionAttributes   *[]ComputerInventoryDataSubsetExtensionAttributes   `json:"extensionAttributes,omitempty"`
	ContentCaching        *ComputerInventoryDataSubsetContentCaching          `json:"contentCaching,omitempty"`
	GroupMemberships      *[]ComputerInventoryDataSubsetGroupMemberships      `json:"groupMemberships,omitempty"`
}

type ComputerInventoryDataSubsetGeneral struct {
	Name              string `json:"name,omitempty"`
	LastIPAddress     string `json:"lastIpAddress,omitempty"`
	LastReportedIP    string `json:"lastReportedIp,omitempty"`
	JamfBinaryVersion string `json:"jamfBinaryVersion,omitempty"`
	Platform          string `json:"platform,omitempty"`
	Barcode1          string `json:"barcode1,omitempty"`
	Barcode2          string `json:"barcode2,omitempty"`
	AssetTag          string `json:"assetTag,omitempty"`
	RemoteManagement  struct {
		Managed            bool   `json:"managed,omitempty"`
		ManagementUsername string `json:"managementUsername,omitempty"`
	} `json:"remoteManagement,omitempty"`
	Supervised bool `json:"supervised,omitempty"`
	MdmCapable struct {
		Capable      bool     `json:"capable,omitempty"`
		CapableUsers []string `json:"capableUsers,omitempty"`
	} `json:"mdmCapable,omitempty"`
	ReportDate           time.Time `json:"reportDate,omitempty"`
	LastContactTime      time.Time `json:"lastContactTime,omitempty"`
	LastCloudBackupDate  time.Time `json:"lastCloudBackupDate,omitempty"`
	LastEnrolledDate     time.Time `json:"lastEnrolledDate,omitempty"`
	MdmProfileExpiration time.Time `json:"mdmProfileExpiration,omitempty"`
	InitialEntryDate     string    `json:"initialEntryDate,omitempty"`
	DistributionPoint    string    `json:"distributionPoint,omitempty"`
	EnrollmentMethod     struct {
		ID         string `json:"id,omitempty"`
		ObjectName string `json:"objectName,omitempty"`
		ObjectType string `json:"objectType,omitempty"`
	} `json:"enrollmentMethod,omitempty"`
	Site struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"site,omitempty"`
	ItunesStoreAccountActive             bool `json:"itunesStoreAccountActive,omitempty"`
	EnrolledViaAutomatedDeviceEnrollment bool `json:"enrolledViaAutomatedDeviceEnrollment,omitempty"`
	UserApprovedMdm                      bool `json:"userApprovedMdm,omitempty"`
	ExtensionAttributes                  []struct {
		DefinitionID string   `json:"definitionId,omitempty"`
		Name         string   `json:"name,omitempty"`
		Description  string   `json:"description,omitempty"`
		Enabled      bool     `json:"enabled,omitempty"`
		MultiValue   bool     `json:"multiValue,omitempty"`
		Values       []string `json:"values,omitempty"`
		DataType     string   `json:"dataType,omitempty"`
		Options      []string `json:"options,omitempty"`
		InputType    string   `json:"inputType,omitempty"`
	} `json:"extensionAttributes,omitempty"`
}

type ComputerInventoryDataSubsetDiskEncryption struct {
	BootPartitionEncryptionDetails struct {
		PartitionName              string `json:"partitionName,omitempty"`
		PartitionFileVault2State   string `json:"partitionFileVault2State,omitempty"`
		PartitionFileVault2Percent int    `json:"partitionFileVault2Percent,omitempty"`
	} `json:"bootPartitionEncryptionDetails,omitempty"`
	IndividualRecoveryKeyValidityStatus string   `json:"individualRecoveryKeyValidityStatus,omitempty"`
	InstitutionalRecoveryKeyPresent     bool     `json:"institutionalRecoveryKeyPresent,omitempty"`
	DiskEncryptionConfigurationName     string   `json:"diskEncryptionConfigurationName,omitempty"`
	FileVault2EnabledUserNames          []string `json:"fileVault2EnabledUserNames,omitempty"`
	FileVault2EligibilityMessage        string   `json:"fileVault2EligibilityMessage,omitempty"`
}

type ComputerInventoryDataSubsetPurchasing struct {
	Leased              bool   `json:"leased,omitempty"`
	Purchased           bool   `json:"purchased,omitempty"`
	PoNumber            string `json:"poNumber,omitempty"`
	PoDate              string `json:"poDate,omitempty"`
	Vendor              string `json:"vendor,omitempty"`
	WarrantyDate        string `json:"warrantyDate,omitempty"`
	AppleCareID         string `json:"appleCareId,omitempty"`
	LeaseDate           string `json:"leaseDate,omitempty"`
	PurchasePrice       string `json:"purchasePrice,omitempty"`
	LifeExpectancy      int    `json:"lifeExpectancy,omitempty"`
	PurchasingAccount   string `json:"purchasingAccount,omitempty"`
	PurchasingContact   string `json:"purchasingContact,omitempty"`
	ExtensionAttributes []struct {
		DefinitionID string   `json:"definitionId,omitempty"`
		Name         string   `json:"name,omitempty"`
		Description  string   `json:"description,omitempty"`
		Enabled      bool     `json:"enabled,omitempty"`
		MultiValue   bool     `json:"multiValue,omitempty"`
		Values       []string `json:"values,omitempty"`
		DataType     string   `json:"dataType,omitempty"`
		Options      []string `json:"options,omitempty"`
		InputType    string   `json:"inputType,omitempty"`
	} `json:"extensionAttributes,omitempty"`
}

type ComputerInventoryDataSubsetApplications struct {
	Name              string `json:"name,omitempty"`
	Path              string `json:"path,omitempty"`
	Version           string `json:"version,omitempty"`
	MacAppStore       bool   `json:"macAppStore,omitempty"`
	SizeMegabytes     int    `json:"sizeMegabytes,omitempty"`
	BundleID          string `json:"bundleId,omitempty"`
	UpdateAvailable   bool   `json:"updateAvailable,omitempty"`
	ExternalVersionID string `json:"externalVersionId,omitempty"`
}

type ComputerInventoryDataSubsetStorage struct {
	BootDriveAvailableSpaceMegabytes int `json:"bootDriveAvailableSpaceMegabytes,omitempty"`
	Disks                            []struct {
		ID            string `json:"id,omitempty"`
		Device        string `json:"device,omitempty"`
		Model         string `json:"model,omitempty"`
		Revision      string `json:"revision,omitempty"`
		SerialNumber  string `json:"serialNumber,omitempty"`
		SizeMegabytes int    `json:"sizeMegabytes,omitempty"`
		SmartStatus   string `json:"smartStatus,omitempty"`
		Type          string `json:"type,omitempty"`
		Partitions    []struct {
			Name                      string `json:"name,omitempty"`
			SizeMegabytes             int    `json:"sizeMegabytes,omitempty"`
			AvailableMegabytes        int    `json:"availableMegabytes,omitempty"`
			PartitionType             string `json:"partitionType,omitempty"`
			PercentUsed               int    `json:"percentUsed,omitempty"`
			FileVault2State           string `json:"fileVault2State,omitempty"`
			FileVault2ProgressPercent int    `json:"fileVault2ProgressPercent,omitempty"`
			LvmManaged                bool   `json:"lvmManaged,omitempty"`
		} `json:"partitions,omitempty"`
	} `json:"disks,omitempty"`
}

type ComputerInventoryDataSubsetUserAndLocation struct {
	Username            string `json:"username,omitempty"`
	Realname            string `json:"realname,omitempty"`
	Email               string `json:"email,omitempty"`
	Position            string `json:"position,omitempty"`
	Phone               string `json:"phone,omitempty"`
	DepartmentID        string `json:"departmentId,omitempty"`
	BuildingID          string `json:"buildingId,omitempty"`
	Room                string `json:"room,omitempty"`
	ExtensionAttributes []struct {
		DefinitionID string   `json:"definitionId,omitempty"`
		Name         string   `json:"name,omitempty"`
		Description  string   `json:"description,omitempty"`
		Enabled      bool     `json:"enabled,omitempty"`
		MultiValue   bool     `json:"multiValue,omitempty"`
		Values       []string `json:"values,omitempty"`
		DataType     string   `json:"dataType,omitempty"`
		Options      []string `json:"options,omitempty"`
		InputType    string   `json:"inputType,omitempty"`
	} `json:"extensionAttributes,omitempty"`
}

type ComputerInventoryDataSubsetConfigurationProfiles struct {
	ID                string    `json:"id,omitempty"`
	Username          string    `json:"username,omitempty"`
	LastInstalled     time.Time `json:"lastInstalled,omitempty"`
	Removable         bool      `json:"removable,omitempty"`
	DisplayName       string    `json:"displayName,omitempty"`
	ProfileIdentifier string    `json:"profileIdentifier,omitempty"`
}

type ComputerInventoryDataSubsetPrinters struct {
	Name     string `json:"name,omitempty"`
	Type     string `json:"type,omitempty"`
	URI      string `json:"uri,omitempty"`
	Location string `json:"location,omitempty"`
}

type ComputerInventoryDataSubsetServices struct {
	Name string `json:"name,omitempty"`
}

type ComputerInventoryDataSubsetHardware struct {
	Make                   string `json:"make,omitempty"`
	Model                  string `json:"model,omitempty"`
	ModelIdentifier        string `json:"modelIdentifier,omitempty"`
	SerialNumber           string `json:"serialNumber,omitempty"`
	ProcessorSpeedMhz      int    `json:"processorSpeedMhz,omitempty"`
	ProcessorCount         int    `json:"processorCount,omitempty"`
	CoreCount              int    `json:"coreCount,omitempty"`
	ProcessorType          string `json:"processorType,omitempty"`
	ProcessorArchitecture  string `json:"processorArchitecture,omitempty"`
	BusSpeedMhz            int    `json:"busSpeedMhz,omitempty"`
	CacheSizeKilobytes     int    `json:"cacheSizeKilobytes,omitempty"`
	NetworkAdapterType     string `json:"networkAdapterType,omitempty"`
	MacAddress             string `json:"macAddress,omitempty"`
	AltNetworkAdapterType  string `json:"altNetworkAdapterType,omitempty"`
	AltMacAddress          string `json:"altMacAddress,omitempty"`
	TotalRAMMegabytes      int    `json:"totalRamMegabytes,omitempty"`
	OpenRAMSlots           int    `json:"openRamSlots,omitempty"`
	BatteryCapacityPercent int    `json:"batteryCapacityPercent,omitempty"`
	SmcVersion             string `json:"smcVersion,omitempty"`
	NicSpeed               string `json:"nicSpeed,omitempty"`
	OpticalDrive           string `json:"opticalDrive,omitempty"`
	BootRom                string `json:"bootRom,omitempty"`
	BleCapable             bool   `json:"bleCapable,omitempty"`
	SupportsIosAppInstalls bool   `json:"supportsIosAppInstalls,omitempty"`
	AppleSilicon           bool   `json:"appleSilicon,omitempty"`
	ExtensionAttributes    []struct {
		DefinitionID string   `json:"definitionId,omitempty"`
		Name         string   `json:"name,omitempty"`
		Description  string   `json:"description,omitempty"`
		Enabled      bool     `json:"enabled,omitempty"`
		MultiValue   bool     `json:"multiValue,omitempty"`
		Values       []string `json:"values,omitempty"`
		DataType     string   `json:"dataType,omitempty"`
		Options      []string `json:"options,omitempty"`
		InputType    string   `json:"inputType,omitempty"`
	} `json:"extensionAttributes,omitempty"`
}

type ComputerInventoryDataSubsetLocalUserAccounts struct {
	UID                            string `json:"uid,omitempty"`
	Username                       string `json:"username,omitempty"`
	FullName                       string `json:"fullName,omitempty"`
	Admin                          bool   `json:"admin,omitempty"`
	HomeDirectory                  string `json:"homeDirectory,omitempty"`
	HomeDirectorySizeMb            int    `json:"homeDirectorySizeMb,omitempty"`
	FileVault2Enabled              bool   `json:"fileVault2Enabled,omitempty"`
	UserAccountType                string `json:"userAccountType,omitempty"`
	PasswordMinLength              int    `json:"passwordMinLength,omitempty"`
	PasswordMaxAge                 int    `json:"passwordMaxAge,omitempty"`
	PasswordMinComplexCharacters   int    `json:"passwordMinComplexCharacters,omitempty"`
	PasswordHistoryDepth           int    `json:"passwordHistoryDepth,omitempty"`
	PasswordRequireAlphanumeric    bool   `json:"passwordRequireAlphanumeric,omitempty"`
	ComputerAzureActiveDirectoryID string `json:"computerAzureActiveDirectoryId,omitempty"`
	UserAzureActiveDirectoryID     string `json:"userAzureActiveDirectoryId,omitempty"`
	AzureActiveDirectoryID         string `json:"azureActiveDirectoryId,omitempty"`
}

type ComputerInventoryDataSubsetCertificates struct {
	CommonName        string    `json:"commonName,omitempty"`
	Identity          bool      `json:"identity,omitempty"`
	ExpirationDate    time.Time `json:"expirationDate,omitempty"`
	Username          string    `json:"username,omitempty"`
	LifecycleStatus   string    `json:"lifecycleStatus,omitempty"`
	CertificateStatus string    `json:"certificateStatus,omitempty"`
}

type ComputerInventoryDataSubsetAttachments struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	FileType  string `json:"fileType,omitempty"`
	SizeBytes int    `json:"sizeBytes,omitempty"`
}

type ComputerInventoryDataSubsetPlugins struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
	Path    string `json:"path,omitempty"`
}

type ComputerInventoryDataSubsetPackageReceipts struct {
	InstalledByJamfPro      []string `json:"installedByJamfPro,omitempty"`
	InstalledByInstallerSwu []string `json:"installedByInstallerSwu,omitempty"`
	Cached                  []string `json:"cached,omitempty"`
}

type ComputerInventoryDataSubsetFonts struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
	Path    string `json:"path,omitempty"`
}

type ComputerInventoryDataSubsetSecurity struct {
	SipStatus             string `json:"sipStatus,omitempty"`
	GatekeeperStatus      string `json:"gatekeeperStatus,omitempty"`
	XprotectVersion       string `json:"xprotectVersion,omitempty"`
	AutoLoginDisabled     bool   `json:"autoLoginDisabled,omitempty"`
	RemoteDesktopEnabled  bool   `json:"remoteDesktopEnabled,omitempty"`
	ActivationLockEnabled bool   `json:"activationLockEnabled,omitempty"`
	RecoveryLockEnabled   bool   `json:"recoveryLockEnabled,omitempty"`
	FirewallEnabled       bool   `json:"firewallEnabled,omitempty"`
	SecureBootLevel       string `json:"secureBootLevel,omitempty"`
	ExternalBootLevel     string `json:"externalBootLevel,omitempty"`
	BootstrapTokenAllowed bool   `json:"bootstrapTokenAllowed,omitempty"`
}

type ComputerInventoryDataSubsetOperatingSystem struct {
	Name                   string `json:"name,omitempty"`
	Version                string `json:"version,omitempty"`
	Build                  string `json:"build,omitempty"`
	ActiveDirectoryStatus  string `json:"activeDirectoryStatus,omitempty"`
	FileVault2Status       string `json:"fileVault2Status,omitempty"`
	SoftwareUpdateDeviceID string `json:"softwareUpdateDeviceId,omitempty"`
	ExtensionAttributes    []struct {
		DefinitionID string   `json:"definitionId,omitempty"`
		Name         string   `json:"name,omitempty"`
		Description  string   `json:"description,omitempty"`
		Enabled      bool     `json:"enabled,omitempty"`
		MultiValue   bool     `json:"multiValue,omitempty"`
		Values       []string `json:"values,omitempty"`
		DataType     string   `json:"dataType,omitempty"`
		Options      []string `json:"options,omitempty"`
		InputType    string   `json:"inputType,omitempty"`
	} `json:"extensionAttributes,omitempty"`
}

type ComputerInventoryDataSubsetLicensedSoftware struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ComputerInventoryDataSubsetIbeacons struct {
	Name string `json:"name,omitempty"`
}

type ComputerInventoryDataSubsetSoftwareUpdates struct {
	Name        string `json:"name,omitempty"`
	Version     string `json:"version,omitempty"`
	PackageName string `json:"packageName,omitempty"`
}

type ComputerInventoryDataSubsetExtensionAttributes struct {
	DefinitionID string   `json:"definitionId,omitempty"`
	Name         string   `json:"name,omitempty"`
	Description  string   `json:"description,omitempty"`
	Enabled      bool     `json:"enabled,omitempty"`
	MultiValue   bool     `json:"multiValue,omitempty"`
	Values       []string `json:"values,omitempty"`
	DataType     string   `json:"dataType,omitempty"`
	Options      []string `json:"options,omitempty"`
	InputType    string   `json:"inputType,omitempty"`
}

type ComputerInventoryDataSubsetContentCaching struct {
	ComputerContentCachingInformationID string `json:"computerContentCachingInformationId,omitempty"`
	Parents                             []struct {
		ContentCachingParentID string `json:"contentCachingParentId,omitempty"`
		Address                string `json:"address,omitempty"`
		Alerts                 struct {
			ContentCachingParentAlertID string        `json:"contentCachingParentAlertId,omitempty"`
			Addresses                   []interface{} `json:"addresses,omitempty"`
			ClassName                   string        `json:"className,omitempty"`
			PostDate                    time.Time     `json:"postDate,omitempty"`
		} `json:"alerts,omitempty"`
		Details struct {
			ContentCachingParentDetailsID string `json:"contentCachingParentDetailsId,omitempty"`
			AcPower                       bool   `json:"acPower,omitempty"`
			CacheSizeBytes                int    `json:"cacheSizeBytes,omitempty"`
			Capabilities                  struct {
				ContentCachingParentCapabilitiesID string `json:"contentCachingParentCapabilitiesId,omitempty"`
				Imports                            bool   `json:"imports,omitempty"`
				Namespaces                         bool   `json:"namespaces,omitempty"`
				PersonalContent                    bool   `json:"personalContent,omitempty"`
				QueryParameters                    bool   `json:"queryParameters,omitempty"`
				SharedContent                      bool   `json:"sharedContent,omitempty"`
				Prioritization                     bool   `json:"prioritization,omitempty"`
			} `json:"capabilities,omitempty"`
			Portable     bool `json:"portable,omitempty"`
			LocalNetwork []struct {
				ContentCachingParentLocalNetworkID string `json:"contentCachingParentLocalNetworkId,omitempty"`
				Speed                              int    `json:"speed,omitempty"`
				Wired                              bool   `json:"wired,omitempty"`
			} `json:"localNetwork,omitempty"`
		} `json:"details,omitempty"`
		GUID    string `json:"guid,omitempty"`
		Healthy bool   `json:"healthy,omitempty"`
		Port    int    `json:"port,omitempty"`
		Version string `json:"version,omitempty"`
	} `json:"parents,omitempty"`
	Alerts []struct {
		CacheBytesLimit      int       `json:"cacheBytesLimit,omitempty"`
		ClassName            string    `json:"className,omitempty"`
		PathPreventingAccess string    `json:"pathPreventingAccess,omitempty"`
		PostDate             time.Time `json:"postDate,omitempty"`
		ReservedVolumeBytes  int       `json:"reservedVolumeBytes,omitempty"`
		Resource             string    `json:"resource,omitempty"`
	} `json:"alerts,omitempty"`
	Activated            bool `json:"activated,omitempty"`
	Active               bool `json:"active,omitempty"`
	ActualCacheBytesUsed int  `json:"actualCacheBytesUsed,omitempty"`
	CacheDetails         []struct {
		ComputerContentCachingCacheDetailsID string `json:"computerContentCachingCacheDetailsId,omitempty"`
		CategoryName                         string `json:"categoryName,omitempty"`
		DiskSpaceBytesUsed                   int    `json:"diskSpaceBytesUsed,omitempty"`
	} `json:"cacheDetails,omitempty"`
	CacheBytesFree                  int64  `json:"cacheBytesFree,omitempty"`
	CacheBytesLimit                 int    `json:"cacheBytesLimit,omitempty"`
	CacheStatus                     string `json:"cacheStatus,omitempty"`
	CacheBytesUsed                  int    `json:"cacheBytesUsed,omitempty"`
	DataMigrationCompleted          bool   `json:"dataMigrationCompleted,omitempty"`
	DataMigrationProgressPercentage int    `json:"dataMigrationProgressPercentage,omitempty"`
	DataMigrationError              struct {
		Code     int    `json:"code,omitempty"`
		Domain   string `json:"domain,omitempty"`
		UserInfo []struct {
			Key   string `json:"key,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"userInfo,omitempty"`
	} `json:"dataMigrationError,omitempty"`
	MaxCachePressureLast1HourPercentage int       `json:"maxCachePressureLast1HourPercentage,omitempty"`
	PersonalCacheBytesFree              int64     `json:"personalCacheBytesFree,omitempty"`
	PersonalCacheBytesLimit             int       `json:"personalCacheBytesLimit,omitempty"`
	PersonalCacheBytesUsed              int       `json:"personalCacheBytesUsed,omitempty"`
	Port                                int       `json:"port,omitempty"`
	PublicAddress                       string    `json:"publicAddress,omitempty"`
	RegistrationError                   string    `json:"registrationError,omitempty"`
	RegistrationResponseCode            int       `json:"registrationResponseCode,omitempty"`
	RegistrationStarted                 time.Time `json:"registrationStarted,omitempty"`
	RegistrationStatus                  string    `json:"registrationStatus,omitempty"`
	RestrictedMedia                     bool      `json:"restrictedMedia,omitempty"`
	ServerGUID                          string    `json:"serverGuid,omitempty"`
	StartupStatus                       string    `json:"startupStatus,omitempty"`
	TetheratorStatus                    string    `json:"tetheratorStatus,omitempty"`
	TotalBytesAreSince                  time.Time `json:"totalBytesAreSince,omitempty"`
	TotalBytesDropped                   int       `json:"totalBytesDropped,omitempty"`
	TotalBytesImported                  int       `json:"totalBytesImported,omitempty"`
	TotalBytesReturnedToChildren        int       `json:"totalBytesReturnedToChildren,omitempty"`
	TotalBytesReturnedToClients         int       `json:"totalBytesReturnedToClients,omitempty"`
	TotalBytesReturnedToPeers           int       `json:"totalBytesReturnedToPeers,omitempty"`
	TotalBytesStoredFromOrigin          int       `json:"totalBytesStoredFromOrigin,omitempty"`
	TotalBytesStoredFromParents         int       `json:"totalBytesStoredFromParents,omitempty"`
	TotalBytesStoredFromPeers           int       `json:"totalBytesStoredFromPeers,omitempty"`
}

type ComputerInventoryDataSubsetGroupMemberships struct {
	GroupID    string `json:"groupId,omitempty"`
	GroupName  string `json:"groupName,omitempty"`
	SmartGroup bool   `json:"smartGroup,omitempty"`
}

type ComputerInventoryDataSubsetName string

const (
	ComputerInventoryDataSubsetNameGeneral               ComputerInventoryDataSubsetName = "GENERAL"
	ComputerInventoryDataSubsetNameLocation              ComputerInventoryDataSubsetName = "DISK_ENCRYPTION"
	ComputerInventoryDataSubsetNamePurchasing            ComputerInventoryDataSubsetName = "PURCHASING"
	ComputerInventoryDataSubsetNameApplications          ComputerInventoryDataSubsetName = "APPLICATIONS"
	ComputerInventoryDataSubsetNameStorage               ComputerInventoryDataSubsetName = "STORAGE"
	ComputerInventoryDataSubsetNameUserAndLocation       ComputerInventoryDataSubsetName = "USER_AND_LOCATION"
	ComputerInventoryDataSubsetNameConfigurationProfiles ComputerInventoryDataSubsetName = "CONFIGURATION_PROFILES"
	ComputerInventoryDataSubsetNamePrinters              ComputerInventoryDataSubsetName = "PRINTERS"
	ComputerInventoryDataSubsetNameServices              ComputerInventoryDataSubsetName = "SERVICES"
	ComputerInventoryDataSubsetNameHardware              ComputerInventoryDataSubsetName = "HARDWARE"
	ComputerInventoryDataSubsetNameLocalUserAccounts     ComputerInventoryDataSubsetName = "LOCAL_USER_ACCOUNTS"
	ComputerInventoryDataSubsetNameCertificates          ComputerInventoryDataSubsetName = "CERTIFICATES"
	ComputerInventoryDataSubsetNameAttachments           ComputerInventoryDataSubsetName = "ATTACHMENTS"
	ComputerInventoryDataSubsetNamePlugins               ComputerInventoryDataSubsetName = "PLUGINS"
	ComputerInventoryDataSubsetNamePackageReceipts       ComputerInventoryDataSubsetName = "PACKAGE_RECEIPTS"
	ComputerInventoryDataSubsetNameFonts                 ComputerInventoryDataSubsetName = "FONTS"
	ComputerInventoryDataSubsetNameSecurity              ComputerInventoryDataSubsetName = "SECURITY"
	ComputerInventoryDataSubsetNameOperatingSystem       ComputerInventoryDataSubsetName = "OPERATING_SYSTEM"
	ComputerInventoryDataSubsetNameLicensedSoftware      ComputerInventoryDataSubsetName = "LICENSED_SOFTWARE"
	ComputerInventoryDataSubsetNameIBeacons              ComputerInventoryDataSubsetName = "IBEACONS"
	ComputerInventoryDataSubsetNameSoftwareUpdates       ComputerInventoryDataSubsetName = "SOFTWARE_UPDATES"
	ComputerInventoryDataSubsetNameExtensionAttributes   ComputerInventoryDataSubsetName = "EXTENSION_ATTRIBUTES"
	ComputerInventoryDataSubsetNameContentCaching        ComputerInventoryDataSubsetName = "CONTENT_CACHING"
	ComputerInventoryDataSubsetNameGroupMemberships      ComputerInventoryDataSubsetName = "GROUP_MEMBERSHIPS"
)

type ComputerInventoriesQuery struct {
	Sections *[]ComputerInventoryDataSubsetName
	Page     int
	PageSize int
	Sort     *[]string
	Filter   string
}

func (c *Client) GetComputerInventories(query *ComputerInventoriesQuery) (*ComputerInventoriesResponse, error) {
	var out *ComputerInventoriesResponse
	additionalHeaders := url.Values{}
	if query != nil {
		if query.Sections != nil {
			for _, section := range *query.Sections {
				additionalHeaders.Add("section", string(section))
			}
		}
		if query.Page != 0 {
			additionalHeaders.Add("page", strconv.Itoa(query.Page))
		}
		if query.PageSize != 0 {
			additionalHeaders.Add("page-size", strconv.Itoa(query.PageSize))
		}
		if query.Sort != nil {
			additionalHeaders.Add("sort", strings.Join(*query.Sort, ","))
		}
		if query.Filter != "" {
			additionalHeaders.Add("filter", query.Filter)
		}
	}
	err := c.DoRequest("GET", uriComputerInventory, nil, &additionalHeaders, &out)
	return out, err
}

func (c *Client) GetComputerInventoryByID(id string, sections ...ComputerInventoryDataSubsetName) (*ComputerInventory, error) {
	var out *ComputerInventory
	uri := fmt.Sprintf("%s/%v", uriComputerInventoryDetail, id)
	additionalHeaders := url.Values{}
	for _, section := range sections {
		additionalHeaders.Add("section", string(section))
	}
	err := c.DoRequest("GET", uri, nil, &additionalHeaders, &out)
	return out, err
}

func (c *Client) DeleteComputerInventoryByID(id string) error {
	var out *ComputerInventory
	uri := fmt.Sprintf("%s/%v", uriComputerInventoryDetail, id)
	additionalHeaders := url.Values{}
	err := c.DoRequest("DELETE", uri, nil, &additionalHeaders, &out)
	return err
}

func (c *Client) GetComputerInventoryDetailByID(id string) (*ComputerInventory, error) {
	var out *ComputerInventory
	uri := fmt.Sprintf("%s/%v", uriComputerInventoryDetail, id)
	err := c.DoRequest("GET", uri, nil, nil, &out)
	return out, err
}

func (c *Client) UpdateComputerInventoryDetailByID(id string, computer ComputerInventory) (*ComputerInventory, error) {
	var out *ComputerInventory
	uri := fmt.Sprintf("%s/%v", uriComputerInventoryDetail, id)
	err := c.DoRequest("PATCH", uri, computer, nil, &out)
	return out, err
}
