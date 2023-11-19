// classicapi_mobile_device_enrollment_groups.go
// Jamf Pro Classic Api - Mobile Device Groups
// API reference: https://developer.jamf.com/jamf-pro/reference/mobiledevicegroups
// Jamf Pro Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriMobileDeviceGroups = "/JSSResource/mobiledevicegroups"

// ResponseMobileDeviceGroupsList represents the response for a list of mobile device groups.
type ResponseMobileDeviceGroupsList struct {
	Size              int                     `xml:"size"`
	MobileDeviceGroup []MobileDeviceGroupItem `xml:"mobile_device_group"`
}

// MobileDeviceGroupItem represents a single mobile device group item.
type MobileDeviceGroupItem struct {
	ID      int    `xml:"id"`
	Name    string `xml:"name"`
	IsSmart bool   `xml:"is_smart"`
}

// ResponseMobileDeviceGroup represents the response for a single mobile device group.
type ResponseMobileDeviceGroup struct {
	ID                    int                             `xml:"id"`
	Name                  string                          `xml:"name"`
	IsSmart               bool                            `xml:"is_smart"`
	Criteria              []MobileDeviceGroupCriteriaItem `xml:"criteria>criterion,omitempty"`
	Site                  MobileDeviceGroupSite           `xml:"site"`
	MobileDevices         []MobileDeviceGroupDeviceItem   `xml:"mobile_devices>mobile_device,omitempty"`
	MobileDeviceAdditions []MobileDeviceGroupDeviceItem   `xml:"mobile_device_additions>mobile_device,omitempty"`
	MobileDeviceDeletions []MobileDeviceGroupDeviceItem   `xml:"mobile_device_deletions>mobile_device,omitempty"`
}

// MobileDeviceGroupCriteriaItem represents a single criterion within a mobile device group.
type MobileDeviceGroupCriteriaItem struct {
	Name         string `xml:"name"`
	Priority     int    `xml:"priority"`
	AndOr        string `xml:"and_or"`
	SearchType   string `xml:"search_type"`
	Value        string `xml:"value"`
	OpeningParen bool   `xml:"opening_paren,omitempty"`
	ClosingParen bool   `xml:"closing_paren,omitempty"`
}

// MobileDeviceGroupSite represents the site information for a mobile device group.
type MobileDeviceGroupSite struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// MobileDeviceGroupDeviceItem represents a single mobile device within a group.
type MobileDeviceGroupDeviceItem struct {
	ID             int    `xml:"id"`
	Name           string `xml:"name"`
	MacAddress     string `xml:"mac_address,omitempty"`
	UDID           string `xml:"udid"`
	WifiMacAddress string `xml:"wifi_mac_address,omitempty"`
	SerialNumber   string `xml:"serial_number,omitempty"`
}

// GetMobileDeviceGroups retrieves a serialized list of mobile device groups.
func (c *Client) GetMobileDeviceGroups() (*ResponseMobileDeviceGroupsList, error) {
	endpoint := uriMobileDeviceGroups

	var groups ResponseMobileDeviceGroupsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &groups)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch mobile device groups: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &groups, nil
}

// GetMobileDeviceGroupsByID retrieves a single mobile device group by its ID.
func (c *Client) GetMobileDeviceGroupsByID(id int) (*ResponseMobileDeviceGroup, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriMobileDeviceGroups, id)

	var group ResponseMobileDeviceGroup
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &group)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch mobile device group by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &group, nil
}

// GetMobileDeviceGroupsByName retrieves a single mobile device group by its name.
func (c *Client) GetMobileDeviceGroupsByName(name string) (*ResponseMobileDeviceGroup, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMobileDeviceGroups, name)

	var group ResponseMobileDeviceGroup
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &group)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch mobile device group by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &group, nil
}

// CreateMobileDeviceGroup creates a new mobile device group on the Jamf Pro server.
func (c *Client) CreateMobileDeviceGroup(group *ResponseMobileDeviceGroup) (*ResponseMobileDeviceGroup, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriMobileDeviceGroups)

	// Set default values for site if not included within request
	if group.Site.ID == 0 && group.Site.Name == "" {
		group.Site = MobileDeviceGroupSite{
			ID:   -1,
			Name: "None",
		}
	}

	// Wrap the group with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_group"`
		*ResponseMobileDeviceGroup
	}{
		ResponseMobileDeviceGroup: group,
	}

	var responseGroup ResponseMobileDeviceGroup
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &responseGroup)
	if err != nil {
		return nil, fmt.Errorf("failed to create mobile device group: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseGroup, nil
}

// UpdateMobileDeviceGroupByID updates a mobile device group by its ID.
func (c *Client) UpdateMobileDeviceGroupByID(id int, group *ResponseMobileDeviceGroup) (*ResponseMobileDeviceGroup, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriMobileDeviceGroups, id)

	// Wrap the group with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_group"`
		*ResponseMobileDeviceGroup
	}{
		ResponseMobileDeviceGroup: group,
	}

	var updatedGroup ResponseMobileDeviceGroup
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedGroup)
	if err != nil {
		return nil, fmt.Errorf("failed to update mobile device group by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedGroup, nil
}

// UpdateMobileDeviceGroupByName updates a mobile device group by its name.
func (c *Client) UpdateMobileDeviceGroupByName(name string, group *ResponseMobileDeviceGroup) (*ResponseMobileDeviceGroup, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMobileDeviceGroups, name)

	// Wrap the group with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_group"`
		*ResponseMobileDeviceGroup
	}{
		ResponseMobileDeviceGroup: group,
	}

	var updatedGroup ResponseMobileDeviceGroup
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedGroup)
	if err != nil {
		return nil, fmt.Errorf("failed to update mobile device group by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedGroup, nil
}

// DeleteMobileDeviceGroupByID deletes a mobile device group by its ID.
func (c *Client) DeleteMobileDeviceGroupByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriMobileDeviceGroups, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete mobile device group by ID: %v", err)
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
		return fmt.Errorf("failed to delete mobile device group by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
