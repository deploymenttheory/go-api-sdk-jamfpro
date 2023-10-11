// ibeacons.go
// Jamf Pro Classic Api
// Classic API requires the structs to support both XML and JSON.

package jamfpro

import (
	"fmt"
)

const uriIbeacons = "/JSSResource/ibeacons"

type ResponseIbeacon struct {
	ID    int    `json:"id" xml:"id"`
	Name  string `json:"name" xml:"name"`
	UUID  string `json:"uuid" xml:"uuid"`
	Major int    `json:"major" xml:"major"`
	Minor int    `json:"minor" xml:"minor"`
}

type ResponseIbeaconList struct {
	Ibeacons []IbeaconListItem `json:"ibeacon" xml:"ibeacon"`
}

type IbeaconListItem struct {
	ID    int    `json:"id" xml:"id"`
	Name  string `json:"name" xml:"name"`
	UUID  string `json:"uuid" xml:"uuid"`
	Major int    `json:"major" xml:"major"`
	Minor int    `json:"minor" xml:"minor"`
}

type IBeaconScope struct {
	Id   int    `xml:"id"`
	Name string `xml:"name"`
}

// GetIbeaconByID retrieves the iBeacon configuration by its ID.
func (c *Client) GetIbeaconByID(id int) (*ResponseIbeacon, error) {
	url := fmt.Sprintf("%s/id/%d", uriIbeacons, id)

	var ibeacon ResponseIbeacon
	if err := c.DoRequest("GET", url, nil, nil, &ibeacon); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &ibeacon, nil
}

// GetIbeaconByName retrieves the iBeacon configuration by its Name.
func (c *Client) GetIbeaconByName(name string) (*ResponseIbeacon, error) {
	url := fmt.Sprintf("%s/name/%s", uriIbeacons, name)

	var ibeacon ResponseIbeacon
	if err := c.DoRequest("GET", url, nil, nil, &ibeacon); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &ibeacon, nil
}

// GetIbeacons retrieves all iBeacon configurations.
func (c *Client) GetIbeacons() (*ResponseIbeaconList, error) {
	url := uriIbeacons

	var ibeacons ResponseIbeaconList
	if err := c.DoRequest("GET", url, nil, nil, &ibeacons); err != nil {
		return nil, fmt.Errorf("failed to fetch all iBeacons: %v", err)
	}

	return &ibeacons, nil
}

// CreateIbeacon creates a new iBeacon configuration.
func (c *Client) CreateIbeacon(ibeacon *ResponseIbeacon) (*ResponseIbeacon, error) {
	url := fmt.Sprintf("%s/id/0", uriIbeacons)

	reqBody := &struct {
		XMLName struct{} `xml:"ibeacon"`
		*ResponseIbeacon
	}{
		ResponseIbeacon: ibeacon,
	}

	var responseIbeacon ResponseIbeacon
	if err := c.DoRequest("POST", url, reqBody, nil, &responseIbeacon); err != nil {
		return nil, fmt.Errorf("failed to create iBeacon: %v", err)
	}

	return &responseIbeacon, nil
}

// UpdateIbeaconById updates an existing iBeacon configuration by its ID.
func (c *Client) UpdateIbeaconById(id int, ibeacon *ResponseIbeacon) (*ResponseIbeacon, error) {
	url := fmt.Sprintf("%s/id/%d", uriIbeacons, id)

	reqBody := &struct {
		XMLName struct{} `xml:"ibeacon"`
		*ResponseIbeacon
	}{
		ResponseIbeacon: ibeacon,
	}

	var responseIbeacon ResponseIbeacon
	if err := c.DoRequest("PUT", url, reqBody, nil, &responseIbeacon); err != nil {
		return nil, fmt.Errorf("failed to update iBeacon by ID: %v", err)
	}

	return &responseIbeacon, nil
}

// UpdateIbeaconByName updates an existing iBeacon configuration by its name.
func (c *Client) UpdateIbeaconByName(name string, ibeacon *ResponseIbeacon) (*ResponseIbeacon, error) {
	url := fmt.Sprintf("%s/name/%s", uriIbeacons, name)

	reqBody := &struct {
		XMLName struct{} `xml:"ibeacon"`
		*ResponseIbeacon
	}{
		ResponseIbeacon: ibeacon,
	}

	var responseIbeacon ResponseIbeacon
	if err := c.DoRequest("PUT", url, reqBody, nil, &responseIbeacon); err != nil {
		return nil, fmt.Errorf("failed to update iBeacon by name: %v", err)
	}

	return &responseIbeacon, nil
}

// DeleteIbeaconById deletes an existing iBeacon configuration by its ID.
func (c *Client) DeleteIbeaconById(id int) error {
	url := fmt.Sprintf("%s/id/%d", uriIbeacons, id)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete iBeacon by ID: %v", err)
	}

	return nil
}

// DeleteIbeaconByName deletes an existing iBeacon configuration by its name.
func (c *Client) DeleteIbeaconByName(name string) error {
	url := fmt.Sprintf("%s/name/%s", uriIbeacons, name)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete iBeacon by name: %v", err)
	}

	return nil
}
