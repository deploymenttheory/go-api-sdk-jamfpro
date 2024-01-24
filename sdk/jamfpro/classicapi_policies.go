// classicapi_policies.go
// Jamf Pro Classic Api - Policies
// api reference: https://developer.jamf.com/jamf-pro/reference/policies
// Jamf Pro Classic Api requires the structs to support an XML data structure.

/*
Shared Resources in this Endpoint
- SharedResourceSite
- SharedResourceSelfServiceIcon
*/

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriPolicies = "/JSSResource/policies"

// List

// Policies List Structs
type ResponsePoliciesList struct {
	Size   int                `xml:"size,omitempty"`
	Policy []PoliciesListItem `xml:"policy,omitempty"`
}

type PoliciesListItem struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

// ResponsePolicyCreateAndUpdate represents the response structure for creating or updating a policy
type ResponsePolicyCreateAndUpdate struct {
	XMLName xml.Name `xml:"policy,omitempty"`
	ID      int      `xml:"id,omitempty"`
}

// Resource

// ResourcePolicy represents the response structure for a single policy
type ResourcePolicy struct {
	General              PolicySubsetGeneral              `xml:"general,omitempty"`
	Scope                PolicySubsetScope                `xml:"scope,omitempty"`
	SelfService          PolicySubsetSelfService          `xml:"self_service,omitempty"`
	PackageConfiguration PolicySubsetPackageConfiguration `xml:"package_configuration,omitempty"`
	Scripts              PolicySubsetScripts              `xml:"scripts,omitempty"`
	Printers             PolicySubsetPrinters             `xml:"printers,omitempty"`
	DockItems            PolicySubsetDockItems            `xml:"dock_items,omitempty"`
	AccountMaintenance   PolicySubsetAccountMaintenance   `xml:"account_maintenance,omitempty"`
	Maintenance          PolicySubsetMaintenance          `xml:"maintenance,omitempty"`
	FilesProcesses       PolicySubsetFilesProcesses       `xml:"files_processes,omitempty"`
	UserInteraction      PolicySubsetUserInteraction      `xml:"user_interaction,omitempty"`
	DiskEncryption       PolicySubsetDiskEncryption       `xml:"disk_encryption,omitempty"`
	Reboot               PolicySubsetReboot               `xml:"reboot,omitempty"`
}

// Subsets & Containers

// General

type PolicySubsetGeneral struct {
	ID                         int                                        `xml:"id,omitempty"`
	Name                       string                                     `xml:"name,omitempty"`
	Enabled                    bool                                       `xml:"enabled,omitempty"`
	Trigger                    string                                     `xml:"trigger,omitempty"`
	TriggerCheckin             bool                                       `xml:"trigger_checkin,omitempty"`
	TriggerEnrollmentComplete  bool                                       `xml:"trigger_enrollment_complete,omitempty"`
	TriggerLogin               bool                                       `xml:"trigger_login,omitempty"`
	TriggerLogout              bool                                       `xml:"trigger_logout,omitempty"`
	TriggerNetworkStateChanged bool                                       `xml:"trigger_network_state_changed,omitempty"`
	TriggerStartup             bool                                       `xml:"trigger_startup,omitempty"`
	TriggerOther               string                                     `xml:"trigger_other,omitempty"`
	Frequency                  string                                     `xml:"frequency,omitempty"`
	RetryEvent                 string                                     `xml:"retry_event,omitempty"`
	RetryAttempts              int                                        `xml:"retry_attempts,omitempty"`
	NotifyOnEachFailedRetry    bool                                       `xml:"notify_on_each_failed_retry,omitempty"`
	LocationUserOnly           bool                                       `xml:"location_user_only,omitempty"`
	TargetDrive                string                                     `xml:"target_drive,omitempty"`
	Offline                    bool                                       `xml:"offline,omitempty"`
	Category                   PolicyCategory                             `xml:"category,omitempty"`
	DateTimeLimitations        PolicySubsetGeneralDateTimeLimitations     `xml:"date_time_limitations,omitempty"`
	NetworkLimitations         PolicySubsetGeneralNetworkLimitations      `xml:"network_limitations,omitempty"`
	OverrideDefaultSettings    PolicySubsetGeneralOverrideDefaultSettings `xml:"override_default_settings,omitempty"`
	NetworkRequirements        string                                     `xml:"network_requirements,omitempty"`
	Site                       SharedResourceSite                         `xml:"site,omitempty"`
}

