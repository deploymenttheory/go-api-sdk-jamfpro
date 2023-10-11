// restrictedSoftware.go
// Jamf Pro Classic Api
// Classic API requires the structs to support both XML and JSON.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriAPIRestrictedSoftware = "/JSSResource/restrictedsoftware"

type ResponseRestrictedSoftware struct {
	General struct {
		ID                    int    `json:"id,omitempty" xml:"id,omitempty"`
		Name                  string `json:"name" xml:"name"`
		ProcessName           string `json:"process_name" xml:"process_name"`
		MatchExactProcessName bool   `json:"match_exact_process_name" xml:"match_exact_process_name"`
		SendNotification      bool   `json:"send_notification" xml:"send_notification"`
		KillProcess           bool   `json:"kill_process" xml:"kill_process"`
		DeleteExecutable      bool   `json:"delete_executable" xml:"delete_executable"`
		DisplayMessage        string `json:"display_message" xml:"display_message"`
		Site                  struct {
			ID   int    `json:"id" xml:"id"`
			Name string `json:"name" xml:"name"`
		} `json:"site" xml:"site"`
	} `json:"general" xml:"general"`
	Scope struct {
		AllComputers bool `json:"all_computers" xml:"all_computers"`
		Computers    []struct {
			Computer struct {
				ID   int    `json:"id" xml:"id"`
				Name string `json:"name" xml:"name"`
			} `json:"computer" xml:"computer"`
		} `json:"computers" xml:"computers"`
		ComputerGroups []struct {
			ComputerGroup struct {
				ID   int    `json:"id" xml:"id"`
				Name string `json:"name" xml:"name"`
			} `json:"computer_group" xml:"computer_group"`
		} `json:"computer_groups" xml:"computer_groups"`
		Buildings []struct {
			Building struct {
				ID   int    `json:"id" xml:"id"`
				Name string `json:"name" xml:"name"`
			} `json:"building" xml:"building"`
		} `json:"buildings" xml:"buildings"`
		Departments []struct {
			Department struct {
				ID   int    `json:"id" xml:"id"`
				Name string `json:"name" xml:"name"`
			} `json:"department" xml:"department"`
		} `json:"departments" xml:"departments"`
		Exclusions struct {
			Computers []struct {
				Computer struct {
					ID   int    `json:"id" xml:"id"`
					Name string `json:"name" xml:"name"`
				} `json:"computer" xml:"computer"`
			} `json:"computers" xml:"computers"`
			ComputerGroups []struct {
				ComputerGroup struct {
					ID   int    `json:"id" xml:"id"`
					Name string `json:"name" xml:"name"`
				} `json:"computer_group" xml:"computer_group"`
			} `json:"computer_groups" xml:"computer_groups"`
			Buildings []struct {
				Building struct {
					ID   int    `json:"id" xml:"id"`
					Name string `json:"name" xml:"name"`
				} `json:"building" xml:"building"`
			} `json:"buildings" xml:"buildings"`
			Departments []struct {
				Department struct {
					ID   int    `json:"id" xml:"id"`
					Name string `json:"name" xml:"name"`
				} `json:"department" xml:"department"`
			} `json:"departments" xml:"departments"`
			Users []struct {
				User struct {
					ID   int    `json:"id" xml:"id"`
					Name string `json:"name" xml:"name"`
				} `json:"user" xml:"user"`
			} `json:"users" xml:"users"`
		} `json:"exclusions" xml:"exclusions"`
	} `json:"scope" xml:"scope"`
}

type ResponseRestrictedSoftwareList struct {
	Size                    int                       `json:"size" xml:"size"`
	RestrictedSoftwareTitle []RestrictedSoftwareTitle `json:"restricted_software_title" xml:"restricted_software_title"`
}

