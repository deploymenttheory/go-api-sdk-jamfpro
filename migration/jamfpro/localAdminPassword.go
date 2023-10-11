package jamfpro

import (
	"fmt"
)

const uriLocalAdminPasswordSettings = "/api/v2/local-admin-password/settings"

type ResponseLocalAdminPasswordSettings struct {
	AutoDeployEnabled        bool `json:"autoDeployEnabled"`
	PasswordRotationTime     int  `json:"passwordRotationTime"`
	AutoRotateEnabled        bool `json:"autoRotateEnabled"`
	AutoRotateExpirationTime int  `json:"autoRotateExpirationTime"`
}

func (c *Client) GetLocalAdminPasswordSettings() (*ResponseLocalAdminPasswordSettings, error) {
	uri := uriLocalAdminPasswordSettings

	var out ResponseLocalAdminPasswordSettings
	err := c.DoRequest("GET", uri, nil, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to get local admin password settings: %v", err)
	}

	return &out, nil
}

func (c *Client) UpdateLocalAdminPasswordSettings(settings *ResponseLocalAdminPasswordSettings) (*ResponseLocalAdminPasswordSettings, error) {
	uri := uriLocalAdminPasswordSettings

	var out ResponseLocalAdminPasswordSettings
	err := c.DoRequest("PUT", uri, settings, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to update local admin password settings: %v", err)
	}

	return &out, nil
}
