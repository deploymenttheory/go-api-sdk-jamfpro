package jamfpro

import (
	"fmt"
)

// Constants
const uriAPIConditionalAccess = "/api/v1/conditional-access/device-compliance-information"

// Structures
type ResponseConditionalAccess struct {
	TotalCount int                            `json:"totalCount,omitempty"`
	Results    []ConditionalAccessDeviceState `json:"results"`
}

type ConditionalAccessDeviceState struct {
	DeviceId                          string                            `json:"deviceId"`
	Applicable                        bool                              `json:"applicable"`
	ComplianceState                   string                            `json:"complianceState"`
	ComplianceVendor                  string                            `json:"complianceVendor"`
	ComplianceVendorDeviceInformation ComplianceVendorDeviceInformation `json:"complianceVendorDeviceInformation"`
}

type ComplianceVendorDeviceInformation struct {
	DeviceIds []string `json:"deviceIds"`
}

// Functions

func (c *Client) GetConditionalAccessComplianceStateDeviceIdByName(deviceType string, name string, deviceID int) (string, error) {
	var id string
	states, err := c.GetConditionalAccessComplianceStateByDeviceTypeAndID(deviceType, deviceID)
	if err != nil {
		return "", err
	}

	for _, v := range states.Results {
		if v.DeviceId == name {
			id = v.DeviceId
			break
		}
	}
	return id, err
}

func (c *Client) GetConditionalAccessComplianceStateByDeviceTypeAndDeviceName(deviceType string, deviceName string) (*ConditionalAccessDeviceState, error) {
	// First, fetch all devices of the given type
	uri := fmt.Sprintf("%s/%s", uriAPIConditionalAccess, deviceType)

	var out ResponseConditionalAccess
	err := c.DoRequest("GET", uri, nil, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to get all Conditional Access Device States for %s: %v", deviceType, err)
	}

	// Iterate through the list to find the one with the desired name
	for _, state := range out.Results {
		if state.DeviceId == deviceName {
			return &state, nil
		}
	}

	return nil, fmt.Errorf("conditional Access Device State with name '%s' not found", deviceName)
}

func (c *Client) GetConditionalAccessComplianceStateByDeviceTypeAndID(deviceType string, deviceID int) (*ResponseConditionalAccess, error) {
	uri := fmt.Sprintf("%s/%s/%d", uriAPIConditionalAccess, deviceType, deviceID)

	var out ResponseConditionalAccess
	err := c.DoRequest("GET", uri, nil, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to get Conditional Access Device State for %s ID %d: %v", deviceType, deviceID, err)
	}

	return &out, nil
}

func (c *Client) GetConditionalAccessComplianceStateByDeviceID(deviceType string, deviceID int) (*ConditionalAccessDeviceState, error) {
	uri := fmt.Sprintf("%s/%s/%d", uriAPIConditionalAccess, deviceType, deviceID)

	var out ConditionalAccessDeviceState
	err := c.DoRequest("GET", uri, nil, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to get Conditional Access Device State by ID: %v", err)
	}

	return &out, nil
}
