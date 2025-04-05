// jamfproapi_jamf_pro_notifications.go
// Jamf Pro Api - JAMF Pro Notifications
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-notifications
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
)

const uriNotifications = "/api/v1/notifications"

// Response

// ResponseNotifications represents the JSON structure for a list of notifications
type ResponseNotifications []ResourceNotification

// ResourceNotification represents a single notification in the Jamf Pro API
type ResourceNotification struct {
	Type   string                 `json:"type"`
	ID     string                 `json:"id"`
	Params map[string]interface{} `json:"params"`
}

// CRUD

// GetNotificationsForUserAndSite retrieves all notifications from the Jamf Pro API
func (c *Client) GetNotificationsForUserAndSite() (*ResponseNotifications, error) {
	endpoint := uriNotifications

	var notifications ResponseNotifications
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &notifications)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch notifications: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &notifications, nil
}
