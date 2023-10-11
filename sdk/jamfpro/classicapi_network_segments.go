// classicapi_network_segments.go
// Jamf Pro Api - Network Segments
// api reference: https://developer.jamf.com/jamf-pro/reference/networksegments
// Jamf Pro API requires the structs to support an XML data structure.

package jamfpro

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

// Functions TODO
