// jamfproapi_jamf_connect.go
// Jamf Pro Api - Jamf Connect
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-connect
// Jamf Pro API requires the structs to support an JSON data structure.

package jamfpro

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

const uriJamfConnect = "/api/v1/jamf-connect"

// List

// Response

// ResponseJamfConnectConfigProfilesList Struct for paginated response for Jamf Connect config profiles
type ResponseJamfConnectConfigProfilesList struct {
	TotalCount int                                `json:"totalCount"`
	Results    []ResourceJamfConnectConfigProfile `json:"results"`
}

// ResponseJamfConnectError struct for Jamf Connect error response
type ResponseJamfConnectError struct {
	HTTPStatus int                `json:"httpStatus"`
	Errors     []JamfConnectError `json:"errors"`
}

type JamfConnectError struct {
	Code        string `json:"code"`
	Field       string `json:"field"`
	Description string `json:"description"`
	ID          string `json:"id"`
}

// Resource

// ResourceJamfConnect  Struct to represent the Jamf Connect settings
type ResourceJamfConnect struct {
	ID             string `json:"id,omitempty"`
	DisplayName    string `json:"displayName,omitempty"`
	Description    string `json:"description,omitempty"`
	Enabled        bool   `json:"enabled"`
	Settings       string `json:"settings,omitempty"`
	Version        string `json:"version,omitempty"`
	LastModified   string `json:"lastModified,omitempty"`
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`
}

// Struct representing a Jamf Connect config profile
type ResourceJamfConnectConfigProfile struct {
	UUID               string `json:"uuid"`
	ProfileID          int    `json:"profileId"`
	ProfileName        string `json:"profileName"`
	ScopeDescription   string `json:"scopeDescription"`
	SiteID             string `json:"siteId"`
	Version            string `json:"version"`
	AutoDeploymentType string `json:"autoDeploymentType"`
}

// ResourceJamfConnectTaskRetry represents the request structure for task retry
type ResourceJamfConnectTaskRetry struct {
	IDs []string `json:"ids"`
}

// ResourceJamfConnectConfigProfileUpdate represents the updateable fields for a Jamf Connect profile
type ResourceJamfConnectConfigProfileUpdate struct {
	JamfConnectVersion string `json:"version,omitempty"`
	AutoDeploymentType string `json:"autoDeploymentType,omitempty"`
}

// CRUD

// GetJamfConnectSettings fetches Jamf Connect settings from Jamf Pro
func (c *Client) GetJamfConnectSettings() (*ResourceJamfConnect, error) {
	endpoint := uriJamfConnect
	var jamfConnect ResourceJamfConnect

	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &jamfConnect)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "jamf connect settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &jamfConnect, nil
}

// GetJamfConnectConfigProfiles gets full list of Jamf Connect config profiles & handles pagination
func (c *Client) GetJamfConnectConfigProfiles(sort_filter string) (*ResponseJamfConnectConfigProfilesList, error) {
	endpoint := fmt.Sprintf("%s/config-profiles", uriJamfConnect)

	resp, err := c.DoPaginatedGet(
		endpoint,
		standardPageSize,
		startingPageNumber,
		sort_filter,
	)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "jamf connect config profiles", err)
	}

	var out ResponseJamfConnectConfigProfilesList
	out.TotalCount = resp.Size

	for _, value := range resp.Results {
		var newObj ResourceJamfConnectConfigProfile
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "jamf connect config profile", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil
}

// UpdateJamfConnectConfigProfileByConfigProfileUUID updates the way the Jamf Connect app gets updated on computers within scope of the associated configuration profile.
// The profile is identified by its UUID.
func (c *Client) UpdateJamfConnectConfigProfileByConfigProfileUUID(configProfileUUID string, profileUpdate *ResourceJamfConnectConfigProfileUpdate) (*ResourceJamfConnectConfigProfile, error) {
	endpoint := fmt.Sprintf("%s/config-profiles/%s", uriJamfConnect, configProfileUUID)
	var updatedProfile ResourceJamfConnectConfigProfile

	resp, err := c.HTTP.DoRequest("PUT", endpoint, profileUpdate, &updatedProfile)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "jamf connect config profile", configProfileUUID, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedProfile, nil
}

// RetryJamfConnectDeploymentTasksByConfigProfileUUID requests a retry of Connect install tasks for a specified computers
// asscoiated with a specific jamf connect configuration profile.
func (c *Client) RetryJamfConnectDeploymentTasksByConfigProfileUUID(configProfileUUID string, computerIDs []string) error {
	endpoint := fmt.Sprintf("%s/deployments/%s/tasks/retry", uriJamfConnect, configProfileUUID)

	requestBody := &ResourceJamfConnectTaskRetry{
		IDs: computerIDs,
	}

	resp, err := c.HTTP.DoRequest("POST", endpoint, requestBody, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedCreate, "retry jamf connect tasks", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
