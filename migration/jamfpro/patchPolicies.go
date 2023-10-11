// patchPolicies.go
// Jamf Pro Api

package jamfpro

import (
	"fmt"
)

const uriAPIPatchPolicies = "/api/v2/patch-policies/"

// PatchPolicyDashboardStatus represents the status of the Patch Policy on the dashboard
type PatchPolicyDashboardStatus struct {
	OnDashboard bool `json:"onDashboard,omitempty" xml:"onDashboard,omitempty"`
}

// ResponsePatchPolicy represents the entire response structure for patch policies
type ResponsePatchPolicy struct {
	TotalCount int                          `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	Results    PatchPolicyDataSubsetResults `json:"results,omitempty" xml:"results,omitempty"`
}

type PatchPolicyDataSubsetResults struct {
	PatchPolicyDetails PatchPolicyDataSubsetPatchPolicyDetails `json:"patchPolicy,omitempty" xml:"patchPolicy,omitempty"`
}

type PatchPolicyDataSubsetPatchPolicyDetails struct {
	ID                           string `json:"id,omitempty" xml:"id,omitempty"`
	Name                         string `json:"name,omitempty" xml:"name,omitempty"`
	Enabled                      bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
	TargetPatchVersion           string `json:"targetPatchVersion,omitempty" xml:"targetPatchVersion,omitempty"`
	DeploymentMethod             string `json:"deploymentMethod,omitempty" xml:"deploymentMethod,omitempty"`
	SoftwareTitleId              string `json:"softwareTitleId,omitempty" xml:"softwareTitleId,omitempty"`
	SoftwareTitleConfigurationId string `json:"softwareTitleConfigurationId,omitempty" xml:"softwareTitleConfigurationId,omitempty"`
	KillAppsDelayMinutes         int    `json:"killAppsDelayMinutes,omitempty" xml:"killAppsDelayMinutes,omitempty"`
	KillAppsMessage              string `json:"killAppsMessage,omitempty" xml:"killAppsMessage,omitempty"`
	Downgrade                    bool   `json:"downgrade,omitempty" xml:"downgrade,omitempty"`
	PatchUnknownVersion          bool   `json:"patchUnknownVersion,omitempty" xml:"patchUnknownVersion,omitempty"`
	NotificationHeader           string `json:"notificationHeader,omitempty" xml:"notificationHeader,omitempty"`
	SelfServiceEnforceDeadline   bool   `json:"selfServiceEnforceDeadline,omitempty" xml:"selfServiceEnforceDeadline,omitempty"`
	SelfServiceDeadline          int    `json:"selfServiceDeadline,omitempty" xml:"selfServiceDeadline,omitempty"`
	InstallButtonText            string `json:"installButtonText,omitempty" xml:"installButtonText,omitempty"`
	SelfServiceDescription       string `json:"selfServiceDescription,omitempty" xml:"selfServiceDescription,omitempty"`
	IconId                       string `json:"iconId,omitempty" xml:"iconId,omitempty"`
	ReminderFrequency            int    `json:"reminderFrequency,omitempty" xml:"reminderFrequency,omitempty"`
	ReminderEnabled              bool   `json:"reminderEnabled,omitempty" xml:"reminderEnabled,omitempty"`
}

type PatchPolicyDataSubsetDashboardStatus struct {
	OnDashboard bool `json:"onDashboard,omitempty" xml:"onDashboard,omitempty"`
}

// GetPatchPolicy retrieves the details of a patch policy
func (c *Client) GetPatchPolicy(policyID string) (*ResponsePatchPolicy, error) {
	// Update the endpoint to include "policy-details/"
	url := fmt.Sprintf("%spolicy-details/%s", uriAPIPatchPolicies, policyID)

	var policy ResponsePatchPolicy
	if err := c.DoRequest("GET", url, nil, nil, &policy); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &policy, nil
}

// GetPatchPolicyDashboardStatus retrieves whether the Patch Policy is on the Dashboard
func (c *Client) GetPatchPolicyDashboardStatus(policyID string) (*PatchPolicyDashboardStatus, error) {
	url := fmt.Sprintf("%s%s/dashboard", uriAPIPatchPolicies, policyID)

	var dashboardStatus PatchPolicyDashboardStatus
	if err := c.DoRequest("GET", url, nil, nil, &dashboardStatus); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &dashboardStatus, nil
}
