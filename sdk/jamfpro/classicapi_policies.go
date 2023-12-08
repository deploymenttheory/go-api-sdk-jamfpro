// classicapi_policies.go
// Jamf Pro Classic Api - Policies
// api reference: https://developer.jamf.com/jamf-pro/reference/policies
// Jamf Pro Classic Api requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriPolicies = "/JSSResource/policies"

// Policies List Structs
type ResponsePoliciesList struct {
	Size   int          `xml:"size"`
	Policy []PolicyItem `xml:"policy"`
}

type PolicyItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// ResponsePolicy represents the response structure for a single policy
type ResponsePolicy struct {
	General              PolicyGeneral              `xml:"general"`
	Scope                PolicyScope                `xml:"scope,omitempty"`
	SelfService          PolicySelfService          `xml:"self_service"`
	PackageConfiguration PolicyPackageConfiguration `xml:"package_configuration,omitempty"`
	Scripts              PolicyScripts              `xml:"scripts,omitempty"`
	Printers             PolicyPrinters             `xml:"printers"`
	DockItems            PolicyDockItems            `xml:"dock_items"`
	AccountMaintenance   PolicyAccountMaintenance   `xml:"account_maintenance"`
	Maintenance          PolicyMaintenance          `xml:"maintenance"`
	FilesProcesses       PolicyFilesProcesses       `xml:"files_processes"`
	UserInteraction      PolicyUserInteraction      `xml:"user_interaction"`
	DiskEncryption       PolicyDiskEncryption       `xml:"disk_encryption"`
	Reboot               PolicyReboot               `xml:"reboot"`
}

// PolicyGeneral represents the general information of a policy
type PolicyGeneral struct {
	ID                         int                       `xml:"id"`
	Name                       string                    `xml:"name"`
	Enabled                    bool                      `xml:"enabled"`
	Trigger                    string                    `xml:"trigger,omitempty"`
	TriggerCheckin             bool                      `xml:"trigger_checkin"`
	TriggerEnrollmentComplete  bool                      `xml:"trigger_enrollment_complete"`
	TriggerLogin               bool                      `xml:"trigger_login"`
	TriggerLogout              bool                      `xml:"trigger_logout"`
	TriggerNetworkStateChanged bool                      `xml:"trigger_network_state_changed"`
	TriggerStartup             bool                      `xml:"trigger_startup"`
	TriggerOther               string                    `xml:"trigger_other,omitempty"`
	Frequency                  string                    `xml:"frequency,omitempty"`
	RetryEvent                 string                    `xml:"retry_event,omitempty"`
	RetryAttempts              int                       `xml:"retry_attempts,omitempty"`
	NotifyOnEachFailedRetry    bool                      `xml:"notify_on_each_failed_retry"`
	LocationUserOnly           bool                      `xml:"location_user_only"`
	TargetDrive                string                    `xml:"target_drive,omitempty"`
	Offline                    bool                      `xml:"offline"`
	Category                   PolicyCategory            `xml:"category,omitempty"`
	DateTimeLimitations        PolicyDateTimeLimitations `xml:"date_time_limitations,omitempty"`
	NetworkLimitations         PolicyNetworkLimitations  `xml:"network_limitations,omitempty"`
	OverrideDefaultSettings    PolicyOverrideSettings    `xml:"override_default_settings,omitempty"`
	NetworkRequirements        string                    `xml:"network_requirements,omitempty"`
	Site                       PolicySite                `xml:"site"`
}

type PolicyCategory struct {
	ID        string `xml:"id,omitempty"`
	Name      string `xml:"name,omitempty"`
	DisplayIn bool   `xml:"display_in,omitempty"`
	FeatureIn bool   `xml:"feature_in,omitempty"`
}

