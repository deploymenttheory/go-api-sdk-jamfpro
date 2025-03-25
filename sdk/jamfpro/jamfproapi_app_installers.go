// jamfproapi_app_installers.go
// Jamf Pro Api - Jamf App Catalog
// api reference: Undocumented
// Jamf Pro Api requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

const uriJamfAppCatalogAppInstaller = "/api/v1/app-installers"

// List

// Struct for paginated response for app installers
type ResponseJamfAppCatalogTitleList struct {
	Size    int                                  `json:"totalCount"`
	Results []ResourceJamfAppCatalogAppInstaller `json:"results"`
}

// Response

type ResponseJamfAppCatalogDeploymentCreateAndUpdate struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

type ResponseJamfAppCatalogDeploymentTermsAndConditionsStatus struct {
	Accepted       bool   `json:"accepted"`
	AcceptanceTime string `json:"acceptanceTime"`
}

// Resource

type ResponseJamfAppCatalogGlobalSettings struct {
	EndUserExperienceSettings JamfAppCatalogDeploymentSubsetNotificationSettings `json:"endUserExperienceSettings"`
}

type ResourceJamfAppCatalogAppInstaller struct {
	ID                       string                                        `json:"id"`
	BundleId                 string                                        `json:"bundleId,omitempty"`
	TitleName                string                                        `json:"titleName,omitempty"`
	Publisher                string                                        `json:"publisher,omitempty"`
	IconUrl                  string                                        `json:"iconUrl,omitempty"`
	Version                  string                                        `json:"version,omitempty"`
	SizeInBytes              int                                           `json:"sizeInBytes,omitempty"`
	MinimumOsVersion         string                                        `json:"minimumOsVersion,omitempty"`
	Language                 string                                        `json:"language,omitempty"`
	AvailabilityDate         string                                        `json:"availabilityDate,omitempty"`
	PackageSigningIdentity   string                                        `json:"packageSigningIdentity,omitempty"`
	InstallerPackageHashType string                                        `json:"installerPackageHashType,omitempty"`
	InstallerPackageHash     string                                        `json:"installerPackageHash,omitempty"`
	ShortVersion             string                                        `json:"shortVersion,omitempty"`
	Architecture             string                                        `json:"architecture,omitempty"`
	OriginalMediaSources     []JamfAppCatalogAppInstallerSubsetMediaSource `json:"originalMediaSources,omitempty"`
	LaunchDaemonIncluded     *bool                                         `json:"launchDaemonIncluded"`
	NotificationAvailable    *bool                                         `json:"notificationAvailable"`
	SuppressAutoUpdate       *bool                                         `json:"suppressAutoUpdate"`
}

// MediaSource struct for the media sources within the JSON response
type JamfAppCatalogAppInstallerSubsetMediaSource struct {
	HashType string `json:"hashType,omitempty"`
	Hash     string `json:"hash,omitempty"`
	Url      string `json:"url,omitempty"`
}

// Struct which represents AppInstallers object JSON from Pro API
type ResourceJamfAppCatalogDeployment struct {
	ID                              string                                             `json:"id"`
	Name                            string                                             `json:"name,omitempty"`
	Enabled                         *bool                                              `json:"enabled"`
	AppTitleId                      string                                             `json:"appTitleId,omitempty"`
	DeploymentType                  string                                             `json:"deploymentType,omitempty"`
	UpdateBehavior                  string                                             `json:"updateBehavior,omitempty"`
	CategoryId                      string                                             `json:"categoryId,omitempty"`
	SiteId                          string                                             `json:"siteId,omitempty"`
	SmartGroupId                    string                                             `json:"smartGroupId,omitempty"`
	InstallPredefinedConfigProfiles *bool                                              `json:"installPredefinedConfigProfiles"`
	TitleAvailableInAis             *bool                                              `json:"titleAvailableInAis"`
	TriggerAdminNotifications       *bool                                              `json:"triggerAdminNotifications"`
	NotificationSettings            JamfAppCatalogDeploymentSubsetNotificationSettings `json:"notificationSettings,omitempty"`
	SelfServiceSettings             JamfAppCatalogDeploymentSubsetSelfServiceSettings  `json:"selfServiceSettings,omitempty"`
	SelectedVersion                 string                                             `json:"selectedVersion,omitempty"`
	LatestAvailableVersion          string                                             `json:"latestAvailableVersion,omitempty"`
	VersionRemoved                  *bool                                              `json:"versionRemoved"`
}

