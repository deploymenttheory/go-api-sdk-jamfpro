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

// Resource

type ResourceAvailableUpdates struct {
	MacOS []string `json:"macOS"`
	IOS   []string `json:"iOS"`
}

// ResourceManagedSoftwareUpdatePlan represents the payload structure for creating a managed software update plan.
type ResourceManagedSoftwareUpdatePlan struct {
	Devices []ManagedSoftwareUpdatePlanDevice `json:"devices"`
	Config  ManagedSoftwareUpdatePlanConfig   `json:"config"`
}

// ManagedSoftwareUpdatePlanDevice defines the structure for device objects in the managed software update plan.
type ManagedSoftwareUpdatePlanDevice struct {
	ObjectType string `json:"objectType"`
	DeviceId   string `json:"deviceId"`
}

// ManagedSoftwareUpdatePlanConfig defines the configuration for a managed software update plan.
type ManagedSoftwareUpdatePlanConfig struct {
	UpdateAction              string `json:"updateAction"`
	VersionType               string `json:"versionType"`
	SpecificVersion           string `json:"specificVersion,omitempty"` // omitempty allows this field to be omitted if empty
	MaxDeferrals              int    `json:"maxDeferrals"`
	ForceInstallLocalDateTime string `json:"forceInstallLocalDateTime"`
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

// Creates Managed Software Update Plan from ResourceManagedSoftwareUpdatePlan struct
func (c *Client) CreateManagedSoftwareUpdatePlan(plan *ResourceManagedSoftwareUpdatePlan) (*ResponseManagedSoftwareUpdatePlanCreate, error) {
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
