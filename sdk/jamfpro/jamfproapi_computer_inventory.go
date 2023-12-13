// jamfproapi_computer_inventory.go
// Jamf Pro Api - Computer Inventory
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-computers-inventory
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/mitchellh/mapstructure"
)

const uriComputersInventory = "/api/v1/computers-inventory" // Define the constant for the computers inventory endpoint

// ResponseComputerInventoryList represents the top-level JSON response structure.
type ResponseComputerInventoryList struct {
	TotalCount int                         `json:"totalCount"`
	Results    []ResourceComputerInventory `json:"results"`
}

// ResponseComputerInventory represents an individual computer from the inventory.
type ResourceComputerInventory struct {
	ID                    string                                                        `json:"id"`
	UDID                  string                                                        `json:"udid"`
	General               ResourceDataComputerInventoryDataSubsetGeneral                `json:"general"`
	DiskEncryption        ResourceDataComputerInventoryDataSubsetDiskEncryption         `json:"diskEncryption"`
	Purchasing            ResourceDataComputerInventoryDataSubsetPurchasing             `json:"purchasing"`
	Applications          []ResourceDataComputerInventoryDataSubsetApplication          `json:"applications"`
	Storage               ResourceDataComputerInventoryDataSubsetStorage                `json:"storage"`
	UserAndLocation       ResourceDataComputerInventoryDataSubsetUserAndLocation        `json:"userAndLocation"`
	ConfigurationProfiles []ResourceDataComputerInventoryDataSubsetConfigurationProfile `json:"configurationProfiles"`
	Printers              []ResourceDataComputerInventoryDataSubsetPrinter              `json:"printers"`
	Services              []ResourceDataComputerInventoryDataSubsetService              `json:"services"`
	Hardware              ResourceDataComputerInventoryDataSubsetHardware               `json:"hardware"`
	LocalUserAccounts     []ResourceDataComputerInventoryDataSubsetLocalUserAccount     `json:"localUserAccounts"`
	Certificates          []ResourceDataComputerInventoryDataSubsetCertificate          `json:"certificates"`
	Attachments           []ResourceDataComputerInventoryDataSubsetAttachment           `json:"attachments"`
	Plugins               []ResourceDataComputerInventoryDataSubsetPlugin               `json:"plugins"`
	PackageReceipts       ResourceDataComputerInventoryDataSubsetPackageReceipts        `json:"packageReceipts"`
	Fonts                 []ResourceDataComputerInventoryDataSubsetFont                 `json:"fonts"`
	Security              ResourceDataComputerInventoryDataSubsetSecurity               `json:"security"`
	OperatingSystem       ResourceDataComputerInventoryDataSubsetOperatingSystem        `json:"operatingSystem"`
	LicensedSoftware      []ResourceDataComputerInventoryDataSubsetLicensedSoftware     `json:"licensedSoftware"`
	Ibeacons              []ResourceDataComputerInventoryDataSubsetIbeacon              `json:"ibeacons"`
	SoftwareUpdates       []ResourceDataComputerInventoryDataSubsetSoftwareUpdate       `json:"softwareUpdates"`
	ExtensionAttributes   []ResourceDataComputerInventoryDataSubsetExtensionAttribute   `json:"extensionAttributes"`
	ContentCaching        ResourceDataComputerInventoryDataSubsetContentCaching         `json:"contentCaching"`
	GroupMemberships      []ResourceDataComputerInventoryDataSubsetGroupMembership      `json:"groupMemberships"`
}

