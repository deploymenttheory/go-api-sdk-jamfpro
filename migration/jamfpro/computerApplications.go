// computerApplications.go
// Jamf Pro Classic Api
// Classic API requires the structs to support both XML and JSON.

package jamfpro

import (
	"fmt"
)

const uriAPIComputerApplications = "/JSSResource/computerapplications"

type ResponseComputerApplication struct {
	Versions        ComputerApplicationDataSubsetVersions        `json:"versions,omitempty" xml:"versions,omitempty"`
	UniqueComputers ComputerApplicationDataSubsetUniqueComputers `json:"unique_computers,omitempty" xml:"unique_computers,omitempty"`
}

type ComputerApplicationDataSubsetVersions struct {
	Version []ComputerApplicationDataSubsetVersionDetail `json:"version,omitempty" xml:"version,omitempty"`
}

type ComputerApplicationDataSubsetVersionDetail struct {
	Number    string                                      `json:"number,omitempty" xml:"number,omitempty"`
	Computers []ComputerApplicationDataSubsetComputerWrap `json:"computers,omitempty" xml:"computers,omitempty"`
}

type ComputerApplicationDataSubsetComputerWrap struct {
	Computer ComputerApplicationDataSubsetComputerDetail `json:"computer,omitempty" xml:"computer,omitempty"`
}

type ComputerApplicationDataSubsetUniqueComputers struct {
	Computer []ComputerApplicationDataSubsetComputerDetail `json:"computer,omitempty" xml:"computer,omitempty"`
}

type ComputerApplicationDataSubsetComputerDetail struct {
	ID           int    `json:"id,omitempty" xml:"id,omitempty"`
	Name         string `json:"name,omitempty" xml:"name,omitempty"`
	UDID         string `json:"udid,omitempty" xml:"udid,omitempty"`
	SerialNumber string `json:"serial_number,omitempty" xml:"serial_number,omitempty"`
	MacAddress   string `json:"mac_address,omitempty" xml:"mac_address,omitempty"`
}

// GetComputerApplicationByName retrieves the Computer Application by its name.
func (c *Client) GetComputerApplicationByName(appName string) (*ResponseComputerApplication, error) {
	url := fmt.Sprintf("%s/application/%s", uriAPIComputerApplications, appName)

	var app ResponseComputerApplication
	if err := c.DoRequest("GET", url, nil, nil, &app); err != nil {
		return nil, fmt.Errorf("failed to get computer application by name: %v", err)
	}

	return &app, nil
}

// GetComputerApplicationsByNameWithDisplayFields retrieves the Computer Applications by name with additional display fields.
func (c *Client) GetComputerApplicationsByNameWithAdditionalDisplayFields(appName string, inventory string) (*ResponseComputerApplication, error) {
	url := fmt.Sprintf("%s/application/%s/inventory/%s", uriAPIComputerApplications, appName, inventory)

	var app ResponseComputerApplication
	if err := c.DoRequest("GET", url, nil, nil, &app); err != nil {
		return nil, fmt.Errorf("failed to get computer applications by name with additional display fields: %v", err)
	}

	return &app, nil
}

// GetComputerApplicationByNameAndVersion retrieves the Computer Application by its name and version.
func (c *Client) GetComputerApplicationByNameAndVersion(appName string, version string) (*ResponseComputerApplication, error) {
	url := fmt.Sprintf("%s/application/%s/version/%s", uriAPIComputerApplications, appName, version)

	var app ResponseComputerApplication
	if err := c.DoRequest("GET", url, nil, nil, &app); err != nil {
		return nil, fmt.Errorf("failed to get computer application by name and version: %v", err)
	}

	return &app, nil
}

// GetComputerApplicationByNameAndVersionWithDisplayFields retrieves the Computer Application by its name and version With additional DisplayFields.
func (c *Client) GetComputerApplicationByNameAndVersionWithAdditionalDisplayFields(appName string, version string, inventory string) (*ResponseComputerApplication, error) {
	url := fmt.Sprintf("%s/application/%s/version/%s/inventory/%s", uriAPIComputerApplications, appName, version, inventory)

	var app ResponseComputerApplication
	if err := c.DoRequest("GET", url, nil, nil, &app); err != nil {
		return nil, fmt.Errorf("failed to get computer application by name, version, and with additional display fields: %v", err)
	}

	return &app, nil
}
