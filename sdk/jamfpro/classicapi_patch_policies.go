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

// ResponsePatchPolicies represents the root element of the patch policy XML.
type ResponsePatchPolicies struct {
	General                      PatchPoliciesDataSubsetGeneral         `xml:"general"`
	Scope                        PatchPoliciesDataSubsetScope           `xml:"scope"`
	UserInteraction              PatchPoliciesDataSubsetUserInteraction `xml:"user_interaction"`
	SoftwareTitleConfigurationID int                                    `xml:"software_title_configuration_id"`
}

// PatchPoliciesDataSubsetGeneral contains general information about the patch.
type PatchPoliciesDataSubsetGeneral struct {
	ID                 int                                  `xml:"id"`
	Name               string                               `xml:"name"`
	Enabled            bool                                 `xml:"enabled"`
	TargetVersion      string                               `xml:"target_version"`
	ReleaseDate        string                               `xml:"release_date"`
	IncrementalUpdates bool                                 `xml:"incremental_updates"`
	Reboot             bool                                 `xml:"reboot"`
	MinimumOS          string                               `xml:"minimum_os"`
	KillApps           []PatchPoliciesDataSubsetKillAppItem `xml:"kill_apps>kill_app"`
	DistributionMethod string                               `xml:"distribution_method"`
	AllowDowngrade     bool                                 `xml:"allow_downgrade"`
	PatchUnknown       bool                                 `xml:"patch_unknown"`
}

// PatchPoliciesDataSubsetKillAppItem represents an item in the KillApps array.
type PatchPoliciesDataSubsetKillAppItem struct {
	KillApp PatchPoliciesDataSubsetKillApp `xml:"kill_app"`
}

// PatchPoliciesDataSubsetKillApp contains the details of an app to kill during patching.
type PatchPoliciesDataSubsetKillApp struct {
	KillAppName     string `xml:"kill_app_name"`
	KillAppBundleID string `xml:"kill_app_bundle_id"`
}

// PatchPoliciesDataSubsetScope represents the scope of the patch policy.
type PatchPoliciesDataSubsetScope struct {
	AllComputers   bool                                   `xml:"all_computers"`
	Computers      []PatchPoliciesDataSubsetComputerItem  `xml:"computers>computer"`
	ComputerGroups []PatchPoliciesDataSubsetComputerGroup `xml:"computer_groups>computer_group"`
	Buildings      []PatchPoliciesDataSubsetBuilding      `xml:"buildings>building"`
	Departments    []PatchPoliciesDataSubsetDepartment    `xml:"departments>department"`
	Limitations    PatchPoliciesDataSubsetLimitations     `xml:"limitations"`
	Exclusions     PatchPoliciesDataSubsetExclusions      `xml:"exclusions"`
}

// PatchPoliciesDataSubsetComputerItem represents a computer in the scope.
type PatchPoliciesDataSubsetComputerItem struct {
	Computer PatchPoliciesDataSubsetComputer `xml:"computer"`
}

// PatchPoliciesDataSubsetComputer contains computer details.
type PatchPoliciesDataSubsetComputer struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
	UDID string `xml:"udid"`
}

// PatchPoliciesDataSubsetComputerGroup represents a computer group in the scope.
type PatchPoliciesDataSubsetComputerGroup struct {
	ComputerGroup PatchPoliciesDataSubsetGroup `xml:"computer_group"`
}

// PatchPoliciesDataSubsetBuilding represents a building in the scope.
type PatchPoliciesDataSubsetBuilding struct {
	Building PatchPoliciesDataSubsetGroup `xml:"building"`
}

// PatchPoliciesDataSubsetDepartment represents a department in the scope.
type PatchPoliciesDataSubsetDepartment struct {
	Department PatchPoliciesDataSubsetGroup `xml:"department"`
}

// PatchPoliciesDataSubsetGroup is a general struct for group elements.
type PatchPoliciesDataSubsetGroup struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// PatchPoliciesDataSubsetLimitations represents limitations in the scope.
type PatchPoliciesDataSubsetLimitations struct {
	NetworkSegments []PatchPoliciesDataSubsetNetworkSegmentItem `xml:"network_segments>network_segment"`
	IBeacons        []PatchPoliciesDataSubsetIBeaconItem        `xml:"ibeacons>ibeacon"`
}

// PatchPoliciesDataSubsetNetworkSegmentItem represents a network segment in limitations.
type PatchPoliciesDataSubsetNetworkSegmentItem struct {
	NetworkSegment PatchPoliciesDataSubsetGroup `xml:"network_segment"`
}

// PatchPoliciesDataSubsetIBeaconItem represents an iBeacon in limitations.
type PatchPoliciesDataSubsetIBeaconItem struct {
	IBeacon PatchPoliciesDataSubsetGroup `xml:"ibeacon"`
}

