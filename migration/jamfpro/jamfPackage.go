// jamfPackage.go
// Jamf Pro API

package jamfpro

import (
	"fmt"
)

const (
	uriAPIJamfPackageV1 = "/api/v1/jamf-package"
	uriAPIJamfPackageV2 = "/api/v2/jamf-package"
)

// JamfPackageV1 Response structure (from your initial request)
type JamfPackageV1 struct {
	ID       string `json:"id"`
	Filename string `json:"filename"`
	Version  string `json:"version"`
	Created  string `json:"created"`
	URL      string `json:"url"`
}

// JamfPackageV2 Response structure (from the new v2 API)
type JamfPackageV2 struct {
	DisplayName       string         `json:"displayName"`
	ReleaseHistoryUrl string         `json:"releaseHistoryUrl"`
	Artifacts         []JamfArtifact `json:"artifacts"`
}

type JamfArtifact struct {
	ID       string `json:"id"`
	Filename string `json:"filename"`
	Version  string `json:"version"`
	Created  string `json:"created"`
	URL      string `json:"url"`
}

// GetJamfPackageV1 retrieves the packages for a given Jamf application using v1 API
func (c *Client) GetJamfPackageV1(application string) ([]JamfPackageV1, error) {
	if application != "protect" && application != "connect" {
		return nil, fmt.Errorf("invalid application value: %s. Supported values are 'protect' and 'connect'", application)
	}

	url := fmt.Sprintf("%s?application=%s", uriAPIJamfPackageV1, application)

	var packagesList []JamfPackageV1
	if err := c.DoRequest("GET", url, nil, nil, &packagesList); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return packagesList, nil
}

// GetJamfPackageV2 retrieves the packages for a given Jamf application using v2 API
func (c *Client) GetJamfPackageV2(application string) (JamfPackageV2, error) {
	if application != "protect" && application != "connect" {
		return JamfPackageV2{}, fmt.Errorf("invalid application value: %s. Supported values are 'protect' and 'connect'", application)
	}

	url := fmt.Sprintf("%s?application=%s", uriAPIJamfPackageV2, application)

	var packageInfo JamfPackageV2
	if err := c.DoRequest("GET", url, nil, nil, &packageInfo); err != nil {
		return JamfPackageV2{}, fmt.Errorf("failed to execute request: %v", err)
	}

	return packageInfo, nil
}
