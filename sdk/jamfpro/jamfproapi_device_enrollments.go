// jamfproapi_device_enrollments.go
// Jamf Pro Api - Device Enrollments
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-device-enrollments
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
	"net/url"

	"github.com/mitchellh/mapstructure"
)

const uriDeviceEnrollments = "/api/v1/device-enrollments"

// List

// ResponseDeviceEnrollmentList represents the response for device enrollments list.
type ResponseDeviceEnrollmentsList struct {
	TotalCount int                        `json:"totalCount"`
	Results    []ResourceDeviceEnrollment `json:"results"`
}

// ResponseDeviceEnrollmentHistory represents the response for device enrollment history
type ResponseDeviceEnrollmentHistory struct {
	TotalCount int                               `json:"totalCount"`
	Results    []ResourceDeviceEnrollmentHistory `json:"results"`
}

// Resource

// ResourceDeviceEnrollment represents a single device enrollment instance.
type ResourceDeviceEnrollment struct {
	ID                    string `json:"id"`
	Name                  string `json:"name"`
	SupervisionIdentityId string `json:"supervisionIdentityId"`
	SiteId                string `json:"siteId"`
	ServerName            string `json:"serverName"`
	ServerUuid            string `json:"serverUuid"`
	AdminId               string `json:"adminId"`
	OrgName               string `json:"orgName"`
	OrgEmail              string `json:"orgEmail"`
	OrgPhone              string `json:"orgPhone"`
	OrgAddress            string `json:"orgAddress"`
	TokenExpirationDate   string `json:"tokenExpirationDate"`
}

// ResourceDeviceEnrollmentHistory represents a single history entry
type ResourceDeviceEnrollmentHistory struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Date     string `json:"date"`
	Note     string `json:"note"`
	Details  string `json:"details"`
}

// ResourceDeviceEnrollmentSync represents a single sync state instance
type ResourceDeviceEnrollmentSync struct {
	SyncState  string `json:"syncState"`
	InstanceID string `json:"instanceId"`
	Timestamp  string `json:"timestamp"`
}

// ResourceDeviceEnrollmentTokenUpload represents the request body for token upload
type ResourceDeviceEnrollmentTokenUpload struct {
	TokenFileName string `json:"tokenFileName,omitempty"` // Optional
	EncodedToken  string `json:"encodedToken"`
}

// ResourceDeviceEnrollmentUpdate represents the request body for updating a device enrollment
type ResourceDeviceEnrollmentUpdate struct {
	Name                  string `json:"name"`                            // Required
	SupervisionIdentityId string `json:"supervisionIdentityId,omitempty"` // Optional
	SiteId                string `json:"siteId,omitempty"`                // Optional
}

// Response

// ResponseDeviceEnrollmentTokenUpload represents the response after token upload
type ResponseDeviceEnrollmentTokenUpload struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// CRUD

// GetDeviceEnrollments retrieves a paginated list of device enrollments.
func (c *Client) GetDeviceEnrollments(params url.Values) (*ResponseDeviceEnrollmentsList, error) {
	resp, err := c.DoPaginatedGet(
		uriDeviceEnrollments,
		standardPageSize,
		startingPageNumber,
		params,
	)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "device enrollments", err)
	}

	var out ResponseDeviceEnrollmentsList
	out.TotalCount = resp.Size

	for _, value := range resp.Results {
		var newObj ResourceDeviceEnrollment
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "device enrollments", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil
}

// GetDeviceEnrollmentByID retrieves a device enrollment by ID.
func (c *Client) GetDeviceEnrollmentByID(id string) (*ResourceDeviceEnrollment, error) {
	endpoint := fmt.Sprintf("%s/%v", uriDeviceEnrollments, id)
	var out ResourceDeviceEnrollment
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "device enrollment", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// GetDeviceEnrollmentByName retrieves a device enrollment by Name.
func (c *Client) GetDeviceEnrollmentByName(name string) (*ResourceDeviceEnrollment, error) {
	enrollments, err := c.GetDeviceEnrollments(url.Values{})
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "device enrollments", err)
	}

	for _, value := range enrollments.Results {
		if value.Name == name {
			return &value, nil
		}
	}

	return nil, fmt.Errorf(errMsgFailedGetByName, "device enrollment", name, errMsgNoName)
}

// GetDeviceEnrollmentHistory retrieves the history for a specific device enrollment
func (c *Client) GetDeviceEnrollmentHistory(id string, params url.Values) (*ResponseDeviceEnrollmentHistory, error) {
	endpoint := fmt.Sprintf("%s/%s/history", uriDeviceEnrollments, id)
	var out ResponseDeviceEnrollmentHistory
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "device enrollment history", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// GetDeviceEnrollmentSyncStates retrieves all sync states for a specific device enrollment instance
func (c *Client) GetDeviceEnrollmentSyncStates(id string) ([]ResourceDeviceEnrollmentSync, error) {
	endpoint := fmt.Sprintf("%s/%s/syncs", uriDeviceEnrollments, id)
	var out []ResourceDeviceEnrollmentSync
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "device enrollment sync states", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return out, nil
}

// CreateDeviceEnrollmentWithMDMServerToken creates a new device enrollment instance using
// The downloaded token base 64 encoded from the MDM server to be used to create a new Device Enrollment Instance.
func (c *Client) CreateDeviceEnrollmentWithMDMServerToken(tokenUpload *ResourceDeviceEnrollmentTokenUpload) (*ResponseDeviceEnrollmentTokenUpload, error) {
	endpoint := fmt.Sprintf("%s/upload-token", uriDeviceEnrollments)

	if tokenUpload.EncodedToken == "" {
		return nil, fmt.Errorf("encoded token is required")
	}

	var out ResponseDeviceEnrollmentTokenUpload
	resp, err := c.HTTP.DoRequest("POST", endpoint, tokenUpload, &out)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "device enrollment", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// UpdateDeviceEnrollmentMetadataByID updates a device enrollment instance with the specified ID
func (c *Client) UpdateDeviceEnrollmentMetadataByID(id string, deviceEnrollment *ResourceDeviceEnrollmentUpdate) (*ResourceDeviceEnrollment, error) {
	endpoint := fmt.Sprintf("%s/%s", uriDeviceEnrollments, id)

	if deviceEnrollment.Name == "" {
		return nil, fmt.Errorf("name is required for updating device enrollment")
	}

	var out ResourceDeviceEnrollment
	resp, err := c.HTTP.DoRequest("PUT", endpoint, deviceEnrollment, &out)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "device enrollment", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// UpdateDeviceEnrollmentMDMServerToken updates an existing device enrollment instance with a new token
func (c *Client) UpdateDeviceEnrollmentMDMServerToken(id string, tokenUpload *ResourceDeviceEnrollmentTokenUpload) (*ResponseDeviceEnrollmentTokenUpload, error) {
	endpoint := fmt.Sprintf("%s/%s/upload-token", uriDeviceEnrollments, id)

	// Validate required fields
	if tokenUpload.EncodedToken == "" {
		return nil, fmt.Errorf("encoded token is required")
	}

	var out ResponseDeviceEnrollmentTokenUpload
	resp, err := c.HTTP.DoRequest("PUT", endpoint, tokenUpload, &out)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "device enrollment token", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// DeleteDeviceEnrollmentByID deletes a device enrollment instance with the specified ID
func (c *Client) DeleteDeviceEnrollmentByID(id string) error {
	endpoint := fmt.Sprintf("%s/%s", uriDeviceEnrollments, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "device enrollment", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
