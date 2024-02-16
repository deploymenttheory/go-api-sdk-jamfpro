// classicapi_computer_groups.go
// Jamf Pro Classic Api - Computer Groups
// api reference: https://developer.jamf.com/jamf-pro/reference/computergroups
// Classic API requires the structs to support an XML data structure.

/*
Shared Resources in this Endpoint:
- SharedResourceSite
- SharedContainerCriteria
*/

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriComputerGroups = "/JSSResource/computergroups"

// List

type ResponseComputerGroupsList struct {
	Size    int                     `xml:"size"`
	Results []ComputerGroupListItem `xml:"computer_group"`
}

type ComputerGroupListItem struct {
	ID      int    `xml:"id,omitempty"`
	Name    string `xml:"name,omitempty"`
	IsSmart bool   `xml:"is_smart,omitempty"`
}

// Resource

type ResourceComputerGroup struct {
	ID        int                           `xml:"id"`
	Name      string                        `xml:"name"`
	IsSmart   bool                          `xml:"is_smart"`
	Site      SharedResourceSite            `xml:"site"`
	Criteria  SharedContainerCriteria       `xml:"criteria"`
	Computers []ComputerGroupSubsetComputer `xml:"computers>computer"`
}

// Subsets & Containers

type ComputerGroupSubsetComputer struct {
	ID            int    `json:"id,omitempty" xml:"id,omitempty"`
	Name          string `json:"name,omitempty" xml:"name,omitempty"`
	SerialNumber  string `json:"serial_number,omitempty" xml:"serial_number,omitempty"`
	MacAddress    string `json:"mac_address,omitempty" xml:"mac_address,omitempty"`
	AltMacAddress string `json:"alt_mac_address,omitempty" xml:"alt_mac_address,omitempty"`
}

// CRUD

// GetComputerGroups gets a list of all computer groups
func (c *Client) GetComputerGroups() (*ResponseComputerGroupsList, error) {
	endpoint := uriComputerGroups

	var computerGroups ResponseComputerGroupsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &computerGroups, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "computer groups", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &computerGroups, nil
}

// GetComputerGroupByID retrieves a computer group by its ID.
func (c *Client) GetComputerGroupByID(id int) (*ResourceComputerGroup, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriComputerGroups, id)

	var group ResourceComputerGroup
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &group, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "computer group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &group, nil
}

// GetComputerGroupByName retrieves a computer group by its name.
func (c *Client) GetComputerGroupByName(name string) (*ResourceComputerGroup, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriComputerGroups, name)

	var group ResourceComputerGroup
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &group, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "computer group", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &group, nil
}

// CreateComputerGroup creates a new computer group.
func (c *Client) CreateComputerGroup(group *ResourceComputerGroup) (*ResourceComputerGroup, error) {
	endpoint := uriComputerGroups

	requestBody := struct {
		XMLName xml.Name `xml:"computer_group"`
		*ResourceComputerGroup
	}{
		ResourceComputerGroup: group,
	}

	var createdGroup ResourceComputerGroup
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &createdGroup, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "computer group", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdGroup, nil
}

// UpdateComputerGroupByID updates an existing computer group by its ID.
func (c *Client) UpdateComputerGroupByID(id int, group *ResourceComputerGroup) (*ResourceComputerGroup, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriComputerGroups, id)

	if group.Site.ID == 0 && group.Site.Name == "" {
		group.Site.ID = -1
		group.Site.Name = "none"
	}

	requestBody := struct {
		XMLName xml.Name `xml:"computer_group"`
		*ResourceComputerGroup
	}{
		ResourceComputerGroup: group,
	}

	var updatedGroup ResourceComputerGroup
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedGroup, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "computer group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedGroup, nil
}

// UpdateComputerGroupByName updates a computer group by its name.
func (c *Client) UpdateComputerGroupByName(name string, group *ResourceComputerGroup) (*ResourceComputerGroup, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriComputerGroups, name)

	if group.Site.ID == 0 && group.Site.Name == "" {
		group.Site.ID = -1
		group.Site.Name = "none"
	}

	requestBody := struct {
		XMLName xml.Name `xml:"computer_group"`
		*ResourceComputerGroup
	}{
		ResourceComputerGroup: group,
	}

	var updatedGroup ResourceComputerGroup
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedGroup, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "computer group", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedGroup, nil
}

// DeleteComputerGroupByID deletes a computer group by its ID.
func (c *Client) DeleteComputerGroupByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriComputerGroups, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil, c.HTTP.Logger)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "computer group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteComputerGroupByName deletes a computer group by its name.
func (c *Client) DeleteComputerGroupByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriComputerGroups, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil, c.HTTP.Logger)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByName, "computer group", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
