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
	Size   int `xml:"size"`
	Policy []struct {
		ID   int    `xml:"id"`
		Name string `xml:"name"`
	} `xml:"policy"`
}

// ResponsePolicyCreateAndUpdate represents the response structure for creating or updating a policy
type ResponsePolicyCreateAndUpdate struct {
	XMLName xml.Name `xml:"policy"`
	ID      int      `xml:"id"`
}

// ResourcePolicy represents the response structure for a single policy
type ResourcePolicy struct {
	General struct {
		ID                         int            `xml:"id"`
		Name                       string         `xml:"name"`
		Enabled                    bool           `xml:"enabled"`
		Trigger                    string         `xml:"trigger,omitempty"`
		TriggerCheckin             bool           `xml:"trigger_checkin"`
		TriggerEnrollmentComplete  bool           `xml:"trigger_enrollment_complete"`
		TriggerLogin               bool           `xml:"trigger_login"`
		TriggerLogout              bool           `xml:"trigger_logout"`
		TriggerNetworkStateChanged bool           `xml:"trigger_network_state_changed"`
		TriggerStartup             bool           `xml:"trigger_startup"`
		TriggerOther               string         `xml:"trigger_other,omitempty"`
		Frequency                  string         `xml:"frequency,omitempty"`
		RetryEvent                 string         `xml:"retry_event,omitempty"`
		RetryAttempts              int            `xml:"retry_attempts,omitempty"`
		NotifyOnEachFailedRetry    bool           `xml:"notify_on_each_failed_retry"`
		LocationUserOnly           bool           `xml:"location_user_only"`
		TargetDrive                string         `xml:"target_drive,omitempty"`
		Offline                    bool           `xml:"offline"`
		Category                   PolicyCategory `xml:"category,omitempty"`
		DateTimeLimitations        struct {
			ActivationDate      string `xml:"activation_date,omitempty"`
			ActivationDateEpoch int    `xml:"activation_date_epoch,omitempty"`
			ActivationDateUTC   string `xml:"activation_date_utc,omitempty"`
			ExpirationDate      string `xml:"expiration_date,omitempty"`
			ExpirationDateEpoch int    `xml:"expiration_date_epoch,omitempty"`
			ExpirationDateUTC   string `xml:"expiration_date_utc,omitempty"`
			NoExecuteOn         []struct {
				Day string `xml:",chardata"`
			} `xml:"no_execute_on>day,omitempty"`
			NoExecuteStart string `xml:"no_execute_start,omitempty"`
			NoExecuteEnd   string `xml:"no_execute_end,omitempty"`
		} `xml:"date_time_limitations,omitempty"`
		NetworkLimitations struct {
			MinimumNetworkConnection string `xml:"minimum_network_connection,omitempty"`
			AnyIPAddress             bool   `xml:"any_ip_address"`
			NetworkSegments          string `xml:"network_segments"`
		} `xml:"network_limitations,omitempty"`
		OverrideDefaultSettings struct {
			TargetDrive       string `xml:"target_drive,omitempty"`
			DistributionPoint string `xml:"distribution_point,omitempty"`
			ForceAfpSmb       bool   `xml:"force_afp_smb,omitempty"`
			SUS               string `xml:"sus,omitempty"`
			NetbootServer     string `xml:"netboot_server,omitempty"`
		} `xml:"override_default_settings,omitempty"`
		NetworkRequirements string `xml:"network_requirements,omitempty"`
		Site                struct {
			ID   int    `xml:"id,omitempty"`
			Name string `xml:"name,omitempty"`
		} `xml:"site"`
	} `xml:"general"`
	Scope struct {
		AllComputers   bool                            `xml:"all_computers"`
		AllJSSUsers    bool                            `xml:"all_jss_users"`
		Computers      []PolicyDataSubsetComputer      `xml:"computers>computer,omitempty"`
		ComputerGroups []PolicyDataSubsetComputerGroup `xml:"computer_groups>computer_group,omitempty"`
		JSSUsers       []PolicyDataSubsetJSSUser       `xml:"jss_users>jss_user,omitempty"`
		JSSUserGroups  []PolicyDataSubsetJSSUserGroup  `xml:"jss_user_groups>jss_user_group,omitempty"`
		Buildings      []PolicyDataSubsetBuilding      `xml:"buildings>building,omitempty"`
		Departments    []PolicyDataSubsetDepartment    `xml:"departments>department,omitempty"`
		LimitToUsers   PolicyLimitToUsers              `xml:"limit_to_users,omitempty"`
		Limitations    struct {
			Users           []PolicyDataSubsetUser           `xml:"users>user,omitempty"`
			UserGroups      []PolicyDataSubsetUserGroup      `xml:"user_groups>user_group,omitempty"`
			NetworkSegments []PolicyDataSubsetNetworkSegment `xml:"network_segments>network_segment,omitempty"`
			IBeacons        []PolicyDataSubsetIBeacon        `xml:"ibeacons>ibeacon,omitempty"`
		} `xml:"limitations,omitempty"`
		Exclusions struct {
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
		} `xml:"exclusions,omitempty"`
	} `xml:"scope,omitempty"`
	SelfService struct {
		UseForSelfService           bool   `xml:"use_for_self_service"`
		SelfServiceDisplayName      string `xml:"self_service_display_name"`
		InstallButtonText           string `xml:"install_button_text"`
		ReinstallButtonText         string `xml:"re_install_button_text"`
		SelfServiceDescription      string `xml:"self_service_description"`
		ForceUsersToViewDescription bool   `xml:"force_users_to_view_description"`
		SelfServiceIcon             struct {
			ID       int    `xml:"id,omitempty"`
			Filename string `xml:"filename,omitempty"`
			URI      string `xml:"uri,omitempty"`
		} `xml:"self_service_icon,omitempty"`
		FeatureOnMainPage     bool `xml:"feature_on_main_page"`
		SelfServiceCategories []struct {
			Category PolicyCategory `xml:"category"`
		} `xml:"self_service_categories"`
	} `xml:"self_service"`
	PackageConfiguration struct {
		Packages []struct {
			ID                int    `xml:"id,omitempty"`
			Name              string `xml:"name,omitempty"`
			Action            string `xml:"action,omitempty"`
			FillUserTemplate  bool   `xml:"fut,omitempty"`
			FillExistingUsers bool   `xml:"feu,omitempty"`
			UpdateAutorun     bool   `xml:"update_autorun,omitempty"`
		} `xml:"packages>package"`
		DistributionPoint string `xml:"distribution_point"`
	} `xml:"package_configuration,omitempty"`
	Scripts struct {
		Size   int `xml:"size"`
		Script []struct {
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
		} `xml:"script"`
	} `xml:"scripts,omitempty"`
	Printers struct {
		Size                 int  `xml:"size"`
		LeaveExistingDefault bool `xml:"leave_existing_default"`
		Printer              []struct {
			ID          int    `xml:"id"`
			Name        string `xml:"name"`
			Action      string `xml:"action"`
			MakeDefault bool   `xml:"make_default"`
		} `xml:"printer"`
	} `xml:"printers"`
	DockItems struct {
		Size     int `xml:"size"`
		DockItem []struct {
			ID     int    `xml:"id"`
			Name   string `xml:"name"`
			Action string `xml:"action"`
		} `xml:"dock_item"`
	} `xml:"dock_items"`
	AccountMaintenance struct {
		Accounts []struct {
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
		} `xml:"accounts>account"`
		DirectoryBindings []struct {
			ID   int    `xml:"id"`
			Name string `xml:"name"`
		} `xml:"directory_bindings>binding"`
		ManagementAccount struct {
			Action                string `xml:"action"`
			ManagedPassword       string `xml:"managed_password"`
			ManagedPasswordLength int    `xml:"managed_password_length"`
		} `xml:"management_account"`
		OpenFirmwareEfiPassword struct {
			OfMode           string `xml:"of_mode"`
			OfPassword       string `xml:"of_password"`
			OfPasswordSHA256 string `xml:"of_password_sha256"`
		} `xml:"open_firmware_efi_password"`
	} `xml:"account_maintenance"`
	Maintenance struct {
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
	} `xml:"maintenance"`
	FilesProcesses struct {
		SearchByPath         string `xml:"search_by_path"`
		DeleteFile           bool   `xml:"delete_file"`
		LocateFile           string `xml:"locate_file"`
		UpdateLocateDatabase bool   `xml:"update_locate_database"`
		SpotlightSearch      string `xml:"spotlight_search"`
		SearchForProcess     string `xml:"search_for_process"`
		KillProcess          bool   `xml:"kill_process"`
		RunCommand           string `xml:"run_command"`
	} `xml:"files_processes"`
	UserInteraction struct {
		MessageStart          string `xml:"message_start"`
		AllowUserToDefer      bool   `xml:"allow_user_to_defer"`
		AllowDeferralUntilUtc string `xml:"allow_deferral_until_utc"`
		AllowDeferralMinutes  int    `xml:"allow_deferral_minutes"`
		MessageFinish         string `xml:"message_finish"`
	} `xml:"user_interaction"`
	DiskEncryption struct {
		Action                                 string `xml:"action"`
		DiskEncryptionConfigurationID          int    `xml:"disk_encryption_configuration_id"`
		AuthRestart                            bool   `xml:"auth_restart"`
		RemediateKeyType                       string `xml:"remediate_key_type,omitempty"`
		RemediateDiskEncryptionConfigurationID int    `xml:"remediate_disk_encryption_configuration_id,omitempty"`
	} `xml:"disk_encryption"`
	Reboot struct {
		Message                     string `xml:"message"`
		StartupDisk                 string `xml:"startup_disk"`
		SpecifyStartup              string `xml:"specify_startup"`
		NoUserLoggedIn              string `xml:"no_user_logged_in"`
		UserLoggedIn                string `xml:"user_logged_in"`
		MinutesUntilReboot          int    `xml:"minutes_until_reboot"`
		StartRebootTimerImmediately bool   `xml:"start_reboot_timer_immediately"`
		FileVault2Reboot            bool   `xml:"file_vault_2_reboot"`
	} `xml:"reboot"`
}

