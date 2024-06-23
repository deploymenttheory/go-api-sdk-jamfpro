// classicapi_restricted_software.go
// Jamf Pro Classic Api - Restricted Software
// API reference: https://developer.jamf.com/jamf-pro/reference/restrictedsoftware
// Jamf Pro Classic API requires the structs to support an XML data structure.

/*
Shared Resources in this Endpoint:
- SharedResourceSite
*/

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriRestrictedSoftware = "/JSSResource/restrictedsoftware"

// Responses

type ResponseRestrictedSoftwareCreateAndUpdate struct {
	ID int `xml:"id"`
}

// Structs for Restricted Software List
type ResponseRestrictedSoftwaresList struct {
	Size               int                          `xml:"size"`
	RestrictedSoftware []RestrictedSoftwareListItem `xml:"restricted_software_title"`
}

type RestrictedSoftwareListItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Resource

// Structs for individual Restricted Software
type ResourceRestrictedSoftware struct {
	General RestrictedSoftwareSubsetGeneral `xml:"general"`
	Scope   RestrictedSoftwareSubsetScope   `xml:"scope"`
}

// Subsets & Containers

// General

type RestrictedSoftwareSubsetGeneral struct {
	ID                    int                 `xml:"id"`
	Name                  string              `xml:"name"`
	ProcessName           string              `xml:"process_name"`
	MatchExactProcessName bool                `xml:"match_exact_process_name"`
	SendNotification      bool                `xml:"send_notification"`
	KillProcess           bool                `xml:"kill_process"`
	DeleteExecutable      bool                `xml:"delete_executable"`
	DisplayMessage        string              `xml:"display_message"`
	Site                  *SharedResourceSite `xml:"site"`
}

// Scope

type RestrictedSoftwareSubsetScope struct {
	AllComputers   bool                                    `xml:"all_computers"`
	Computers      []RestrictedSoftwareSubsetScopeEntity   `xml:"computers>computer"`
	ComputerGroups []RestrictedSoftwareSubsetScopeEntity   `xml:"computer_groups>computer_group"`
	Buildings      []RestrictedSoftwareSubsetScopeEntity   `xml:"buildings>building"`
	Departments    []RestrictedSoftwareSubsetScopeEntity   `xml:"departments>department"`
	Exclusions     RestrictedSoftwareSubsetScopeExclusions `xml:"exclusions"`
}

type RestrictedSoftwareSubsetScopeExclusions struct {
	Computers      []RestrictedSoftwareSubsetScopeEntity `xml:"computers>computer"`
	ComputerGroups []RestrictedSoftwareSubsetScopeEntity `xml:"computer_groups>computer_group"`
	Buildings      []RestrictedSoftwareSubsetScopeEntity `xml:"buildings>building"`
	Departments    []RestrictedSoftwareSubsetScopeEntity `xml:"departments>department"`
	Users          []RestrictedSoftwareSubsetScopeEntity `xml:"users>user"`
}

// Shared

type RestrictedSoftwareSubsetScopeEntity struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// CRUD

// GetRestrictedSoftwares retrieves a list of all restricted software.
func (c *Client) GetRestrictedSoftwares() (*ResponseRestrictedSoftwaresList, error) {
	endpoint := uriRestrictedSoftware

	var restrictedSoftwaresList ResponseRestrictedSoftwaresList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &restrictedSoftwaresList)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "restricted softwares", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &restrictedSoftwaresList, nil
}

// GetRestrictedSoftwareByID fetches the details of a specific restricted software entry by its ID.
func (c *Client) GetRestrictedSoftwareByID(id int) (*ResourceRestrictedSoftware, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriRestrictedSoftware, id)

	var restrictedSoftware ResourceRestrictedSoftware
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &restrictedSoftware)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "restricted software", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &restrictedSoftware, nil
}

// GetRestrictedSoftwareByName retrieves the details of a specific restricted software entry by its name.
func (c *Client) GetRestrictedSoftwareByName(name string) (*ResourceRestrictedSoftware, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriRestrictedSoftware, name)

	var restrictedSoftware ResourceRestrictedSoftware
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &restrictedSoftware)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "restricted software", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &restrictedSoftware, nil
}

// CreateRestrictedSoftware creates a new restricted software entry in Jamf Pro.
func (c *Client) CreateRestrictedSoftware(restrictedSoftware *ResourceRestrictedSoftware) (*ResponseRestrictedSoftwareCreateAndUpdate, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriRestrictedSoftware, restrictedSoftware.General.ID)

	requestBody := struct {
		XMLName xml.Name `xml:"restricted_software"`
		*ResourceRestrictedSoftware
	}{
		ResourceRestrictedSoftware: restrictedSoftware,
	}

	var responseRestrictedSoftware ResponseRestrictedSoftwareCreateAndUpdate
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &responseRestrictedSoftware)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "restricted software", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseRestrictedSoftware, nil
}

// UpdateRestrictedSoftwareByID updates an existing restricted software entry by its ID.
func (c *Client) UpdateRestrictedSoftwareByID(id int, restrictedSoftware *ResourceRestrictedSoftware) (*ResponseRestrictedSoftwareCreateAndUpdate, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriRestrictedSoftware, id)

	requestBody := struct {
		XMLName xml.Name `xml:"restricted_software"`
		*ResourceRestrictedSoftware
	}{
		ResourceRestrictedSoftware: restrictedSoftware,
	}

	var responseRestrictedSoftware ResponseRestrictedSoftwareCreateAndUpdate
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responseRestrictedSoftware)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "restricted software", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseRestrictedSoftware, nil
}

// UpdateRestrictedSoftwareByName updates an existing restricted software entry by its name.
func (c *Client) UpdateRestrictedSoftwareByName(name string, restrictedSoftware *ResourceRestrictedSoftware) (*ResponseRestrictedSoftwareCreateAndUpdate, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriRestrictedSoftware, name)

	requestBody := struct {
		XMLName xml.Name `xml:"restricted_software"`
		*ResourceRestrictedSoftware
	}{
		ResourceRestrictedSoftware: restrictedSoftware,
	}

	var responseRestrictedSoftware ResponseRestrictedSoftwareCreateAndUpdate
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &responseRestrictedSoftware)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "restricted software", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseRestrictedSoftware, nil
}

// DeleteRestrictedSoftwareByID deletes a restricted software entry by its ID.
func (c *Client) DeleteRestrictedSoftwareByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriRestrictedSoftware, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "restricted software", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteRestrictedSoftwareByName deletes a restricted software entry by its name.
func (c *Client) DeleteRestrictedSoftwareByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriRestrictedSoftware, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByName, "restricted software", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