// JamfAppCatalogDeploymentSubsetNotificationSettings struct represents the notification settings within ResourceJamfAppCatalogAppInstaller
type JamfAppCatalogDeploymentSubsetNotificationSettings struct {
	NotificationMessage  string `json:"notificationMessage,omitempty"`
	NotificationInterval int    `json:"notificationInterval,omitempty"`
	DeadlineMessage      string `json:"deadlineMessage,omitempty"`
	Deadline             int    `json:"deadline,omitempty"`
	QuitDelay            int    `json:"quitDelay,omitempty"`
	CompleteMessage      string `json:"completeMessage,omitempty"`
	Relaunch             *bool  `json:"relaunch"`
	Suppress             *bool  `json:"suppress,omitempty"`
}

// JamfAppCatalogDeploymentSubsetSelfServiceSettings struct represents the self-service settings within ResourceJamfAppCatalogAppInstaller
type JamfAppCatalogDeploymentSubsetSelfServiceSettings struct {
	IncludeInFeaturedCategory   *bool                                    `json:"includeInFeaturedCategory"`
	IncludeInComplianceCategory *bool                                    `json:"includeInComplianceCategory"`
	ForceViewDescription        *bool                                    `json:"forceViewDescription"`
	Description                 string                                   `json:"description,omitempty"`
	Categories                  []JamfAppCatalogDeploymentSubsetCategory `json:"categories,omitempty"`
}

// Category struct represents the categories within SelfServiceSettings
type JamfAppCatalogDeploymentSubsetCategory struct {
	ID       string `json:"id,omitempty"`
	Featured *bool  `json:"featured,omitempty"`
}

// CRUD

// GetJamfAppCatalogAppInstallerTermsAndConditionsStatus returns the terms and conditions status for the Jamf App Catalog
func (c *Client) GetJamfAppCatalogAppInstallerTermsAndConditionsStatus() (*ResponseJamfAppCatalogDeploymentTermsAndConditionsStatus, error) {
	endpoint := fmt.Sprintf("%s/terms-and-conditions", uriJamfAppCatalogAppInstaller)

	var globalSettings ResponseJamfAppCatalogDeploymentTermsAndConditionsStatus
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &globalSettings)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "Jamf App Catalog Titles", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &globalSettings, nil
}

// AcceptJamfAppCatalogAppInstallerTermsAndConditions accepts the terms and conditions for the Jamf App Catalog
// This is required on an account by account basis.
func (c *Client) AcceptJamfAppCatalogAppInstallerTermsAndConditions() (*ResponseJamfAppCatalogDeploymentTermsAndConditionsStatus, error) {
	endpoint := fmt.Sprintf("%s/terms-and-conditions/accept", uriJamfAppCatalogAppInstaller)

	var globalSettings ResponseJamfAppCatalogDeploymentTermsAndConditionsStatus
	resp, err := c.HTTP.DoRequest("POST", endpoint, nil, &globalSettings)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "Jamf App Catalog Titles", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &globalSettings, nil
}

// Gets full list of Get Jamf App Catalog App Installer Titles & handles pagination
func (c *Client) GetJamfAppCatalogAppInstallerTitles(sort_filter string) (*ResponseJamfAppCatalogTitleList, error) {
	resp, err := c.DoPaginatedGet(
		uriJamfAppCatalogAppInstaller+"/titles",
		standardPageSize,
		startingPageNumber,
		sort_filter,
	)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "Jamf App Catalog Titles", err)
	}

	var out ResponseJamfAppCatalogTitleList
	out.Size = resp.Size

	for _, value := range resp.Results {
		var newObj ResourceJamfAppCatalogAppInstaller
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "Jamf App Catalog Titles", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil

}

// GetJamfAppCatalogAppInstallerByTitleID retrieves by title ID & returns ResourceJamfAppCatalogAppInstaller
func (c *Client) GetJamfAppCatalogAppInstallerByTitleID(id string) (*ResourceJamfAppCatalogAppInstaller, error) {
	endpoint := fmt.Sprintf("%s/titles/%s", uriJamfAppCatalogAppInstaller, id)
	var appInstaller ResourceJamfAppCatalogAppInstaller
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &appInstaller)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "Jamf App Catalog Title", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &appInstaller, nil
}

// Retrieves Jamf App Catalog installer deployment by provided ID & returns ResourceJamfAppCatalogDeployment
func (c *Client) GetJamfAppCatalogAppInstallerByID(id string) (*ResourceJamfAppCatalogDeployment, error) {
	endpoint := fmt.Sprintf("%s/deployments/%s", uriJamfAppCatalogAppInstaller, id)
	var appInstaller ResourceJamfAppCatalogDeployment
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &appInstaller)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "jamf app catalog deployments", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &appInstaller, nil
}