type PolicyCategory struct {
	ID        int    `xml:"id,omitempty"`
	Name      string `xml:"name,omitempty"`
	DisplayIn bool   `xml:"display_in,omitempty"`
	FeatureIn bool   `xml:"feature_in,omitempty"`
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
func (c *Client) GetPolicyByID(id int) (*ResourcePolicy, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriPolicies, id)

	var policyDetails ResourcePolicy
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
func (c *Client) GetPolicyByName(name string) (*ResourcePolicy, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriPolicies, name)

	var policyDetails ResourcePolicy
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
func (c *Client) CreatePolicyByID(policy *ResourcePolicy) (*ResponsePolicyCreateAndUpdate, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriPolicies, policy.General.ID)

	// Wrap the policy with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"policy"`
		*ResourcePolicy
	}{
		ResourcePolicy: policy,
	}

	var responsePolicy ResponsePolicyCreateAndUpdate
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &responsePolicy)
	if err != nil {
		return nil, fmt.Errorf("failed to create policy: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	// Return the ID of the newly created policy
	return &responsePolicy, nil
}

// UpdatePolicyByID updates an existing policy by its ID.
func (c *Client) UpdatePolicyByID(id int, policy *ResourcePolicy) (*ResponsePolicyCreateAndUpdate, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriPolicies, id)

	// Wrap the policy with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"policy"`
		*ResourcePolicy
	}{
		ResourcePolicy: policy,
	}

	var response ResponsePolicyCreateAndUpdate
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to update policy: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdatePolicyByName updates an existing policy by its name.
func (c *Client) UpdatePolicyByName(name string, policy *ResourcePolicy) (*ResponsePolicyCreateAndUpdate, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriPolicies, name)

	// Wrap the policy with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"policy"`
		*ResourcePolicy
	}{
		ResourcePolicy: policy,
	}

	var response ResponsePolicyCreateAndUpdate
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to update policy: %v", err)
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
