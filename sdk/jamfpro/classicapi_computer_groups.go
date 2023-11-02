// classicapi_computer_groups.go
// Jamf Pro Classic Api - Computer Groups
// api reference: https://developer.jamf.com/jamf-pro/reference/computergroups
// Classic API requires the structs to support an XML data structure.
package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriComputerGroups = "/JSSResource/computergroups"

type ComputerGroupsListResponse struct {
	Size    int                     `xml:"size"`
	Results []ComputerGroupListItem `xml:"computer_group"`
}

type ComputerGroupListItem struct {
	ID      int    `xml:"id,omitempty"`
	Name    string `xml:"name,omitempty"`
	IsSmart bool   `xml:"is_smart,omitempty"`
}

type ResponseComputerGroup struct {
	ID        int                         `xml:"id"`
	Name      string                      `xml:"name"`
	IsSmart   bool                        `xml:"is_smart"`
	Site      ComputerGroupSite           `xml:"site"`
	Criteria  []ComputerGroupCriterion    `xml:"criteria>criterion"`
	Computers []ComputerGroupComputerItem `xml:"computers>computer"`
}

type CriterionContainer struct {
	Size      int                    `xml:"size"`
	Criterion ComputerGroupCriterion `xml:"criterion"`
}

type ComputerContainer struct {
	Size     int                       `xml:"size"`
	Computer ComputerGroupComputerItem `xml:"computer"`
}

type ComputerGroupSite struct {
	ID   int    `json:"id,omitempty" xml:"id,omitempty"`
	Name string `json:"name,omitempty" xml:"name,omitempty"`
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

type ComputerGroupComputerItem struct {
	ID            int    `json:"id,omitempty" xml:"id,omitempty"`
	Name          string `json:"name,omitempty" xml:"name,omitempty"`
	SerialNumber  string `json:"serial_number,omitempty" xml:"serial_number,omitempty"`
	MacAddress    string `json:"mac_address,omitempty" xml:"mac_address,omitempty"`
	AltMacAddress string `json:"alt_mac_address,omitempty" xml:"alt_mac_address,omitempty"`
}

type DeviceGroupAndOr string

const (
	And DeviceGroupAndOr = "and"
	Or  DeviceGroupAndOr = "or"
)

// GetComputerGroups gets a list of all computer groups
func (c *Client) GetComputerGroups() (*ComputerGroupsListResponse, error) {
	endpoint := uriComputerGroups

	var computerGroups ComputerGroupsListResponse
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &computerGroups)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all Computer Groups: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &computerGroups, nil
}

// GetComputerGroupByID retrieves a computer group by its ID.
func (c *Client) GetComputerGroupByID(id int) (*ResponseComputerGroup, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriComputerGroups, id)

	var group ResponseComputerGroup
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &group)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Computer Group by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &group, nil
}

// GetComputerGroupByName retrieves a computer group by its name.
func (c *Client) GetComputerGroupByName(name string) (*ResponseComputerGroup, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriComputerGroups, name)

	var group ResponseComputerGroup
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &group)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Computer Group by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &group, nil
}

// CreateComputerGroup creates a new computer group.
func (c *Client) CreateComputerGroup(group *ResponseComputerGroup) (*ResponseComputerGroup, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriComputerGroups) // Using ID 0 for creation as per the pattern

	// Check if site is not provided and set default values
	if group.Site.ID == 0 && group.Site.Name == "" {
		group.Site = ComputerGroupSite{
			ID:   -1,
			Name: "None",
		}
	}

	// Wrap the group request with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"computer_group"`
		*ResponseComputerGroup
	}{
		ResponseComputerGroup: group,
	}

	var createdGroup ResponseComputerGroup
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &createdGroup)
	if err != nil {
		return nil, fmt.Errorf("failed to create Computer Group: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdGroup, nil
}

// UpdateComputerGroupByID updates an existing computer group by its ID.
func (c *Client) UpdateComputerGroupByID(id int, group *ResponseComputerGroup) (*ResponseComputerGroup, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriComputerGroups, id)

	// Check if site is not provided and set default values
	if group.Site.ID == 0 && group.Site.Name == "" {
		group.Site = ComputerGroupSite{
			ID:   -1,
			Name: "None",
		}
	}

	// Wrap the group request with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"computer_group"`
		*ResponseComputerGroup
	}{
		ResponseComputerGroup: group,
	}

	var updatedGroup ResponseComputerGroup
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedGroup)
	if err != nil {
		return nil, fmt.Errorf("failed to update Computer Group by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedGroup, nil
}

// UpdateComputerGroupByName updates a computer group by its name.
func (c *Client) UpdateComputerGroupByName(name string, group *ResponseComputerGroup) (*ResponseComputerGroup, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriComputerGroups, name)

	// Check if site is not provided and set default values
	if group.Site.ID == 0 && group.Site.Name == "" {
		group.Site = ComputerGroupSite{
			ID:   -1,
			Name: "None",
		}
	}

	// Wrap the group request with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"computer_group"`
		*ResponseComputerGroup
	}{
		ResponseComputerGroup: group,
	}

	var updatedGroup ResponseComputerGroup
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedGroup)
	if err != nil {
		return nil, fmt.Errorf("failed to update Computer Group by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedGroup, nil
}

// DeleteComputerGroupByID deletes a computer group by its ID.
func (c *Client) DeleteComputerGroupByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriComputerGroups, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete Computer Group by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteComputerGroupByName deletes a computer group by its name.
func (c *Client) DeleteComputerGroupByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriComputerGroups, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete Computer Group by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