// GetJamfAppCatalogAppInstallerByName retrieves deployment by name & returns ResourceJamfAppCatalogDeployment
func (c *Client) GetJamfAppCatalogAppInstallerByName(name string) (*ResourceJamfAppCatalogDeployment, error) {
	// First, get all deployments
	endpoint := fmt.Sprintf("%s/deployments", uriJamfAppCatalogAppInstaller)
	var deployments struct {
		Results []ResourceJamfAppCatalogDeployment `json:"results"`
	}

	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &deployments)
	if err != nil {
		return nil, fmt.Errorf("failed to get deployments: %v", err)
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	for _, deployment := range deployments.Results {
		if deployment.Name == name {
			return &deployment, nil
		}
	}

	return nil, fmt.Errorf("no deployment found with name: %s", name)
}

// GetJamfAppCatalogAppInstallerByTitleName retrieves title by name & returns ResourceJamfAppCatalogAppInstaller
func (c *Client) GetJamfAppCatalogAppInstallerByTitleName(name string) (*ResourceJamfAppCatalogAppInstaller, error) {
	titles, err := c.GetJamfAppCatalogAppInstallerTitles("")
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "Jamf App Catalog Titles", err)
	}

	var titleID string
	for _, value := range titles.Results {
		if value.TitleName == name {
			titleID = value.ID
			break
		}
	}

	if titleID == "" {
		return nil, fmt.Errorf(errMsgFailedGetByName, "Jamf App Catalog Title", name, errMsgNoName)
	}

	fullResource, err := c.GetJamfAppCatalogAppInstallerByTitleID(titleID)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "Jamf App Catalog Title", titleID, err)
	}

	return fullResource, nil
}

// GetJamfAppCatalogAppInstallerGlobalSettings retrieves the global end-user experience settings for the App Catalog
func (c *Client) GetJamfAppCatalogAppInstallerGlobalSettings() (*JamfAppCatalogDeploymentSubsetNotificationSettings, error) {
	endpoint := fmt.Sprintf("%s/global-settings", uriJamfAppCatalogAppInstaller)

	var response ResponseJamfAppCatalogGlobalSettings
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "Jamf App Catalog Global Settings", err)
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response.EndUserExperienceSettings, nil
}

// UpdateJamfAppCatalogAppInstallerGlobalSettings updates the global end-user experience settings for the Jamf App Catalog
func (c *Client) UpdateJamfAppCatalogAppInstallerGlobalSettings(payload *ResponseJamfAppCatalogGlobalSettings) (*ResponseJamfAppCatalogGlobalSettings, error) {
	endpoint := fmt.Sprintf("%s/global-settings", uriJamfAppCatalogAppInstaller)
	var response ResponseJamfAppCatalogGlobalSettings

	resp, err := c.HTTP.DoRequest("PUT", endpoint, payload, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "Jamf App Catalog Global Settings", "global-settings", err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// Creates Jamf App Catalog Deployment from ResponseJamfAppCatalogDeploymentCreateAndUpdate struct
func (c *Client) CreateJamfAppCatalogAppInstallerDeployment(payload *ResourceJamfAppCatalogDeployment) (*ResponseJamfAppCatalogDeploymentCreateAndUpdate, error) {
	endpoint := uriJamfAppCatalogAppInstaller + "/deployments"
	var response ResponseJamfAppCatalogDeploymentCreateAndUpdate

	resp, err := c.HTTP.DoRequest("POST", endpoint, payload, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "jamf app catalog deployment", err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateJamfAppCatalogDeploymentByID updates Jamf App Catalog Deployment by provided ID & returns ResourceJamfAppCatalogDeployment
func (c *Client) UpdateJamfAppCatalogAppInstallerDeploymentByID(id string, payload *ResourceJamfAppCatalogDeployment) (*ResourceJamfAppCatalogDeployment, error) {
	endpoint := fmt.Sprintf("%s/deployments/%s", uriJamfAppCatalogAppInstaller, id)
	var response ResourceJamfAppCatalogDeployment
	resp, err := c.HTTP.DoRequest("PUT", endpoint, payload, &response)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "script", id, err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &response, nil

}

// DeleteJamfAppCatalogDeploymentByID deletes deployment by provided ID & returns ResponseJamfAppCatalogDeploymentCreateAndUpdate
func (c *Client) DeleteJamfAppCatalogAppInstallerDeploymentByID(id string) error {
	endpoint := fmt.Sprintf("%s/deployments/%s", uriJamfAppCatalogAppInstaller, id)
	var response interface{}
	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, &response)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "script", id, err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return nil
}
