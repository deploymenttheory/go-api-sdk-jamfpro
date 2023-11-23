// jamfproapi_device_enrollments.go
// Jamf Pro Api - Device Enrollments
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-device-enrollments
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

const uriDeviceEnrollments = "/api/v1/device-enrollments"

// ResponseDeviceEnrollmentList represents the response for device enrollments list.
type ResponseDeviceEnrollmentList struct {
	TotalCount int                `json:"totalCount"`
	Results    []DeviceEnrollment `json:"results"`
}

// DeviceEnrollment represents a single device enrollment instance.
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

// GetDeviceEnrollments retrieves a paginated list of device enrollments.
func (c *Client) GetDeviceEnrollments(sort []string) (*ResponseDeviceEnrollmentList, error) {
	const maxPageSize = 2000 // Assuming 2000 is a suitable limit for this API
	var allEnrollments []DeviceEnrollment

	page := 0
	for {
		// Construct the endpoint with query parameters for the current page
		params := url.Values{
			"page":      []string{strconv.Itoa(page)},
			"page-size": []string{strconv.Itoa(maxPageSize)},
		}
		if len(sort) > 0 {
			params.Add("sort", url.QueryEscape(strings.Join(sort, ",")))
		}
		endpointWithParams := fmt.Sprintf("%s?%s", uriDeviceEnrollments, params.Encode())

		// Fetch the device enrollments for the current page
		var response ResponseDeviceEnrollmentList
		resp, err := c.HTTP.DoRequest("GET", endpointWithParams, nil, &response)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch device enrollments: %v", err)
		}

		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}

		// Add the fetched enrollments to the total list
		allEnrollments = append(allEnrollments, response.Results...)

		// Check if all enrollments have been fetched
		if len(allEnrollments) >= response.TotalCount {
			break
		}

		// Increment page number for the next iteration
		page++
	}

	// Return the combined list of all device enrollments
	return &ResponseDeviceEnrollmentList{
		TotalCount: len(allEnrollments),
		Results:    allEnrollments,
	}, nil
}
