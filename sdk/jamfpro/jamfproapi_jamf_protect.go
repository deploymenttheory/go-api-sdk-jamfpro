// jamfproapi_jamf_protect.go
package jamfpro

import (
	"fmt"
	"net/url"

	"github.com/mitchellh/mapstructure"
)

const (
	uriJamfProtect            = "/api/v1/jamf-protect"
	uriJamfProtectSync        = "/api/v1/jamf-protect/plans/sync"
	uriJamfProtectRegister    = "/api/v1/jamf-protect/register"
	uriJamfProtectDeployments = "/api/v1/jamf-protect/deployments"
	uriJamfProtectHistory     = "/api/v1/jamf-protect/history"
	uriJamfProtectPlans       = "/api/v1/jamf-protect/plans"
)

// Response Structs
type ResponseJamfProtectSettings struct {
	ID             string `json:"id"`
	ProtectURL     string `json:"protectUrl"`
	SyncStatus     string `json:"syncStatus"`
	APIClientID    string `json:"apiClientId"`
	AutoInstall    bool   `json:"autoInstall"`
	LastSyncTime   string `json:"lastSyncTime"`
	APIClientName  string `json:"apiClientName"`
	RegistrationID string `json:"registrationId"`
}

type ResponseJamfProtectDeploymentTasksList struct {
	Results    []ResourceJamfProtectDeploymentTask `json:"results"`
	TotalCount int                                 `json:"totalCount"`
}

type ResponseJamfProtectHistoryList struct {
	TotalCount int                          `json:"totalCount"`
	Results    []ResourceJamfProtectHistory `json:"results"`
}

type ResponseJamfProtectPlansList struct {
	TotalCount int                       `json:"totalCount"`
	Results    []ResourceJamfProtectPlan `json:"results"`
}

// Resource Structs
type ResourceJamfProtectDeploymentTask struct {
	ID           string `json:"id"`
	Status       string `json:"status"`
	Updated      string `json:"updated"`
	Version      string `json:"version"`
	ComputerID   string `json:"computerId"`
	ComputerName string `json:"computerName"`
}

type ResourceJamfProtectHistory struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Date     string `json:"date"`
	Note     string `json:"note"`
	Details  string `json:"details"`
}

