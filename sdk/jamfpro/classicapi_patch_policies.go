// classicapi_patch_policies.go
// Jamf Pro Classic Api  - Patch Policies
// api reference: https://developer.jamf.com/jamf-pro/reference/patchpolicies
// Jamf Pro Classic Api requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

// Constant for the Patch Policies endpoint
const uriPatchPolicies = "/JSSResource/patchpolicies"

// ResourcePatchPolicies represents the root element of the patch policy XML.
type ResourcePatchPolicies struct {
	General struct {
		ID                 int    `xml:"id"`
		Name               string `xml:"name"`
		Enabled            bool   `xml:"enabled"`
		TargetVersion      string `xml:"target_version"`
		ReleaseDate        string `xml:"release_date"`
		IncrementalUpdates bool   `xml:"incremental_updates"`
		Reboot             bool   `xml:"reboot"`
		MinimumOS          string `xml:"minimum_os"`
		KillApps           []struct {
			KillApp struct {
				KillAppName     string `xml:"kill_app_name"`
				KillAppBundleID string `xml:"kill_app_bundle_id"`
			} `xml:"kill_app"`
		} `xml:"kill_apps>kill_app"`
		DistributionMethod string `xml:"distribution_method"`
		AllowDowngrade     bool   `xml:"allow_downgrade"`
		PatchUnknown       bool   `xml:"patch_unknown"`
	} `xml:"general"`
	Scope struct {
		AllComputers   bool                               `xml:"all_computers"`
		Computers      []PatchPoliciesSubsetComputerItem  `xml:"computers>computer"`
		ComputerGroups []PatchPoliciesSubsetComputerGroup `xml:"computer_groups>computer_group"`
		Buildings      []PatchPoliciesSubsetBuilding      `xml:"buildings>building"`
		Departments    []PatchPoliciesSubsetDepartment    `xml:"departments>department"`
		Limitations    struct {
			NetworkSegments []PatchPoliciesSubsetNetworkSegmentItem `xml:"network_segments>network_segment"`
			IBeacons        []PatchPoliciesSubsetIBeaconItem        `xml:"ibeacons>ibeacon"`
		} `xml:"limitations"`
		Exclusions struct {
			Computers       []PatchPoliciesSubsetComputerItem       `xml:"computers>computer"`
			ComputerGroups  []PatchPoliciesSubsetComputerGroup      `xml:"computer_groups>computer_group"`
			Buildings       []PatchPoliciesSubsetBuilding           `xml:"buildings>building"`
			Departments     []PatchPoliciesSubsetDepartment         `xml:"departments>department"`
			NetworkSegments []PatchPoliciesSubsetNetworkSegmentItem `xml:"network_segments>network_segment"`
			IBeacons        []PatchPoliciesSubsetIBeaconItem        `xml:"ibeacons>ibeacon"`
		} `xml:"exclusions"`
	} `xml:"scope"`
	UserInteraction struct {
		InstallButtonText      string `xml:"install_button_text"`
		SelfServiceDescription string `xml:"self_service_description"`
		SelfServiceIcon        struct {
			ID       int    `xml:"id"`
			Filename string `xml:"filename"`
			URI      string `xml:"uri"`
		} `xml:"self_service_icon"`
		Notifications struct {
			NotificationEnabled bool   `xml:"notification_enabled"`
			NotificationType    string `xml:"notification_type"`
			NotificationSubject string `xml:"notification_subject"`
			NotificationMessage string `xml:"notification_message"`
			Reminders           struct {
				NotificationRemindersEnabled  bool `xml:"notification_reminders_enabled"`
				NotificationReminderFrequency int  `xml:"notification_reminder_frequency"`
			} `xml:"reminders"`
		} `xml:"notifications"`
		Deadlines struct {
			DeadlineEnabled bool `xml:"deadline_enabled"`
			DeadlinePeriod  int  `xml:"deadline_period"`
		} `xml:"deadlines"`
		GracePeriod struct {
			GracePeriodDuration       int    `xml:"grace_period_duration"`
			NotificationCenterSubject string `xml:"notification_center_subject"`
			Message                   string `xml:"message"`
		} `xml:"grace_period"`
	} `xml:"user_interaction"`
	SoftwareTitleConfigurationID int `xml:"software_title_configuration_id"`
}

