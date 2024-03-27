// jamfproapi_managed_software_updates.go
// Jamf Pro Api - Managed Software Updates
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-available-updates
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

const uriManagedSoftwareUpdates = "/api/v1/managed-software-updates"

// Structs

// List

type ResponseManagedSoftwareUpdateList struct {
	AvailableUpdates ResourceAvailableUpdates `json:"availableUpdates"`
}

type ResponseManagedSoftwareUpdatePlanList struct {
	TotalCount int                                     `json:"totalCount"`
	Results    []ResourceManagedSoftwareUpdatePlanList `json:"results"`
}

type ResourceManagedSoftwareUpdatePlanList struct {
	PlanUuid                  string                                    `json:"planUuid,omitempty"`
	Device                    ManagedSoftwareUpdatePlanListSubsetDevice `json:"device,omitempty"`
	UpdateAction              string                                    `json:"updateAction,omitempty"`
	VersionType               string                                    `json:"versionType,omitempty"`
	SpecificVersion           string                                    `json:"specificVersion,omitempty"`
	MaxDeferrals              int                                       `json:"maxDeferrals,omitempty"`
	ForceInstallLocalDateTime string                                    `json:"forceInstallLocalDateTime,omitempty"`
	RecipeId                  string                                    `json:"recipeId,omitempty"`
	Status                    ManagedSoftwareUpdatePlanListSubsetStatus `json:"status,omitempty"`
}

type ManagedSoftwareUpdatePlanListSubsetDevice struct {
	DeviceId   string `json:"deviceId,omitempty"`
	ObjectType string `json:"objectType,omitempty"`
	Href       string `json:"href,omitempty"`
}

type ManagedSoftwareUpdatePlanListSubsetStatus struct {
	State        string   `json:"state,omitempty"`
	ErrorReasons []string `json:"errorReasons"`
}

// Response

type ResponseManagedSoftwareUpdatePlanCreate struct {
	Plans []ManagedSoftwareUpdatePlanCreateSubsetPlan `json:"plans"`
}

type ManagedSoftwareUpdatePlanCreateSubsetPlan struct {
	Device ManagedSoftwareUpdatePlanCreateSubsetDevice `json:"device"`
	PlanID string                                      `json:"planId"`
	Href   string                                      `json:"href"`
}

type ManagedSoftwareUpdatePlanCreateSubsetDevice struct {
	DeviceID   string `json:"deviceId"`
	ObjectType string `json:"objectType"`
	Href       string `json:"href"`
}

type ResponseManagedSoftwareUpdateFeatureToggle struct {
	Toggle                       bool `json:"toggle"`
	ForceInstallLocalDateEnabled bool `json:"forceInstallLocalDateEnabled"`
	DssEnabled                   bool `json:"dssEnabled"`
	RecipeEnabled                bool `json:"recipeEnabled"`
}

// Resource

type ResourceAvailableUpdates struct {
	MacOS []string `json:"macOS"`
	IOS   []string `json:"iOS"`
}

// ResourceManagedSoftwareUpdatePlan represents the payload structure for creating a managed software update plan.
type ResourceManagedSoftwareUpdatePlan struct {
	Devices []ManagedSoftwareUpdatePlanDevice `json:"devices,omitempty"`
	Config  ManagedSoftwareUpdatePlanConfig   `json:"config,omitempty"`
}

// ManagedSoftwareUpdatePlanDevice defines the structure for device objects in the managed software update plan.
type ManagedSoftwareUpdatePlanDevice struct {
	ObjectType string `json:"objectType,omitempty"`
	DeviceId   string `json:"deviceId,omitempty"`
}

// ManagedSoftwareUpdatePlanConfig defines the configuration for a managed software update plan.
type ManagedSoftwareUpdatePlanConfig struct {
	UpdateAction              string `json:"updateAction,omitempty"`
	VersionType               string `json:"versionType,omitempty"`
	SpecificVersion           string `json:"specificVersion,omitempty"`
	MaxDeferrals              int    `json:"maxDeferrals,omitempty"`
	ForceInstallLocalDateTime string `json:"forceInstallLocalDateTime,omitempty"`
}

// ResourceManagedSoftwareUpdateFeatureToggle represents the payload for updating the feature toggle.
type ResourceManagedSoftwareUpdateFeatureToggle struct {
	Toggle bool `json:"toggle"`
}

// CRUD

// GetManagedSoftwareUpdates retrieves a list of all available managed software updates
func (c *Client) GetManagedSoftwareUpdates() (*ResponseManagedSoftwareUpdateList, error) {
	endpoint := fmt.Sprintf("%s/available-updates", uriManagedSoftwareUpdates)

	var updateList ResponseManagedSoftwareUpdateList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &updateList)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "self service settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updateList, nil
}

// GetManagedSoftwareUpdatePlans retrieves a list of all available managed software updates
func (c *Client) GetManagedSoftwareUpdatePlans(sort_filter string) (*ResponseManagedSoftwareUpdatePlanList, error) {
	resp, err := c.DoPaginatedGet(
		uriManagedSoftwareUpdates+"/plans",
		standardPageSize,
		startingPageNumber,
		sort_filter,
	)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "managed software update plans", err)
	}

	var out ResponseManagedSoftwareUpdatePlanList
	out.TotalCount = resp.Size

	for _, value := range resp.Results {
		var newObj ResourceManagedSoftwareUpdatePlanList
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "script", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil

}

// CreateManagedSoftwareUpdatePlanByDeviceID Creates Managed Software Update Plan by Device ID
func (c *Client) CreateManagedSoftwareUpdatePlanByDeviceID(plan *ResourceManagedSoftwareUpdatePlan) (*ResponseManagedSoftwareUpdatePlanCreate, error) {
	endpoint := uriManagedSoftwareUpdates + "/plans"
	var responseManagedSoftwareUpdatePlanCreate ResponseManagedSoftwareUpdatePlanCreate

	resp, err := c.HTTP.DoRequest("POST", endpoint, plan, &responseManagedSoftwareUpdatePlanCreate)
	if err != nil {
		return nil, fmt.Errorf("failed to create managed software update plan: %v", err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &responseManagedSoftwareUpdatePlanCreate, nil
}

// GetManagedSoftwareUpdateFeatureToggle retrieves the current managed software update feature toggle settings
func (c *Client) GetManagedSoftwareUpdateFeatureToggle() (*ResourceManagedSoftwareUpdateFeatureToggle, error) {
	endpoint := fmt.Sprintf("%s/plans/feature-toggle", uriManagedSoftwareUpdates)

	var featureToggle ResourceManagedSoftwareUpdateFeatureToggle
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &featureToggle)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "managed software update feature toggle", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &featureToggle, nil
}

// UpdateManagedSoftwareUpdateFeatureToggle updates the feature toggle for managed software updates.
func (c *Client) UpdateManagedSoftwareUpdateFeatureToggle(payload *ResourceManagedSoftwareUpdateFeatureToggle) (*ResponseManagedSoftwareUpdateFeatureToggle, error) {
	endpoint := fmt.Sprintf("%s/plans/feature-toggle", uriManagedSoftwareUpdates)

	// Define a variable to hold the response
	var response ResponseManagedSoftwareUpdateFeatureToggle

	// Perform the request and unmarshal the response
	resp, err := c.HTTP.DoRequest("PUT", endpoint, payload, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to update managed software update feature toggle: %v", err)
	}

	// Ensure the response body gets closed
	if resp != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}
