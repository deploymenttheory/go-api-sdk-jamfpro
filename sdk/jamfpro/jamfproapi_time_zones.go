// jamfproapi_time_zones.go
// Jamf Pro Api - Time Zones
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-time-zones
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
)

const uriTimeZones = "/api/v1/time-zones"

// Resource structure
type ResourceTimeZone struct {
	ZoneId      string `json:"zoneId"`
	Region      string `json:"region"`
	DisplayName string `json:"displayName"`
}

// CRUD

// GetTimeZones retrieves the list of available time zones from Jamf Pro.
func (c *Client) GetTimeZones() ([]ResourceTimeZone, error) {
	endpoint := uriTimeZones

	var timeZones []ResourceTimeZone
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &timeZones)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "time zones", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return timeZones, nil
}
