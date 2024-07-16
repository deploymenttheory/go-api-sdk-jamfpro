// classicapi_network_segments.go
// Jamf Pro Classic Api  - Network Segments
// api reference: https://developer.jamf.com/jamf-pro/reference/networksegments
// Jamf Pro Classic Api requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriNetworkSegments = "/JSSResource/networksegments"

// List

// ResponseNetworkSegmentList represents the response for a list of Network Segments.
type ResponseNetworkSegmentList struct {
	Size    int `xml:"size"`
	Results []struct {
		ID              int    `xml:"id"`
		Name            string `xml:"name"`
		StartingAddress string `xml:"starting_address"`
		EndingAddress   string `xml:"ending_address"`
	} `xml:"network_segment"`
}

type ResponseNetworkSegmentCreatedAndUpdated struct {
	ID int `json:"id,omitempty" xml:"id,omitempty"`
}

// Resource

// ResourceNetworkSegment represents the response structure for a Network Segment.
type ResourceNetworkSegment struct {
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

// CRUD

// GetNetworkSegments retrieves a list of network segments.
func (c *Client) GetNetworkSegments() (*ResponseNetworkSegmentList, error) {
	endpoint := uriNetworkSegments

	var segments ResponseNetworkSegmentList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &segments)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "network segments", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &segments, nil
}

// GetNetworkSegmentByID retrieves a specific network segment by its ID.
func (c *Client) GetNetworkSegmentByID(id int) (*ResourceNetworkSegment, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriNetworkSegments, id)

	var segment ResourceNetworkSegment
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &segment)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "network segment", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &segment, nil
}

// GetNetworkSegmentByName retrieves a specific network segment by its name.
func (c *Client) GetNetworkSegmentByName(name string) (*ResourceNetworkSegment, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriNetworkSegments, name)

	var segment ResourceNetworkSegment
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &segment)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "network segment", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &segment, nil
}

// CreateNetworkSegment creates a new network segment on the Jamf Pro server.
func (c *Client) CreateNetworkSegment(segment *ResourceNetworkSegment) (*ResponseNetworkSegmentCreatedAndUpdated, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriNetworkSegments)

	requestBody := struct {
		XMLName xml.Name `xml:"network_segment"`
		*ResourceNetworkSegment
	}{
		ResourceNetworkSegment: segment,
	}

	var responseSegment ResponseNetworkSegmentCreatedAndUpdated
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &responseSegment)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "network segment", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseSegment, nil
}

// UpdateNetworkSegmentByID updates a specific network segment by its ID.
func (c *Client) UpdateNetworkSegmentByID(id int, segment *ResourceNetworkSegment) (*ResponseNetworkSegmentCreatedAndUpdated, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriNetworkSegments, id)

	requestBody := struct {
		XMLName xml.Name `xml:"network_segment"`
		*ResourceNetworkSegment
	}{
		ResourceNetworkSegment: segment,
	}

	var responseSegment ResponseNetworkSegmentCreatedAndUpdated
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responseSegment)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "network segment", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseSegment, nil
}

// UpdateNetworkSegmentByName updates a specific network segment by its name.
func (c *Client) UpdateNetworkSegmentByName(name string, segment *ResourceNetworkSegment) (*ResponseNetworkSegmentCreatedAndUpdated, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriNetworkSegments, name)

	requestBody := struct {
		XMLName xml.Name `xml:"network_segment"`
		*ResourceNetworkSegment
	}{
		ResourceNetworkSegment: segment,
	}

	var responseSegment ResponseNetworkSegmentCreatedAndUpdated
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responseSegment)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "network segment", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseSegment, nil
}

// DeleteNetworkSegmentByID deletes a policy by its ID.
func (c *Client) DeleteNetworkSegmentByID(id string) error {
	endpoint := fmt.Sprintf("%s/id/%s", uriNetworkSegments, id)
	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "network segment", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteNetworkSegmentByName deletes a policy by its name.
func (c *Client) DeleteNetworkSegmentByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriNetworkSegments, name)
	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByName, "network segment", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
