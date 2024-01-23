// jamfproapi_patch_software_title_configurations.go
// Jamf Pro Api - Patch Software Title Configurations
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v2-patch-software-title-configurations-id
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
)

const uriPatchSoftwareTitleConfigurations = "/api/v2/patch-software-title-configurations"

// Response

// Response format struct for create function
type ResponsePatchSoftwareTitleConfigurationCreateAndUpdate struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// Resource

// Struct which represents SoftwareTitleConfiguration object JSON from Pro API
type ResourcePatchSoftwareTitleConfiguration struct {
	DisplayName            string                                                `json:"displayName"`
	CategoryId             string                                                `json:"categoryId"`
	SiteId                 string                                                `json:"siteId"`
	UINotifications        bool                                                  `json:"uiNotifications"`
	EmailNotifications     bool                                                  `json:"emailNotifications"`
	SoftwareTitleId        string                                                `json:"softwareTitleId"`
	JamfOfficial           bool                                                  `json:"jamfOfficial"`
	ExtensionAttributes    []SoftwareTitleConfigurationsSubsetExtensionAttribute `json:"extensionAttributes"`
	SoftwareTitleName      string                                                `json:"softwareTitleName"`
	SoftwareTitleNameId    string                                                `json:"softwareTitleNameId"`
	SoftwareTitlePublisher string                                                `json:"softwareTitlePublisher"`
	PatchSourceName        string                                                `json:"patchSourceName"`
	PatchSourceEnabled     bool                                                  `json:"patchSourceEnabled"`
	ID                     string                                                `json:"id"`
	Packages               []SoftwareTitleConfigurationsSubsetPackage            `json:"packages"`
}

type SoftwareTitleConfigurationsSubsetExtensionAttribute struct {
	Accepted bool   `json:"accepted"`
	EaId     string `json:"eaId"`
}

type SoftwareTitleConfigurationsSubsetPackage struct {
	PackageId   string `json:"packageId"`
	Version     string `json:"version"`
	DisplayName string `json:"displayName"`
}

// CRUD

// GetPatchSoftwareTitleConfigurations retrieves a list of all Jamf App Catalog apps.
func (c *Client) GetPatchSoftwareTitleConfigurations() ([]ResourcePatchSoftwareTitleConfiguration, error) {
	endpoint := uriPatchSoftwareTitleConfigurations

	var jamfAppCatalogApps []ResourcePatchSoftwareTitleConfiguration
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &jamfAppCatalogApps)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "Jamf App Catalog apps", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return jamfAppCatalogApps, nil
}

// Retrieves PatchSoftwareTitleConfiguration from provided ID & returns ResourcePatchSoftwareTitleConfiguration
func (c *Client) GetPatchSoftwareTitleConfigurationByID(id string) (*ResourcePatchSoftwareTitleConfiguration, error) {
	endpoint := fmt.Sprintf("%s/%s", uriPatchSoftwareTitleConfigurations, id)
	var softwareTitleConfiguration ResourcePatchSoftwareTitleConfiguration
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &softwareTitleConfiguration)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "patch software title configuration", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &softwareTitleConfiguration, nil
}

// Retrieves PatchSoftwareTitleConfiguration by display Name by leveraging GetPatchSoftwareTitleConfigurations(), returns ResourcePatchSoftwareTitleConfiguration
func (c *Client) GetPatchSoftwareTitleConfigurationByName(name string) (*ResourcePatchSoftwareTitleConfiguration, error) {
	patchSoftwareTitleConfiguration, err := c.GetPatchSoftwareTitleConfigurations()
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "patch software title configuration", err)
	}

	for _, config := range patchSoftwareTitleConfiguration {
		if config.DisplayName == name {
			return &config, nil
		}
	}

	return nil, fmt.Errorf(errMsgFailedGetByName, "patch software title configuration", name, errMsgNoName)
}

// Creates PatchSoftwareTitleConfiguration from ResourcePatchSoftwareTitleConfiguration struct
func (c *Client) CreatePatchSoftwareTitleConfiguration(PatchSoftwareTitleConfiguration *ResourcePatchSoftwareTitleConfiguration) (*ResponsePatchSoftwareTitleConfigurationCreateAndUpdate, error) {
	endpoint := uriPatchSoftwareTitleConfigurations
	var ResponsePatchSoftwareTitleConfigurationCreate ResponsePatchSoftwareTitleConfigurationCreateAndUpdate

	resp, err := c.HTTP.DoRequest("POST", endpoint, PatchSoftwareTitleConfiguration, &ResponsePatchSoftwareTitleConfigurationCreate)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "Patch Software Title Configuration", err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &ResponsePatchSoftwareTitleConfigurationCreate, nil
}
