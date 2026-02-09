// jamfproapi_cloud_distribution_point.go
// Jamf Pro Api - Cloud Distribution Point
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-distribution-point
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
)

// Responses

const uriCloudDistributionPointV1 = "/api/v1/cloud-distribution-point"

// ResourceCloudDistributionPointV1 represents the writable payload used for both
// POST and PATCH /v1/cloud-distribution-point. Fields marked with pointers are
// only serialized when explicitly set, allowing partial updates while still
// letting callers reuse the same struct for creates.
type ResourceCloudDistributionPointV1 struct {
	CdnType                 string `json:"cdnType"`
	Master                  bool   `json:"master"`
	Username                string `json:"username,omitempty"`
	Password                string `json:"password,omitempty"`
	Directory               string `json:"directory,omitempty"`
	UploadUrl               string `json:"uploadUrl,omitempty"`
	DownloadUrl             string `json:"downloadUrl,omitempty"`
	SecondaryAuthRequired   *bool  `json:"secondaryAuthRequired,omitempty"`
	SecondaryAuthStatusCode *int   `json:"secondaryAuthStatusCode,omitempty"`
	SecondaryAuthTimeToLive *int   `json:"secondaryAuthTimeToLive,omitempty"`
	RequireSignedUrls       *bool  `json:"requireSignedUrls,omitempty"`
	KeyPairId               string `json:"keyPairId,omitempty"`
	ExpirationSeconds       *int   `json:"expirationSeconds,omitempty"`
	PrivateKey              string `json:"privateKey,omitempty"`
}

// ResponseCloudDistributionPointV1 models the read-only fields returned when
// fetching the cloud distribution point configuration.
type ResponseCloudDistributionPointV1 struct {
	HasConnectionSucceeded  bool   `json:"hasConnectionSucceeded"`
	Message                 string `json:"message"`
	InventoryId             string `json:"inventoryId"`
	CdnType                 string `json:"cdnType"`
	Master                  bool   `json:"master"`
	Username                string `json:"username"`
	Directory               string `json:"directory"`
	CdnUrl                  string `json:"cdnUrl"`
	UploadUrl               string `json:"uploadUrl"`
	DownloadUrl             string `json:"downloadUrl"`
	SecondaryAuthRequired   bool   `json:"secondaryAuthRequired"`
	SecondaryAuthStatusCode int    `json:"secondaryAuthStatusCode"`
	SecondaryAuthTimeToLive int    `json:"secondaryAuthTimeToLive"`
	RequireSignedUrls       bool   `json:"requireSignedUrls"`
	KeyPairId               string `json:"keyPairId"`
	ExpirationSeconds       int    `json:"expirationSeconds"`
	PrivateKey              any    `json:"privateKey"`
}

type ResourceCloudDistributionPointUploadCapabilityV1 struct {
	ID   bool `json:"principalDistributionTechnology"`
	Name bool `json:"directUploadCapable"`
}

type ResourceCloudDistributionPointTestConnectionV1 struct {
	HasConnectionSucceeded bool   `json:"hasConnectionSucceeded"`
	Message                string `json:"message"`
}

// GetCloudDistributionPointV1 retrieves the default server configuration for the Cloud Distribution Point.
func (c *Client) GetCloudDistributionPointV1() (*ResponseCloudDistributionPointV1, error) {
	endpoint := uriCloudDistributionPointV1

	var cloudDistributionPoint ResponseCloudDistributionPointV1
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &cloudDistributionPoint)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "Cloud Distribution Point", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &cloudDistributionPoint, nil
}

// CreateCloudDistributionPointV1 provisions a CDN-backed cloud distribution point configuration.
func (c *Client) CreateCloudDistributionPointV1(payload *ResourceCloudDistributionPointV1) (*ResponseCloudDistributionPointV1, error) {
	endpoint := uriCloudDistributionPointV1

	var created ResponseCloudDistributionPointV1
	resp, err := c.HTTP.DoRequest("POST", endpoint, payload, &created)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "cloud distribution point", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &created, nil
}

// UpdateCloudDistributionPointV1 performs a partial update against the current cloud distribution point configuration.
func (c *Client) UpdateCloudDistributionPointV1(payload *ResourceCloudDistributionPointV1) (*ResponseCloudDistributionPointV1, error) {
	endpoint := uriCloudDistributionPointV1

	var updated ResponseCloudDistributionPointV1
	resp, err := c.HTTP.DoRequest("PATCH", endpoint, payload, &updated)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdate, "cloud distribution point", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updated, nil
}

// DeleteCloudDistributionPointV1 removes the active cloud distribution point configuration.
func (c *Client) DeleteCloudDistributionPointV1() error {
	endpoint := uriCloudDistributionPointV1

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDelete, "cloud distribution point", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// GetCloudDistributionPointUploadCapabilityV1 retrieves the upload capability for the Cloud Distribution Point.
func (c *Client) GetCloudDistributionPointUploadCapabilityV1() (*ResourceCloudDistributionPointUploadCapabilityV1, error) {
	endpoint := uriCloudDistributionPointV1 + "/upload-capability"

	var cloudDistributionPointUploadCapability ResourceCloudDistributionPointUploadCapabilityV1
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &cloudDistributionPointUploadCapability)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "Cloud Distribution Point", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &cloudDistributionPointUploadCapability, nil
}

// GetCloudDistributionPointTestConnectionV1 retrieves the test connection status for the Cloud Distribution Point.
func (c *Client) GetCloudDistributionPointTestConnectionV1() (*ResourceCloudDistributionPointTestConnectionV1, error) {
	endpoint := uriCloudDistributionPointV1 + "/test-connection"

	var cloudDistributionPointTestConnection ResourceCloudDistributionPointTestConnectionV1
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &cloudDistributionPointTestConnection)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "Cloud Distribution Point", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &cloudDistributionPointTestConnection, nil
}