type PolicySubsetGeneralDateTimeLimitations struct {
	ActivationDate      string                                              `xml:"activation_date,omitempty"`
	ActivationDateEpoch int                                                 `xml:"activation_date_epoch,omitempty"`
	ActivationDateUTC   string                                              `xml:"activation_date_utc,omitempty"`
	ExpirationDate      string                                              `xml:"expiration_date,omitempty"`
	ExpirationDateEpoch int                                                 `xml:"expiration_date_epoch,omitempty"`
	ExpirationDateUTC   string                                              `xml:"expiration_date_utc,omitempty"`
	NoExecuteOn         []PolicySubsetGeneralDateTimeLimitationsNoExecuteOn `xml:"no_execute_on>day,omitempty"`
	NoExecuteStart      string                                              `xml:"no_execute_start,omitempty"`
	NoExecuteEnd        string                                              `xml:"no_execute_end,omitempty"`
}

type PolicySubsetGeneralDateTimeLimitationsNoExecuteOn struct {
	Day string `xml:",chardata,omitempty"`
}

type PolicySubsetGeneralNetworkLimitations struct {
	MinimumNetworkConnection string `xml:"minimum_network_connection,omitempty"`
	AnyIPAddress             bool   `xml:"any_ip_address,omitempty"`
	NetworkSegments          string `xml:"network_segments,omitempty"`
}

type PolicySubsetGeneralOverrideDefaultSettings struct {
	TargetDrive       string `xml:"target_drive,omitempty"`
	DistributionPoint string `xml:"distribution_point,omitempty"`
	ForceAfpSmb       bool   `xml:"force_afp_smb,omitempty"`
	SUS               string `xml:"sus,omitempty"`
	NetbootServer     string `xml:"netboot_server,omitempty"`
}

// Scope

type PolicySubsetScope struct {
	AllComputers   bool                            `xml:"all_computers,omitempty"`
	AllJSSUsers    bool                            `xml:"all_jss_users,omitempty"`
	Computers      []PolicyDataSubsetComputer      `xml:"computers>computer,omitempty"`
	ComputerGroups []PolicyDataSubsetComputerGroup `xml:"computer_groups>computer_group,omitempty"`
	JSSUsers       []PolicyDataSubsetJSSUser       `xml:"jss_users>jss_user,omitempty"`
	JSSUserGroups  []PolicyDataSubsetJSSUserGroup  `xml:"jss_user_groups>jss_user_group,omitempty"`
	Buildings      []PolicyDataSubsetBuilding      `xml:"buildings>building,omitempty"`
	Departments    []PolicyDataSubsetDepartment    `xml:"departments>department,omitempty"`
	LimitToUsers   PolicyLimitToUsers              `xml:"limit_to_users,omitempty"`
	Limitations    PolicySubsetScopeLimitations    `xml:"limitations,omitempty"`
	Exclusions     PolicySubsetScopeExclusions     `xml:"exclusions,omitempty"`
}

type PolicySubsetScopeLimitations struct {
	Users           []PolicyDataSubsetUser           `xml:"users>user,omitempty"`
	UserGroups      []PolicyDataSubsetUserGroup      `xml:"user_groups>user_group,omitempty"`
	NetworkSegments []PolicyDataSubsetNetworkSegment `xml:"network_segments>network_segment,omitempty"`
	IBeacons        []PolicyDataSubsetIBeacon        `xml:"ibeacons>ibeacon,omitempty"`
}