type RestrictedSoftwareTitle struct {
	ID   int    `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
}

// GetRestrictedSoftwareByID gets a restricted software by its id
func (c *Client) GetRestrictedSoftwareByID(id int) (*ResponseRestrictedSoftware, error) {
	url := fmt.Sprintf("%s/id/%d", uriAPIRestrictedSoftware, id)

	var software ResponseRestrictedSoftware
	if err := c.DoRequest("GET", url, nil, nil, &software); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &software, nil
}

// GetRestrictedSoftwareByName gets a restricted software by its name
func (c *Client) GetRestrictedSoftwareByName(name string) (*ResponseRestrictedSoftware, error) {
	url := fmt.Sprintf("%s/name/%s", uriAPIRestrictedSoftware, name)

	var software ResponseRestrictedSoftware
	if err := c.DoRequest("GET", url, nil, nil, &software); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &software, nil
}

// GetRestrictedSoftwares gets a list of all restricted softwares
func (c *Client) GetRestrictedSoftwares() ([]ResponseRestrictedSoftwareList, error) {
	url := uriAPIRestrictedSoftware

	var softwareList []ResponseRestrictedSoftwareList
	if err := c.DoRequest("GET", url, nil, nil, &softwareList); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return softwareList, nil
}

// CreateRestrictedSoftware creates a new Jamf Pro Restricted Software.
func (c *Client) CreateRestrictedSoftware(software *ResponseRestrictedSoftware) (*ResponseRestrictedSoftware, error) {
	url := fmt.Sprintf("%s/id/0", uriAPIRestrictedSoftware) // ID 0 is typically used for creation in many APIs

	// Check if Site ID is not defined or set to -1, then set Site Name to "None"
	if software.General.Site.ID == -1 || software.General.Site.ID == 0 {
		software.General.Site.Name = "None"
	}

	reqBody := &struct {
		XMLName xml.Name `xml:"restricted_software"`
		*ResponseRestrictedSoftware
	}{
		ResponseRestrictedSoftware: software,
	}

	var responseSoftware ResponseRestrictedSoftware
	if err := c.DoRequest("POST", url, reqBody, nil, &responseSoftware); err != nil {
		return nil, fmt.Errorf("failed to create restricted software: %v", err)
	}

	return &responseSoftware, nil
}

// UpdateRestrictedSoftwareByName updates an existing Jamf Pro Restricted Software by Name.
func (c *Client) UpdateRestrictedSoftwareByName(name string, software *ResponseRestrictedSoftware) (*ResponseRestrictedSoftware, error) {
	url := fmt.Sprintf("%s/name/%s", uriAPIRestrictedSoftware, name)

	// Check if Site ID is not defined or set to -1, then set Site Name to "None"
	if software.General.Site.ID == -1 || software.General.Site.ID == 0 {
		software.General.Site.Name = "None"
	}

	reqBody := &struct {
		XMLName xml.Name `xml:"restricted_software"`
		*ResponseRestrictedSoftware
	}{
		ResponseRestrictedSoftware: software,
	}

	var responseSoftware ResponseRestrictedSoftware
	if err := c.DoRequest("PUT", url, reqBody, nil, &responseSoftware); err != nil {
		return nil, fmt.Errorf("failed to update restricted software by Name: %v", err)
	}

	return &responseSoftware, nil
}

// UpdateRestrictedSoftwareByID updates an existing Jamf Pro Restricted Software by ID.
func (c *Client) UpdateRestrictedSoftwareByID(id int, software *ResponseRestrictedSoftware) (*ResponseRestrictedSoftware, error) {
	url := fmt.Sprintf("%s/id/%d", uriAPIRestrictedSoftware, id)

	// Check if Site ID is not defined or set to -1, then set Site Name to "None"
	if software.General.Site.ID == -1 || software.General.Site.ID == 0 {
		software.General.Site.Name = "None"
	}

	reqBody := &struct {
		XMLName xml.Name `xml:"restricted_software"`
		*ResponseRestrictedSoftware
	}{
		ResponseRestrictedSoftware: software,
	}

	var responseSoftware ResponseRestrictedSoftware
	if err := c.DoRequest("PUT", url, reqBody, nil, &responseSoftware); err != nil {
		return nil, fmt.Errorf("failed to update restricted software by ID: %v", err)
	}

	return &responseSoftware, nil
}

// DeleteRestrictedSoftwareByID deletes an existing Jamf Pro Restricted Software by ID.
func (c *Client) DeleteRestrictedSoftwareByID(id int) error {
	url := fmt.Sprintf("%s/id/%d", uriAPIRestrictedSoftware, id)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete restricted software by ID: %v", err)
	}

	return nil
}

// DeleteRestrictedSoftwareByName deletes an existing Jamf Pro Restricted Software by Name.
func (c *Client) DeleteRestrictedSoftwareByName(name string) error {
	url := fmt.Sprintf("%s/name/%s", uriAPIRestrictedSoftware, name)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete restricted software by Name: %v", err)
	}

	return nil
}
