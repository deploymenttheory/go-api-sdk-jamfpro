package jamfpro

import (
	"fmt"
)

const uriPolicys = "/JSSResource/policies"

type Policy struct {
	General              PolicyGeneral              `xml:"general"`
	Scope                PolicyScope                `xml:"scope,omitempty"`
	SelfService          PolicySelfService          `xml:"self_service"`
	PackageConfiguration PolicyPackageConfiguration `xml:"package_configuration,omitempty"`
	ScriptsConfiguration PolicyScripts              `xml:"scripts,omitempty"`
	Reboot               PolicyReboot               `xml:"reboot"`
	Maintenance          PolicyMaintenance          `xml:"maintenance"`
	FilesAndProcesses    PolicyFilesAndProcesses    `xml:"files_processes"`
	UserInteraction      PolicyUserInteraction      `xml:"user_interaction"`
}

type PolicyGeneral struct {
	ID                         int                                  `xml:"id"`
	Name                       string                               `xml:"name"`
	Enabled                    bool                                 `xml:"enabled"`
	Trigger                    string                               `xml:"trigger"`
	TriggerCheckin             bool                                 `xml:"trigger_checkin"`
	TriggerEnrollmentComplete  bool                                 `xml:"trigger_enrollment_complete"`
	TriggerLogin               bool                                 `xml:"trigger_login"`
	TriggerLogout              bool                                 `xml:"trigger_logout"`
	TriggerNetworkStateChanged bool                                 `xml:"trigger_network_state_changed"`
	TriggerStartup             bool                                 `xml:"trigger_startup"`
	TriggerOther               string                               `xml:"trigger_other"`
	Frequency                  string                               `xml:"frequency"`
	RetryEvent                 string                               `xml:"retry_event"`
	RetryAttempts              int                                  `xml:"retry_attempts"`
	NotifyOnEachFailedRetry    bool                                 `xml:"notify_on_each_failed_retry"`
	LocationUserOnly           bool                                 `xml:"location_user_only"`
	TargetDrive                string                               `xml:"target_drive"`
	Offline                    bool                                 `xml:"offline"`
	Category                   GeneralCategory                      `xml:"category,omitempty"`
	DateTimeLimitations        PolicyGeneralDateTimeLimitations     `xml:"date_time_limitations"`
	NetworkLimitations         PolicyGeneralNetworkLimitations      `xml:"network_limitations"`
	OverrideDefaultSettings    PolicyGeneralOverrideDefaultSettings `xml:"override_default_settings"`
	NetworkRequirements        string                               `xml:"network_requirements"`
	Site                       Site                                 `xml:"site"`
}

// TODO Get types and test
type PolicyGeneralDateTimeLimitations struct {
	ActivationDate      string `xml:"activation_date"`
	ActivationDateEpoch int    `xml:"activation_date_epoch"`
	ActivationDateUtc   string `xml:"activation_date_utc"`
	ExpirationDate      string `xml:"expiration_date"`
	ExpirationDateEpoch int    `xml:"expiration_date_epoch"`
	ExpirationDateUtc   string `xml:"expiration_date_utc"`
	NoExecuteOn         string `xml:"no_execute_on"`
	NoExecuteStart      string `xml:"no_execute_start"`
	NoExecuteEnd        string `xml:"no_execute_end"`
}

// TODO Get types and test
type PolicyGeneralNetworkLimitations struct {
	MinimumNetworkConnection string `xml:"minimum_network_connection"`
	AnyIpAddress             bool   `xml:"any_ip_address"`
	NetworkSegments          string `xml:"network_segments"`
}

// TODO Get types and test
type PolicyGeneralOverrideDefaultSettings struct {
	TargetDrive       string `xml:"target_drive"`
	DistributionPoint string `xml:"distribution_point"`
	ForceAfpSmb       bool   `xml:"force_afp_smb"`
	Sus               string `xml:"sus"`
	NetbootServer     string `xml:"netboot_server"`
}

type PolicyScope struct {
	AllComputers   bool                        `xml:"all_computers"`
	Computers      []ComputerScope             `xml:"computers>computer,omitempty"`
	ComputerGroups []ComputerGroupListResponse `xml:"computer_groups>computer_group,omitempty"`
	Buildings      []BuildingScope             `xml:"buildings>building,omitempty"`
	Departments    []DepartmentScope           `xml:"departments>department,omitempty"`
	Exclusions     PolicyScopeExclusions       `xml:"exclusions,omitempty"`
}

type PolicySelfService struct {
	UseForSelfService           bool                  `xml:"use_for_self_service"`
	SelfServiceDisplayName      string                `xml:"self_service_display_name"`
	InstallButtonText           string                `xml:"install_button_text"`
	ReinstallButtonText         string                `xml:"reinstall_button_text"`
	SelfServiceDescription      string                `xml:"self_service_description"`
	ForceUsersToViewDescription bool                  `xml:"force_users_to_view_description"`
	SelfServiceIcon             SelfServiceIcon       `xml:"self_service_icon,omitempty"`
	FeatureOnMainPage           bool                  `xml:"feature_on_main_page"`
	SelfServiceCategories       []SelfServiceCategory `xml:"self_service_categories>category,omitempty"`
}

