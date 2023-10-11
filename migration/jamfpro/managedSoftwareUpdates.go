// managedSoftwareUpdates.go
// Jamf Pro API
// Jamf Pro API requires the structs to support JSON.
// Introduced in Jamf Pro 10.50
// WIP - Currently Incomplete

package jamfpro

import (
	"fmt"
	"strings"
)

const uriManagedSoftwareUpdates = "/api/v1/managed-software-updates"

// Struct for ResponseManagedSoftwareUpdates
type ResponseManagedSoftwareUpdates struct {
	AvailableUpdates *ManagedSoftwareUpdates `json:"availableUpdates,omitempty"`
}

type ManagedSoftwareUpdates struct {
	MacOS []string `json:"macOS,omitempty"`
	IOS   []string `json:"iOS,omitempty"`
}

type ResponseManagedSoftwareUpdatePlans struct {
	TotalCount int                  `json:"totalCount"`
	Results    []SoftwareUpdatePlan `json:"results"`
}

type SoftwareUpdatePlan struct {
	PlanUUID                  string           `json:"planUuid"`
	Device                    DeviceDetail     `json:"device"`
	UpdateAction              string           `json:"updateAction"`
	VersionType               string           `json:"versionType"`
	SpecificVersion           string           `json:"specificVersion,omitempty"`
	MaxDeferrals              int              `json:"maxDeferrals"`
	ForceInstallLocalDateTime string           `json:"forceInstallLocalDateTime,omitempty"`
	Status                    UpdatePlanStatus `json:"status"`
}

type DeviceDetail struct {
	DeviceID   string `json:"deviceId"`
	ObjectType string `json:"objectType,omitempty"`
	Href       string `json:"href,omitempty"`
}

type UpdatePlanStatus struct {
	State        string   `json:"state,omitempty"`
	ErrorReasons []string `json:"errorReasons,omitempty"`
}

// Define a struct for the query parameters
type ManagedSoftwareUpdateQueryParams struct {
	Page     int      `json:"page"`
	PageSize int      `json:"page-size"`
	Sort     []string `json:"sort"`
	Filter   string   `json:"filter"`
}

// Response for creating managed software update plans
type ResponseCreateManagedSoftwareUpdatePlan struct {
	Plans []CreatedPlan `json:"plans"`
}

type CreatedPlan struct {
	Device DeviceDetail `json:"device"`
	PlanID string       `json:"planId"`
	Href   string       `json:"href"`
}

// Body parameters for creating managed software update plans
type CreateManagedSoftwareUpdatePlanParams struct {
	Devices []DeviceForPlan `json:"devices"`
	Config  PlanConfig      `json:"config"`
}

type DeviceForPlan struct {
	DeviceID   string `json:"deviceId"`
	ObjectType string `json:"objectType"`
}

type PlanConfig struct {
	UpdateAction              string `json:"updateAction"`
	VersionType               string `json:"versionType"`
	SpecificVersion           string `json:"specificVersion,omitempty"`
	MaxDeferrals              int    `json:"maxDeferrals"`
	ForceInstallLocalDateTime string `json:"forceInstallLocalDateTime,omitempty"`
}

type ResponseFeatureToggleStatus struct {
	Toggle                       bool `json:"toggle"`
	ForceInstallLocalDateEnabled bool `json:"forceInstallLocalDateEnabled,omitempty"`
	DssEnabled                   bool `json:"dssEnabled,omitempty"`
}

type ManagedSoftwareUpdateDeviceGroup struct {
	GroupID    string `json:"groupId"`
	ObjectType string `json:"objectType"`
}

type ManagedSoftwareUpdateDeviceGroupItem struct {
	Group  ManagedSoftwareUpdateDeviceGroup `json:"group"`
	Config PlanConfig                       `json:"config"`
}

type ResponseManagedSoftwareUpdatePlansForDeviceGroup struct {
	TotalCount int                                `json:"totalCount"`
	Results    []SoftwareUpdatePlanForDeviceGroup `json:"results"`
}

type SoftwareUpdatePlanForDeviceGroup struct {
	PlanUUID                  string           `json:"planUuid"`
	Device                    DeviceDetail     `json:"device"`
	UpdateAction              string           `json:"updateAction"`
	VersionType               string           `json:"versionType"`
	SpecificVersion           string           `json:"specificVersion,omitempty"`
	MaxDeferrals              int              `json:"maxDeferrals"`
	ForceInstallLocalDateTime string           `json:"forceInstallLocalDateTime,omitempty"`
	Status                    UpdatePlanStatus `json:"status"`
}

// Functions

// GetAvailableManagedSoftwareUpdates retrieves the available managed software updates for macOS and iOS
func (c *Client) GetAvailableManagedSoftwareUpdates() (*ResponseManagedSoftwareUpdates, error) {
	url := fmt.Sprintf("%s/available-updates", uriManagedSoftwareUpdates)
	var availableUpdates ResponseManagedSoftwareUpdates
	if err := c.DoRequest("GET", url, nil, nil, &availableUpdates); err != nil {
		return nil, fmt.Errorf("failed to get available managed software updates: %v", err)
	}

	return &availableUpdates, nil
}