// PatchPoliciesDataSubsetExclusions represents exclusions in the scope.
type PatchPoliciesDataSubsetExclusions struct {
	Computers       []PatchPoliciesDataSubsetComputerItem       `xml:"computers>computer"`
	ComputerGroups  []PatchPoliciesDataSubsetComputerGroup      `xml:"computer_groups>computer_group"`
	Buildings       []PatchPoliciesDataSubsetBuilding           `xml:"buildings>building"`
	Departments     []PatchPoliciesDataSubsetDepartment         `xml:"departments>department"`
	NetworkSegments []PatchPoliciesDataSubsetNetworkSegmentItem `xml:"network_segments>network_segment"`
	IBeacons        []PatchPoliciesDataSubsetIBeaconItem        `xml:"ibeacons>ibeacon"`
}

// PatchPoliciesDataSubsetUserInteraction contains user interaction information.
type PatchPoliciesDataSubsetUserInteraction struct {
	InstallButtonText      string                                 `xml:"install_button_text"`
	SelfServiceDescription string                                 `xml:"self_service_description"`
	SelfServiceIcon        PatchPoliciesDataSubsetSelfServiceIcon `xml:"self_service_icon"`
	Notifications          PatchPoliciesDataSubsetNotifications   `xml:"notifications"`
	Deadlines              PatchPoliciesDataSubsetDeadlines       `xml:"deadlines"`
	GracePeriod            PatchPoliciesDataSubsetGracePeriod     `xml:"grace_period"`
}

// PatchPoliciesDataSubsetSelfServiceIcon represents an icon in self-service.
type PatchPoliciesDataSubsetSelfServiceIcon struct {
	ID       int    `xml:"id"`
	Filename string `xml:"filename"`
	URI      string `xml:"uri"`
}

// PatchPoliciesDataSubsetNotifications represents notifications settings.
type PatchPoliciesDataSubsetNotifications struct {
	NotificationEnabled bool                             `xml:"notification_enabled"`
	NotificationType    string                           `xml:"notification_type"`
	NotificationSubject string                           `xml:"notification_subject"`
	NotificationMessage string                           `xml:"notification_message"`
	Reminders           PatchPoliciesDataSubsetReminders `xml:"reminders"`
}

// PatchPoliciesDataSubsetReminders represents reminder settings.
type PatchPoliciesDataSubsetReminders struct {
	NotificationRemindersEnabled  bool `xml:"notification_reminders_enabled"`
	NotificationReminderFrequency int  `xml:"notification_reminder_frequency"`
}

// PatchPoliciesDataSubsetDeadlines represents deadline settings.
type PatchPoliciesDataSubsetDeadlines struct {
	DeadlineEnabled bool `xml:"deadline_enabled"`
	DeadlinePeriod  int  `xml:"deadline_period"`
}

// PatchPoliciesDataSubsetGracePeriod represents grace period settings.
type PatchPoliciesDataSubsetGracePeriod struct {
	GracePeriodDuration       int    `xml:"grace_period_duration"`
	NotificationCenterSubject string `xml:"notification_center_subject"`
	Message                   string `xml:"message"`
}

// GetPatchPoliciesByID retrieves the details of a patch policy by its ID.
func (c *Client) GetPatchPoliciesByID(id int) (*ResponsePatchPolicies, error) {
	// Construct the endpoint URL using the constant and the provided ID
	endpoint := fmt.Sprintf("%s/id/%d", uriPatchPolicies, id)

	var patchPolicyDetails ResponsePatchPolicies
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
func (c *Client) GetPatchPolicyByIDAndDataSubset(id int, subset string) (*ResponsePatchPolicies, error) {
	endpoint := fmt.Sprintf("%s/id/%d/subset/%s", uriPatchPolicies, id, subset)

	var patchPolicySubset ResponsePatchPolicies
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
func (c *Client) CreatePatchPolicy(policy *ResponsePatchPolicies, softwareTitleConfigID int) (*ResponsePatchPolicies, error) {
	endpoint := fmt.Sprintf("%s/softwaretitleconfig/id/%d", uriPatchPolicies, softwareTitleConfigID)

	requestBody := struct {
		XMLName xml.Name `xml:"patch_policy"`
		*ResponsePatchPolicies
	}{
		ResponsePatchPolicies: policy,
	}

	var responsePolicy ResponsePatchPolicies
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
func (c *Client) UpdatePatchPolicy(policy *ResponsePatchPolicies, softwareTitleConfigID int) (*ResponsePatchPolicies, error) {
	endpoint := fmt.Sprintf("%s/softwaretitleconfig/id/%d", uriPatchPolicies, softwareTitleConfigID)

	requestBody := struct {
		XMLName xml.Name `xml:"patch_policy"`
		*ResponsePatchPolicies
	}{
		ResponsePatchPolicies: policy,
	}

	var responsePolicy ResponsePatchPolicies
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