type PolicySubsetScopeExclusions struct {
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

// Self Service

type PolicySubsetSelfService struct {
	UseForSelfService           bool                                `xml:"use_for_self_service,omitempty"`
	SelfServiceDisplayName      string                              `xml:"self_service_display_name,omitempty"`
	InstallButtonText           string                              `xml:"install_button_text,omitempty"`
	ReinstallButtonText         string                              `xml:"re_install_button_text,omitempty"`
	SelfServiceDescription      string                              `xml:"self_service_description,omitempty"`
	ForceUsersToViewDescription bool                                `xml:"force_users_to_view_description,omitempty"`
	SelfServiceIcon             SharedResourceSelfServiceIcon       `xml:"self_service_icon,omitempty"`
	FeatureOnMainPage           bool                                `xml:"feature_on_main_page,omitempty"`
	SelfServiceCategories       []PolicySubsetSelfServiceCategories `xml:"self_service_categories,omitempty"`
}

type PolicySubsetSelfServiceCategories struct {
	Category PolicyCategory `xml:"category,omitempty"`
}

// Package Configuration

type PolicySubsetPackageConfiguration struct {
	Packages          []PolicySubsetPackageConfigurationPackage `xml:"packages>package,omitempty"`
	DistributionPoint string                                    `xml:"distribution_point,omitempty"`
}

type PolicySubsetPackageConfigurationPackage struct {
	ID                int    `xml:"id,omitempty"`
	Name              string `xml:"name,omitempty"`
	Action            string `xml:"action,omitempty"`
	FillUserTemplate  bool   `xml:"fut,omitempty"`
	FillExistingUsers bool   `xml:"feu,omitempty"`
	UpdateAutorun     bool   `xml:"update_autorun,omitempty"`
}

// Scripts

type PolicySubsetScripts struct {
	Size   int                  `xml:"size,omitempty"`
	Script []PolicySubsetScript `xml:"script,omitempty"`
}

type PolicySubsetScript struct {
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

// Printers

type PolicySubsetPrinters struct {
	Size                 int                   `xml:"size,omitempty"`
	LeaveExistingDefault bool                  `xml:"leave_existing_default,omitempty"`
	Printer              []PolicySubsetPrinter `xml:"printer,omitempty"`
}

type PolicySubsetPrinter struct {
	ID          int    `xml:"id,omitempty"`
	Name        string `xml:"name,omitempty"`
	Action      string `xml:"action,omitempty"`
	MakeDefault bool   `xml:"make_default,omitempty"`
}

// Dock Items

type PolicySubsetDockItems struct {
	Size     int                    `xml:"size,omitempty"`
	DockItem []PolicySubsetDockItem `xml:"dock_item,omitempty"`
}

type PolicySubsetDockItem struct {
	ID     int    `xml:"id,omitempty"`
	Name   string `xml:"name,omitempty"`
	Action string `xml:"action,omitempty"`
}

// Account Maintenance

type PolicySubsetAccountMaintenance struct {
	Accounts                []PolicySubsetAccountMaintenanceAccount               `xml:"accounts>account,omitempty"`
	DirectoryBindings       []PolicySubsetAccountMaintenanceDirectoryBindings     `xml:"directory_bindings>binding,omitempty"`
	ManagementAccount       PolicySubsetAccountMaintenanceManagementAccount       `xml:"management_account,omitempty"`
	OpenFirmwareEfiPassword PolicySubsetAccountMaintenanceOpenFirmwareEfiPassword `xml:"open_firmware_efi_password,omitempty"`
}

type PolicySubsetAccountMaintenanceAccount struct {
	Action                 string `xml:"action,omitempty"`
	Username               string `xml:"username,omitempty"`
	Realname               string `xml:"realname,omitempty"`
	Password               string `xml:"password,omitempty"`
	ArchiveHomeDirectory   bool   `xml:"archive_home_directory,omitempty"`
	ArchiveHomeDirectoryTo string `xml:"archive_home_directory_to,omitempty"`
	Home                   string `xml:"home,omitempty"`
	Hint                   string `xml:"hint,omitempty"`
	Picture                string `xml:"picture,omitempty"`
	Admin                  bool   `xml:"admin,omitempty"`
	FilevaultEnabled       bool   `xml:"filevault_enabled,omitempty"`
	PasswordSha256         string `xml:"password_sha256,omitempty"`
}

type PolicySubsetAccountMaintenanceDirectoryBindings struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type PolicySubsetAccountMaintenanceManagementAccount struct {
	Action                string `xml:"action,omitempty"`
	ManagedPassword       string `xml:"managed_password,omitempty"`
	ManagedPasswordLength int    `xml:"managed_password_length,omitempty"`
}

type PolicySubsetAccountMaintenanceOpenFirmwareEfiPassword struct {
	OfMode           string `xml:"of_mode,omitempty"`
	OfPassword       string `xml:"of_password,omitempty"`
	OfPasswordSHA256 string `xml:"of_password_sha256,omitempty"`
}

// Maintenance

type PolicySubsetMaintenance struct {
	Recon                    bool `xml:"recon,omitempty"`
	ResetName                bool `xml:"reset_name,omitempty"`
	InstallAllCachedPackages bool `xml:"install_all_cached_packages,omitempty"`
	Heal                     bool `xml:"heal,omitempty"`
	Prebindings              bool `xml:"prebindings,omitempty"`
	Permissions              bool `xml:"permissions,omitempty"`
	Byhost                   bool `xml:"byhost,omitempty"`
	SystemCache              bool `xml:"system_cache,omitempty"`
	UserCache                bool `xml:"user_cache,omitempty"`
	Verify                   bool `xml:"verify,omitempty"`
}

// Files Processes

type PolicySubsetFilesProcesses struct {
	SearchByPath         string `xml:"search_by_path,omitempty"`
	DeleteFile           bool   `xml:"delete_file,omitempty"`
	LocateFile           string `xml:"locate_file,omitempty"`
	UpdateLocateDatabase bool   `xml:"update_locate_database,omitempty"`
	SpotlightSearch      string `xml:"spotlight_search,omitempty"`
	SearchForProcess     string `xml:"search_for_process,omitempty"`
	KillProcess          bool   `xml:"kill_process,omitempty"`
	RunCommand           string `xml:"run_command,omitempty"`
}

// User Interaction

type PolicySubsetUserInteraction struct {
	MessageStart          string `xml:"message_start,omitempty"`
	AllowUserToDefer      bool   `xml:"allow_user_to_defer,omitempty"`
	AllowDeferralUntilUtc string `xml:"allow_deferral_until_utc,omitempty"`
	AllowDeferralMinutes  int    `xml:"allow_deferral_minutes,omitempty"`
	MessageFinish         string `xml:"message_finish,omitempty"`
}

// Disk Encryption

type PolicySubsetDiskEncryption struct {
	Action                                 string `xml:"action,omitempty"`
	DiskEncryptionConfigurationID          int    `xml:"disk_encryption_configuration_id,omitempty"`
	AuthRestart                            bool   `xml:"auth_restart,omitempty"`
	RemediateKeyType                       string `xml:"remediate_key_type,omitempty"`
	RemediateDiskEncryptionConfigurationID int    `xml:"remediate_disk_encryption_configuration_id,omitempty"`
}

// Reboot

type PolicySubsetReboot struct {
	Message                     string `xml:"message,omitempty"`
	StartupDisk                 string `xml:"startup_disk,omitempty"`
	SpecifyStartup              string `xml:"specify_startup,omitempty"`
	NoUserLoggedIn              string `xml:"no_user_logged_in,omitempty"`
	UserLoggedIn                string `xml:"user_logged_in,omitempty"`
	MinutesUntilReboot          int    `xml:"minutes_until_reboot,omitempty"`
	StartRebootTimerImmediately bool   `xml:"start_reboot_timer_immediately,omitempty"`
	FileVault2Reboot            bool   `xml:"file_vault_2_reboot,omitempty"`
}

// Shared

type PolicyCategory struct {
	ID        int    `xml:"id,omitempty"`
	Name      string `xml:"name,omitempty"`
	DisplayIn bool   `xml:"display_in,omitempty"`
	FeatureIn bool   `xml:"feature_in,omitempty"`
}

type PolicyDataSubsetComputer struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
	UDID string `xml:"udid,omitempty"`
}

type PolicyDataSubsetComputerGroup struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type PolicyDataSubsetJSSUser struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type PolicyDataSubsetJSSUserGroup struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}
type PolicyDataSubsetBuilding struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type PolicyDataSubsetDepartment struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type PolicyLimitToUsers struct {
	UserGroups []string `xml:"user_groups>user_group,omitempty"`
}

