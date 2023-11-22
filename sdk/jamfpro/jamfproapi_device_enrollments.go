// jamfproapi_device_enrollments.go
// Jamf Pro Api - Device Enrollments
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-device-enrollments
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
	"net/url"
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
// If page and pageSize are both set to -1, it retrieves all pages.
func (c *Client) GetDeviceEnrollments(page, pageSize int, sort []string) (*ResponseDeviceEnrollmentList, error) {
	var allResults []DeviceEnrollment
	currentPage := 1
	totalFetched := 0

	for {
		endpoint := uriDeviceEnrollments
		params := url.Values{}

		if page != -1 && pageSize != -1 {
			params.Add("page", fmt.Sprintf("%d", currentPage))
			params.Add("page-size", fmt.Sprintf("%d", pageSize))
		} else if pageSize > 0 {
			params.Add("page-size", fmt.Sprintf("%d", pageSize))
		}

		if len(sort) > 0 {
			params.Add("sort", url.QueryEscape(strings.Join(sort, ",")))
		}

		endpointWithParams := fmt.Sprintf("%s?%s", endpoint, params.Encode())

		var response ResponseDeviceEnrollmentList
		resp, err := c.HTTP.DoRequest("GET", endpointWithParams, nil, &response)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch device enrollments: %v", err)
		}

		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}

		allResults = append(allResults, response.Results...)
		totalFetched += len(response.Results)

		if totalFetched >= response.TotalCount || (page != -1 && currentPage == page) {
			break
		}
		currentPage++
	}

	return &ResponseDeviceEnrollmentList{
		TotalCount: len(allResults),
		Results:    allResults,
	}, nil
}
