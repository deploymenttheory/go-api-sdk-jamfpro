// jamfproapi_mdm.go
// Jamf Pro Api - MDM Commands
// api reference: https://developer.jamf.com/jamf-pro/reference/post_preview-mdm-commands
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import "fmt"

const uriMDM = "/api/v2/mdm/commands"

// Structs

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

// SendMDMCommandForCreationAndQueuing sends an MDM command for creation and queuing
func (c *Client) SendMDMCommandForCreationAndQueuing(MDMCommand *ResourceMDMCommandRequest) (*ResponseMDMCommand, error) {
	endpoint := uriMDM
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
