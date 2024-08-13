// jamfproapi_mdm.go
// Jamf Pro Api - MDM Commands
// api reference: https://developer.jamf.com/jamf-pro/reference/post_preview-mdm-commands
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import "fmt"

const uriMDMCommands = "/api/v2/mdm/commands"
const uriMDMDeployPackage = "/api/v1/deploy-package"
const uriMDMProfileRenewal = "/api/v1/mdm/renew-profile"

// MDM Commands

// ResourceMDMCommandRequest represents the overall request structure for the MDM command
type ResourceMDMCommandRequest struct {
	CommandData CommandData  `json:"commandData"`
	ClientData  []ClientData `json:"clientData"`
}

// ClientData represents the client data structure in the request
type ClientData struct {
	ManagementID string `json:"managementId"`
}

// CommandData represents the command data structure in the request
type CommandData struct {
	CommandType string `json:"commandType"`
	// Delete_User
	UserName       string `json:"userName,omitempty"`
	ForceDeletion  bool   `json:"forceDeletion,omitempty"`
	DeleteAllUsers bool   `json:"deleteAllUsers,omitempty"`
	// Enable_Lost_Mode
	LostModeMessage  string `json:"lostModeMessage,omitempty"`
	LostModePhone    string `json:"lostModePhone,omitempty"`
	LostModeFootnote string `json:"lostModeFootnote,omitempty"`
	// Erase_Device
	ReturnToService        *ReturnToService `json:"returnToService,omitempty"`
	PreserveDataPlan       bool             `json:"preserveDataPlan,omitempty"`
	DisallowProximitySetup bool             `json:"disallowProximitySetup,omitempty"`
	PIN                    string           `json:"pin,omitempty"`
	ObliterationBehavior   string           `json:"obliterationBehavior,omitempty"`
	// Restart_Device
	RebuildKernelCache bool     `json:"rebuildKernelCache,omitempty"`
	KextPaths          []string `json:"kextPaths,omitempty"`
	NotifyUser         bool     `json:"notifyUser,omitempty"`
	// Settings
	ApplicationAttributes     *ApplicationAttributes     `json:"applicationAttributes,omitempty"`
	SharedDeviceConfiguration *SharedDeviceConfiguration `json:"sharedDeviceConfiguration,omitempty"`
	ApplicationConfiguration  *ApplicationConfiguration  `json:"applicationConfiguration,omitempty"`
	SoftwareUpdateSettings    *SoftwareUpdateSettings    `json:"softwareUpdateSettings,omitempty"`
	BootstrapTokenAllowed     bool                       `json:"bootstrapTokenAllowed,omitempty"`
	Bluetooth                 bool                       `json:"bluetooth,omitempty"`
	AppAnalytics              string                     `json:"appAnalytics,omitempty"`
	DiagnosticSubmission      string                     `json:"diagnosticSubmission,omitempty"`
	DataRoaming               string                     `json:"dataRoaming,omitempty"`
	VoiceRoaming              string                     `json:"voiceRoaming,omitempty"`
	PersonalHotspot           string                     `json:"personalHotspot,omitempty"`
	MaximumResidentUsers      int                        `json:"maximumResidentUsers,omitempty"`
	DeviceName                string                     `json:"deviceName,omitempty"`
	TimeZone                  string                     `json:"timeZone,omitempty"`
	PasscodeLockGracePeriod   int                        `json:"passcodeLockGracePeriod,omitempty"`
	// Set_Auto_Admin_Password
	GUID     string `json:"guid,omitempty"`
	Password string `json:"password,omitempty"`
}

// ReturnToService represents the return to service structure in the erase device command
type ReturnToService struct {
	Enabled         bool   `json:"enabled"`
	MDMProfileData  string `json:"mdmProfileData,omitempty"`
	WifiProfileData string `json:"wifiProfileData,omitempty"`
}

// ApplicationAttributes represents the application attributes structure in the settings command
type ApplicationAttributes struct {
	VpnUuid               string   `json:"vpnUuid"`
	AssociatedDomains     []string `json:"associatedDomains"`
	Removable             bool     `json:"removable"`
	EnableDirectDownloads bool     `json:"enableDirectDownloads"`
	Identifier            string   `json:"identifier"`
}