// General represents the 'general' section of a result.
type ResourceDataComputerInventoryDataSubsetGeneral struct {
	Name                                 string                                                      `json:"name"`
	LastIpAddress                        string                                                      `json:"lastIpAddress"`
	LastReportedIp                       string                                                      `json:"lastReportedIp"`
	JamfBinaryVersion                    string                                                      `json:"jamfBinaryVersion"`
	Platform                             string                                                      `json:"platform"`
	Barcode1                             string                                                      `json:"barcode1"`
	Barcode2                             string                                                      `json:"barcode2"`
	AssetTag                             string                                                      `json:"assetTag"`
	RemoteManagement                     ResourceDataComputerInventoryDataSubsetRemoteManagement     `json:"remoteManagement"`
	Supervised                           bool                                                        `json:"supervised"`
	MdmCapable                           ResourceDataComputerInventoryDataSubsetMdmCapable           `json:"mdmCapable"`
	ReportDate                           string                                                      `json:"reportDate"`
	LastContactTime                      string                                                      `json:"lastContactTime"`
	LastCloudBackupDate                  string                                                      `json:"lastCloudBackupDate"`
	LastEnrolledDate                     string                                                      `json:"lastEnrolledDate"`
	MdmProfileExpiration                 string                                                      `json:"mdmProfileExpiration"`
	InitialEntryDate                     string                                                      `json:"initialEntryDate"`
	DistributionPoint                    string                                                      `json:"distributionPoint"`
	EnrollmentMethod                     ResourceDataComputerInventoryDataSubsetEnrollmentMethod     `json:"enrollmentMethod"`
	Site                                 ResourceDataComputerInventoryDataSubsetSite                 `json:"site"`
	ItunesStoreAccountActive             bool                                                        `json:"itunesStoreAccountActive"`
	EnrolledViaAutomatedDeviceEnrollment bool                                                        `json:"enrolledViaAutomatedDeviceEnrollment"`
	UserApprovedMdm                      bool                                                        `json:"userApprovedMdm"`
	DeclarativeDeviceManagementEnabled   bool                                                        `json:"declarativeDeviceManagementEnabled"`
	ExtensionAttributes                  []ResourceDataComputerInventoryDataSubsetExtensionAttribute `json:"extensionAttributes"`
	ManagementId                         string                                                      `json:"managementId"`
}

// RemoteManagement represents the 'remoteManagement' section of 'general'.
type ResourceDataComputerInventoryDataSubsetRemoteManagement struct {
	Managed            bool   `json:"managed"`
	ManagementUsername string `json:"managementUsername"`
}

// MdmCapable represents the 'mdmCapable' section of 'general'.
type ResourceDataComputerInventoryDataSubsetMdmCapable struct {
	Capable      bool     `json:"capable"`
	CapableUsers []string `json:"capableUsers"`
}

// EnrollmentMethod represents the 'enrollmentMethod' section of 'general'.
type ResourceDataComputerInventoryDataSubsetEnrollmentMethod struct {
	ID         string `json:"id"`
	ObjectName string `json:"objectName"`
	ObjectType string `json:"objectType"`
}