type PolicyScopeExclusions struct {
	Computers      []ComputerScope             `xml:"computers>computer,omitempty"`
	ComputerGroups []ComputerGroupListResponse `xml:"computer_groups>computer_group,omitempty"`
	Buildings      []BuildingScope             `xml:"buildings>building,omitempty"`
	Departments    []DepartmentScope           `xml:"departments>department,omitempty"`
}

type PolicyPackageConfiguration struct {
	Packages []PolicyPackageConfigurationPackage `xml:"packages>package,omitempty"`
}

type PolicyPackageConfigurationPackage struct {
	ID                int    `xml:"id,omitempty"`
	Name              string `xml:"name,omitempty"`
	Action            string `xml:"action,omitempty"`
	FillUserTemplate  bool   `xml:"fut,omitempty"`
	FillExistingUsers bool   `xml:"feu,omitempty"`
	UpdateAutorun     bool   `xml:"update_autorun,omitempty"`
}

type PolicyScripts struct {
	Scripts []PolicyScript `xml:"script,omitempty"`
}

type PolicyScript struct {
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

type PolicyFilesAndProcesses struct {
	SearchByPath         string `xml:"search_by_path"`
	DeleteFile           bool   `xml:"delete_file"`
	LocateFile           string `xml:"locate_file"`
	UpdateLocateDatabase bool   `xml:"update_locate_database"`
	SpotlightSearch      string `xml:"spotlight_search"`
	SearchForProcess     string `xml:"search_for_process"`
	KillProcess          bool   `xml:"kill_process"`
	RunCommand           string `xml:"run_command"`
}

type PolicyUserInteraction struct {
	MessageStart          string `xml:"message_start"`
	AllowUsersToDefer     bool   `xml:"allow_users_to_defer"`
	AllowDeferralUntilUtc string `xml:"allow_deferral_until_utc"`
	AllowDeferralMinutes  int    `xml:"allow_deferral_minutes"`
	MessageFinish         string `xml:"message_finish"`
}

type PolicyListResponse struct {
	Size    int              `xml:"size"`
	Results []PolicyListItem `xml:"policy"`
}

type PolicyListItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

func (c *Client) GetPolicyIdByName(name string) (int, error) {
	var id int
	d, err := c.GetPolicies()
	if err != nil {
		return -1, err
	}

	for _, v := range d.Results {
		if v.Name == name {
			id = v.ID
			break
		}
	}
	return id, err
}

func (c *Client) GetPolicyByName(name string) (*Policy, error) {
	id, err := c.GetPolicyIdByName(name)
	if err != nil {
		return nil, err
	}

	return c.GetPolicy(id)
}

func (c *Client) GetPolicy(id int) (*Policy, error) {
	var out *Policy
	uri := fmt.Sprintf("%s/id/%v", uriPolicys, id)
	err := c.DoRequest("GET", uri, nil, nil, &out)

	return out, err
}

func (c *Client) GetPolicies() (*PolicyListResponse, error) {
	out := &PolicyListResponse{}
	err := c.DoRequest("GET", uriPolicys, nil, nil, out)

	return out, err
}

func (c *Client) CreatePolicy(d *Policy) (int, error) {

	// Setting defaults, per jamf unwritten requirements :/

	// Category Defaults
	if d.General.Category.ID == "" || d.General.Category.Name == "" {
		d.General.Category.ID = "-1"
		d.General.Category.Name = "No category assigned"
	}

	// Reboot Defaults
	if d.Reboot.StartupDisk == "" {
		d.Reboot.StartupDisk = "Current Startup Disk"
	}
	if d.Reboot.NoUserLoggedIn == "" {
		d.Reboot.NoUserLoggedIn = "Do not restart"
	}
	if d.Reboot.UserLoggedIn == "" {
		d.Reboot.UserLoggedIn = "Do not restart"
	}
	if d.Reboot.MinutesUntilReboot == 0 {
		d.Reboot.MinutesUntilReboot = 5
	}

	// Package Defaults
	if len(d.PackageConfiguration.Packages) != 0 {
		for i := range d.PackageConfiguration.Packages {
			if d.PackageConfiguration.Packages[i].Action == "" {
				d.PackageConfiguration.Packages[i].Action = "INSTALL"
			}
		}
	}

	// Script Defaults
	if len(d.ScriptsConfiguration.Scripts) != 0 {
		for i := range d.ScriptsConfiguration.Scripts {
			if d.ScriptsConfiguration.Scripts[i].Priority == "" {
				d.ScriptsConfiguration.Scripts[i].Priority = "After"
			}
		}
	}

	res := &PolicyListItem{}
	reqBody := &struct {
		*Policy
		XMLName struct{} `xml:"policy"`
	}{
		Policy: d,
	}

	err := c.DoRequest("POST", fmt.Sprintf("%s/id/0", uriPolicys), reqBody, nil, res)
	return res.ID, err
}

func (c *Client) UpdatePolicy(d *Policy) (int, error) {

	res := &PolicyListItem{}
	uri := fmt.Sprintf("%s/id/%v", uriPolicys, d.General.ID)
	reqBody := &struct {
		*Policy
		XMLName struct{} `xml:"policy"`
	}{
		Policy: d,
	}

	err := c.DoRequest("PUT", uri, reqBody, nil, res)

	return res.ID, err
}

func (c *Client) DeletePolicy(id int) (int, error) {

	policy := &Policy{}
	uri := fmt.Sprintf("%s/id/%v", uriPolicys, id)
	err := c.DoRequest("DELETE", uri, nil, nil, policy)

	return policy.General.ID, err
}
