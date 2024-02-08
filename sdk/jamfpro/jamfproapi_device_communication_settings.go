// jamfproapi_device_communication_settings.go
// Jamf Pro Api - Device Communication Settings
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-device-communication-settings
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import "fmt"

const uriDeviceCommunicationSettings = "/api/v1/device-communication-settings"

// structs

type ResourceDeviceCommunicationSettings struct {
	AutoRenewMobileDeviceMdmProfileWhenCaRenewed                  bool `json:"autoRenewMobileDeviceMdmProfileWhenCaRenewed"`
	AutoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiring bool `json:"autoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiring"`
	AutoRenewComputerMdmProfileWhenCaRenewed                      bool `json:"autoRenewComputerMdmProfileWhenCaRenewed"`
	AutoRenewComputerMdmProfileWhenDeviceIdentityCertExpiring     bool `json:"autoRenewComputerMdmProfileWhenDeviceIdentityCertExpiring"`
	MdmProfileMobileDeviceExpirationLimitInDays                   int  `json:"mdmProfileMobileDeviceExpirationLimitInDays"`
	MdmProfileComputerExpirationLimitInDays                       int  `json:"mdmProfileComputerExpirationLimitInDays"`
}

// CRUD

// Gets device communication settings
func (c *Client) GetDeviceCommunicationSettings() (*ResourceDeviceCommunicationSettings, error) {
	endpoint := uriDeviceCommunicationSettings
	var out ResourceDeviceCommunicationSettings
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "device communication settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// Updates device communication settings
func (c *Client) UpdateDeviceCommunicationSettings(updatedSettings ResourceDeviceCommunicationSettings) (*ResourceDeviceCommunicationSettings, error) {
	endpoint := uriDeviceCommunicationSettings
	var out ResourceDeviceCommunicationSettings
	resp, err := c.HTTP.DoRequest("PUT", endpoint, updatedSettings, &out, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdate, "device communication settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil

}