type PolicyDateTimeLimitations struct {
	ActivationDate      string            `xml:"activation_date,omitempty"`
	ActivationDateEpoch int64             `xml:"activation_date_epoch,omitempty"`
	ActivationDateUTC   string            `xml:"activation_date_utc,omitempty"`
	ExpirationDate      string            `xml:"expiration_date,omitempty"`
	ExpirationDateEpoch int64             `xml:"expiration_date_epoch,omitempty"`
	ExpirationDateUTC   string            `xml:"expiration_date_utc,omitempty"`
	NoExecuteOn         PolicyNoExecuteOn `xml:"no_execute_on,omitempty"`
	NoExecuteStart      string            `xml:"no_execute_start,omitempty"`
	NoExecuteEnd        string            `xml:"no_execute_end,omitempty"`
}

type PolicyNoExecuteOn struct {
	Day string `xml:"day,omitempty"`
}

type PolicyNetworkLimitations struct {
	MinimumNetworkConnection string `xml:"minimum_network_connection,omitempty"`
	AnyIPAddress             bool   `xml:"any_ip_address"`
	NetworkSegments          string `xml:"network_segments"`
}

type PolicyOverrideSettings struct {
	TargetDrive       string `xml:"target_drive,omitempty"`
	DistributionPoint string `xml:"distribution_point,omitempty"`
	ForceAfpSmb       bool   `xml:"force_afp_smb,omitempty"`
	SUS               string `xml:"sus,omitempty"`
	NetbootServer     string `xml:"netboot_server,omitempty"`
}

type PolicySite struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

// PolicyScope represents the scope of the policy
type PolicyScope struct {
	AllComputers   bool                            `xml:"all_computers"`
	AllJSSUsers    bool                            `xml:"all_jss_users"`
	Computers      []PolicyDataSubsetComputer      `xml:"computers>computer,omitempty"`
	ComputerGroups []PolicyDataSubsetComputerGroup `xml:"computer_groups>computer_group,omitempty"`
	JSSUsers       []PolicyDataSubsetJSSUser       `xml:"jss_users>jss_user,omitempty"`
	JSSUserGroups  []PolicyDataSubsetJSSUserGroup  `xml:"jss_user_groups>jss_user_group,omitempty"`
	Buildings      []PolicyDataSubsetBuilding      `xml:"buildings>building,omitempty"`
	Departments    []PolicyDataSubsetDepartment    `xml:"departments>department,omitempty"`
	LimitToUsers   PolicyLimitToUsers              `xml:"limit_to_users,omitempty"`
	Limitations    PolicyLimitations               `xml:"limitations,omitempty"`
	Exclusions     PolicyExclusions                `xml:"exclusions,omitempty"`
}

type PolicyDataSubsetComputer struct {
	ID   int    `xml:"id"`
	Name string `xml:"name,omitempty"`
	UDID string `xml:"udid,omitempty"`
}

type PolicyDataSubsetComputerGroup struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type PolicyDataSubsetJSSUser struct {
	ID   int    `xml:"id"`
	Name string `xml:"name,omitempty"`
}

type PolicyDataSubsetJSSUserGroup struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}
type PolicyDataSubsetBuilding struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type PolicyDataSubsetDepartment struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type PolicyLimitToUsers struct {
	UserGroups []string `xml:"user_groups>user_group,omitempty"`
}

type PolicyLimitations struct {
	Users           []PolicyDataSubsetUser           `xml:"users>user,omitempty"`
	UserGroups      []PolicyDataSubsetUserGroup      `xml:"user_groups>user_group,omitempty"`
	NetworkSegments []PolicyDataSubsetNetworkSegment `xml:"network_segments>network_segment,omitempty"`
	IBeacons        []PolicyDataSubsetIBeacon        `xml:"ibeacons>ibeacon,omitempty"`
}

type PolicyDataSubsetUser struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type PolicyDataSubsetUserGroup struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type PolicyDataSubsetNetworkSegment struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
	UID  string `xml:"uid,omitempty"`
}

