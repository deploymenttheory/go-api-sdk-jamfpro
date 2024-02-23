package jamfpro

import (
	"fmt"
)

const uriDeviceCommunicationSettings = "/api/v1/device-communication-settings"

type ResponseDeviceCommunicationSettings struct {
	AutoRenewMobileDeviceMdmProfileWhenCaRenewed                  bool `json:"autoRenewMobileDeviceMdmProfileWhenCaRenewed"`
	AutoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiring bool `json:"autoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiring"`
	AutoRenewComputerMdmProfileWhenCaRenewed                      bool `json:"autoRenewComputerMdmProfileWhenCaRenewed"`
	AutoRenewComputerMdmProfileWhenDeviceIdentityCertExpiring     bool `json:"autoRenewComputerMdmProfileWhenDeviceIdentityCertExpiring"`
	MdmProfileMobileDeviceExpirationLimitInDays                   int  `json:"mdmProfileMobileDeviceExpirationLimitInDays"`
	MdmProfileComputerExpirationLimitInDays                       int  `json:"mdmProfileComputerExpirationLimitInDays"`
}

func (c *Client) GetDeviceCommunicationSettings() (*ResponseDeviceCommunicationSettings, error) {
	uri := uriDeviceCommunicationSettings

	var out ResponseDeviceCommunicationSettings
	err := c.DoRequest("GET", uri, nil, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to get device communication settings: %v", err)
	}

	return &out, nil
}

func (c *Client) UpdateDeviceCommunicationSettings(
	autoRenewMobileDeviceMdmProfileWhenCaRenewed *bool,
	autoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiring *bool,
	autoRenewComputerMdmProfileWhenCaRenewed *bool,
	autoRenewComputerMdmProfileWhenDeviceIdentityCertExpiring *bool,
	mdmProfileMobileDeviceExpirationLimitInDays *int,
	mdmProfileComputerExpirationLimitInDays *int) (*ResponseDeviceCommunicationSettings, error) {

	in := struct {
		AutoRenewMobileDeviceMdmProfileWhenCaRenewed                  *bool `json:"autoRenewMobileDeviceMdmProfileWhenCaRenewed"`
		AutoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiring *bool `json:"autoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiring"`
		AutoRenewComputerMdmProfileWhenCaRenewed                      *bool `json:"autoRenewComputerMdmProfileWhenCaRenewed"`
		AutoRenewComputerMdmProfileWhenDeviceIdentityCertExpiring     *bool `json:"autoRenewComputerMdmProfileWhenDeviceIdentityCertExpiring"`
		MdmProfileMobileDeviceExpirationLimitInDays                   *int  `json:"mdmProfileMobileDeviceExpirationLimitInDays"`
		MdmProfileComputerExpirationLimitInDays                       *int  `json:"mdmProfileComputerExpirationLimitInDays"`
	}{
		AutoRenewMobileDeviceMdmProfileWhenCaRenewed:                  autoRenewMobileDeviceMdmProfileWhenCaRenewed,
		AutoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiring: autoRenewMobileDeviceMdmProfileWhenDeviceIdentityCertExpiring,
		AutoRenewComputerMdmProfileWhenCaRenewed:                      autoRenewComputerMdmProfileWhenCaRenewed,
		AutoRenewComputerMdmProfileWhenDeviceIdentityCertExpiring:     autoRenewComputerMdmProfileWhenDeviceIdentityCertExpiring,
		MdmProfileMobileDeviceExpirationLimitInDays:                   mdmProfileMobileDeviceExpirationLimitInDays,
		MdmProfileComputerExpirationLimitInDays:                       mdmProfileComputerExpirationLimitInDays,
	}

	var out ResponseDeviceCommunicationSettings
	uri := uriDeviceCommunicationSettings
	err := c.DoRequest("PUT", uri, in, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to update device communication settings: %v", err)
	}

	return &out, nil
}