// Site represents the 'site' section of 'general'.
type ResourceDataComputerInventoryDataSubsetSite struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ExtensionAttribute represents a generic extension attribute.
type ResourceDataComputerInventoryDataSubsetExtensionAttribute struct {
	DefinitionId string   `json:"definitionId"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Enabled      bool     `json:"enabled"`
	MultiValue   bool     `json:"multiValue"`
	Values       []string `json:"values"`
	DataType     string   `json:"dataType"`
	Options      []string `json:"options"`
	InputType    string   `json:"inputType"`
}

// ComputerInventoryDataSubsetDiskEncryption represents the 'diskEncryption' section of a result.
type ResourceDataComputerInventoryDataSubsetDiskEncryption struct {
	BootPartitionEncryptionDetails      ResourceDataComputerInventoryDataSubsetBootPartitionEncryptionDetails `json:"bootPartitionEncryptionDetails"`
	IndividualRecoveryKeyValidityStatus string                                                                `json:"individualRecoveryKeyValidityStatus"`
	InstitutionalRecoveryKeyPresent     bool                                                                  `json:"institutionalRecoveryKeyPresent"`
	DiskEncryptionConfigurationName     string                                                                `json:"diskEncryptionConfigurationName"`
	FileVault2EnabledUserNames          []string                                                              `json:"fileVault2EnabledUserNames"`
	FileVault2EligibilityMessage        string                                                                `json:"fileVault2EligibilityMessage"`
}

// BootPartitionEncryptionDetails represents the details of disk encryption.
type ResourceDataComputerInventoryDataSubsetBootPartitionEncryptionDetails struct {
	PartitionName              string `json:"partitionName"`
	PartitionFileVault2State   string `json:"partitionFileVault2State"`
	PartitionFileVault2Percent int    `json:"partitionFileVault2Percent"`
}

// Purchasing represents the 'purchasing' section of a result.
type ResourceDataComputerInventoryDataSubsetPurchasing struct {
	Leased              bool                                                        `json:"leased"`
	Purchased           bool                                                        `json:"purchased"`
	PoNumber            string                                                      `json:"poNumber"`
	PoDate              string                                                      `json:"poDate"`
	Vendor              string                                                      `json:"vendor"`
	WarrantyDate        string                                                      `json:"warrantyDate"`
	AppleCareId         string                                                      `json:"appleCareId"`
	LeaseDate           string                                                      `json:"leaseDate"`
	PurchasePrice       string                                                      `json:"purchasePrice"`
	LifeExpectancy      int                                                         `json:"lifeExpectancy"`
	PurchasingAccount   string                                                      `json:"purchasingAccount"`
	PurchasingContact   string                                                      `json:"purchasingContact"`
	ExtensionAttributes []ResourceDataComputerInventoryDataSubsetExtensionAttribute `json:"extensionAttributes"`
}

// ComputerInventoryDataSubsetApplication represents an individual application in the 'applications' array.
type ResourceDataComputerInventoryDataSubsetApplication struct {
	Name              string `json:"name"`
	Path              string `json:"path"`
	Version           string `json:"version"`
	MacAppStore       bool   `json:"macAppStore"`
	SizeMegabytes     int    `json:"sizeMegabytes"`
	BundleId          string `json:"bundleId"`
	UpdateAvailable   bool   `json:"updateAvailable"`
	ExternalVersionId string `json:"externalVersionId"`
}

// Storage represents the 'storage' section of a result.
type ResourceDataComputerInventoryDataSubsetStorage struct {
	BootDriveAvailableSpaceMegabytes int                                           `json:"bootDriveAvailableSpaceMegabytes"`
	Disks                            []ResourceDataComputerInventoryDataSubsetDisk `json:"disks"`
}

// ComputerInventoryDataSubsetDisk represents a storage disk.
type ResourceDataComputerInventoryDataSubsetDisk struct {
	ID            string                                             `json:"id"`
	Device        string                                             `json:"device"`
	Model         string                                             `json:"model"`
	Revision      string                                             `json:"revision"`
	SerialNumber  string                                             `json:"serialNumber"`
	SizeMegabytes int                                                `json:"sizeMegabytes"`
	SmartStatus   string                                             `json:"smartStatus"`
	Type          string                                             `json:"type"`
	Partitions    []ResourceDataComputerInventoryDataSubsetPartition `json:"partitions"`
}

// Partition represents a partition of a disk.
type ResourceDataComputerInventoryDataSubsetPartition struct {
	Name                      string `json:"name"`
	SizeMegabytes             int    `json:"sizeMegabytes"`
	AvailableMegabytes        int    `json:"availableMegabytes"`
	PartitionType             string `json:"partitionType"`
	PercentUsed               int    `json:"percentUsed"`
	FileVault2State           string `json:"fileVault2State"`
	FileVault2ProgressPercent int    `json:"fileVault2ProgressPercent"`
	LvmManaged                bool   `json:"lvmManaged"`
}

// UserAndLocation represents the 'userAndLocation' section of a result.
type ResourceDataComputerInventoryDataSubsetUserAndLocation struct {
	Username            string                                                      `json:"username"`
	Realname            string                                                      `json:"realname"`
	Email               string                                                      `json:"email"`
	Position            string                                                      `json:"position"`
	Phone               string                                                      `json:"phone"`
	DepartmentId        string                                                      `json:"departmentId"`
	BuildingId          string                                                      `json:"buildingId"`
	Room                string                                                      `json:"room"`
	ExtensionAttributes []ResourceDataComputerInventoryDataSubsetExtensionAttribute `json:"extensionAttributes"`
}

// ConfigurationProfile represents a configuration profile.
type ResourceDataComputerInventoryDataSubsetConfigurationProfile struct {
	ID                string `json:"id"`
	Username          string `json:"username"`
	LastInstalled     string `json:"lastInstalled"`
	Removable         bool   `json:"removable"`
	DisplayName       string `json:"displayName"`
	ProfileIdentifier string `json:"profileIdentifier"`
}

// Printer represents a printer device.
type ResourceDataComputerInventoryDataSubsetPrinter struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	URI      string `json:"uri"`
	Location string `json:"location"`
}

// Service represents a service.
type ResourceDataComputerInventoryDataSubsetService struct {
	Name string `json:"name"`
}

// Hardware represents the hardware details of a computer.
type ResourceDataComputerInventoryDataSubsetHardware struct {
	Make                   string                                                      `json:"make"`
	Model                  string                                                      `json:"model"`
	ModelIdentifier        string                                                      `json:"modelIdentifier"`
	SerialNumber           string                                                      `json:"serialNumber"`
	ProcessorSpeedMhz      int                                                         `json:"processorSpeedMhz"`
	ProcessorCount         int                                                         `json:"processorCount"`
	CoreCount              int                                                         `json:"coreCount"`
	ProcessorType          string                                                      `json:"processorType"`
	ProcessorArchitecture  string                                                      `json:"processorArchitecture"`
	BusSpeedMhz            int                                                         `json:"busSpeedMhz"`
	CacheSizeKilobytes     int                                                         `json:"cacheSizeKilobytes"`
	NetworkAdapterType     string                                                      `json:"networkAdapterType"`
	MacAddress             string                                                      `json:"macAddress"`
	AltNetworkAdapterType  string                                                      `json:"altNetworkAdapterType"`
	AltMacAddress          string                                                      `json:"altMacAddress"`
	TotalRamMegabytes      int                                                         `json:"totalRamMegabytes"`
	OpenRamSlots           int                                                         `json:"openRamSlots"`
	BatteryCapacityPercent int                                                         `json:"batteryCapacityPercent"`
	SmcVersion             string                                                      `json:"smcVersion"`
	NicSpeed               string                                                      `json:"nicSpeed"`
	OpticalDrive           string                                                      `json:"opticalDrive"`
	BootRom                string                                                      `json:"bootRom"`
	BleCapable             bool                                                        `json:"bleCapable"`
	SupportsIosAppInstalls bool                                                        `json:"supportsIosAppInstalls"`
	AppleSilicon           bool                                                        `json:"appleSilicon"`
	ExtensionAttributes    []ResourceDataComputerInventoryDataSubsetExtensionAttribute `json:"extensionAttributes"`
}

// LocalUserAccount represents a local user account on the computer.
type ResourceDataComputerInventoryDataSubsetLocalUserAccount struct {
	UID                            string `json:"uid"`
	UserGuid                       string `json:"userGuid"`
	Username                       string `json:"username"`
	FullName                       string `json:"fullName"`
	Admin                          bool   `json:"admin"`
	HomeDirectory                  string `json:"homeDirectory"`
	HomeDirectorySizeMb            int    `json:"homeDirectorySizeMb"`
	FileVault2Enabled              bool   `json:"fileVault2Enabled"`
	UserAccountType                string `json:"userAccountType"`
	PasswordMinLength              int    `json:"passwordMinLength"`
	PasswordMaxAge                 int    `json:"passwordMaxAge"`
	PasswordMinComplexCharacters   int    `json:"passwordMinComplexCharacters"`
	PasswordHistoryDepth           int    `json:"passwordHistoryDepth"`
	PasswordRequireAlphanumeric    bool   `json:"passwordRequireAlphanumeric"`
	ComputerAzureActiveDirectoryId string `json:"computerAzureActiveDirectoryId"`
	UserAzureActiveDirectoryId     string `json:"userAzureActiveDirectoryId"`
	AzureActiveDirectoryId         string `json:"azureActiveDirectoryId"`
}

// Certificate represents a security certificate.
type ResourceDataComputerInventoryDataSubsetCertificate struct {
	CommonName        string `json:"commonName"`
	Identity          bool   `json:"identity"`
	ExpirationDate    string `json:"expirationDate"`
	Username          string `json:"username"`
	LifecycleStatus   string `json:"lifecycleStatus"`
	CertificateStatus string `json:"certificateStatus"`
	SubjectName       string `json:"subjectName"`
	SerialNumber      string `json:"serialNumber"`
	Sha1Fingerprint   string `json:"sha1Fingerprint"`
	IssuedDate        string `json:"issuedDate"`
}

// Attachment represents an attachment.
type ResourceDataComputerInventoryDataSubsetAttachment struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	FileType  string `json:"fileType"`
	SizeBytes int    `json:"sizeBytes"`
}

// Plugin represents a software plugin.
type ResourceDataComputerInventoryDataSubsetPlugin struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Path    string `json:"path"`
}

// PackageReceipts represents the package receipts.
type ResourceDataComputerInventoryDataSubsetPackageReceipts struct {
	InstalledByJamfPro      []string `json:"installedByJamfPro"`
	InstalledByInstallerSwu []string `json:"installedByInstallerSwu"`
	Cached                  []string `json:"cached"`
}

// Font represents a font installed on the computer.
type ResourceDataComputerInventoryDataSubsetFont struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Path    string `json:"path"`
}

// Security represents the security settings of the computer.
type ResourceDataComputerInventoryDataSubsetSecurity struct {
	SipStatus             string `json:"sipStatus"`
	GatekeeperStatus      string `json:"gatekeeperStatus"`
	XprotectVersion       string `json:"xprotectVersion"`
	AutoLoginDisabled     bool   `json:"autoLoginDisabled"`
	RemoteDesktopEnabled  bool   `json:"remoteDesktopEnabled"`
	ActivationLockEnabled bool   `json:"activationLockEnabled"`
	RecoveryLockEnabled   bool   `json:"recoveryLockEnabled"`
	FirewallEnabled       bool   `json:"firewallEnabled"`
	SecureBootLevel       string `json:"secureBootLevel"`
	ExternalBootLevel     string `json:"externalBootLevel"`
	BootstrapTokenAllowed bool   `json:"bootstrapTokenAllowed"`
}

// OperatingSystem represents the operating system details of the computer.
type ResourceDataComputerInventoryDataSubsetOperatingSystem struct {
	Name                     string                                                      `json:"name"`
	Version                  string                                                      `json:"version"`
	Build                    string                                                      `json:"build"`
	SupplementalBuildVersion string                                                      `json:"supplementalBuildVersion"`
	RapidSecurityResponse    string                                                      `json:"rapidSecurityResponse"`
	ActiveDirectoryStatus    string                                                      `json:"activeDirectoryStatus"`
	FileVault2Status         string                                                      `json:"fileVault2Status"`
	SoftwareUpdateDeviceId   string                                                      `json:"softwareUpdateDeviceId"`
	ExtensionAttributes      []ResourceDataComputerInventoryDataSubsetExtensionAttribute `json:"extensionAttributes"`
}

// LicensedSoftware represents licensed software.
type ResourceDataComputerInventoryDataSubsetLicensedSoftware struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Ibeacon represents an iBeacon.
type ResourceDataComputerInventoryDataSubsetIbeacon struct {
	Name string `json:"name"`
}

// SoftwareUpdate represents a software update.
type ResourceDataComputerInventoryDataSubsetSoftwareUpdate struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	PackageName string `json:"packageName"`
}

// ContentCaching represents the content caching details.
type ResourceDataComputerInventoryDataSubsetContentCaching struct {
	ComputerContentCachingInformationId string                                                                  `json:"computerContentCachingInformationId"`
	Parents                             []ResourceDataComputerInventoryDataSubsetContentCachingParent           `json:"parents"`
	Alerts                              []ResourceDataComputerInventoryDataSubsetContentCachingAlert            `json:"alerts"` // Corrected to slice
	Activated                           bool                                                                    `json:"activated"`
	Active                              bool                                                                    `json:"active"`
	ActualCacheBytesUsed                int                                                                     `json:"actualCacheBytesUsed"`
	CacheDetails                        []ResourceDataComputerInventoryDataSubsetContentCachingCacheDetail      `json:"cacheDetails"`
	CacheBytesFree                      int                                                                     `json:"cacheBytesFree"`
	CacheBytesLimit                     int                                                                     `json:"cacheBytesLimit"`
	CacheStatus                         string                                                                  `json:"cacheStatus"`
	CacheBytesUsed                      int                                                                     `json:"cacheBytesUsed"`
	DataMigrationCompleted              bool                                                                    `json:"dataMigrationCompleted"`
	DataMigrationProgressPercentage     int                                                                     `json:"dataMigrationProgressPercentage"`
	DataMigrationError                  ResourceDataComputerInventoryDataSubsetContentCachingDataMigrationError `json:"dataMigrationError"`
	MaxCachePressureLast1HourPercentage int                                                                     `json:"maxCachePressureLast1HourPercentage"`
	PersonalCacheBytesFree              int                                                                     `json:"personalCacheBytesFree"`
	PersonalCacheBytesLimit             int                                                                     `json:"personalCacheBytesLimit"`
	PersonalCacheBytesUsed              int                                                                     `json:"personalCacheBytesUsed"`
	Port                                int                                                                     `json:"port"`
	PublicAddress                       string                                                                  `json:"publicAddress"`
	RegistrationError                   string                                                                  `json:"registrationError"`
	RegistrationResponseCode            int                                                                     `json:"registrationResponseCode"`
	RegistrationStarted                 string                                                                  `json:"registrationStarted"`
	RegistrationStatus                  string                                                                  `json:"registrationStatus"`
	RestrictedMedia                     bool                                                                    `json:"restrictedMedia"`
	ServerGuid                          string                                                                  `json:"serverGuid"`
	StartupStatus                       string                                                                  `json:"startupStatus"`
	TetheratorStatus                    string                                                                  `json:"tetheratorStatus"`
	TotalBytesAreSince                  string                                                                  `json:"totalBytesAreSince"`
	TotalBytesDropped                   int64                                                                   `json:"totalBytesDropped"`
	TotalBytesImported                  int64                                                                   `json:"totalBytesImported"`
	TotalBytesReturnedToChildren        int64                                                                   `json:"totalBytesReturnedToChildren"`
	TotalBytesReturnedToClients         int64                                                                   `json:"totalBytesReturnedToClients"`
	TotalBytesReturnedToPeers           int64                                                                   `json:"totalBytesReturnedToPeers"`
	TotalBytesStoredFromOrigin          int64                                                                   `json:"totalBytesStoredFromOrigin"`
	TotalBytesStoredFromParents         int64                                                                   `json:"totalBytesStoredFromParents"`
	TotalBytesStoredFromPeers           int64                                                                   `json:"totalBytesStoredFromPeers"`
}

// ContentCachingParent represents a parent in the content caching details.
type ResourceDataComputerInventoryDataSubsetContentCachingParent struct {
	ContentCachingParentId string                                                       `json:"contentCachingParentId"`
	Address                string                                                       `json:"address"`
	Alerts                 ResourceDataComputerInventoryDataSubsetContentCachingAlert   `json:"alerts"` // Changed from slice to struct
	Details                ResourceDataComputerInventoryDataSubsetContentCachingDetails `json:"details"`
	Guid                   string                                                       `json:"guid"`
	Healthy                bool                                                         `json:"healthy"`
	Port                   int                                                          `json:"port"`
	Version                string                                                       `json:"version"`
}

// ContentCachingAlert represents an alert in the content caching details.
type ResourceDataComputerInventoryDataSubsetContentCachingAlert struct {
	ContentCachingParentAlertId string   `json:"contentCachingParentAlertId"`
	Addresses                   []string `json:"addresses"`
	ClassName                   string   `json:"className"`
	PostDate                    string   `json:"postDate"`
}

// ContentCachingDetails represents the details of content caching.
type ResourceDataComputerInventoryDataSubsetContentCachingDetails struct {
	ContentCachingParentDetailsId string                                                              `json:"contentCachingParentDetailsId"`
	AcPower                       bool                                                                `json:"acPower"`
	CacheSizeBytes                int64                                                               `json:"cacheSizeBytes"`
	Capabilities                  ResourceDataComputerInventoryDataSubsetContentCachingCapabilities   `json:"capabilities"`
	Portable                      bool                                                                `json:"portable"`
	LocalNetwork                  []ResourceDataComputerInventoryDataSubsetContentCachingLocalNetwork `json:"localNetwork"`
}

// ContentCachingCapabilities represents the capabilities in content caching details.
type ResourceDataComputerInventoryDataSubsetContentCachingCapabilities struct {
	ContentCachingParentCapabilitiesId string `json:"contentCachingParentCapabilitiesId"`
	Imports                            bool   `json:"imports"`
	Namespaces                         bool   `json:"namespaces"`
	PersonalContent                    bool   `json:"personalContent"`
	QueryParameters                    bool   `json:"queryParameters"`
	SharedContent                      bool   `json:"sharedContent"`
	Prioritization                     bool   `json:"prioritization"`
}

// ContentCachingLocalNetwork represents a local network in content caching details.
type ResourceDataComputerInventoryDataSubsetContentCachingLocalNetwork struct {
	ContentCachingParentLocalNetworkId string `json:"contentCachingParentLocalNetworkId"`
	Speed                              int    `json:"speed"`
	Wired                              bool   `json:"wired"`
}

// ContentCachingCacheDetail represents cache details in content caching.
type ResourceDataComputerInventoryDataSubsetContentCachingCacheDetail struct {
	ComputerContentCachingCacheDetailsId string `json:"computerContentCachingCacheDetailsId"`
	CategoryName                         string `json:"categoryName"`
	DiskSpaceBytesUsed                   int64  `json:"diskSpaceBytesUsed"`
}

// ContentCachingDataMigrationError represents a data migration error in content caching.
type ResourceDataComputerInventoryDataSubsetContentCachingDataMigrationError struct {
	Code     int                                                             `json:"code"`
	Domain   string                                                          `json:"domain"`
	UserInfo []ResourceDataComputerInventoryDataSubsetContentCachingUserInfo `json:"userInfo"`
}

// ContentCachingUserInfo represents user info in content caching data migration error.
type ResourceDataComputerInventoryDataSubsetContentCachingUserInfo struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// GroupMembership represents a group membership.
type ResourceDataComputerInventoryDataSubsetGroupMembership struct {
	GroupId    string `json:"groupId"`
	GroupName  string `json:"groupName"`
	SmartGroup bool   `json:"smartGroup"`
}

// FileVaultInventoryList represents the paginated FileVault inventory response.
type ResourceDataFileVaultInventoryList struct {
	TotalCount int                              `json:"totalCount"`
	Results    []ResourceDataFileVaultInventory `json:"results"`
}

// FileVaultInventory represents the FileVault information for a single computer.
type ResourceDataFileVaultInventory struct {
	ComputerId                          string                                                                `json:"computerId"`
	Name                                string                                                                `json:"name"`
	PersonalRecoveryKey                 string                                                                `json:"personalRecoveryKey"`
	BootPartitionEncryptionDetails      ResourceDataComputerInventoryDataSubsetBootPartitionEncryptionDetails `json:"bootPartitionEncryptionDetails"`
	IndividualRecoveryKeyValidityStatus string                                                                `json:"individualRecoveryKeyValidityStatus"`
	InstitutionalRecoveryKeyPresent     bool                                                                  `json:"institutionalRecoveryKeyPresent"`
	DiskEncryptionConfigurationName     string                                                                `json:"diskEncryptionConfigurationName"`
}

type ResponseRecoveryLockPassword struct {
	RecoveryLockPassword string `json:"recoveryLockPassword"`
}

type ResponseUploadAttachment struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// GetComputersInventory retrieves all computer inventory information with optional sorting and section filters.
func (c *Client) GetComputersInventory(sort_filter string) (*ResponseComputerInventoryList, error) {
	resp, err := c.DoPaginatedGet(
		uriComputersInventory,
		standardPageSize,
		startingPageNumber,
		"",
	)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "computers-inventories", err)
	}

	var out ResponseComputerInventoryList
	out.TotalCount = resp.Size

	for _, value := range resp.Results {
		var newObj ResourceComputerInventory
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "computer-inventory", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil
}

// GetComputerInventoryByID retrieves a specific computer's inventory information by its ID.
func (c *Client) GetComputerInventoryByID(id string) (*ResourceComputerInventory, error) {
	endpoint := fmt.Sprintf("%s/%s", uriComputersInventory, id)

	// Fetch the computer inventory by ID
	var responseInventory ResourceComputerInventory
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &responseInventory)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "computer inventory", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseInventory, nil
}

// GetComputerInventoryByName retrieves a specific computer's inventory information by its name.
func (c *Client) GetComputerInventoryByName(name string) (*ResourceComputerInventory, error) {
	inventories, err := c.GetComputersInventory("")
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "computer inventory", err)
	}

	for _, inventory := range inventories.Results {
		if inventory.General.Name == name {
			return &inventory, nil
		}
	}

	return nil, fmt.Errorf(errMsgFailedGetByName, "computer inventory", name, err)
}

// UpdateComputerInventoryByID updates a specific computer's inventory information by its ID.
func (c *Client) UpdateComputerInventoryByID(id string, inventoryUpdate *ResourceComputerInventory) (*ResourceComputerInventory, error) {
	endpoint := fmt.Sprintf("%s/%s", uriComputersInventory, id)

	var updatedInventory ResourceComputerInventory
	resp, err := c.HTTP.DoRequest("PATCH", endpoint, inventoryUpdate, &updatedInventory)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "computer inventory", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedInventory, nil
}

// DeleteComputerInventoryByID deletes a computer's inventory information by its ID.
func (c *Client) DeleteComputerInventoryByID(id string) error {
	endpoint := fmt.Sprintf("%s/%s", uriComputersInventory, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "computer-iventory", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

////////// PROGRESS TO HERE

// GetComputersFileVaultInventory retrieves all computer inventory filevault information.
func (c *Client) GetComputersFileVaultInventory() (*ResourceDataFileVaultInventoryList, error) {
	var allInventories []ResourceDataFileVaultInventory

	page := 0
	for {
		params := url.Values{
			"page":      []string{strconv.Itoa(page)},
			"page-size": []string{strconv.Itoa(maxPageSize)},
		}

		endpointWithParams := fmt.Sprintf("%s/filevault?%s", uriComputersInventory, params.Encode())

		// Fetch the FileVault inventory for the current page
		var responseInventories ResourceDataFileVaultInventoryList
		resp, err := c.HTTP.DoRequest("GET", endpointWithParams, nil, &responseInventories)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch FileVault inventory: %v", err)
		}

		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}

		// Add the fetched inventories to the total list
		allInventories = append(allInventories, responseInventories.Results...)

		// Check if all inventories have been fetched
		if len(allInventories) >= responseInventories.TotalCount {
			break
		}

		// Increment page number for the next iteration
		page++
	}

	// Return the combined list of all FileVault inventories
	return &ResourceDataFileVaultInventoryList{
		TotalCount: len(allInventories),
		Results:    allInventories,
	}, nil
}

// GetComputerFileVaultInventoryByID returns file vault details by the computer ID.
func (c *Client) GetComputerFileVaultInventoryByID(id string) (*ResourceDataFileVaultInventory, error) {
	// Construct the endpoint URL using the provided ID
	endpoint := fmt.Sprintf("%s/%s/filevault", uriComputersInventory, id)

	// Fetch the FileVault inventory by ID
	var fileVaultInventory ResourceDataFileVaultInventory
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &fileVaultInventory)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch FileVault inventory by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &fileVaultInventory, nil
}

// GetComputerRecoveryLockPasswordByID returns a computer recover lock password by the computer ID.
func (c *Client) GetComputerRecoveryLockPasswordByID(id string) (*ResponseRecoveryLockPassword, error) {
	endpoint := fmt.Sprintf("%s/%s/view-recovery-lock-password", uriComputersInventory, id)

	var recoveryLockPasswordResponse ResponseRecoveryLockPassword
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &recoveryLockPasswordResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Recovery Lock password by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &recoveryLockPasswordResponse, nil
}

// UploadAttachmentAndAssignToComputerByID uploads a file attachment to a computer by computer ID.
func (c *Client) UploadAttachmentAndAssignToComputerByID(id, filePath string) (*ResponseUploadAttachment, error) {
	endpoint := fmt.Sprintf("%s/%s/attachments", uriComputersInventory, id)

	// Construct the files map
	files := map[string]string{
		"file": filePath, // Assuming 'file' is the form field name for the file uploads
	}

	// Initialize the response struct
	var uploadResponse ResponseUploadAttachment

	// Call DoMultipartRequest with the method, endpoint, files, and the response struct
	resp, err := c.HTTP.DoMultipartRequest("POST", endpoint, nil, files, &uploadResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to upload attachment and assign to computer: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	// Return the response struct pointer
	return &uploadResponse, nil
}

// DeleteAttachmentByIDAndComputerID deletes a computer's inventory attached by computer ID
// and the computer's attachment ID. Multiple attachments can be assigned to a single computer resource.
func (c *Client) DeleteAttachmentByIDAndComputerID(computerID, attachmentID string) error {
	// Construct the endpoint URL using the provided computerID and attachmentID
	endpoint := fmt.Sprintf("%s/%s/attachments/%s", uriComputersInventory, computerID, attachmentID)

	// Make a DELETE request to the endpoint
	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete attachment: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	// Check if the DELETE operation was successful
	// Typical success codes for DELETE are 200 (OK), 202 (Accepted), or 204 (No Content)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("failed to delete attachment, status code: %d", resp.StatusCode)
	}

	return nil
}
