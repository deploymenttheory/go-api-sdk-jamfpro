// jamfproapi_managed_software_updates.go
// Jamf Pro Api - Managed Software Updates (BETA)
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

// Response

type ResponseManagedSoftwareUpdatePlanList struct {
	TotalCount int                                 `json:"totalCount"`
	Results    []ResponseManagedSoftwareUpdatePlan `json:"results"`
}

type ResponseManagedSoftwareUpdatePlan struct {
	PlanUuid                  string                                        `json:"planUuid,omitempty"`
	Device                    ResponseManagedSoftwareUpdatePlanSubsetDevice `json:"device,omitempty"`
	UpdateAction              string                                        `json:"updateAction,omitempty"`
	VersionType               string                                        `json:"versionType,omitempty"`
	SpecificVersion           string                                        `json:"specificVersion,omitempty"`
	BuildVersion              string                                        `json:"buildVersion,omitempty"`
	MaxDeferrals              int                                           `json:"maxDeferrals,omitempty"`
	ForceInstallLocalDateTime string                                        `json:"forceInstallLocalDateTime,omitempty"`
	RecipeId                  string                                        `json:"recipeId,omitempty"`
	Status                    ResponseManagedSoftwareUpdatePlanSubsetStatus `json:"status,omitempty"`
}

type ResponseManagedSoftwareUpdatePlanSubsetDevice struct {
	DeviceId   string `json:"deviceId,omitempty"`
	ObjectType string `json:"objectType,omitempty"`
	Href       string `json:"href,omitempty"`
}

type ResponseManagedSoftwareUpdatePlanSubsetStatus struct {
	State        string   `json:"state,omitempty"`
	ErrorReasons []string `json:"errorReasons"`
}

// ResponseDeclarationsList represents the response structure for the list of declarations.
type ResponseDeclarationsList struct {
	Declarations []ResourceDeclaration `json:"declarations"`
}

// ResourceDeclaration represents the structure of a single declaration associated with a managed software update plan.
type ResourceDeclaration struct {
	UUID        string `json:"uuid"`
	PayloadJson string `json:"payloadJson"`
	Type        string `json:"type"`
	Group       string `json:"group"`
}

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

// ResponseManagedSoftwareUpdatePlansFeatureToggleStatus represents the response structure for the feature toggle status.
type ResponseManagedSoftwareUpdatePlansFeatureToggleStatus struct {
	ToggleOn  *FeatureEnablementToggleStatus `json:"toggleOn"`
	ToggleOff *FeatureEnablementToggleStatus `json:"toggleOff"`
}

// FeatureToggleStatusDetail represents the detailed status of the feature toggle (on/off).
type FeatureEnablementToggleStatus struct {
	StartTime                string  `json:"startTime"`
	EndTime                  string  `json:"endTime"`
	ElapsedTime              int     `json:"elapsedTime"`
	State                    string  `json:"state"`
	TotalRecords             int64   `json:"totalRecords"`
	ProcessedRecords         int64   `json:"processedRecords"`
	PercentComplete          float64 `json:"percentComplete"`
	FormattedPercentComplete string  `json:"formattedPercentComplete"`
	ExitState                string  `json:"exitState"`
	ExitMessage              string  `json:"exitMessage"`
}

// Resource

type ResourceAvailableUpdates struct {
	MacOS []string `json:"macOS"`
	IOS   []string `json:"iOS"`
}

// ResourceManagedSoftwareUpdatePlan represents the payload structure for creating a managed software update plan.
type ResourceManagedSoftwareUpdatePlan struct {
	Devices []ResourcManagedSoftwareUpdatePlanObject `json:"devices,omitempty"`
	Group   ResourcManagedSoftwareUpdatePlanObject   `json:"group,omitempty"`
	Config  ResourcManagedSoftwareUpdatePlanConfig   `json:"config,omitempty"`
}

// ResourcManagedSoftwareUpdatePlanDevice defines the structure for device objects in the managed software update plan.
type ResourcManagedSoftwareUpdatePlanObject struct {
	ObjectType string `json:"objectType"`
	DeviceId   string `json:"deviceId,omitempty"`
	GroupId    string `json:"groupId,omitempty"`
}

