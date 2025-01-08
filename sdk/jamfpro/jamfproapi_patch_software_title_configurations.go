// jamfproapi_patch_software_title_configurations.go
// Jamf Pro Api - Patch Software Title Configurations
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v2-patch-software-title-configurations-id
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import "fmt"

const uriPatchSoftwareTitleConfigurations = "/api/v2/patch-software-title-configurations"

// Structs

// List

type ResponsePatchSoftwareTitleConfigurationList []ResourcePatchSoftwareTitleConfiguration

// Resource

type ResourcePatchSoftwareTitleConfiguration struct {
	ID                     string                                                    `json:"id"`
	DisplayName            string                                                    `json:"displayName"`
	CategoryID             string                                                    `json:"categoryId"`
	SiteID                 string                                                    `json:"siteId"`
	UiNotifications        bool                                                      `json:"uiNotifications"`
	EmailNotifications     bool                                                      `json:"emailNotifications"`
	SoftwareTitleID        string                                                    `json:"softwareTitleId"`
	ExtensionAttributes    []PatchSoftwareTitleConfigurationSubsetExtensionAttribute `json:"extensionAttributes"`
	SoftwareTitleName      string                                                    `json:"softwareTitleName"`
	SoftwareTitleNameId    string                                                    `json:"softwareTitleNameId"`
	SoftwareTitlePublisher string                                                    `json:"softwareTitlePublisher"`
	JamfOfficial           bool                                                      `json:"jamfOfficial"`
	PatchSourceName        string                                                    `json:"patchSourceName"`
	PatchSourceEnabled     bool                                                      `json:"patchSourceEnabled"`
	Packages               []PatchSoftwareTitleConfigurationSubsetPackage            `json:"packages"`
}

// Subsets

type PatchSoftwareTitleConfigurationSubsetExtensionAttribute struct {
	Accepted bool   `json:"accepted"`
	EaID     string `json:"eaId"`
}

type PatchSoftwareTitleConfigurationSubsetPackage struct {
	PackageId   string `json:"packageId"`
	Version     string `json:"version"`
	DisplayName string `json:"displayName"`
}

// Response

type ResponsePatchSoftwareTitleConfigurationCreate struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// CRUD

// GetPatchSoftwareTitleConfigurations retrieves list of PatchSoftwareTitleConfigurations
func (c *Client) GetPatchSoftwareTitleConfigurations() (*ResponsePatchSoftwareTitleConfigurationList, error) {
	endpoint := uriPatchSoftwareTitleConfigurations
	var out ResponsePatchSoftwareTitleConfigurationList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "patch software title configurations", err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &out, nil

}

// GetPatchSoftwareTitleConfigurationById retrieves a singular PatchSoftwareTitleConfiguration from a given ID
func (c *Client) GetPatchSoftwareTitleConfigurationById(id string) (*ResourcePatchSoftwareTitleConfiguration, error) {
	endpoint := fmt.Sprintf("%s/%s", uriPatchSoftwareTitleConfigurations, id)
	var out ResourcePatchSoftwareTitleConfiguration
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "patch software title configuration", id, err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// GetPatchSoftwareTitleConfigurationByName retrieves a patch software title configuration by Name.
func (c *Client) GetPatchSoftwareTitleConfigurationByName(name string) (*ResourcePatchSoftwareTitleConfiguration, error) {
	patchSoftwareTitles, err := c.GetPatchSoftwareTitleConfigurations()
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "patch software title configuration", err)
	}

	for _, value := range *patchSoftwareTitles {
		if value.DisplayName == name {
			return &value, nil
		}
	}

	return nil, fmt.Errorf(errMsgFailedGetByName, "patch software title configuration", name, errMsgNoName)
}

// CreatePatchSoftwareTitleConfiguration Creates a new PatchSoftwareTitleConfiguration
func (c *Client) CreatePatchSoftwareTitleConfiguration(configuration ResourcePatchSoftwareTitleConfiguration) (*ResponsePatchSoftwareTitleConfigurationCreate, error) {
	endpoint := uriPatchSoftwareTitleConfigurations
	var out ResponsePatchSoftwareTitleConfigurationCreate
	resp, err := c.HTTP.DoRequest("POST", endpoint, configuration, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "patch software title configuration", err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// UpdatePatchSoftwareTitleConfigurationById Updates a single PatchSoftwareTitleConfiguration with given ID
func (c *Client) UpdatePatchSoftwareTitleConfigurationById(id string, updatedConfiguration ResourcePatchSoftwareTitleConfiguration) (*ResponsePatchSoftwareTitleConfigurationCreate, error) {
	endpoint := fmt.Sprintf("%s/%s", uriPatchSoftwareTitleConfigurations, id)
	var out ResponsePatchSoftwareTitleConfigurationCreate
	resp, err := c.HTTP.DoRequest("PATCH", endpoint, updatedConfiguration, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "patch software title configuration", err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// DeletePatchSoftwareTitleConfigurationById deletes a PatchSoftwareTitleConfiguration with given ID
func (c *Client) DeletePatchSoftwareTitleConfigurationById(id string) error {
	endpoint := fmt.Sprintf("%s/%s", uriPatchSoftwareTitleConfigurations, id)
	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "patch software title configuration", id, err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return nil
}