type PolicyDataSubsetIBeacon struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type PolicyExclusions struct {
	Computers       []PolicyDataSubsetComputer       `xml:"computers>computer,omitempty"`
	ComputerGroups  []PolicyDataSubsetComputerGroup  `xml:"computer_groups>computer_group,omitempty"`
	Users           []PolicyDataSubsetUser           `xml:"users>user,omitempty"`
	UserGroups      []PolicyDataSubsetUserGroup      `xml:"user_groups>user_group,omitempty"`
	Buildings       []PolicyDataSubsetBuilding       `xml:"buildings>building,omitempty"`
	Departments     []PolicyDataSubsetDepartment     `xml:"departments>department,omitempty"`
	NetworkSegments []PolicyDataSubsetNetworkSegment `xml:"network_segments>network_segment,omitempty"`
	JSSUsers        []PolicyDataSubsetJSSUser        `xml:"jss_users>jss_user,omitempty"`
	JSSUserGroups   []PolicyDataSubsetJSSUserGroup   `xml:"jss_user_groups>jss_user_group,omitempty"`
	IBeacons        []PolicyDataSubsetIBeacon        `xml:"ibeacons>ibeacon,omitempty"`
}

// PolicySelfService represents the self service settings of a policy
type PolicySelfService struct {
	UseForSelfService           bool                        `xml:"use_for_self_service"`
	SelfServiceDisplayName      string                      `xml:"self_service_display_name"`
	InstallButtonText           string                      `xml:"install_button_text"`
	ReinstallButtonText         string                      `xml:"re_install_button_text"`
	SelfServiceDescription      string                      `xml:"self_service_description"`
	ForceUsersToViewDescription bool                        `xml:"force_users_to_view_description"`
	SelfServiceIcon             PolicySelfServiceIcon       `xml:"self_service_icon,omitempty"`
	FeatureOnMainPage           bool                        `xml:"feature_on_main_page"`
	SelfServiceCategories       []PolicySelfServiceCategory `xml:"self_service_categories"`
}

type PolicySelfServiceIcon struct {
	ID       int    `xml:"id,omitempty"`
	Filename string `xml:"filename,omitempty"`
	URI      string `xml:"uri,omitempty"`
}

type PolicySelfServiceCategory struct {
	Category PolicyCategory `xml:"category"`
}

// PolicyPackageConfiguration represents the package configuration settings of a policy
type PolicyPackageConfiguration struct {
	Packages          []PolicyPackage `xml:"packages>package"`
	DistributionPoint string          `xml:"distribution_point"`
}

type PolicyPackage struct {
	ID                int    `xml:"id,omitempty"`
	Name              string `xml:"name,omitempty"`
	Action            string `xml:"action,omitempty"`
	FillUserTemplate  bool   `xml:"fut,omitempty"`
	FillExistingUsers bool   `xml:"feu,omitempty"`
	UpdateAutorun     bool   `xml:"update_autorun,omitempty"`
}

// PolicyScripts represents the scripts settings of a policy
type PolicyScripts struct {
	Size   int                `xml:"size"`
	Script []PolicyScriptItem `xml:"script"`
}

type PolicyScriptItem struct {
	ID          string `xml:"id,omitempty"`
	Name        string `xml:"name,omitempty"`
	Priority    string `xml:"priority,omitempty"`
	Parameter4  string `xml:"parameter4,omitempty"`
	Parameter5  string `xml:"parameter5,omitempty"`
	Parameter6  string `xml:"parameter6,omitempty"`
	Parameter7  string `xml:"parameter7,omitempty"`
	Parameter8  string `xml:"parameter8,omitempty"`
	Parameter9  string `xml:"parameter9,omitempty"`
	Parameter10 string `xml:"parameter10,omitempty"`
	Parameter11 string `xml:"parameter11,omitempty"`
}

// PolicyPrinters represents the printers settings of a policy
type PolicyPrinters struct {
	Size                 int                 `xml:"size"`
	LeaveExistingDefault bool                `xml:"leave_existing_default"`
	Printer              []PolicyPrinterItem `xml:"printer"`
}

type PolicyPrinterItem struct {
	ID          int    `xml:"id"`
	Name        string `xml:"name"`
	Action      string `xml:"action"`
	MakeDefault bool   `xml:"make_default"`
}

