// jamfproapi_cloud_distribution_point.go
// Jamf Pro Api - Cloud Distribution Point
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-distribution-point
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
)

// Responses

const uriCloudDistributionPoint = "/api/v1/cloud-distribution-point"

// Resource
type ResourceCloudDistributionPoint struct {
	HasConnectionSucceeded  bool        `json:"hasConnectionSucceeded"`
	Message                 string      `json:"message"`
	InventoryId             string      `json:"inventoryId"`
	CdnType                 string      `json:"cdnType"`
	Master                  bool        `json:"master"`
	Username                string      `json:"username"`
	Directory               string      `json:"directory"`
	CdnUrl                  string      `json:"cdnUrl"`
	UploadUrl               string      `json:"uploadUrl"`
	DownloadUrl             string      `json:"downloadUrl"`
	SecondaryAuthRequired   bool        `json:"secondaryAuthRequired"`
	SecondaryAuthStatusCode int         `json:"secondaryAuthStatusCode"`
	SecondaryAuthTimeToLive int         `json:"secondaryAuthTimeToLive"`
	RequireSignedUrls       bool        `json:"requireSignedUrls"`
	KeyPairId               string      `json:"keyPairId"`
	ExpirationSeconds       int         `json:"expirationSeconds"`
	PrivateKey              interface{} `json:"privateKey"`
}

type ResourceCloudDistributionPointUploadCapability struct {
	ID   bool `json:"principalDistributionTechnology"`
	Name bool `json:"directUploadCapable"`
}

type ResourceCloudDistributionPointTestConnection struct {
	HasConnectionSucceeded bool   `json:"hasConnectionSucceeded"`
	Message                string `json:"message"`
}

// GetCloudDistributionPoint retrieves the default server configuration for the Cloud Distribution Point.
func (c *Client) GetCloudDistributionPoint() (*ResourceCloudDistributionPoint, error) {
	endpoint := uriCloudDistributionPoint

	var cloudDistributionPoint ResourceCloudDistributionPoint
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &cloudDistributionPoint)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "Cloud Distribution Point", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &cloudDistributionPoint, nil
}

// GetCloudDistributionPointUploadCapability retrieves the upload capability for the Cloud Distribution Point.
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

// GetCloudDistributionPointTestConnection retrieves the test connection status for the Cloud Distribution Point.
func (c *Client) GetCloudDistributionPointTestConnection() (*ResourceCloudDistributionPointTestConnection, error) {
	endpoint := uriCloudDistributionPoint + "/test-connection"

	var cloudDistributionPointTestConnection ResourceCloudDistributionPointTestConnection
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &cloudDistributionPointTestConnection)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "Cloud Distribution Point", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &cloudDistributionPointTestConnection, nil
}
