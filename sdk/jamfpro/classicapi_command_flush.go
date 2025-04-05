// classicapi_command_flush.go
// Jamf Pro Classic Api - Command Flush
// api reference: https://developer.jamf.com/jamf-pro/reference/createcommandflushwithidandstatus
// Classic API requires the structs to support an XML data structure.

package jamfpro

import "fmt"

const uriCommandFlush = "/JSSResource/commandflush"

// ClearFailedMDMCommandsByComputerID clears all failed MDM commands for a specific computer.
func (c *Client) ClearFailedComputerMDMCommandsByComputerID(computerId string) error {
	endpoint := fmt.Sprintf("%s/computers/id/%s/status/Failed", uriCommandFlush, computerId)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedActionByID, "clear failed MDM commands", computerId, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// ClearFailedComputerMDMCommandsByComputerGroupID clears all failed MDM commands for a specific
// smart or static computer group.
func (c *Client) ClearFailedComputerMDMCommandsByComputerGroupID(computerGroupId string) error {
	endpoint := fmt.Sprintf("%s/computergroups/id/%s/status/Failed", uriCommandFlush, computerGroupId)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedActionByID, "clear failed MDM commands", computerGroupId, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// ClearFailedMobileDeviceMDMCommandsByMobileDeviceID clears all failed MDM commands
// for a specific mobile device.
func (c *Client) ClearFailedMobileDeviceMDMCommandsByMobileDeviceID(mobileDeviceID string) error {
	endpoint := fmt.Sprintf("%s/mobiledevices/id/%s/status/Failed", uriCommandFlush, mobileDeviceID)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedActionByID, "clear failed MDM commands", mobileDeviceID, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// ClearFailedMobileDeviceMDMCommandsByMobileDeviceGroupID clears all failed MDM commands
// for a specific mobile device group.
func (c *Client) ClearFailedMobileDeviceMDMCommandsByMobileDeviceGroupID(mobileDeviceGroupId string) error {
	endpoint := fmt.Sprintf("%s/mobiledevicegroups/id/%s/status/Failed", uriCommandFlush, mobileDeviceGroupId)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedActionByID, "clear failed MDM commands", mobileDeviceGroupId, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// ClearPendingMDMCommandsByComputerID clears all Pending MDM commands for a specific computer.
func (c *Client) ClearPendingComputerMDMCommandsByComputerID(computerId string) error {
	endpoint := fmt.Sprintf("%s/computers/id/%s/status/Pending", uriCommandFlush, computerId)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedActionByID, "clear Pending MDM commands", computerId, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// ClearPendingComputerMDMCommandsByComputerGroupID clears all Pending MDM commands for a specific
// smart or static computer group.
func (c *Client) ClearPendingComputerMDMCommandsByComputerGroupID(computerGroupId string) error {
	endpoint := fmt.Sprintf("%s/computergroups/id/%s/status/Pending", uriCommandFlush, computerGroupId)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedActionByID, "clear Pending MDM commands", computerGroupId, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// ClearPendingMobileDeviceMDMCommandsByMobileDeviceID clears all Pending MDM commands
// for a specific mobile device.
func (c *Client) ClearPendingMobileDeviceMDMCommandsByMobileDeviceID(mobileDeviceID string) error {
	endpoint := fmt.Sprintf("%s/mobiledevices/id/%s/status/Pending", uriCommandFlush, mobileDeviceID)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedActionByID, "clear Pending MDM commands", mobileDeviceID, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// ClearPendingMobileDeviceMDMCommandsByMobileDeviceGroupID clears all Pending MDM commands
// for a specific mobile device group.
func (c *Client) ClearPendingMobileDeviceMDMCommandsByMobileDeviceGroupID(mobileDeviceGroupId string) error {
	endpoint := fmt.Sprintf("%s/mobiledevicegroups/id/%s/status/Pending", uriCommandFlush, mobileDeviceGroupId)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedActionByID, "clear Pending MDM commands", mobileDeviceGroupId, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// ClearFailedAndPendingMDMCommandsByComputerID clears all failed & pending MDM commands for a specific computer.
func (c *Client) ClearFailedAndPendingComputerMDMCommandsByComputerID(computerId string) error {
	endpoint := fmt.Sprintf("%s/computers/id/%s/status/Pending%2BFailed", uriCommandFlush, computerId)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedActionByID, "clear failed & pending MDM commands", computerId, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// ClearFailedAndPendingComputerMDMCommandsByComputerGroupID clears all failed & pending MDM commands for a specific
// smart or static computer group.
func (c *Client) ClearFailedAndPendingComputerMDMCommandsByComputerGroupID(computerGroupId string) error {
	endpoint := fmt.Sprintf("%s/computergroups/id/%s/status/Pending%2BFailed", uriCommandFlush, computerGroupId)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedActionByID, "clear failed & pending MDM commands", computerGroupId, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// ClearFailedAndPendingMobileDeviceMDMCommandsByMobileDeviceID clears all failed & pending MDM commands
// for a specific mobile device.
func (c *Client) ClearFailedAndPendingMobileDeviceMDMCommandsByMobileDeviceID(mobileDeviceID string) error {
	endpoint := fmt.Sprintf("%s/mobiledevices/id/%s/status/Pending%2BFailed", uriCommandFlush, mobileDeviceID)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedActionByID, "clear failed & pending MDM commands", mobileDeviceID, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// ClearFailedAndPendingMobileDeviceMDMCommandsByMobileDeviceGroupID clears all failed MDM commands
// for a specific mobile device group.
func (c *Client) ClearFailedAndPendingMobileDeviceMDMCommandsByMobileDeviceGroupID(mobileDeviceGroupId string) error {
	endpoint := fmt.Sprintf("%s/mobiledevicegroups/id/%s/status/Pending%2BFailed", uriCommandFlush, mobileDeviceGroupId)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedActionByID, "clear failed MDM commands", mobileDeviceGroupId, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
