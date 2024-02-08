// jamfproapi_package.go
// Jamf Pro Api - Package
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-package
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import "fmt"

const uriPackageV1 = "/api/v1/jamf-package"
const uriPackageV2 = "/api/v2/jamf-package"

// Structs

// List

// V1

type ResponsePackageListV1 struct {
	Results []ResponsePackageInfoShared
}

type ResponsePackageInfoShared struct {
	ID       string `json:"id"`
	Filename string `json:"filename"`
	Version  string `json:"version"`
	Created  string `json:"created"`
	Url      string `json:"url"`
}

// V2

// List

type ResponsePackageV2 struct {
	DisplayName       string `json:"displayName"`
	ReleaseHistoryUrl string `json:"releaseHistoryUrl"`
	Artifacts         []ResponsePackageInfoShared
}

// CRUD

// GetPackageInfoByApplicationV1 Returns a list of package info from the v1 api
func (c *Client) GetPackageInfoByApplicationV1(application string) (*ResponsePackageListV1, error) {
	endpoint := fmt.Sprintf("%s?application=%s", uriPackageV1, application)
	var out ResponsePackageListV1
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByString, "jamf package", "application name", application, err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &out, nil

}

// GetPackageInfoByApplicationV1 Returns a list of package info from the v2 api with more info
func (c *Client) GetPackageInfoByApplicationV2(application string) (*ResponsePackageV2, error) {
	endpoint := fmt.Sprintf("%s?application=%s", uriPackageV2, application)
	var out ResponsePackageV2
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByString, "jamf package", "application name", application, err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}
