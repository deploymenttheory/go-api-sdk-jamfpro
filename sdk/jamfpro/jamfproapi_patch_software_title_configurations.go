// jamfproapi_patch_software_title_configurations.go
// Jamf Pro Api - Patch Software Title Configurations
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v2-patch-software-title-configurations-id
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
	"net/url"

	"github.com/mitchellh/mapstructure"
)

const uriPatchSoftwareTitleConfigurations = "/api/v2/patch-software-title-configurations"

// Structs

// List

type ResponsePatchSoftwareTitleConfigurationList []ResourcePatchSoftwareTitleConfiguration

// Resource

type ResourcePatchSoftwareTitleConfiguration struct {
	ID                     string                                                    `json:"id,omitempty"`
	DisplayName            string                                                    `json:"displayName"`
	SoftwareTitleID        string                                                    `json:"softwareTitleId"`
	CategoryID             string                                                    `json:"categoryId,omitempty"`
	SiteID                 string                                                    `json:"siteId,omitempty"`
	UiNotifications        bool                                                      `json:"uiNotifications,omitempty"`
	EmailNotifications     bool                                                      `json:"emailNotifications,omitempty"`
	ExtensionAttributes    []PatchSoftwareTitleConfigurationSubsetExtensionAttribute `json:"extensionAttributes,omitempty"`
	SoftwareTitleName      string                                                    `json:"softwareTitleName,omitempty"`
	SoftwareTitleNameId    string                                                    `json:"softwareTitleNameId,omitempty"`
	SoftwareTitlePublisher string                                                    `json:"softwareTitlePublisher,omitempty"`
	JamfOfficial           bool                                                      `json:"jamfOfficial,omitempty"`
	PatchSourceName        string                                                    `json:"patchSourceName,omitempty"`
	PatchSourceEnabled     bool                                                      `json:"patchSourceEnabled,omitempty"`
	Packages               []PatchSoftwareTitleConfigurationSubsetPackage            `json:"packages,omitempty"`
}

// Resource struct for a patch software title definition
type ResourcePatchSoftwareTitleDefinition struct {
	Version                string                                      `json:"version"`
	MinimumOperatingSystem string                                      `json:"minimumOperatingSystem"`
	ReleaseDate            string                                      `json:"releaseDate"`
	RebootRequired         bool                                        `json:"rebootRequired"`
	KillApps               []PatchSoftwareTitleDefinitionSubsetKillApp `json:"killApps"`
	Standalone             bool                                        `json:"standalone"`
	AbsoluteOrderID        string                                      `json:"absoluteOrderId"`
}

// Resource struct for patch software title dependency
type ResourcePatchSoftwareTitleDependency struct {
	SmartGroupID   string `json:"smartGroupId"`
	SmartGroupName string `json:"smartGroupName"`
}

// Resource struct for patch software title extension attribute
type ResourcePatchSoftwareTitleExtensionAttribute struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
}

// Subsets

type PatchSoftwareTitleConfigurationSubsetExtensionAttribute struct {
	Accepted bool   `json:"accepted,omitempty"`
	EaID     string `json:"eaId,omitempty"`
}

type PatchSoftwareTitleConfigurationSubsetPackage struct {
	PackageId   string `json:"packageId,omitempty"`
	Version     string `json:"version,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}

type PatchSoftwareTitleDefinitionSubsetKillApp struct {
	AppName string `json:"appName"`
}

// Response

type ResponsePatchSoftwareTitleConfigurationCreate struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// Response struct for patch software title definitions list
type ResponsePatchSoftwareTitleDefinitionsList struct {
	TotalCount int                                    `json:"totalCount"`
	Results    []ResourcePatchSoftwareTitleDefinition `json:"results"`
}

// Response struct for patch software title configuration dependencies
type ResponsePatchSoftwareTitleDependenciesList struct {
	TotalCount int                                    `json:"totalCount"`
	Results    []ResourcePatchSoftwareTitleDependency `json:"results"`
}

// ResponsePatchSoftwareTitleExtensionAttributesList represents a list of extension attributes
type ResponsePatchSoftwareTitleExtensionAttributesList []ResourcePatchSoftwareTitleExtensionAttribute

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

// GetPatchSoftwareTitleDefinitions retrieves patch software title definitions with the supplied id
func (c *Client) GetPatchSoftwareTitleDefinitions(id string, params url.Values) (*ResponsePatchSoftwareTitleDefinitionsList, error) {
	if id == "" {
		return nil, fmt.Errorf("patch software title id cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/%s/definitions", uriPatchSoftwareTitleConfigurations, id)

	resp, err := c.DoPaginatedGet(endpoint, params)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "patch software title definitions", err)
	}

	var out ResponsePatchSoftwareTitleDefinitionsList
	out.TotalCount = resp.Size

	for _, value := range resp.Results {
		var newObj ResourcePatchSoftwareTitleDefinition
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "patch software title definition", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil
}

// GetPatchSoftwareTitleDependencies retrieves list of dependencies for a patch software title configuration
func (c *Client) GetPatchSoftwareTitleDependencies(id string) (*ResponsePatchSoftwareTitleDependenciesList, error) {
	if id == "" {
		return nil, fmt.Errorf("patch software title configuration id cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/%s/dependencies", uriPatchSoftwareTitleConfigurations, id)

	var out ResponsePatchSoftwareTitleDependenciesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "patch software title dependencies", err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// GetPatchSoftwareTitleExtensionAttributes retrieves extension attributes for a patch software title configuration
func (c *Client) GetPatchSoftwareTitleExtensionAttributes(id string) (*ResponsePatchSoftwareTitleExtensionAttributesList, error) {
	if id == "" {
		return nil, fmt.Errorf("patch software title id cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/%s/extension-attributes", uriPatchSoftwareTitleConfigurations, id)

	var out ResponsePatchSoftwareTitleExtensionAttributesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "patch software title extension attributes", err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &out, nil
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

// UpdatePatchSoftwareTitleConfigurationById updates a single PatchSoftwareTitleConfiguration with the given ID
func (c *Client) UpdatePatchSoftwareTitleConfigurationById(id string, updatedConfiguration ResourcePatchSoftwareTitleConfiguration) (*ResourcePatchSoftwareTitleConfiguration, error) {
	endpoint := fmt.Sprintf("%s/%s", uriPatchSoftwareTitleConfigurations, id)

	var out ResourcePatchSoftwareTitleConfiguration

	resp, err := c.HTTP.DoRequest("PATCH", endpoint, updatedConfiguration, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to update patch software title configuration: %v", err)
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
