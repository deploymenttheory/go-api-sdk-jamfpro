// classicapi_ibeacons.go
// Jamf Pro Classic Api - iBeacons
// api reference: https://developer.jamf.com/jamf-pro/reference/ibeacons
// Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriIbeacons = "/JSSResource/ibeacons"

// ResponseIBeaconsList represents the response structure for a list of iBeacons.
type ResponseIBeaconsList struct {
	Size     int                `xml:"size"`
	IBeacons []ResourceIBeacons `xml:"ibeacon"`
}

// ResponseIBeacons represents the structure of an individual iBeacon.
type ResourceIBeacons struct {
	ID    int    `xml:"id"`
	Name  string `xml:"name"`
	UUID  string `xml:"uuid"`
	Major int    `xml:"major,omitempty"`
	Minor int    `xml:"minor,omitempty"`
}

// GetIBeacons retrieves a list of all iBeacons registered in Jamf Pro.
// It returns a serialized list of iBeacon details including ID, name, UUID, major, and minor values.
func (c *Client) GetIBeacons() (*ResponseIBeaconsList, error) {
	endpoint := uriIbeacons

	var iBeacons ResponseIBeaconsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &iBeacons)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch iBeacons: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &iBeacons, nil
}

// GetIBeaconByID fetches the details of a specific iBeacon by its ID.
// It returns the iBeacon's ID, name, UUID, major, and minor values.
func (c *Client) GetIBeaconByID(id int) (*ResourceIBeacons, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriIbeacons, id)
	var beacon ResourceIBeacons
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &beacon)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch iBeacon by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &beacon, nil
}

// GetIBeaconByName fetches the details of a specific iBeacon by its name.
// It returns the iBeacon's ID, name, UUID, major, and minor values.
func (c *Client) GetIBeaconByName(name string) (*ResourceIBeacons, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriIbeacons, name)
	var beacon ResourceIBeacons
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &beacon)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch iBeacon by Name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &beacon, nil
}

// CreateIBeacon creates a new iBeacon in Jamf Pro.
func (c *Client) CreateIBeacon(beacon *ResourceIBeacons) (*ResourceIBeacons, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriIbeacons) // '0' typically used for creation in APIs

	// The requestBody struct should mirror the ResponseIBeacons struct
	requestBody := struct {
		XMLName xml.Name `xml:"ibeacon"`
		*ResourceIBeacons
	}{
		ResourceIBeacons: beacon,
	}

	var response ResourceIBeacons
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to create iBeacon: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateIBeaconByID updates an existing iBeacon by its ID in Jamf Pro.
func (c *Client) UpdateIBeaconByID(id int, beacon *ResourceIBeacons) (*ResourceIBeacons, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriIbeacons, id)

	requestBody := struct {
		XMLName xml.Name `xml:"ibeacon"`
		*ResourceIBeacons
	}{
		ResourceIBeacons: beacon,
	}

	var response ResourceIBeacons
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to update iBeacon by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateIBeaconByName updates an existing iBeacon by its name in Jamf Pro.
func (c *Client) UpdateIBeaconByName(name string, beacon *ResourceIBeacons) (*ResourceIBeacons, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriIbeacons, name)

	requestBody := struct {
		XMLName xml.Name `xml:"ibeacon"`
		*ResourceIBeacons
	}{
		ResourceIBeacons: beacon,
	}

	var response ResourceIBeacons
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to update iBeacon by Name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// DeleteIBeaconByID deletes an iBeacon by its ID in Jamf Pro.
func (c *Client) DeleteIBeaconByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriIbeacons, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete iBeacon by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteIBeaconByName deletes an iBeacon by its name in Jamf Pro.
func (c *Client) DeleteIBeaconByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriIbeacons, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete iBeacon by Name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
