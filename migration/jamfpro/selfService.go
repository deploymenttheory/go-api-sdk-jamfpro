// selfService.go
package jamfpro

import (
	"encoding/json"
	"fmt"
	"os"
)

const uriSelfServiceSettings = "/api/v1/self-service/settings"

// SelfServiceSettings represents the settings of Self Service.
type SelfServiceSettings struct {
	InstallSettings       InstallSettings       `json:"installSettings"`
	LoginSettings         LoginSettings         `json:"loginSettings"`
	ConfigurationSettings ConfigurationSettings `json:"configurationSettings"`
}

type InstallSettings struct {
	InstallAutomatically bool   `json:"installAutomatically"`
	InstallLocation      string `json:"installLocation"`
}

type LoginSettings struct {
	UserLoginLevel  string `json:"userLoginLevel"`
	AllowRememberMe bool   `json:"allowRememberMe"`
	AuthType        string `json:"authType"`
}

type ConfigurationSettings struct {
	NotificationsEnabled  bool   `json:"notificationsEnabled"`
	AlertUserApprovedMdm  bool   `json:"alertUserApprovedMdm"`
	DefaultLandingPage    string `json:"defaultLandingPage"`
	DefaultHomeCategoryId int32  `json:"defaultHomeCategoryId"`
	BookmarksName         string `json:"bookmarksName"`
}

// GetSelfServiceSettings retrieves the Self Service settings.
func (c *Client) GetSelfServiceSettings(exportFilePath string) (*SelfServiceSettings, error) {
	var out *SelfServiceSettings

	// Call DoRequest with the GET method and desired output structure
	err := c.DoRequest("GET", uriSelfServiceSettings, nil, nil, &out)
	if err != nil {
		return nil, err
	}

	if exportFilePath != "" {
		// Marshal the output structure to JSON for exporting to a file
		responseData, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal Self Service settings: %v", err)
		}

		file, err := os.Create(exportFilePath)
		if err != nil {
			return nil, fmt.Errorf("failed to create file: %v", err)
		}
		defer file.Close()

		_, err = file.Write(responseData)
		if err != nil {
			return nil, fmt.Errorf("failed to write to file: %v", err)
		}
		fmt.Printf("Response data exported to %s\n", exportFilePath)
	}

	return out, nil
}

// UpdateSelfServiceSettings updates the Self Service settings.
func (c *Client) UpdateSelfServiceSettings(settings *SelfServiceSettings) (*SelfServiceSettings, error) {
	var out *SelfServiceSettings

	// Call DoRequest with the PUT method, the data to update, and a place to store the response
	err := c.DoRequest("PUT", uriSelfServiceSettings, settings, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to update Self Service settings: %v", err)
	}

	return out, nil
}
