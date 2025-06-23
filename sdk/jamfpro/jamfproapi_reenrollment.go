// jamfproapi_reenrollment.go
// Jamf Pro Api - Re-enrollment (preview)
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-reenrollment
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import "fmt"

const uriReenrollmentSettings = "/api/v1/reenrollment"

// Structs

// Resource

type ResourceReenrollmentSettings struct {
	FlushPolicyHistory              bool   `json:"isFlushPolicyHistoryEnabled"`
	FlushLocationInformation        bool   `json:"isFlushLocationInformationEnabled"`
	FlushLocationInformationHistory bool   `json:"isFlushLocationInformationHistoryEnabled"`
	FlushExtensionAttributes        bool   `json:"isFlushExtensionAttributesEnabled"`
	FlushSoftwareUpdatePlans        bool   `json:"isFlushSoftwareUpdatePlansEnabled"`
	FlushMdmQueue                   string `json:"flushMDMQueue"`
}

// CRUD

func (c *Client) GetReenrollmentSettings() (*ResourceReenrollmentSettings, error) {
	endpoint := uriReenrollmentSettings

	var out ResourceReenrollmentSettings
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "re-enrollment settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

func (c *Client) UpdateReenrollmentSettings(settingsUpdate ResourceReenrollmentSettings) (*ResourceReenrollmentSettings, error) {
	endpoint := uriReenrollmentSettings

	var out ResourceReenrollmentSettings
	resp, err := c.HTTP.DoRequest("PUT", endpoint, settingsUpdate, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdate, "re-enrollment settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil

}
