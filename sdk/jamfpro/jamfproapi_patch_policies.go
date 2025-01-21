// jamfproapi_patch_policies.go
// Jamf Pro Api - Patch Policies On Dashboard
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v2-patch-policies
// Jamf Pro Api requires the structs to support an JSON data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

const (
	uriPatchPoliciesJamfProAPI = "/api/v2/patch-policies"
	uriPatchPoliciesClassicAPI = "/JSSResource/patchpolicies"
)

// List

// ResponsePatchPoliciesList represents the paginated response for patch policies
type ResponsePatchPoliciesList struct {
	TotalCount int                             `json:"totalCount"`
	Results    []ResourcePatchPolicyJamfProAPI `json:"results"`
}

// Response

// ResponsePatchPolicyCreate represents the response when creating a patch policy
type ResponsePatchPolicyCreate struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// ResponsePatchPolicyDashboardStatus represents the response for checking if a patch policy is on the dashboard
type ResponsePatchPolicyDashboardStatus struct {
	OnDashboard bool `json:"onDashboard"`
}

// Resource

// ResourcePatchPolicy represents a Patch Policy object from Pro API
type ResourcePatchPolicyJamfProAPI struct {
	ID                           string `json:"id"`
	Name                         string `json:"name"`
	Enabled                      bool   `json:"enabled"`
	TargetPatchVersion           string `json:"targetPatchVersion"`
	DeploymentMethod             string `json:"deploymentMethod"`
	SoftwareTitleId              string `json:"softwareTitleId"`
	SoftwareTitleConfigurationId string `json:"softwareTitleConfigurationId"`
	KillAppsDelayMinutes         int    `json:"killAppsDelayMinutes"`
	KillAppsMessage              string `json:"killAppsMessage"`
	Downgrade                    bool   `json:"downgrade"`
	PatchUnknownVersion          bool   `json:"patchUnknownVersion"`
	NotificationHeader           string `json:"notificationHeader"`
	SelfServiceEnforceDeadline   bool   `json:"selfServiceEnforceDeadline"`
	SelfServiceDeadline          int    `json:"selfServiceDeadline"`
	InstallButtonText            string `json:"installButtonText"`
	SelfServiceDescription       string `json:"selfServiceDescription"`
	IconId                       string `json:"iconId"`
	ReminderFrequency            int    `json:"reminderFrequency"`
	ReminderEnabled              bool   `json:"reminderEnabled"`
}

// ResourcePatchPolicyCreateRequest represents the XML structure for creating a patch policy
type ResourcePatchPolicyClassicAPI struct {
	General                      ResourcePatchPolicyCreateRequestGeneral         `xml:"general"`
	Scope                        ResourcePatchPolicyCreateRequestScope           `xml:"scope"`
	UserInteraction              ResourcePatchPolicyCreateRequestUserInteraction `xml:"user_interaction"`
	SoftwareTitleConfigurationID string                                          `xml:"software_title_configuration_id"`
}

type ResourcePatchPolicyCreateRequestGeneral struct {
	Name               string                                   `xml:"name"`
	Enabled            bool                                     `xml:"enabled"`
	TargetVersion      string                                   `xml:"target_version"`
	ReleaseDate        string                                   `xml:"release_date"`
	IncrementalUpdates bool                                     `xml:"incremental_updates"`
	Reboot             bool                                     `xml:"reboot"`
	MinimumOS          string                                   `xml:"minimum_os"`
	KillApps           ResourcePatchPolicyCreateRequestKillApps `xml:"kill_apps"`
	DistributionMethod string                                   `xml:"distribution_method"`
	AllowDowngrade     bool                                     `xml:"allow_downgrade"`
	PatchUnknown       bool                                     `xml:"patch_unknown"`
}

type ResourcePatchPolicyCreateRequestKillApps struct {
	KillApp []ResourcePatchPolicyCreateRequestKillApp `xml:"kill_app"`
}

type ResourcePatchPolicyCreateRequestKillApp struct {
	KillAppName     string `xml:"kill_app_name"`
	KillAppBundleID string `xml:"kill_app_bundle_id"`
}

type ResourcePatchPolicyCreateRequestScope struct {
	AllComputers bool `xml:"all_computers"`
}

type ResourcePatchPolicyCreateRequestUserInteraction struct {
	InstallButtonText      string                                        `xml:"install_button_text"`
	SelfServiceDescription string                                        `xml:"self_service_description"`
	Notifications          ResourcePatchPolicyCreateRequestNotifications `xml:"notifications"`
	Deadlines              ResourcePatchPolicyCreateRequestDeadlines     `xml:"deadlines"`
	GracePeriod            ResourcePatchPolicyCreateRequestGracePeriod   `xml:"grace_period"`
}

type ResourcePatchPolicyCreateRequestNotifications struct {
	Enabled   bool                                      `xml:"notification_enabled"`
	Type      string                                    `xml:"notification_type"`
	Subject   string                                    `xml:"notification_subject"`
	Message   string                                    `xml:"notification_message"`
	Reminders ResourcePatchPolicyCreateRequestReminders `xml:"reminders"`
}

