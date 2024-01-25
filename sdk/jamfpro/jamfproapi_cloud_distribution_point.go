// jamfproapi_cloud_distribution_point.go
// Jamf Pro Api - Cloud Distribution Point
// api reference:
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
)

// Responses

const uriCloudDistributionPoint = "/v1/cloud-distribution-point"

type ResponseCloudDistributionPointList struct {
	TotalCount int                              `json:"totalCount"`
	Results    []ResourceCloudDistributionPoint `json:"results"`
}

type ResponseCloudDistributionPointCreate struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// Resource

type ResourceCloudDistributionPoint struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetCloudDistributionPointUploadCapability retrieves the default server configuration for the Cloud Identity Provider.
func (c *Client) GetCloudDistributionPointUploadCapability() (*ResourceCloudDistributionPoint, error) {
	endpoint := uriCloudDistributionPoint + "/upload-capability"

	var cloudDistributionPointUploadCapability ResourceCloudDistributionPoint
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &cloudDistributionPointUploadCapability)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "Cloud Distribution Point", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &cloudDistributionPointUploadCapability, nil
}