// PatchPoliciesSubsetComputerItem represents a computer in the scope.
type PatchPoliciesSubsetComputerItem struct {
	Computer struct {
		ID   int    `xml:"id"`
		Name string `xml:"name"`
		UDID string `xml:"udid"`
	} `xml:"computer"`
}

// PatchPoliciesSubsetComputerGroup represents a computer group in the scope.
type PatchPoliciesSubsetComputerGroup struct {
	ComputerGroup PatchPoliciesSubsetGroup `xml:"computer_group"`
}

// PatchPoliciesSubsetBuilding represents a building in the scope.
type PatchPoliciesSubsetBuilding struct {
	Building PatchPoliciesSubsetGroup `xml:"building"`
}

// PatchPoliciesSubsetDepartment represents a department in the scope.
type PatchPoliciesSubsetDepartment struct {
	Department PatchPoliciesSubsetGroup `xml:"department"`
}

// PatchPoliciesSubsetGroup is a general struct for group elements.
type PatchPoliciesSubsetGroup struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// PatchPoliciesSubsetNetworkSegmentItem represents a network segment in limitations.
type PatchPoliciesSubsetNetworkSegmentItem struct {
	NetworkSegment PatchPoliciesSubsetGroup `xml:"network_segment"`
}

// PatchPoliciesSubsetIBeaconItem represents an iBeacon in limitations.
type PatchPoliciesSubsetIBeaconItem struct {
	IBeacon PatchPoliciesSubsetGroup `xml:"ibeacon"`
}

// GetPatchPoliciesByID retrieves the details of a patch policy by its ID.
func (c *Client) GetPatchPoliciesByID(id int) (*ResourcePatchPolicies, error) {
	// Construct the endpoint URL using the constant and the provided ID
	endpoint := fmt.Sprintf("%s/id/%d", uriPatchPolicies, id)

	var patchPolicyDetails ResourcePatchPolicies
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &patchPolicyDetails)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch patch policy by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &patchPolicyDetails, nil
}

// GetPatchPolicyByIDAndDataSubset retrieves a specific subset of data for a patch policy by its ID.
func (c *Client) GetPatchPolicyByIDAndDataSubset(id int, subset string) (*ResourcePatchPolicies, error) {
	endpoint := fmt.Sprintf("%s/id/%d/subset/%s", uriPatchPolicies, id, subset)

	var patchPolicySubset ResourcePatchPolicies
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &patchPolicySubset)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch patch policy subset by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &patchPolicySubset, nil
}

// CreatePatchPolicy creates a new patch policy.
func (c *Client) CreatePatchPolicy(policy *ResourcePatchPolicies, softwareTitleConfigID int) (*ResourcePatchPolicies, error) {
	endpoint := fmt.Sprintf("%s/softwaretitleconfig/id/%d", uriPatchPolicies, softwareTitleConfigID)

	requestBody := struct {
		XMLName xml.Name `xml:"patch_policy"`
		*ResourcePatchPolicies
	}{
		ResourcePatchPolicies: policy,
	}

	var responsePolicy ResourcePatchPolicies
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &responsePolicy)
	if err != nil {
		return nil, fmt.Errorf("failed to create patch policy: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responsePolicy, nil
}

// UpdatePatchPolicy creates a new patch policy.
func (c *Client) UpdatePatchPolicy(policy *ResourcePatchPolicies, softwareTitleConfigID int) (*ResourcePatchPolicies, error) {
	endpoint := fmt.Sprintf("%s/softwaretitleconfig/id/%d", uriPatchPolicies, softwareTitleConfigID)

	requestBody := struct {
		XMLName xml.Name `xml:"patch_policy"`
		*ResourcePatchPolicies
	}{
		ResourcePatchPolicies: policy,
	}

	var responsePolicy ResourcePatchPolicies
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responsePolicy)
	if err != nil {
		return nil, fmt.Errorf("failed to create patch policy: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responsePolicy, nil
}

// DeletePatchPolicyByID deletes a patch policy by its ID.
func (c *Client) DeletePatchPolicyByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriPatchPolicies, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete patch policy by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
