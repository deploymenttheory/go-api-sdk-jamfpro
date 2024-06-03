// jamfproapi_login_customization.go
// Jamf Pro Api - Login Customization
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-login-customization
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
)

const uriLoginCustomization = "/api/v1/login-customization"

// Resource

// ResourceLoginCustomization represents the structure of the response for login customization.
type ResourceLoginCustomization struct {
	RampInstance            bool   `json:"rampInstance"`
	IncludeCustomDisclaimer bool   `json:"includeCustomDisclaimer"`
	DisclaimerHeading       string `json:"disclaimerHeading"`
	DisclaimerMainText      string `json:"disclaimerMainText"`
	ActionText              string `json:"actionText"`
}

// CRUD

/*
Function: GetLoginCustomization
Method: GET
Path: /api/v1/login-customization
Description: Gets the login customization settings.
Parameters: None
Returns: ResourceLoginCustomization - The login customization settings.
Example:

	customization, err := client.GetLoginCustomization()
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Println(customization)

Errors: Returns an error if the request fails.
*/
func (c *Client) GetLoginCustomization() (*ResourceLoginCustomization, error) {
	endpoint := uriLoginCustomization

	var loginCustomization ResourceLoginCustomization
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &loginCustomization)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "login customization", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &loginCustomization, nil
}

/*
Function: UpdateLoginCustomization
Method: PUT
Path: /api/v1/login-customization
Description: Updates the login customization settings.
Parameters:
  - loginCustomizationUpdate (*ResourceLoginCustomization): The updated login customization settings.

Returns: ResourceLoginCustomization - The updated login customization settings.
Example:

	updatedSettings := &jamfpro.ResourceLoginCustomization{
	    RampInstance: false,
	    IncludeCustomDisclaimer: true,
	    DisclaimerHeading: "Updated Disclaimer Header",
	    DisclaimerMainText: "Updated disclaimer main text",
	    ActionText: "Accept",
	}
	updated, err := client.UpdateLoginCustomization(updatedSettings)
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Println(updated)

Errors: Returns an error if the request fails or if the resource cannot be updated.
*/
func (c *Client) UpdateLoginCustomization(loginCustomizationUpdate *ResourceLoginCustomization) (*ResourceLoginCustomization, error) {
	endpoint := uriLoginCustomization

	var updatedLoginCustomization ResourceLoginCustomization
	resp, err := c.HTTP.DoRequest("PUT", endpoint, loginCustomizationUpdate, &updatedLoginCustomization)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdate, "login customization", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedLoginCustomization, nil
}
