package jamfpro

import (
	"fmt"
)

const uriDeviceEnrollments = "/api/v1/device-enrollments"

type ResponseDeviceEnrollment struct {
	TotalCount int                `json:"totalCount"`
	Results    []DeviceEnrollment `json:"results"`
}

type DeviceEnrollment struct {
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

type UpdateDeviceEnrollmentTokenRequest struct {
	TokenFileName string `json:"tokenFileName"`
	EncodedToken  string `json:"encodedToken"`
}

type InstanceSyncState struct {
	SyncState  string `json:"syncState"`
	InstanceID string `json:"instanceId"`
	Timestamp  string `json:"timestamp"`
}

type DeviceAssignedToEnrollment struct {
	ID                                string    `json:"id"`
	DeviceEnrollmentProgramInstanceId string    `json:"deviceEnrollmentProgramInstanceId"`
	PrestageId                        string    `json:"prestageId"`
	SerialNumber                      string    `json:"serialNumber"`
	Description                       string    `json:"description"`
	Model                             string    `json:"model"`
	Color                             string    `json:"color"`
	AssetTag                          string    `json:"assetTag"`
	ProfileStatus                     string    `json:"profileStatus"`
	SyncState                         SyncState `json:"syncState"`
	ProfileAssignTime                 string    `json:"profileAssignTime"`
	ProfilePushTime                   string    `json:"profilePushTime"`
	DeviceAssignedDate                string    `json:"deviceAssignedDate"`
}

type SyncState struct {
	ID           int    `json:"id"`
	SerialNumber string `json:"serialNumber"`
	ProfileUUID  string `json:"profileUUID"`
	SyncStatus   string `json:"syncStatus"`
	FailureCount int    `json:"failureCount"`
	Timestamp    int    `json:"timestamp"`
}

func (c *Client) GetDeviceEnrollmentIdByName(name string) (string, error) {
	var id string
	enrollments, err := c.GetDeviceEnrollments()
	if err != nil {
		return "", err
	}

	for _, v := range enrollments.Results {
		if v.Name == name {
			id = v.ID
			break
		}
	}
	return id, err
}

func (c *Client) GetDeviceEnrollmentByName(name string) (*DeviceEnrollment, error) {
	allEnrollmentsResponse, err := c.GetDeviceEnrollments()
	if err != nil {
		return nil, err
	}

	for _, enrollment := range allEnrollmentsResponse.Results {
		if enrollment.Name == name {
			return &enrollment, nil
		}
	}

	return nil, fmt.Errorf("device enrollment with name '%s' not found", name)
}

func (c *Client) GetDeviceEnrollments() (*ResponseDeviceEnrollment, error) {
	uri := fmt.Sprintf("%s?page=0&page-size=100&sort=id%%3Aasc", uriDeviceEnrollments)

	var out ResponseDeviceEnrollment
	err := c.DoRequest("GET", uri, nil, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to get device enrollments: %v", err)
	}

	return &out, nil
}

func (c *Client) GetDeviceEnrollmentByID(enrollmentID string) (*DeviceEnrollment, error) {
	uri := fmt.Sprintf("%s/%s", uriDeviceEnrollments, enrollmentID)

	var out DeviceEnrollment
	err := c.DoRequest("GET", uri, nil, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to get device enrollment by ID: %v", err)
	}

	return &out, nil
}

// GetJamfProDeviceEnrollmentPublicKey retrieves the public key for device enrollments.
func (c *Client) GetJamfProDeviceEnrollmentPublicKey() (string, error) {
	uri := fmt.Sprintf("%s/public-key", uriDeviceEnrollments)

	// Call the specialized function
	response, err := c.DoRawRequest(uri, nil)
	if err != nil {
		return "", fmt.Errorf("failed to get Jamf Pro Device Enrollment public key: %v", err)
	}

	return response, nil
}

func (c *Client) GetDevicesAssignedToDeviceEnrollmentID(enrollmentID string) ([]DeviceAssignedToEnrollment, error) {
	uri := fmt.Sprintf("%s/%s/devices", uriDeviceEnrollments, enrollmentID)

	var out struct {
		TotalCount int                          `json:"totalCount"`
		Results    []DeviceAssignedToEnrollment `json:"results"`
	}
	err := c.DoRequest("GET", uri, nil, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to get devices assigned to device enrollment ID: %v", err)
	}

	return out.Results, nil
}

func (c *Client) GetLatestSyncStateForDeviceEnrollmentInstance(enrollmentID string) (*InstanceSyncState, error) {
	uri := fmt.Sprintf("%s/%s/syncs/latest", uriDeviceEnrollments, enrollmentID)

	var out InstanceSyncState
	err := c.DoRequest("GET", uri, nil, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest sync state for device enrollment instance: %v", err)
	}

	return &out, nil
}

func (c *Client) GetAllSyncStateForDeviceEnrollmentInstance() ([]InstanceSyncState, error) {
	uri := fmt.Sprintf("%s/syncs", uriDeviceEnrollments)

	var out []InstanceSyncState
	err := c.DoRequest("GET", uri, nil, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to get all instance sync states for all Device Enrollment Instances: %v", err)
	}

	return out, nil
}

func (c *Client) CreateDeviceEnrollmentInstanceWithToken(tokenFileName string, encodedToken string) (string, error) {
	uri := fmt.Sprintf("%s/upload-token", uriDeviceEnrollments)

	payload := struct {
		TokenFileName string `json:"tokenFileName"`
		EncodedToken  string `json:"encodedToken"`
	}{
		TokenFileName: tokenFileName,
		EncodedToken:  encodedToken,
	}

	var out struct {
		ID   string `json:"id"`
		Href string `json:"href"`
	}
	err := c.DoRequest("POST", uri, payload, nil, &out)
	if err != nil {
		return "", fmt.Errorf("failed to create device enrollment instance with token: %v", err)
	}

	return out.ID, nil
}

func (c *Client) UpdateDeviceEnrollment(d *DeviceEnrollment) (*DeviceEnrollment, error) {
	uri := fmt.Sprintf("%s/%s", uriDeviceEnrollments, d.ID)
	updatedEnrollment := &DeviceEnrollment{}

	err := c.DoRequest("PUT", uri, d, nil, updatedEnrollment)
	if err != nil {
		return nil, fmt.Errorf("failed to update device enrollment: %v", err)
	}

	return updatedEnrollment, nil
}

// UpdateDeviceEnrollmentInstanceWithToken updates a device enrollment instance with the provided token.
func (c *Client) UpdateDeviceEnrollmentInstanceWithToken(enrollmentID string, tokenFileName string, encodedToken string) (*DeviceEnrollment, error) {
	uri := fmt.Sprintf("%s/%s/upload-token", uriDeviceEnrollments, enrollmentID)

	payload := &UpdateDeviceEnrollmentTokenRequest{
		TokenFileName: tokenFileName,
		EncodedToken:  encodedToken,
	}

	var out DeviceEnrollment
	err := c.DoRequest("PUT", uri, payload, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to update device enrollment instance with token: %v", err)
	}

	return &out, nil
}

func (c *Client) DeleteDeviceEnrollment(enrollmentID string) error {
	uri := fmt.Sprintf("%s/%s", uriDeviceEnrollments, enrollmentID)

	err := c.DoRequest("DELETE", uri, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete device enrollment: %v", err)
	}

	return nil
}
