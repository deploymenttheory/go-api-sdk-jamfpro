// networkSegments.go
// Jamf Pro Classic Api
// Classic API requires the structs to support both XML and JSON.

package jamfpro

import (
	"fmt"
)

const uriNetworkSegments = "/JSSResource/networksegments"

type ResponseNetworkSegment struct {
	ID                  int    `json:"id" xml:"id"`
	Name                string `json:"name" xml:"name"`
	StartingAddress     string `json:"starting_address" xml:"starting_address"`
	EndingAddress       string `json:"ending_address" xml:"ending_address"`
	DistributionServer  string `json:"distribution_server,omitempty" xml:"distribution_server,omitempty"`
	DistributionPoint   string `json:"distribution_point,omitempty" xml:"distribution_point,omitempty"`
	URL                 string `json:"url,omitempty" xml:"url,omitempty"`
	SWUServer           string `json:"swu_server,omitempty" xml:"swu_server,omitempty"`
	Building            string `json:"building,omitempty" xml:"building,omitempty"`
	Department          string `json:"department,omitempty" xml:"department,omitempty"`
	OverrideBuildings   bool   `json:"override_buildings" xml:"override_buildings"`
	OverrideDepartments bool   `json:"override_departments" xml:"override_departments"`
}

type NetworkSegmentList struct {
	Size    int                      `json:"size" xml:"size"`
	Results []ResponseNetworkSegment `json:"network_segment" xml:"network_segment"`
}

type NetworkSegmentScope struct {
	ID   int    `xml:"id"`
	UID  string `xml:"uid,omitempty"`
	Name int    `xml:"name"`
}

// GetNetworkSegmentByID retrieves the Network Segment by its ID
func (c *Client) GetNetworkSegmentByID(id int) (*ResponseNetworkSegment, error) {
	url := fmt.Sprintf("%s/id/%d", uriNetworkSegments, id)

	var segment ResponseNetworkSegment
	if err := c.DoRequest("GET", url, nil, nil, &segment); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &segment, nil
}

// GetNetworkSegmentByName retrieves the Network Segment by its Name
func (c *Client) GetNetworkSegmentByName(name string) (*ResponseNetworkSegment, error) {
	url := fmt.Sprintf("%s/name/%s", uriNetworkSegments, name)

	var segment ResponseNetworkSegment
	if err := c.DoRequest("GET", url, nil, nil, &segment); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &segment, nil
}

// GetNetworkSegments retrieves all Network Segments
func (c *Client) GetNetworkSegments() (*NetworkSegmentList, error) {
	url := uriNetworkSegments

	var segments NetworkSegmentList
	if err := c.DoRequest("GET", url, nil, nil, &segments); err != nil {
		return nil, fmt.Errorf("failed to fetch all Network Segments: %v", err)
	}

	return &segments, nil
}

// CreateNetworkSegment creates a new Network Segment
func (c *Client) CreateNetworkSegment(segment *ResponseNetworkSegment) (*ResponseNetworkSegment, error) {
	url := fmt.Sprintf("%s/id/0", uriNetworkSegments)

	// Construct a custom request body structure for proper XML serialization
	reqBody := &struct {
		XMLName struct{} `xml:"network_segment"`
		*ResponseNetworkSegment
	}{
		ResponseNetworkSegment: segment,
	}

	// Execute the request
	var responseSegment ResponseNetworkSegment
	if err := c.DoRequest("POST", url, reqBody, nil, &responseSegment); err != nil {
		return nil, fmt.Errorf("failed to create Network Segment: %v", err)
	}

	return &responseSegment, nil
}

// UpdateNetworkSegmentById updates an existing Network Segment by its ID
func (c *Client) UpdateNetworkSegmentById(id int, segment *ResponseNetworkSegment) (*ResponseNetworkSegment, error) {
	url := fmt.Sprintf("%s/id/%d", uriNetworkSegments, id)

	// Construct a custom request body structure for proper XML serialization
	reqBody := &struct {
		XMLName struct{} `xml:"network_segment"`
		*ResponseNetworkSegment
	}{
		ResponseNetworkSegment: segment,
	}
	// Execute the request
	var responseSegment ResponseNetworkSegment
	if err := c.DoRequest("PUT", url, reqBody, nil, &responseSegment); err != nil {
		return nil, fmt.Errorf("failed to update Network Segment: %v", err)
	}

	return &responseSegment, nil
}

// UpdateNetworkSegmentByName updates an existing Network Segment by its name
func (c *Client) UpdateNetworkSegmentByName(name string, segment *ResponseNetworkSegment) (*ResponseNetworkSegment, error) {
	url := fmt.Sprintf("%s/name/%s", uriNetworkSegments, name)

	// Construct a custom request body structure for proper XML serialization
	reqBody := &struct {
		XMLName struct{} `xml:"network_segment"`
		*ResponseNetworkSegment
	}{
		ResponseNetworkSegment: segment,
	}

	// Execute the request
	var responseSegment ResponseNetworkSegment
	if err := c.DoRequest("PUT", url, reqBody, nil, &responseSegment); err != nil {
		return nil, fmt.Errorf("failed to update Network Segment by name: %v", err)
	}

	return &responseSegment, nil
}

// DeleteNetworkSegmentById deletes an existing Network Segment by its ID
func (c *Client) DeleteNetworkSegmentById(id int) error {
	url := fmt.Sprintf("%s/id/%d", uriNetworkSegments, id)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete Network Segment: %v", err)
	}

	return nil
}

// DeleteNetworkSegmentByName deletes an existing Network Segment by its name
func (c *Client) DeleteNetworkSegmentByName(name string) error {
	url := fmt.Sprintf("%s/name/%s", uriNetworkSegments, name)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete Network Segment by name: %v", err)
	}

	return nil
}