type PolicyDataSubsetUser struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type PolicyDataSubsetUserGroup struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type PolicyDataSubsetNetworkSegment struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
	UID  string `xml:"uid,omitempty"`
}

type PolicyDataSubsetIBeacon struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

// CRUD

// GetPolicies retrieves a list of all policies.
func (c *Client) GetPolicies() (*ResponsePoliciesList, error) {
	endpoint := uriPolicies

	var policiesList ResponsePoliciesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &policiesList)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "policies", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &policiesList, nil
}

// GetPolicyByID retrieves the details of a policy by its ID.
func (c *Client) GetPolicyByID(id int) (*ResourcePolicy, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriPolicies, id)

	var policyDetails ResourcePolicy
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &policyDetails)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "policy", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &policyDetails, nil
}

// GetPolicyByName retrieves a policy by its name.
func (c *Client) GetPolicyByName(name string) (*ResourcePolicy, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriPolicies, name)

	var policyDetails ResourcePolicy
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &policyDetails)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "policy", name, err)
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
		return nil, fmt.Errorf(errMsgFailedGetByCategory, "policies", category, err)
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
		return nil, fmt.Errorf(errMsgFailedGetByType, "policies", createdBy, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &policiesList, nil
}

// CreatePolicy creates a new policy.
func (c *Client) CreatePolicy(policy *ResourcePolicy) (*ResponsePolicyCreateAndUpdate, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriPolicies, policy.General.ID)

	requestBody := struct {
		XMLName xml.Name `xml:"policy,omitempty"`
		*ResourcePolicy
	}{
		ResourcePolicy: policy,
	}

	var responsePolicy ResponsePolicyCreateAndUpdate
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &responsePolicy)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "policy", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responsePolicy, nil
}

// UpdatePolicyByID updates an existing policy by its ID.
func (c *Client) UpdatePolicyByID(id int, policy *ResourcePolicy) (*ResponsePolicyCreateAndUpdate, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriPolicies, id)

	requestBody := struct {
		XMLName xml.Name `xml:"policy,omitempty"`
		*ResourcePolicy
	}{
		ResourcePolicy: policy,
	}

	var response ResponsePolicyCreateAndUpdate
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "policy", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdatePolicyByName updates an existing policy by its name.
func (c *Client) UpdatePolicyByName(name string, policy *ResourcePolicy) (*ResponsePolicyCreateAndUpdate, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriPolicies, name)

	requestBody := struct {
		XMLName xml.Name `xml:"policy,omitempty"`
		*ResourcePolicy
	}{
		ResourcePolicy: policy,
	}

	var response ResponsePolicyCreateAndUpdate
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "policy", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// DeletePolicyByID deletes a policy by its ID.
func (c *Client) DeletePolicyByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriPolicies, id)
	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "policy", id, err)
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
		return fmt.Errorf(errMsgFailedDeleteByName, "policy", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
