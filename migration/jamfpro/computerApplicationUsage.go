// computerApplicationUsage.go
// Jamf Pro Classic Api
// Classic API requires the structs to support both XML and JSON.

package jamfpro

import (
	"fmt"
)

const uriAPIComputerApplicationUsage = "/JSSResource/computerapplicationusage"

type ResponseComputerApplicationUsage struct {
	Usage ComputerApplicationUsageDetail `json:"usage,omitempty" xml:"usage,omitempty"`
}

type ComputerApplicationUsageDetail struct {
	Date string                  `json:"date,omitempty" xml:"date,omitempty"`
	Apps []ComputerAppUsageEntry `json:"apps,omitempty" xml:"apps,omitempty"`
}

type ComputerAppUsageEntry struct {
	App ComputerAppUsageDetail `json:"app,omitempty" xml:"app,omitempty"`
}

type ComputerAppUsageDetail struct {
	Name       string `json:"name,omitempty" xml:"name,omitempty"`
	Version    string `json:"version,omitempty" xml:"version,omitempty"`
	Foreground int    `json:"foreground,omitempty" xml:"foreground,omitempty"` // Number of minutes application was in the foreground
	Open       int    `json:"open,omitempty" xml:"open,omitempty"`             // Number of minutes the application was open
}

// GetComputerApplicationUsageByComputerID retrieves the Computer Application Usage by its Computer ID.
func (c *Client) GetComputerApplicationUsageByComputerID(id string, startDate string, endDate string) (*ResponseComputerApplicationUsage, error) {
	url := fmt.Sprintf("%s/id/%s/%s_%s", uriAPIComputerApplicationUsage, id, startDate, endDate)

	var usage ResponseComputerApplicationUsage
	if err := c.DoRequest("GET", url, nil, nil, &usage); err != nil {
		return nil, fmt.Errorf("failed to get computer application usage by computer ID: %v", err)
	}

	return &usage, nil
}

// GetComputerApplicationUsageByComputerName retrieves the Computer Application Usage by its Computer Name.
func (c *Client) GetComputerApplicationUsageByComputerName(name string, startDate string, endDate string) (*ResponseComputerApplicationUsage, error) {
	url := fmt.Sprintf("%s/name/%s/%s_%s", uriAPIComputerApplicationUsage, name, startDate, endDate)

	var usage ResponseComputerApplicationUsage
	if err := c.DoRequest("GET", url, nil, nil, &usage); err != nil {
		return nil, fmt.Errorf("failed to get computer application usage by computer name: %v", err)
	}

	return &usage, nil
}

// GetComputerApplicationUsageByComputerSerialNumber retrieves the Computer Application Usage by its Computer Serial Number.
func (c *Client) GetComputerApplicationUsageByComputerSerialNumber(serialNumber string, startDate string, endDate string) (*ResponseComputerApplicationUsage, error) {
	url := fmt.Sprintf("%s/serialnumber/%s/%s_%s", uriAPIComputerApplicationUsage, serialNumber, startDate, endDate)

	var usage ResponseComputerApplicationUsage
	if err := c.DoRequest("GET", url, nil, nil, &usage); err != nil {
		return nil, fmt.Errorf("failed to get computer application usage by computer serial number: %v", err)
	}

	return &usage, nil
}

// GetComputerApplicationUsageByComputerMacAddress retrieves the Computer Application Usage by its Computer MAC address.
func (c *Client) GetComputerApplicationUsageByComputerMacAddress(macAddress string, startDate string, endDate string) (*ResponseComputerApplicationUsage, error) {
	url := fmt.Sprintf("%s/macaddress/%s/%s_%s", uriAPIComputerApplicationUsage, macAddress, startDate, endDate)

	var usage ResponseComputerApplicationUsage
	if err := c.DoRequest("GET", url, nil, nil, &usage); err != nil {
		return nil, fmt.Errorf("failed to get computer application usage by computer MAC address: %v", err)
	}

	return &usage, nil
}
