// jamfproapi_user_sessions.go
// Jamf Pro Api - User Sessions
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-user-sessions-active
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import "fmt"

const uriUserSessions = "/api/v1/user-sessions"

// Resource

// ResourceActiveUserSession represents an active (currently logged in) user session.
type ResourceActiveUserSession struct {
	Username         string `json:"username,omitempty"`
	SessionID        string `json:"sessionId,omitempty"`
	LastAccessedTime string `json:"lastAccessedTime,omitempty"`
	CreationTime     string `json:"creationTime,omitempty"`
	UserAgent        string `json:"userAgent,omitempty"`
	IPAddress        string `json:"ipAddress,omitempty"`
}

// ResponseActiveUsersCount represents the count of currently logged in users.
type ResponseActiveUsersCount struct {
	ActiveUserCount int `json:"activeUserCount"`
}

// GetActiveUserSessions retrieves the list of currently active user sessions.
func (c *Client) GetActiveUserSessions() ([]ResourceActiveUserSession, error) {
	endpoint := fmt.Sprintf("%s/active", uriUserSessions)

	var out []ResourceActiveUserSession
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "active user sessions", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return out, nil
}

// GetActiveUsersCount retrieves the count of currently logged in users.
func (c *Client) GetActiveUsersCount() (*ResponseActiveUsersCount, error) {
	endpoint := fmt.Sprintf("%s/count", uriUserSessions)

	var out ResponseActiveUsersCount
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "active users count", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}