// PolicyDockItems represents the dock items settings of a policy
type PolicyDockItems struct {
	Size     int              `xml:"size"`
	DockItem []PolicyDockItem `xml:"dock_item"`
}

type PolicyDockItem struct {
	ID     int    `xml:"id"`
	Name   string `xml:"name"`
	Action string `xml:"action"`
}

// PolicyAccountMaintenance represents the account maintenance settings of a policy
type PolicyAccountMaintenance struct {
	Accounts                []PolicyAccount               `xml:"accounts>account"`
	DirectoryBindings       []PolicyDirectoryBinding      `xml:"directory_bindings>binding"`
	ManagementAccount       PolicyManagementAccount       `xml:"management_account"`
	OpenFirmwareEfiPassword PolicyOpenFirmwareEfiPassword `xml:"open_firmware_efi_password"`
}

type PolicyAccount struct {
	Action                 string `xml:"action"`
	Username               string `xml:"username"`
	Realname               string `xml:"realname"`
	Password               string `xml:"password"`
	ArchiveHomeDirectory   bool   `xml:"archive_home_directory"`
	ArchiveHomeDirectoryTo string `xml:"archive_home_directory_to"`
	Home                   string `xml:"home"`
	Hint                   string `xml:"hint"`
	Picture                string `xml:"picture"`
	Admin                  bool   `xml:"admin"`
	FilevaultEnabled       bool   `xml:"filevault_enabled"`
	PasswordSha256         string `xml:"password_sha256"`
}

type PolicyDirectoryBinding struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type PolicyManagementAccount struct {
	Action                string `xml:"action"`
	ManagedPassword       string `xml:"managed_password"`
	ManagedPasswordLength int    `xml:"managed_password_length"`
}

type PolicyOpenFirmwareEfiPassword struct {
	OfMode           string `xml:"of_mode"`
	OfPassword       string `xml:"of_password"`
	OfPasswordSHA256 string `xml:"of_password_sha256"`
}

// PolicyMaintenance represents the maintenance settings of a policy
type PolicyMaintenance struct {
	Recon                    bool `xml:"recon"`
	ResetName                bool `xml:"reset_name"`
	InstallAllCachedPackages bool `xml:"install_all_cached_packages"`
	Heal                     bool `xml:"heal"`
	Prebindings              bool `xml:"prebindings"`
	Permissions              bool `xml:"permissions"`
	Byhost                   bool `xml:"byhost"`
	SystemCache              bool `xml:"system_cache"`
	UserCache                bool `xml:"user_cache"`
	Verify                   bool `xml:"verify"`
}

// PolicyFilesProcesses represents the files and processes settings of a policy
type PolicyFilesProcesses struct {
	SearchByPath         string `xml:"search_by_path"`
	DeleteFile           bool   `xml:"delete_file"`
	LocateFile           string `xml:"locate_file"`
	UpdateLocateDatabase bool   `xml:"update_locate_database"`
	SpotlightSearch      string `xml:"spotlight_search"`
	SearchForProcess     string `xml:"search_for_process"`
	KillProcess          bool   `xml:"kill_process"`
	RunCommand           string `xml:"run_command"`
}

// PolicyUserInteraction represents the user interaction settings of a policy
type PolicyUserInteraction struct {
	MessageStart          string `xml:"message_start"`
	AllowUserToDefer      bool   `xml:"allow_user_to_defer"`
	AllowDeferralUntilUtc string `xml:"allow_deferral_until_utc"`
	AllowDeferralMinutes  int    `xml:"allow_deferral_minutes"`
	MessageFinish         string `xml:"message_finish"`
}

// PolicyDiskEncryption represents the disk encryption settings of a policy
type PolicyDiskEncryption struct {
	Action                                 string `xml:"action"`
	DiskEncryptionConfigurationID          int    `xml:"disk_encryption_configuration_id"`
	AuthRestart                            bool   `xml:"auth_restart"`
	RemediateKeyType                       string `xml:"remediate_key_type,omitempty"`
	RemediateDiskEncryptionConfigurationID int    `xml:"remediate_disk_encryption_configuration_id,omitempty"`
}

