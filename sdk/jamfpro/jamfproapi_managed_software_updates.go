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
	TotalCount int                                 `json:"totalCount"`
	Results    []ResourceManagedSoftwareUpdatePlan `json:"results"`
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

type ResourceManagedSoftwareUpdatePlan struct {
	PlanUuid                  string                                `json:"planUuid"`
	Device                    ManagedSoftwareUpdatePlanSubsetDevice `json:"device"`
	UpdateAction              string                                `json:"updateAction"`
	VersionType               string                                `json:"versionType"`
	SpecificVersion           string                                `json:"specificVersion"`
	MaxDeferrals              int                                   `json:"maxDeferrals"`
	ForceInstallLocalDateTime string                                `json:"forceInstallLocalDateTime"`
	RecipeId                  string                                `json:"recipeId"`
	Status                    ManagedSoftwareUpdatePlanSubsetStatus `json:"status"`
}

type ManagedSoftwareUpdatePlanSubsetDevice struct {
	DeviceId   string `json:"deviceId"`
	ObjectType string `json:"objectType"`
	Href       string `json:"href"`
}

type ManagedSoftwareUpdatePlanSubsetStatus struct {
	State        string   `json:"state"`
	ErrorReasons []string `json:"errorReasons"`
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
		var newObj ResourceManagedSoftwareUpdatePlan
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
	var ResponseManagedSoftwareUpdatePlanCreate ResponseManagedSoftwareUpdatePlanCreate

	resp, err := c.HTTP.DoRequest("POST", endpoint, plan, &ResponseManagedSoftwareUpdatePlanCreate)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "script", err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &ResponseManagedSoftwareUpdatePlanCreate, nil
}