// SharedDeviceConfiguration represents the shared device configuration structure in the settings command
type SharedDeviceConfiguration struct {
	QuotaSize     int `json:"quotaSize"`
	ResidentUsers int `json:"residentUsers"`
}

// ApplicationConfiguration represents the application configuration structure in the settings command
type ApplicationConfiguration struct {
	Configuration string `json:"configuration"`
	Identifier    string `json:"identifier"`
}

// SoftwareUpdateSettings represents the software update settings structure in the settings command
type SoftwareUpdateSettings struct {
	RecommendationCadence string `json:"recommendationCadence"`
}

// ResponseMDMCommand represents the response structure for the MDM command
type ResponseMDMCommand struct {
	ID   string `json:"id,omitempty"`
	Href string `json:"href"`
}

// Deploy Package

// ResourceDeployPackage represents the request structure for deploying a package
type ResourceDeployPackage struct {
	Manifest         PackageManifest `json:"manifest"`
	InstallAsManaged bool            `json:"installAsManaged"`
	Devices          []int           `json:"devices"`
	GroupID          string          `json:"groupId"`
}

// PackageManifest represents the package manifest structure in the deploy package command
type PackageManifest struct {
	HashType         string `json:"hashType"`
	URL              string `json:"url"`
	Hash             string `json:"hash"`
	DisplayImageURL  string `json:"displayImageUrl"`
	FullSizeImageURL string `json:"fullSizeImageUrl"`
	BundleID         string `json:"bundleId"`
	BundleVersion    string `json:"bundleVersion"`
	Subtitle         string `json:"subtitle"`
	Title            string `json:"title"`
	SizeInBytes      int    `json:"sizeInBytes"`
}

// ResponseDeployPackage represents the response structure for deploying a package
type ResponseDeployPackage struct {
	QueuedCommands []QueuedCommand             `json:"queuedCommands"`
	Errors         []SharedResourceErrorDetail `json:"errors"`
}

// QueuedCommand represents the details of a queued command in the response
type QueuedCommand struct {
	Device      int    `json:"device"`
	CommandUUID string `json:"commandUuid"`
}

// Renew Profile

// ResourceMDMProfileRenewal represents the request structure for renewing MDM profiles
type ResourceMDMProfileRenewal struct {
	UDIDs []string `json:"udids"`
}

// ResponseMDMProfileRenewal represents the response structure for renewing MDM profiles
type ResponseMDMProfileRenewal struct {
	UDIDsNotProcessed struct {
		UDIDs []string `json:"udids"`
	} `json:"udidsNotProcessed"`
}

// SendMDMCommandForCreationAndQueuing sends an MDM command for creation and queuing
func (c *Client) SendMDMCommandForCreationAndQueuing(MDMCommand *ResourceMDMCommandRequest) (*ResponseMDMCommand, error) {
	endpoint := uriMDMCommands
	var responseMDMCommand ResponseMDMCommand

	resp, err := c.HTTP.DoRequest("POST", endpoint, MDMCommand, &responseMDMCommand)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "send MDM Command", err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &responseMDMCommand, nil
}

// SendMDMCommandForPackageDeployment deploys a package using an MDM command
func (c *Client) SendMDMCommandForPackageDeployment(deployPackageRequest *ResourceDeployPackage) (*ResponseDeployPackage, error) {
	endpoint := uriMDMDeployPackage + "?verbose=true"
	var responseDeployPackage ResponseDeployPackage

	resp, err := c.HTTP.DoRequest("POST", endpoint, deployPackageRequest, &responseDeployPackage)
	if err != nil {
		return nil, fmt.Errorf("failed to deploy package: %v", err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &responseDeployPackage, nil
}

// SendMDMCommandForMDMProfileRenewal renews MDM profiles for specified UDIDs
func (c *Client) SendMDMCommandForMDMProfileRenewal(renewProfileRequest *ResourceMDMProfileRenewal) (*ResponseMDMProfileRenewal, error) {
	endpoint := uriMDMProfileRenewal
	var responseMDMProfileRenewal ResponseMDMProfileRenewal

	resp, err := c.HTTP.DoRequest("POST", endpoint, renewProfileRequest, &responseMDMProfileRenewal)
	if err != nil {
		return nil, fmt.Errorf("failed to renew MDM profile: %v", err)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	return &responseMDMProfileRenewal, nil
}
