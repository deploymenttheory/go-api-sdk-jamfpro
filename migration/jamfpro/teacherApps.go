// teacherApps.go
// Jamf Pro Api
// API requires the structs to support JSON.

package jamfpro

import (
	"fmt"
)

const uriTeacherApp = "/api/v1/teacher-app"

type JamfTeacherAppSettings struct {
	IsEnabled                   bool              `json:"isEnabled"`
	TimezoneId                  string            `json:"timezoneId"`
	AutoClear                   string            `json:"autoClear"`
	MaxRestrictionLengthSeconds int               `json:"maxRestrictionLengthSeconds"`
	DisplayNameType             string            `json:"displayNameType"`
	Features                    TeacherAppFeature `json:"features"`
	SafelistedApps              []SafelistedApp   `json:"safelistedApps"`
}

type TeacherAppFeature struct {
	IsAllowAppLock         bool `json:"isAllowAppLock"`
	IsAllowWebLock         bool `json:"isAllowWebLock"`
	IsAllowRestrictions    bool `json:"isAllowRestrictions"`
	IsAllowAttentionScreen bool `json:"isAllowAttentionScreen"`
	IsAllowClearPasscode   bool `json:"isAllowClearPasscode"`
}

type SafelistedApp struct {
	Name     string `json:"name"`
	BundleId string `json:"bundleId"`
}

type TeacherAppHistoryResponse struct {
	TotalCount int                     `json:"totalCount"`
	Results    []TeacherAppHistoryItem `json:"results"`
}

type TeacherAppHistoryItem struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Date     string `json:"date"`
	Note     string `json:"note"`
	Details  string `json:"details"`
}

// GetTeacherAppSettings retrieves the Jamf Teacher app settings
func (c *Client) GetTeacherAppSettings() (*JamfTeacherAppSettings, error) {
	url := uriTeacherApp

	var settings JamfTeacherAppSettings
	if err := c.DoRequest("GET", url, nil, nil, &settings); err != nil {
		return nil, fmt.Errorf("failed to fetch Jamf Teacher app settings: %v", err)
	}

	return &settings, nil
}

// GetTeacherAppSettingsHistory retrieves the Jamf Teacher app settings history
func (c *Client) GetTeacherAppSettingsHistory(page, pageSize int, sort, filter string) (*TeacherAppHistoryResponse, error) {
	url := fmt.Sprintf("%s/history?page=%d&page-size=%d&sort=%s&filter=%s", uriTeacherApp, page, pageSize, sort, filter)

	var history TeacherAppHistoryResponse
	if err := c.DoRequest("GET", url, nil, nil, &history); err != nil {
		return nil, fmt.Errorf("failed to fetch Jamf Teacher app settings history: %v", err)
	}

	return &history, nil
}
