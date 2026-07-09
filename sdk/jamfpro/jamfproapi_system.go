// jamfproapi_system.go
// Jamf Pro Api - System
// api reference: https://developer.jamf.com/jamf-pro/reference/post_v1-system-platform-initialize
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import "fmt"

const uriSystemPlatformInitialize = "/api/v1/system/platform-initialize"

// ResourceSystemPlatformInitialize represents the payload for initial Jamf Pro platform setup.
type ResourceSystemPlatformInitialize struct {
	ActivationCode  string `json:"activationCode"`
	InstitutionName string `json:"institutionName"`
	EulaAccepted    bool   `json:"eulaAccepted"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	JssURL          string `json:"jssUrl"`
}

// InitializeSystemPlatform performs the initial platform setup of a Jamf Pro server.
func (c *Client) InitializeSystemPlatform(request *ResourceSystemPlatformInitialize) error {
	endpoint := uriSystemPlatformInitialize

	resp, err := c.HTTP.DoRequest("POST", endpoint, request, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedAction, "initialize system platform", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
