// jamfproapi_service_discovery_enrollment.go
// Jamf Pro API - Service Discovery Enrollment
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-service-discovery-enrollment-well-known-settings
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
	"net/http"
)

const uriServiceDiscoveryEnrollmentWellKnownSettingsV1 = "/api/v1/service-discovery-enrollment/well-known-settings"

// ResourceServiceDiscoveryWellKnownSettingV1 represents a single well-known setting entry.
type ResourceServiceDiscoveryWellKnownSettingV1 struct {
	OrgName        string `json:"orgName,omitempty"`
	ServerUUID     string `json:"serverUuid"`
	EnrollmentType string `json:"enrollmentType"`
}

// ResponseServiceDiscoveryEnrollmentWellKnownSettingsV1 represents the response structure for retrieving all well-known service discovery settings.
type ResponseServiceDiscoveryEnrollmentWellKnownSettingsV1 struct {
	WellKnownSettings []ResourceServiceDiscoveryWellKnownSettingV1 `json:"wellKnownSettings"`
}

// GetServiceDiscoveryEnrollmentWellKnownSettingsV1 retrieves all well-known service discovery settings.
func (c *Client) GetServiceDiscoveryEnrollmentWellKnownSettingsV1() (*ResponseServiceDiscoveryEnrollmentWellKnownSettingsV1, error) {
	endpoint := uriServiceDiscoveryEnrollmentWellKnownSettingsV1

	var response ResponseServiceDiscoveryEnrollmentWellKnownSettingsV1
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "service discovery well-known settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateServiceDiscoveryEnrollmentWellKnownSettingsV1 updates the enrollment types for all organizations.
func (c *Client) UpdateServiceDiscoveryEnrollmentWellKnownSettingsV1(request ResponseServiceDiscoveryEnrollmentWellKnownSettingsV1) error {
	endpoint := uriServiceDiscoveryEnrollmentWellKnownSettingsV1

	resp, err := c.HTTP.DoRequest("PUT", endpoint, request, nil)
	if err != nil && resp == nil {
		return fmt.Errorf(errMsgFailedUpdate, "service discovery well-known settings", err)
	}

	if resp == nil {
		return fmt.Errorf("failed to update service discovery well-known settings: received nil response")
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to update service discovery well-known settings: unexpected status code %d", resp.StatusCode)
	}

	return nil
}
