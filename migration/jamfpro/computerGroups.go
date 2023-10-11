package jamfpro

import (
	"fmt"
	"sort"
)

const uriComputerGroups = "/JSSResource/computergroups"

type ComputerGroupsResponse struct {
	Size    int                         `xml:"size"`
	Results []ComputerGroupListResponse `xml:"computer_group"`
}

type ComputerGroupListResponse struct {
	ID      int    `xml:"id,omitempty"`
	Name    string `xml:"name,omitempty"`
	IsSmart bool   `xml:"is_smart,omitempty"`
}

type ComputerGroup struct {
	ID           int                          `xml:"id"`
	Name         string                       `xml:"name"`
	IsSmart      bool                         `xml:"is_smart"`
	Site         Site                         `xml:"site"`
	Criteria     []ComputerGroupCriterion     `xml:"criteria>criterion"`
	CriteriaSize int                          `xml:"criteria>size"`
	Computers    []ComputerGroupComputerEntry `xml:"computers>computer"`
	ComputerSize int                          `xml:"computers>size"`
}

type ComputerGroupRequest struct {
	Name      string                       `xml:"name"`
	IsSmart   bool                         `xml:"is_smart"`
	Site      Site                         `xml:"site"`
	Criteria  []ComputerGroupCriterion     `xml:"criteria>criterion"`
	Computers []ComputerGroupComputerEntry `xml:"computers>computer,omitempty"`
}

type ComputerGroupComputerEntry struct {
	ID           int    `json:"id,omitempty" xml:"id,omitempty"`
	Name         string `json:"name,omitempty" xml:"name,omitempty"`
	SerialNumber string `json:"serial_number,omitempty" xml:"serial_number,omitempty"`
}

type ComputerGroupCriterion struct {
	Name         string           `xml:"name"`
	Priority     int              `xml:"priority"`
	AndOr        DeviceGroupAndOr `xml:"and_or"`
	SearchType   string           `xml:"search_type"`
	SearchValue  string           `xml:"value"`
	OpeningParen bool             `xml:"opening_paren"`
	ClosingParen bool             `xml:"closing_paren"`
}

type DeviceGroupAndOr string

const (
	And DeviceGroupAndOr = "and"
	Or  DeviceGroupAndOr = "or"
)

func (c *Client) GetComputerGroupIdByName(name string) (int, error) {
	var id int
	d, err := c.GetComputerGroups()
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

func (c *Client) GetComputerGroupByName(name string) (*ComputerGroup, error) {
	id, err := c.GetComputerGroupIdByName(name)
	if err != nil {
		return nil, err
	}

	return c.GetComputerGroup(id)
}

func (c *Client) GetComputerGroup(id int) (*ComputerGroup, error) {
	var out *ComputerGroup
	uri := fmt.Sprintf("%s/id/%v", uriComputerGroups, id)
	err := c.DoRequest("GET", uri, nil, nil, &out)

	return out, err
}

func (c *Client) GetComputerGroups() (*ComputerGroupsResponse, error) {
	out := &ComputerGroupsResponse{}
	err := c.DoRequest("GET", uriComputerGroups, nil, nil, out)

	return out, err
}

func (c *Client) CreateComputerGroup(d *ComputerGroupRequest) (int, error) {

	group := &ComputerGroup{}
	d.Criteria = validateComputerGroupCriteriaOrder(d.Criteria)
	reqBody := &struct {
		*ComputerGroupRequest
		XMLName struct{} `xml:"computer_group"`
	}{
		ComputerGroupRequest: d,
	}

	err := c.DoRequest("POST", fmt.Sprintf("%s/id/0", uriComputerGroups), reqBody, nil, group)
	return group.ID, err
}

func (c *Client) UpdateComputerGroup(d *ComputerGroup) (int, error) {

	group := &ComputerGroup{}
	d.Criteria = validateComputerGroupCriteriaOrder(d.Criteria)
	uri := fmt.Sprintf("%s/id/%v", uriComputerGroups, d.ID)
	reqBody := &struct {
		*ComputerGroup
		XMLName struct{} `xml:"computer_group"`
	}{
		ComputerGroup: d,
	}

	err := c.DoRequest("PUT", uri, reqBody, nil, group)

	return group.ID, err
}

func (c *Client) DeleteComputerGroup(id int) (int, error) {

	group := &ComputerGroup{}
	uri := fmt.Sprintf("%s/id/%v", uriComputerGroups, id)
	err := c.DoRequest("DELETE", uri, nil, nil, group)

	return group.ID, err
}

func validateComputerGroupCriteriaOrder(criteria []ComputerGroupCriterion) []ComputerGroupCriterion {
	sort.Slice(criteria, func(i, j int) bool {
		return criteria[i].Priority < criteria[j].Priority
	})
	return criteria
}
