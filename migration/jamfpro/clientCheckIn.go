// clientCheckIn.go
// Jamf Pro Api
// API requires the structs to support JSON.

package jamfpro

import (
	"fmt"
)

const uriAPICheckIn = "/api/v3/check-in"

type ResponseClientCheckIn struct {
	CheckInFrequency                 int  `json:"checkInFrequency"`
	CreateHooks                      bool `json:"createHooks"`
	HookLog                          bool `json:"hookLog"`
	HookPolicies                     bool `json:"hookPolicies"`
	CreateStartupScript              bool `json:"createStartupScript"`
	StartupLog                       bool `json:"startupLog"`
	StartupPolicies                  bool `json:"startupPolicies"`
	StartupSsh                       bool `json:"startupSsh"`
	EnableLocalConfigurationProfiles bool `json:"enableLocalConfigurationProfiles"`
}

type ClientCheckInHistoryResult struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Date     string `json:"date"`
	Note     string `json:"note"`
	Details  string `json:"details"`
}

type ResponseCheckInHistoryNote struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

type ResponseClientCheckInHistory struct {
	TotalCount int                          `json:"totalCount"`
	Results    []ClientCheckInHistoryResult `json:"results"`
}

func (c *Client) GetClientCheckIn() (*ResponseClientCheckIn, error) {
	var out ResponseClientCheckIn
	err := c.DoRequest("GET", uriAPICheckIn, nil, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to get check-in data: %v", err)
	}

	return &out, nil
}

func (c *Client) UpdateclientCheckIn(d *ResponseClientCheckIn) (*ResponseClientCheckIn, error) {
	updatedCheckIn := &ResponseClientCheckIn{}
	err := c.DoRequest("PUT", uriAPICheckIn, d, nil, updatedCheckIn)
	if err != nil {
		return nil, fmt.Errorf("failed to update check-in data: %v", err)
	}

	return updatedCheckIn, nil
}

func (c *Client) GetClientCheckInHistory(page int, pageSize int, sort string, filter string) (*ResponseClientCheckInHistory, error) {
	uri := fmt.Sprintf("%s/history?page=%d&page-size=%d&sort=%s&filter=%s", uriAPICheckIn, page, pageSize, sort, filter)

	var out ResponseClientCheckInHistory
	err := c.DoRequest("GET", uri, nil, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to get client check-in history: %v", err)
	}

	return &out, nil
}

func (c *Client) AddNoteToCheckInHistory(note string) (*ResponseCheckInHistoryNote, error) {
	uri := fmt.Sprintf("%s/history", uriAPICheckIn)

	payload := struct {
		Note string `json:"note"`
	}{
		Note: note,
	}

	newNote := &ResponseCheckInHistoryNote{}

	err := c.DoRequest("POST", uri, &payload, nil, newNote)
	if err != nil {
		return nil, fmt.Errorf("failed to add note to check-in history: %v", err)
	}

	return newNote, nil
}
