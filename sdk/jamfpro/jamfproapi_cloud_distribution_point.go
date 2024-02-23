// jamfproapi_cloud_distribution_point.go
// Jamf Pro Api - Cloud Distribution Point
// api reference: undocumented
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
)

// Responses

const uriCloudDistributionPoint = "/api/v1/cloud-distribution-point"

type ResponseCloudDistributionPointCreate struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// Resource

type ResourceCloudDistributionPointUploadCapability struct {
	ID   bool `json:"principalDistributionTechnology"`
	Name bool `json:"directUploadCapable"`
}

// GetCloudDistributionPointUploadCapability retrieves the default server configuration for the Cloud Identity Provider.
func (c *Client) GetCloudDistributionPointUploadCapability() (*ResourceCloudDistributionPointUploadCapability, error) {
	endpoint := uriCloudDistributionPoint + "/upload-capability"

	var cloudDistributionPointUploadCapability ResourceCloudDistributionPointUploadCapability
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &cloudDistributionPointUploadCapability)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "Cloud Distribution Point", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &cloudDistributionPointUploadCapability, nil
}