type ResourcePatchPolicyCreateRequestReminders struct {
	Enabled   bool `xml:"notification_reminders_enabled"`
	Frequency int  `xml:"notification_reminder_frequency"`
}

type ResourcePatchPolicyCreateRequestDeadlines struct {
	Enabled bool `xml:"deadline_enabled"`
	Period  int  `xml:"deadline_period"`
}

type ResourcePatchPolicyCreateRequestGracePeriod struct {
	Duration            int    `xml:"grace_period_duration"`
	NotificationSubject string `xml:"notification_center_subject"`
	Message             string `xml:"message"`
}

// CRUD

// GetPatchPolicies gets the full list of patch policies & handles pagination
func (c *Client) GetPatchPolicies(sortFilter string) (*ResponsePatchPoliciesList, error) {
	resp, err := c.DoPaginatedGet(
		uriPatchPoliciesJamfProAPI+"/policy-details",
		standardPageSize,
		startingPageNumber,
		sortFilter,
	)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "patch policies", err)
	}

	var out ResponsePatchPoliciesList
	out.TotalCount = resp.Size

	for _, value := range resp.Results {
		var newObj ResourcePatchPolicyJamfProAPI
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "patch policy", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil
}

// GetPatchPolicyByID retrieves a specific patch policy by ID
func (c *Client) GetPatchPolicyByID(id string) (*ResourcePatchPolicyJamfProAPI, error) {
	if id == "" {
		return nil, fmt.Errorf("patch policy ID cannot be empty")
	}

	policies, err := c.GetPatchPolicies("")
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "patch policy", err)
	}

	for _, policy := range policies.Results {
		if policy.ID == id {
			return &policy, nil
		}
	}

	return nil, fmt.Errorf("no patch policy found with ID: %s", id)
}

// GetPatchPolicyByName retrieves a specific patch policy by name
func (c *Client) GetPatchPolicyByName(name string) (*ResourcePatchPolicyJamfProAPI, error) {
	if name == "" {
		return nil, fmt.Errorf("patch policy name cannot be empty")
	}

	policies, err := c.GetPatchPolicies("")
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "patch policy", err)
	}

	for _, policy := range policies.Results {
		if policy.Name == name {
			return &policy, nil
		}
	}

	return nil, fmt.Errorf("no patch policy found with name: %s", name)
}

// GetPatchPolicyDashboardStatus checks if a patch policy is on the dashboard
func (c *Client) GetPatchPolicyDashboardStatus(id string) (*ResponsePatchPolicyDashboardStatus, error) {
	endpoint := fmt.Sprintf("%s/%s/dashboard", uriPatchPoliciesJamfProAPI, id)

	var response ResponsePatchPolicyDashboardStatus
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to check patch policy dashboard status: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// CreatePatchPolicy creates a new patch policy for a specific software title configuration
func (c *Client) CreatePatchPolicy(softwareTitleConfigID string, policyRequest *ResourcePatchPolicyClassicAPI) error {
	endpoint := fmt.Sprintf("%s/softwaretitleconfig/id/%s", uriPatchPoliciesClassicAPI, softwareTitleConfigID)

	requestBody := struct {
		XMLName xml.Name `xml:"patch_policy"`
		*ResourcePatchPolicyClassicAPI
	}{
		ResourcePatchPolicyClassicAPI: policyRequest,
	}

	resp, err := c.HTTP.DoRequest("POST", endpoint, requestBody, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedCreate, "patch policy", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// UpdatePatchPolicyByID updates an existing patch policy
func (c *Client) UpdatePatchPolicyByID(softwareTitleConfigID string, policyRequest *ResourcePatchPolicyClassicAPI) error {
	endpoint := fmt.Sprintf("%s/softwaretitleconfig/id/%s", uriPatchPoliciesClassicAPI, softwareTitleConfigID)

	requestBody := struct {
		XMLName xml.Name `xml:"patch_policy"`
		*ResourcePatchPolicyClassicAPI
	}{
		ResourcePatchPolicyClassicAPI: policyRequest,
	}

	resp, err := c.HTTP.DoRequest("PUT", endpoint, requestBody, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedCreate, "patch policy", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeletePatchPolicyByID deletes a patch policy by ID
func (c *Client) DeletePatchPolicyByID(id string) error {
	if id == "" {
		return fmt.Errorf("patch policy ID cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/id/%s", uriPatchPoliciesClassicAPI, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "patch policy", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// AddPatchPolicyToDashboard adds a patch policy to the dashboard
func (c *Client) AddPatchPolicyToDashboard(id string) error {
	endpoint := fmt.Sprintf("%s/%s/dashboard", uriPatchPoliciesJamfProAPI, id)

	resp, err := c.HTTP.DoRequest("POST", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to add patch policy to dashboard: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	// Check if the response status code indicates success
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to add patch policy to dashboard: unexpected status code %d", resp.StatusCode)
	}

	return nil
}

// DeletePatchPolicyFromDashboard removes a patch policy from the dashboard
func (c *Client) DeletePatchPolicyFromDashboard(id string) error {
	endpoint := fmt.Sprintf("%s/%s/dashboard", uriPatchPoliciesJamfProAPI, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to remove patch policy from dashboard: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
