// jamfproapi_jamf_protect.go
// Jamf Pro Api - Jamf Protect
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-protect
// Jamf Pro API requires the structs to support an JSON data structure.

package jamfpro

import (
	"fmt"
	"net/url"

	"github.com/mitchellh/mapstructure"
)

const uriJamfProtect = "/api/v1/jamf-protect"

// Response Structs
type ResponseJamfProtectHistoryList struct {
	TotalCount int                          `json:"totalCount"`
	Results    []ResourceJamfProtectHistory `json:"results"`
}

type ResponseJamfProtectPlansList struct {
	TotalCount int                       `json:"totalCount"`
	Results    []ResourceJamfProtectPlan `json:"results"`
}

// Resource Structs
type ResourceJamfProtectIntegrationSettings struct {
	ID             string `json:"id"`
	APIClientID    string `json:"apiClientId"`
	APIClientName  string `json:"apiClientName"`
	RegistrationID string `json:"registrationId"`
	ProtectURL     string `json:"protectUrl"`
	LastSyncTime   string `json:"lastSyncTime"`
	SyncStatus     string `json:"syncStatus"`
	AutoInstall    bool   `json:"autoInstall"`
}

type ResourceJamfProtectIntegrationRequest struct {
	AutoInstall bool `json:"autoInstall"`
}

type ResourceJamfProtectRetryRequest struct {
	IDs []string `json:"ids"`
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

type ResourceJamfProtectHistory struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Date     string `json:"date"`
	Note     string `json:"note"`
	Details  string `json:"details"`
}

// Structs for Jamf Protect API registration
type ResourceJamfProtectRegisterRequest struct {
	ProtectURL string `json:"protectUrl"`
	ClientID   string `json:"clientId"`
	Password   string `json:"password"`
}

type ResourceJamfProtectRegisterResponse struct {
	ID             string `json:"id"`
	APIClientID    string `json:"apiClientId"`
	APIClientName  string `json:"apiClientName"`
	RegistrationID string `json:"registrationId"`
	ProtectURL     string `json:"protectUrl"`
	LastSyncTime   string `json:"lastSyncTime"`
	SyncStatus     string `json:"syncStatus"`
	AutoInstall    bool   `json:"autoInstall"`
}

// CRUD

// GetJamfProtectIntegrationSettings retrieves the current Jamf Protect integration settings
func (c *Client) GetJamfProtectIntegrationSettings() (*ResourceJamfProtectIntegrationSettings, error) {
	endpoint := uriJamfProtect
	var out ResourceJamfProtectIntegrationSettings
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "jamf protect integration settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// UpdateJamfProtectIntegrationSettings updates the Jamf Protect integration settings
func (c *Client) UpdateJamfProtectIntegrationSettings(updatedSettings ResourceJamfProtectIntegrationSettings) (*ResourceJamfProtectIntegrationSettings, error) {
	endpoint := uriJamfProtect
	var out ResourceJamfProtectIntegrationSettings
	resp, err := c.HTTP.DoRequest("PUT", endpoint, updatedSettings, &out)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdate, "jamf protect integration settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// CreateJamfProtectIntegration creates a new Jamf Protect integration
func (c *Client) CreateJamfProtectIntegration(autoInstall bool) (*ResourceJamfProtectIntegrationSettings, error) {
	endpoint := uriJamfProtect
	requestBody := ResourceJamfProtectIntegrationRequest{
		AutoInstall: autoInstall,
	}
	var out ResourceJamfProtectIntegrationSettings
	resp, err := c.HTTP.DoRequest("PUT", endpoint, requestBody, &out)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "jamf protect integration", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
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

// New function to request a retry of Protect install tasks
func (c *Client) RetryJamfProtectInstallTasks(deploymentID string, taskIDs []string) error {
	endpoint := fmt.Sprintf("%s/deployments/%s/tasks/retry", uriJamfProtect, deploymentID)

	requestBody := ResourceJamfProtectRetryRequest{
		IDs: taskIDs,
	}

	resp, err := c.HTTP.DoRequest("POST", endpoint, requestBody, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedUpdate, "retry Jamf Protect install tasks", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// GetJamfProtectHistory retrieves the history of Jamf Protect actions
func (c *Client) GetJamfProtectHistory(params url.Values) (*ResponseJamfProtectHistoryList, error) {
	endpoint := fmt.Sprintf("%s/history", uriJamfProtect)

	resp, err := c.DoPaginatedGet(endpoint, params)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "Jamf Protect history", err)
	}

	var out ResponseJamfProtectHistoryList
	out.TotalCount = resp.Size

	for _, value := range resp.Results {
		var newObj ResourceJamfProtectHistory
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "Jamf Protect history", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil
}

// GetJamfProtectPlans retrieves all previously synced Jamf Protect Plans with their associated configuration profile information
func (c *Client) GetJamfProtectPlans(params url.Values) (*ResponseJamfProtectPlansList, error) {
	endpoint := fmt.Sprintf("%s/plans", uriJamfProtect)

	resp, err := c.DoPaginatedGet(endpoint, params)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "Jamf Protect plans", err)
	}

	var out ResponseJamfProtectPlansList
	out.TotalCount = resp.Size

	for _, value := range resp.Results {
		var newObj ResourceJamfProtectPlan
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "Jamf Protect plan", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil
}

// SyncJamfProtectPlans syncs plans with Jamf Protect
func (c *Client) SyncJamfProtectPlans() error {
	endpoint := fmt.Sprintf("%s/plans/sync", uriJamfProtect)

	resp, _ := c.HTTP.DoRequest("POST", endpoint, nil, nil)
	if resp == nil {
		return fmt.Errorf("failed to sync Jamf Protect plans: no response received")
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to sync Jamf Protect plans: unexpected status code %d", resp.StatusCode)
	}

	return nil
}

// CreateJamfProtectAPIConfiguration registers a Jamf Protect API configuration with Jamf Pro
func (c *Client) CreateJamfProtectAPIConfiguration(config ResourceJamfProtectRegisterRequest) (*ResourceJamfProtectRegisterResponse, error) {
	endpoint := fmt.Sprintf("%s/register", uriJamfProtect)

	var response ResourceJamfProtectRegisterResponse
	resp, err := c.HTTP.DoRequest("POST", endpoint, config, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "register Jamf Protect API configuration", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}