// PolicyReboot represents the reboot settings of a policy
type PolicyReboot struct {
	Message                     string `xml:"message"`
	StartupDisk                 string `xml:"startup_disk"`
	SpecifyStartup              string `xml:"specify_startup"`
	NoUserLoggedIn              string `xml:"no_user_logged_in"`
	UserLoggedIn                string `xml:"user_logged_in"`
	MinutesUntilReboot          int    `xml:"minutes_until_reboot"`
	StartRebootTimerImmediately bool   `xml:"start_reboot_timer_immediately"`
	FileVault2Reboot            bool   `xml:"file_vault_2_reboot"`
}

// GetPolicies retrieves a list of all policies.
func (c *Client) GetPolicies() (*ResponsePoliciesList, error) {
	endpoint := uriPolicies

	var policiesList ResponsePoliciesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &policiesList)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all policies: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &policiesList, nil
}

// GetPolicyByID retrieves the details of a policy by its ID.
func (c *Client) GetPolicyByID(id int) (*ResponsePolicy, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriPolicies, id)

	var policyDetails ResponsePolicy
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &policyDetails)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch policy by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &policyDetails, nil
}

// GetPolicyByName retrieves a policy by its name.
func (c *Client) GetPolicyByName(name string) (*ResponsePolicy, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriPolicies, name)

	var policyDetails ResponsePolicy
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &policyDetails)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch policy by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &policyDetails, nil
}

// GetPolicyByCategory retrieves policies by their category.
func (c *Client) GetPolicyByCategory(category string) (*ResponsePoliciesList, error) {
	endpoint := fmt.Sprintf("%s/category/%s", uriPolicies, category)

	var policiesList ResponsePoliciesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &policiesList)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch policies by category: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &policiesList, nil
}

// GetPoliciesByType retrieves policies by the type of entity that created them.
// The createdBy param can be either the value 'casper' which refers to Casper Remote. Or the value 'jss', which refers to policies created in the GUI or via the API.
func (c *Client) GetPoliciesByType(createdBy string) (*ResponsePoliciesList, error) {
	endpoint := fmt.Sprintf("%s/createdBy/%s", uriPolicies, createdBy)

	var policiesList ResponsePoliciesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &policiesList)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch policies by type: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &policiesList, nil
}

// CreatePolicy creates a new policy.
func (c *Client) CreatePolicyByID(policy *ResponsePolicy) (*ResponsePolicy, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriPolicies, policy.General.ID)

	// Wrap the policy with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"policy"`
		*ResponsePolicy
	}{
		ResponsePolicy: policy,
	}

	var responsePolicy ResponsePolicy
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &responsePolicy)
	if err != nil {
		return nil, fmt.Errorf("failed to create policy: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responsePolicy, nil
}

// UpdatePolicyByID updates an existing policy by its ID.
func (c *Client) UpdatePolicyByID(id int, policy *ResponsePolicy) (*ResponsePolicy, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriPolicies, id)

	// Wrap the policy with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"policy"`
		*ResponsePolicy
	}{
		ResponsePolicy: policy,
	}

	var updatedPolicy ResponsePolicy
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedPolicy)
	if err != nil {
		return nil, fmt.Errorf("failed to update policy: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedPolicy, nil
}

// UpdatePolicyByName updates an existing policy by its name.
func (c *Client) UpdatePolicyByName(name string, policy *ResponsePolicy) (*ResponsePolicy, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriPolicies, name)

	// Wrap the policy with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"policy"`
		*ResponsePolicy
	}{
		ResponsePolicy: policy,
	}

	var updatedPolicy ResponsePolicy
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedPolicy)
	if err != nil {
		return nil, fmt.Errorf("failed to update policy: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedPolicy, nil
}

// DeletePolicyByID deletes a policy by its ID.
func (c *Client) DeletePolicyByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriPolicies, id)
	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete policy: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeletePolicyByName deletes a policy by its name.
func (c *Client) DeletePolicyByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriPolicies, name)
	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete policy: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
