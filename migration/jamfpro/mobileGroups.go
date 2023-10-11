package jamfpro

import (
	"fmt"

	"github.com/bradfitz/slice"
)

const uriMobileDeviceGroups = "/JSSResource/mobiledevicegroups"

type MobileDeviceGroupsResponse struct {
	Size    int                             `xml:"size"`
	Results []MobileDeviceGroupListResponse `xml:"mobile_device_group"`
}

type MobileDeviceGroupListResponse struct {
	ID      int    `xml:"id,omitempty"`
	Name    string `xml:"name,omitempty"`
	IsSmart bool   `xml:"is_smart,omitempty"`
}

type MobileDeviceGroup struct {
	ID           int                            `xml:"id"`
	Name         string                         `xml:"name"`
	IsSmart      bool                           `xml:"is_smart"`
	Site         Site                           `xml:"site"`
	Criteria     []MobileDeviceGroupCriterion   `xml:"criteria>criterion"`
	CriteriaSize int                            `xml:"criteria>size"`
	Devices      []MobileDeviceGroupDeviceEntry `xml:"mobile_devices>mobile_device"`
	DeviceSize   int                            `xml:"mobile_devices>size"`
}

type MobileDeviceGroupRequest struct {
	Name     string                         `xml:"name"`
	IsSmart  bool                           `xml:"is_smart"`
	Site     Site                           `xml:"site"`
	Criteria []MobileDeviceGroupCriterion   `xml:"criteria>criterion"`
	Devices  []MobileDeviceGroupDeviceEntry `xml:"mobile_devices>mobile_device,omitempty"`
}

type MobileDeviceGroupDeviceEntry struct {
	ID           int    `json:"id,omitempty" xml:"id,omitempty"`
	Name         string `json:"name,omitempty" xml:"name,omitempty"`
	SerialNumber string `json:"serial_number,omitempty" xml:"serial_number,omitempty"`
}

type MobileDeviceGroupCriterion struct {
	Name         string           `xml:"name"`
	Priority     int              `xml:"priority"`
	AndOr        DeviceGroupAndOr `xml:"and_or"`
	SearchType   string           `xml:"search_type"`
	SearchValue  string           `xml:"value"`
	OpeningParen bool             `xml:"opening_paren"`
	ClosingParen bool             `xml:"closing_paren"`
}

func (c *Client) GetMobileDeviceGroupIdByName(name string) (int, error) {
	var id int
	d, err := c.GetMobileDeviceGroups()
	if err != nil {
		return -1, err
	}

	for _, v := range d.Results {
		if v.Name == name {
			id = v.ID
			break
		}
	}
	return id, err
}

func (c *Client) GetMobileDeviceGroupByName(name string) (*MobileDeviceGroup, error) {
	id, err := c.GetMobileDeviceGroupIdByName(name)
	if err != nil {
		return nil, err
	}

	return c.GetMobileDeviceGroup(id)
}

func (c *Client) GetMobileDeviceGroup(id int) (*MobileDeviceGroup, error) {
	var out *MobileDeviceGroup
	uri := fmt.Sprintf("%s/id/%v", uriMobileDeviceGroups, id)
	err := c.DoRequest("GET", uri, nil, nil, &out)

	return out, err
}

func (c *Client) GetMobileDeviceGroups() (*MobileDeviceGroupsResponse, error) {
	out := &MobileDeviceGroupsResponse{}
	err := c.DoRequest("GET", uriMobileDeviceGroups, nil, nil, out)

	return out, err
}

func (c *Client) CreateMobileDeviceGroup(d *MobileDeviceGroupRequest) (int, error) {

	group := &MobileDeviceGroup{}
	d.Criteria = validateMobileDeviceGroupCriteriaOrder(d.Criteria)
	reqBody := &struct {
		*MobileDeviceGroupRequest
		XMLName struct{} `xml:"mobile_device_group"`
	}{
		MobileDeviceGroupRequest: d,
	}

	err := c.DoRequest("POST", fmt.Sprintf("%s/id/0", uriMobileDeviceGroups), reqBody, nil, group)
	return group.ID, err
}

func (c *Client) UpdateMobileDeviceGroup(d *MobileDeviceGroup) (int, error) {

	group := &MobileDeviceGroup{}
	d.Criteria = validateMobileDeviceGroupCriteriaOrder(d.Criteria)
	uri := fmt.Sprintf("%s/id/%v", uriMobileDeviceGroups, d.ID)
	reqBody := &struct {
		*MobileDeviceGroup
		XMLName struct{} `xml:"mobile_device_group"`
	}{
		MobileDeviceGroup: d,
	}

	err := c.DoRequest("PUT", uri, reqBody, nil, group)

	return group.ID, err
}

func (c *Client) DeleteMobileDeviceGroup(id int) (int, error) {

	group := &MobileDeviceGroup{}
	uri := fmt.Sprintf("%s/id/%v", uriMobileDeviceGroups, id)
	err := c.DoRequest("DELETE", uri, nil, nil, group)

	return group.ID, err
}

func validateMobileDeviceGroupCriteriaOrder(criteria []MobileDeviceGroupCriterion) []MobileDeviceGroupCriterion {
	slice.Sort(criteria[:], func(i, j int) bool {
		return criteria[i].Priority < criteria[j].Priority
	})
	return criteria
}