type ResourceJamfProtectPlan struct {
	UUID             string `json:"uuid"`
	ID               string `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	ProfileID        int    `json:"profileId"`
	ProfileName      string `json:"profileName"`
	ScopeDescription string `json:"scopeDescription"`
}

// Request Structs
type ResourceJamfProtectRegistration struct {
	ProtectURL string `json:"protectUrl"`
	ClientID   string `json:"clientId"`
	Password   string `json:"password"`
}

type ResourceJamfProtectSettings struct {
	AutoInstall bool `json:"autoInstall"`
}

type ResourceJamfProtectDeploymentRequest struct {
	PlanID          string   `json:"planId"`
	TargetComputers []string `json:"targetComputers"`
}

type ResourceJamfProtectHistoryNote struct {
	Note    string `json:"note"`
	Details string `json:"details"`
}

// CRUD Operations

// GetJamfProtectSettings retrieves the current Jamf Protect integration settings
func (c *Client) GetJamfProtectSettings() (*ResponseJamfProtectSettings, error) {
	endpoint := uriJamfProtect
	var out ResponseJamfProtectSettings

	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "jamf protect settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// UpdateJamfProtectSettings updates the Jamf Protect integration settings
func (c *Client) UpdateJamfProtectSettings(settings ResourceJamfProtectSettings) (*ResponseJamfProtectSettings, error) {
	endpoint := uriJamfProtect
	var out ResponseJamfProtectSettings

	resp, err := c.HTTP.DoRequest("PUT", endpoint, settings, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdate, "jamf protect settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// RegisterJamfProtect registers a new Jamf Protect integration
func (c *Client) RegisterJamfProtect(registration ResourceJamfProtectRegistration) (*ResponseJamfProtectSettings, error) {
	endpoint := uriJamfProtectRegister
	var out ResponseJamfProtectSettings

	resp, err := c.HTTP.DoRequest("POST", endpoint, registration, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "jamf protect registration", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// SyncJamfProtectPlans initiates a sync of Jamf Protect plans
func (c *Client) SyncJamfProtectPlans() error {
	endpoint := uriJamfProtectSync

	resp, _ := c.HTTP.DoRequest("POST", endpoint, nil, nil)

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to sync Jamf Protect plans: unexpected status code %d", resp.StatusCode)
	}

	return nil
}

// CreateJamfProtectIntegration creates a complete Jamf Protect integration
func (c *Client) CreateJamfProtectIntegration(registration ResourceJamfProtectRegistration, autoInstall bool) (*ResponseJamfProtectSettings, error) {
	_, err := c.RegisterJamfProtect(registration)
	if err != nil {
		return nil, fmt.Errorf("failed to register Jamf Protect: %v", err)
	}

	settings := ResourceJamfProtectSettings{
		AutoInstall: autoInstall,
	}

	updateResp, err := c.UpdateJamfProtectSettings(settings)
	if err != nil {
		return nil, fmt.Errorf("failed to update settings: %v", err)
	}

	err = c.SyncJamfProtectPlans()
	if err != nil {
		return nil, fmt.Errorf("failed to sync plans: %v", err)
	}

	return updateResp, nil
}

// GetJamfProtectDeploymentTasks retrieves a list of tasks for a specific deployment
func (c *Client) GetJamfProtectDeploymentTasks(deploymentID string, params url.Values) (*ResponseJamfProtectDeploymentTasksList, error) {
	endpoint := fmt.Sprintf("%s/%s/tasks", uriJamfProtectDeployments, deploymentID)
	var out ResponseJamfProtectDeploymentTasksList

	resp, err := c.DoPaginatedGet(endpoint, params)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "jamf protect deployment tasks", err)
	}

	out.TotalCount = resp.Size

	for _, value := range resp.Results {
		var newObj ResourceJamfProtectDeploymentTask
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "jamf protect deployment task", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil
}

// RetryJamfProtectDeploymentTasks requests a retry of Protect install tasks for a specific deployment
func (c *Client) RetryJamfProtectDeploymentTasks(deploymentID string) error {
	endpoint := fmt.Sprintf("%s/%s/tasks/retry", uriJamfProtectDeployments, deploymentID)

	resp, err := c.HTTP.DoRequest("POST", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedCreate, "retry jamf protect deployment tasks", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to retry jamf protect deployment tasks: unexpected status code %d", resp.StatusCode)
	}

	return nil
}

// GetJamfProtectHistory retrieves the history of Jamf Protect actions
func (c *Client) GetJamfProtectHistory(params url.Values) (*ResponseJamfProtectHistoryList, error) {
	endpoint := uriJamfProtectHistory

	resp, err := c.DoPaginatedGet(endpoint, params)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "jamf protect history", err)
	}

	var out ResponseJamfProtectHistoryList
	out.TotalCount = resp.Size

	for _, value := range resp.Results {
		var newObj ResourceJamfProtectHistory
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "jamf protect history", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil
}

// CreateJamfProtectHistoryNote adds a note to the Jamf Protect history
func (c *Client) CreateJamfProtectHistoryNote(note ResourceJamfProtectHistoryNote) error {
	endpoint := uriJamfProtectHistory

	resp, err := c.HTTP.DoRequest("POST", endpoint, note, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedCreate, "jamf protect history note", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// GetJamfProtectPlans retrieves all Jamf Protect plans
func (c *Client) GetJamfProtectPlans(params url.Values) (*ResponseJamfProtectPlansList, error) {
	endpoint := uriJamfProtectPlans

	resp, err := c.DoPaginatedGet(endpoint, params)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "jamf protect plans", err)
	}

	var out ResponseJamfProtectPlansList
	out.TotalCount = resp.Size

	for _, value := range resp.Results {
		var newObj ResourceJamfProtectPlan
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "jamf protect plan", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil
}

// DeleteJamfProtectIntegration deletes the current Jamf Protect integration
func (c *Client) DeleteJamfProtectIntegration() error {
	endpoint := uriJamfProtect

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDelete, "jamf protect integration", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
