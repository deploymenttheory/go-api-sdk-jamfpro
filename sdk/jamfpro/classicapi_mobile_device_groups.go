// classicapi_mobile_device_enrollment_groups.go
// Jamf Pro Classic Api - Mobile Device Groups
// API reference: https://developer.jamf.com/jamf-pro/reference/mobiledevicegroups
// Jamf Pro Classic API requires the structs to support an XML data structure.

/*
Shared Resources in this Endpoint:
- SharedResourceSite
- SharedSubsetCriteria
*/

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriMobileDeviceGroups = "/JSSResource/mobiledevicegroups"

// List

// ResponseMobileDeviceGroupsList represents the response for a list of mobile device groups.
type ResponseMobileDeviceGroupsList struct {
	Size              int                          `xml:"size"`
	MobileDeviceGroup []MobileDeviceGroupsListItem `xml:"mobile_device_group"`
}

type MobileDeviceGroupsListItem struct {
	ID      int    `xml:"id"`
	Name    string `xml:"name"`
	IsSmart bool   `xml:"is_smart"`
}

// Resource

// ResourceMobileDeviceGroup represents the response for a single mobile device group.
type ResourceMobileDeviceGroup struct {
	ID                    int                                  `xml:"id"`
	Name                  string                               `xml:"name"`
	IsSmart               bool                                 `xml:"is_smart"`
	Criteria              *SharedContainerCriteria             `xml:"criteria,omitempty"`
	Site                  *SharedResourceSite                  `xml:"site"`
	MobileDevices         *[]MobileDeviceGroupSubsetDeviceItem `xml:"mobile_devices>mobile_device,omitempty"`
	MobileDeviceAdditions *[]MobileDeviceGroupSubsetDeviceItem `xml:"mobile_device_additions>mobile_device,omitempty"`
	MobileDeviceDeletions *[]MobileDeviceGroupSubsetDeviceItem `xml:"mobile_device_deletions>mobile_device,omitempty"`
}

// MobileDeviceGroupDeviceItem represents a single mobile device within a group.
type MobileDeviceGroupSubsetDeviceItem struct {
	ID             int    `xml:"id"`
	Name           string `xml:"name"`
	MacAddress     string `xml:"mac_address,omitempty"`
	UDID           string `xml:"udid"`
	WifiMacAddress string `xml:"wifi_mac_address,omitempty"`
	SerialNumber   string `xml:"serial_number,omitempty"`
}

// CRUD

// GetMobileDeviceGroups retrieves a serialized list of mobile device groups.
func (c *Client) GetMobileDeviceGroups() (*ResponseMobileDeviceGroupsList, error) {
	endpoint := uriMobileDeviceGroups

	var groups ResponseMobileDeviceGroupsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &groups)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "mobile device groups", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &groups, nil
}

// GetMobileDeviceGroupsByID retrieves a single mobile device group by its ID.
func (c *Client) GetMobileDeviceGroupByID(id string) (*ResourceMobileDeviceGroup, error) {
	endpoint := fmt.Sprintf("%s/id/%s", uriMobileDeviceGroups, id)

	var group ResourceMobileDeviceGroup
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &group)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "mobile device group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &group, nil
}

// GetMobileDeviceGroupsByName retrieves a single mobile device group by its name.
func (c *Client) GetMobileDeviceGroupByName(name string) (*ResourceMobileDeviceGroup, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMobileDeviceGroups, name)

	var group ResourceMobileDeviceGroup
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &group)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "mobile device group", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &group, nil
}

// CreateMobileDeviceGroup creates a new mobile device group on the Jamf Pro server.
func (c *Client) CreateMobileDeviceGroup(group *ResourceMobileDeviceGroup) (*ResourceMobileDeviceGroup, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriMobileDeviceGroups)

	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_group"`
		*ResourceMobileDeviceGroup
	}{
		ResourceMobileDeviceGroup: group,
	}

	var responseGroup ResourceMobileDeviceGroup
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &responseGroup)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "mobile device group", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseGroup, nil
}

// UpdateMobileDeviceGroupByID updates a mobile device group by its ID.
func (c *Client) UpdateMobileDeviceGroupByID(id string, group *ResourceMobileDeviceGroup) (*ResourceMobileDeviceGroup, error) {
	endpoint := fmt.Sprintf("%s/id/%s", uriMobileDeviceGroups, id)

	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_group"`
		*ResourceMobileDeviceGroup
	}{
		ResourceMobileDeviceGroup: group,
	}

	var updatedGroup ResourceMobileDeviceGroup
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedGroup)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "mobile device group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedGroup, nil
}

// UpdateMobileDeviceGroupByName updates a mobile device group by its name.
func (c *Client) UpdateMobileDeviceGroupByName(name string, group *ResourceMobileDeviceGroup) (*ResourceMobileDeviceGroup, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMobileDeviceGroups, name)

	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_group"`
		*ResourceMobileDeviceGroup
	}{
		ResourceMobileDeviceGroup: group,
	}

	var updatedGroup ResourceMobileDeviceGroup
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedGroup)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "mobile device group", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedGroup, nil
}

// DeleteMobileDeviceGroupByID deletes a mobile device group by its ID.
func (c *Client) DeleteMobileDeviceGroupByID(id string) error {
	endpoint := fmt.Sprintf("%s/id/%s", uriMobileDeviceGroups, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "mobile device group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteMobileDeviceGroupByName deletes a mobile device group by its name.
func (c *Client) DeleteMobileDeviceGroupByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriMobileDeviceGroups, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByName, "mobile device group", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