// ResourcManagedSoftwareUpdatePlanConfig defines the configuration for a managed software update plan.
type ResourcManagedSoftwareUpdatePlanConfig struct {
	UpdateAction              string `json:"updateAction"`
	VersionType               string `json:"versionType"`
	SpecificVersion           string `json:"specificVersion,omitempty"`
	BuildVersion              string `json:"buildVersion,omitempty"`
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
		var newObj ResponseManagedSoftwareUpdatePlan
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
	var response ResponseManagedSoftwareUpdateFeatureToggle

	resp, err := c.HTTP.DoRequest("PUT", endpoint, payload, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to update managed software update feature toggle: %v", err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// ForceStopManagedSoftwareUpdateFeatureToggleProcess forcefully stops any ongoing or stalled feature-toggle processes.
// This "Break Glass" endpoint should not be used under nominal conditions.
func (c *Client) ForceStopManagedSoftwareUpdateFeatureToggleProcess() (*SharedResourcResponseError, error) {
	endpoint := fmt.Sprintf("%s/plans/feature-toggle/abandon", uriManagedSoftwareUpdates)
	var responseError SharedResourcResponseError

	resp, err := c.HTTP.DoRequest("POST", endpoint, nil, &responseError)
	if err != nil {
		return nil, fmt.Errorf("failed to forcefully abandon feature-toggle process: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseError, nil
}

// GetManagedSoftwareUpdatePlansFeatureToggleStatus retrieves the background status of the Feature Toggle.
func (c *Client) GetManagedSoftwareUpdatePlansFeatureToggleStatus() (*ResponseManagedSoftwareUpdatePlansFeatureToggleStatus, error) {
	endpoint := fmt.Sprintf("%s/plans/feature-toggle/status", uriManagedSoftwareUpdates)

	var status ResponseManagedSoftwareUpdatePlansFeatureToggleStatus
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &status)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve feature toggle status: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &status, nil
}

// CreateManagedSoftwareUpdatePlanByDeviceGroupID creates a managed software update plan by group ID
func (c *Client) CreateManagedSoftwareUpdatePlanByGroupID(plan *ResourceManagedSoftwareUpdatePlan) (*ResponseManagedSoftwareUpdatePlanCreate, error) {
	endpoint := uriManagedSoftwareUpdates + "/plans/group"
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

// GetManagedSoftwareUpdatePlansByGroupID retrieves managed software update plans for a specific group ID.
func (c *Client) GetManagedSoftwareUpdatePlansByGroupID(groupId string, groupType string) (*ResponseManagedSoftwareUpdatePlanList, error) {
	endpoint := fmt.Sprintf("%s/plans/group/%s?group-type=%s", uriManagedSoftwareUpdates, groupId, groupType)

	var responseManagedSoftwareUpdatePlanList ResponseManagedSoftwareUpdatePlanList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &responseManagedSoftwareUpdatePlanList)
	if err != nil {
		return nil, fmt.Errorf("failed to get managed software update plans: %v", err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &responseManagedSoftwareUpdatePlanList, nil
}

// GetManagedSoftwareUpdatePlanByUUID retrieves a Managed Software Update Plan by its UUID.
func (c *Client) GetManagedSoftwareUpdatePlanByUUID(UUID string) (*ResponseManagedSoftwareUpdatePlan, error) {
	endpoint := fmt.Sprintf("%s/plans/%s", uriManagedSoftwareUpdates, UUID)

	var planDetail ResponseManagedSoftwareUpdatePlan
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &planDetail)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve managed software update plan with ID %s: %v", UUID, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &planDetail, nil
}

// GetDeclarationsByManagedSoftwareUpdatePlanUUID retrieves all Declarations associated with a Managed Software Update Plan by its UUID.
func (c *Client) GetDeclarationsByManagedSoftwareUpdatePlanUUID(UUID string) (*ResponseDeclarationsList, error) {
	endpoint := fmt.Sprintf("%s/plans/%s/declarations", uriManagedSoftwareUpdates, UUID)

	var declarationsList ResponseDeclarationsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &declarationsList)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve declarations for managed software update plan with ID %s: %v", UUID, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &declarationsList, nil
}
