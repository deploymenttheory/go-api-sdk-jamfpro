// jamfproapi_impact_alert_notification_settings.go
// Jamf Pro API - Impact Alert Notification Settings
// api reference: /v1/impact-alert-notification-settings
package jamfpro

import (
	"fmt"
)

const (
	uriImpactAlertNotificationSettings = "/api/v1/impact-alert-notification-settings"
)

// Structs

type ResourceImpactAlertNotificationSettings struct {
	ScopeableObjectsAlertEnabled             bool `json:"scopeableObjectsAlertEnabled"`
	ScopeableObjectsConfirmationCodeEnabled  bool `json:"scopeableObjectsConfirmationCodeEnabled"`
	DeployableObjectsAlertEnabled            bool `json:"deployableObjectsAlertEnabled"`
	DeployableObjectsConfirmationCodeEnabled bool `json:"deployableObjectsConfirmationCodeEnabled"`
}

// CRUD

// GetImpactAlertNotificationSettings retrieves the current Impact Alert Notification settings
func (c *Client) GetImpactAlertNotificationSettings() (*ResourceImpactAlertNotificationSettings, error) {
	endpoint := uriImpactAlertNotificationSettings
	var out ResourceImpactAlertNotificationSettings

	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "impact alert notification settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// UpdateImpactAlertNotificationSettings updates the Impact Alert Notification settings
func (c *Client) UpdateImpactAlertNotificationSettings(settings ResourceImpactAlertNotificationSettings) error {
	endpoint := uriImpactAlertNotificationSettings

	resp, _ := c.HTTP.DoRequest("PUT", endpoint, settings, nil)

	if resp == nil {
		return fmt.Errorf("failed to update Impact Alert Notification settings: received nil response")
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to update Impact Alert Notification settings: unexpected status code %d", resp.StatusCode)
	}

	return nil
}