// ManagedSoftwareUpdateQueryParams defines the possible query parameters for fetching managed software update plans.
func (c *Client) GetManagedSoftwareUpdatePlans(params ManagedSoftwareUpdateQueryParams) (*ResponseManagedSoftwareUpdatePlans, error) {
	// Construct the base URL
	url := fmt.Sprintf("%s/plans/0", uriManagedSoftwareUpdates)

	// Add query parameters to the URL
	queryParams := make([]string, 0)

	if params.Page != 0 {
		queryParams = append(queryParams, fmt.Sprintf("page=%d", params.Page))
	}
	if params.PageSize != 0 {
		queryParams = append(queryParams, fmt.Sprintf("page-size=%d", params.PageSize))
	}
	if len(params.Sort) > 0 {
		for _, sortParam := range params.Sort {
			queryParams = append(queryParams, fmt.Sprintf("sort=%s", sortParam))
		}
	}
	if params.Filter != "" {
		queryParams = append(queryParams, fmt.Sprintf("filter=%s", params.Filter))
	}

	if len(queryParams) > 0 {
		url = fmt.Sprintf("%s?%s", url, strings.Join(queryParams, "&"))
	}

	var updatePlans ResponseManagedSoftwareUpdatePlans
	if err := c.DoRequest("GET", url, nil, nil, &updatePlans); err != nil {
		return nil, fmt.Errorf("failed to get managed software update plans: %v", err)
	}

	return &updatePlans, nil
}

// CreateManagedSoftwareUpdatePlan creates a new software update plan
func (c *Client) CreateManagedSoftwareUpdatePlan(params CreateManagedSoftwareUpdatePlanParams) (*ResponseCreateManagedSoftwareUpdatePlan, error) {
	url := fmt.Sprintf("%s/plans", uriManagedSoftwareUpdates)

	var managedSoftwareUpdatePlan ResponseCreateManagedSoftwareUpdatePlan
	if err := c.DoRequest("POST", url, params, nil, &managedSoftwareUpdatePlan); err != nil {
		return nil, fmt.Errorf("failed to create managed software update plan: %v", err)
	}

	return &managedSoftwareUpdatePlan, nil
}

// GetManagedSoftwareUpdateFeatureToggleStatus retrieves the status of the feature toggle
func (c *Client) GetManagedSoftwareUpdateFeatureStatus() (*ResponseFeatureToggleStatus, error) {
	url := fmt.Sprintf("%s/plans/feature-toggle", uriManagedSoftwareUpdates)

	var toggleStatus ResponseFeatureToggleStatus
	if err := c.DoRequest("GET", url, nil, nil, &toggleStatus); err != nil {
		return nil, fmt.Errorf("failed to get feature toggle status: %v", err)
	}

	return &toggleStatus, nil
}

// UpdateManagedSoftwareUpdateFeatureStatus updates the status of the feature toggle
func (c *Client) UpdateManagedSoftwareUpdateFeatureStatus(statusToUpdate *ResponseFeatureToggleStatus) (*ResponseFeatureToggleStatus, error) {
	url := fmt.Sprintf("%s/plans/feature-toggle", uriManagedSoftwareUpdates)

	var updatedFeatureToggleStatus ResponseFeatureToggleStatus
	if err := c.DoRequest("PUT", url, statusToUpdate, nil, &updatedFeatureToggleStatus); err != nil {
		return nil, fmt.Errorf("failed to update feature toggle status: %v", err)
	}

	return &updatedFeatureToggleStatus, nil
}

// CreateManagedSoftwareUpdatePlanForDeviceGroup creates a software update plan for a device group
func (c *Client) CreateManagedSoftwareUpdatePlanForDeviceGroup(params ManagedSoftwareUpdateDeviceGroupItem) (*ResponseCreateManagedSoftwareUpdatePlan, error) {
	url := fmt.Sprintf("%s/plans/group", uriManagedSoftwareUpdates)

	var managedSoftwareUpdatePlan ResponseCreateManagedSoftwareUpdatePlan
	if err := c.DoRequest("POST", url, params, nil, &managedSoftwareUpdatePlan); err != nil {
		return nil, fmt.Errorf("failed to create managed software update plan for device group: %v", err)
	}

	return &managedSoftwareUpdatePlan, nil
}

// GetManagedSoftwareUpdatePlansForDeviceGroup retrieves the managed software update plans for a specific device group
func (c *Client) GetManagedSoftwareUpdatePlansForDeviceGroup(id string, groupType string) (*ResponseManagedSoftwareUpdatePlansForDeviceGroup, error) {
	url := fmt.Sprintf("%s/plans/group/%s?group-type=%s", uriManagedSoftwareUpdates, id, groupType)

	var plans ResponseManagedSoftwareUpdatePlansForDeviceGroup
	if err := c.DoRequest("GET", url, nil, nil, &plans); err != nil {
		return nil, fmt.Errorf("failed to get managed software update plans for device group: %v", err)
	}

	return &plans, nil
}
