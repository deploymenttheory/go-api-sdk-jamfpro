// classicapi_computer_checkin.go
// Jamf Pro Classic Api - Computer Checkin
// api reference: https://developer.jamf.com/jamf-pro/reference/computercheckin
// Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriComputerCheckin = "/JSSResource/computercheckin"

// Struct for the computer check-in settings response

// Resource

type ResourceComputerCheckin struct {
	CheckInFrequency                 int  `xml:"check_in_frequency"`
	CreateStartupScript              bool `xml:"create_startup_script"`
	LogStartupEvent                  bool `xml:"log_startup_event"`
	CheckForPoliciesAtStartup        bool `xml:"check_for_policies_at_startup"`
	ApplyComputerLevelManagedPrefs   bool `xml:"apply_computer_level_managed_preferences"`
	EnsureSSHIsEnabled               bool `xml:"ensure_ssh_is_enabled"`
	CreateLoginLogoutHooks           bool `xml:"create_login_logout_hooks"`
	LogUsername                      bool `xml:"log_username"`
	CheckForPoliciesAtLoginLogout    bool `xml:"check_for_policies_at_login_logout"`
	ApplyUserLevelManagedPreferences bool `xml:"apply_user_level_managed_preferences"`
	HideRestorePartition             bool `xml:"hide_restore_partition"`
	PerformLoginActionsInBackground  bool `xml:"perform_login_actions_in_background"`
	DisplayStatusToUser              bool `xml:"display_status_to_user"`
}

// CRUD

// GetComputerCheckinInformation gets the jamf pro computer check-in settings
func (c *Client) GetComputerCheckinInformation() (*ResourceComputerCheckin, error) {
	endpoint := uriComputerCheckin

	var checkinSettings ResourceComputerCheckin
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &checkinSettings)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Computer Checkin settings: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &checkinSettings, nil
}

// UpdateComputerCheckinInformation updates the jamf pro computer check-in settings
func (c *Client) UpdateComputerCheckinInformation(settings *ResourceComputerCheckin) error {
	endpoint := uriComputerCheckin

	requestBody := struct {
		XMLName xml.Name `xml:"computer_check_in"`
		*ResourceComputerCheckin
	}{
		ResourceComputerCheckin: settings,
	}

	var handleResponse struct{}

	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &handleResponse)
	if err != nil {
		return fmt.Errorf("failed to update Computer Checkin settings: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
