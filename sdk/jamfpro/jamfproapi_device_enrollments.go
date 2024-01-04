// jamfproapi_device_enrollments.go
// Jamf Pro Api - Device Enrollments
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-device-enrollments
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

const uriDeviceEnrollments = "/api/v1/device-enrollments"

// List

// ResponseDeviceEnrollmentList represents the response for device enrollments list.
type ResponseDeviceEnrollmentsList struct {
	TotalCount int                        `json:"totalCount"`
	Results    []ResourceDeviceEnrollment `json:"results"`
}

// Resource

// DeviceEnrollment represents a single device enrollment instance.
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

// CRUD

// GetDeviceEnrollments retrieves a paginated list of device enrollments.
func (c *Client) GetDeviceEnrollments(sort_filter string) (*ResponseDeviceEnrollmentsList, error) {
	resp, err := c.DoPaginatedGet(
		uriDeviceEnrollments,
		standardPageSize,
		startingPageNumber,
		sort_filter,
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
