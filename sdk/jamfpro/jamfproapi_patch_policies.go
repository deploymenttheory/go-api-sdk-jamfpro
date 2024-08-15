// jamfproapi_patch_policies.go
// Jamf Pro Api - Patch Policies On Dashboard
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v2-patch-policies
// Jamf Pro Api requires the structs to support an JSON data structure.

package jamfpro

import (
	"fmt"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

const uriPatchPoliciesJamfProAPI = "/api/v2/patch-policies"

// List

// ResponsePatchPoliciesList represents the paginated response for patch policies
type ResponsePatchPoliciesList struct {
	TotalCount int                   `json:"totalCount"`
	Results    []ResourcePatchPolicy `json:"results"`
}

// Response

// ResponsePatchPolicyDashboardStatus represents the response for checking if a patch policy is on the dashboard
type ResponsePatchPolicyDashboardStatus struct {
	OnDashboard bool `json:"onDashboard"`
}

// Resource

// ResourcePatchPolicy represents a Patch Policy object from Pro API
type ResourcePatchPolicy struct {
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

// ResponsePatchPolicyCreate represents the response when creating a patch policy
type ResponsePatchPolicyCreate struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

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
		var newObj ResourcePatchPolicy
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "patch policy", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil
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

// RemovePatchPolicyFromDashboard removes a patch policy from the dashboard
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
