// classicapi_computer_history.go
// Jamf Pro Classic Api - Computer History
// api reference: https://developer.jamf.com/jamf-pro/reference/computerhistory
// Classic API requires the structs to support an XML data structure.

package jamfpro

import "fmt"

const uriComputerHistory = "/JSSResource/computerhistory"

// ResourceComputerHistory represents the detailed historical information of a computer.
type ResourceComputerHistory struct {
	ComputerID              int                `json:"computer_id" xml:"computer_id"`
	GeneralInfo             GeneralInfo        `json:"general_info,omitempty" xml:"general_info,omitempty"`
	HardwareDetails         Hardware           `json:"hardware_details,omitempty" xml:"hardware_details,omitempty"`
	SoftwareUpdates         []UpdateInfo       `json:"software_updates,omitempty" xml:"software_updates,omitempty"`
	NetworkInfo             NetworkInfo        `json:"network_info,omitempty" xml:"network_info,omitempty"`
	ComputerUsageLogs       []UsageLog         `json:"computer_usage_logs,omitempty" xml:"computer_usage_logs,omitempty"`
	Audits                  []AuditEntry       `json:"audits,omitempty" xml:"audits,omitempty"`
	PolicyLogs              []PolicyLog        `json:"policy_logs,omitempty" xml:"policy_logs,omitempty"`
	CasperRemoteLogs        []RemoteLog        `json:"casper_remote_logs,omitempty" xml:"casper_remote_logs,omitempty"`
	ScreenSharingLogs       []ScreenSharingLog `json:"screen_sharing_logs,omitempty" xml:"screen_sharing_logs,omitempty"`
	CasperImagingLogs       []ImagingLog       `json:"casper_imaging_logs,omitempty" xml:"casper_imaging_logs,omitempty"`
	Commands                CommandStatuses    `json:"commands" xml:"commands"`
	UserLocation            []UserLocation     `json:"user_location,omitempty" xml:"user_location,omitempty"`
	MacAppStoreApplications AppStatuses        `json:"mac_app_store_applications" xml:"mac_app_store_applications"`
}

type GeneralInfo struct {
	ID           int    `json:"id,omitempty" xml:"id,omitempty"`
	Name         string `json:"name,omitempty" xml:"name,omitempty"`
	Serial       string `json:"serial,omitempty" xml:"serial,omitempty"`
	MAC          string `json:"mac,omitempty" xml:"mac,omitempty"`
	OSVersion    string `json:"os_version,omitempty" xml:"os_version,omitempty"`
	UDID         string `json:"udid,omitempty" xml:"udid,omitempty"`
	SerialNumber string `json:"serial_number,omitempty" xml:"serial_number,omitempty"`
	MACAddress   string `json:"mac_address,omitempty" xml:"mac_address,omitempty"`
}

type Hardware struct {
	RAM     string `json:"ram,omitempty" xml:"ram,omitempty"`
	Storage string `json:"storage,omitempty" xml:"storage,omitempty"`
	CPU     string `json:"cpu,omitempty" xml:"cpu,omitempty"`
}

type UpdateInfo struct {
	SoftwareName    string `json:"software_name,omitempty" xml:"software_name,omitempty"`
	SoftwareVersion string `json:"software_version,omitempty" xml:"software_version,omitempty"`
	InstalledDate   string `json:"installed_date,omitempty" xml:"installed_date,omitempty"`
}

type NetworkInfo struct {
	IPAddress  string `json:"ip_address,omitempty" xml:"ip_address,omitempty"`
	SubnetMask string `json:"subnet_mask,omitempty" xml:"subnet_mask,omitempty"`
	Gateway    string `json:"gateway,omitempty" xml:"gateway,omitempty"`
	DNS        string `json:"dns,omitempty" xml:"dns,omitempty"`
}

type UsageLog struct {
	EventLog EventLog `json:"usage_log" xml:"usage_log"`
}

type AuditEntry struct {
	Audit EventLog `json:"audit" xml:"audit"`
}

type PolicyLog struct {
	Policy EventLog `json:"policy_log" xml:"policy_log"`
}

type RemoteLog struct {
	CasperRemote EventLog `json:"casper_remote_log" xml:"casper_remote_log"`
}

type ScreenSharingLog struct {
	ScreenSharing EventLog `json:"screen_sharing_log" xml:"screen_sharing_log"`
}

type ImagingLog struct {
	CasperImaging EventLog `json:"casper_imaging_log" xml:"casper_imaging_log"`
}

type EventLog struct {
	DateTime      string `json:"date_time,omitempty" xml:"date_time,omitempty"`
	DateTimeEpoch int64  `json:"date_time_epoch,omitempty" xml:"date_time_epoch,omitempty"`
	DateTimeUTC   string `json:"date_time_utc,omitempty" xml:"date_time_utc,omitempty"`
	Event         string `json:"event,omitempty" xml:"event,omitempty"`
	Username      string `json:"username,omitempty" xml:"username,omitempty"`
	Status        string `json:"status,omitempty" xml:"status,omitempty"`
	Details       string `json:"details,omitempty" xml:"details,omitempty"`
}

type CommandStatuses struct {
	Completed []CompletedCommand `json:"completed,omitempty" xml:"completed,omitempty"`
	Pending   []PendingCommand   `json:"pending,omitempty" xml:"pending,omitempty"`
	Failed    []FailedCommand    `json:"failed,omitempty" xml:"failed,omitempty"`
}

type CompletedCommand struct {
	Command struct {
		Name           string `json:"name,omitempty" xml:"name,omitempty"`
		Completed      string `json:"completed,omitempty" xml:"completed,omitempty"`
		CompletedEpoch int64  `json:"completed_epoch,omitempty" xml:"completed_epoch,omitempty"`
		CompletedUTC   string `json:"completed_utc,omitempty" xml:"completed_utc,omitempty"`
		Username       string `json:"username,omitempty" xml:"username,omitempty"`
	} `json:"command" xml:"command"`
}

type PendingCommand struct {
	Command struct {
		Name          string `json:"name,omitempty" xml:"name,omitempty"`
		Status        string `json:"status,omitempty" xml:"status,omitempty"`
		Issued        string `json:"issued,omitempty" xml:"issued,omitempty"`
		IssuedEpoch   int64  `json:"issued_epoch,omitempty" xml:"issued_epoch,omitempty"`
		IssuedUTC     string `json:"issued_utc,omitempty" xml:"issued_utc,omitempty"`
		LastPush      string `json:"last_push,omitempty" xml:"last_push,omitempty"`
		LastPushEpoch int64  `json:"last_push_epoch,omitempty" xml:"last_push_epoch,omitempty"`
		LastPushUTC   string `json:"last_push_utc,omitempty" xml:"last_push_utc,omitempty"`
		Username      string `json:"username,omitempty" xml:"username,omitempty"`
	} `json:"command" xml:"command"`
}

type FailedCommand struct {
	Command struct {
		Name        string `json:"name,omitempty" xml:"name,omitempty"`
		Status      string `json:"status,omitempty" xml:"status,omitempty"`
		Issued      string `json:"issued,omitempty" xml:"issued,omitempty"`
		IssuedEpoch int64  `json:"issued_epoch,omitempty" xml:"issued_epoch,omitempty"`
		IssuedUTC   string `json:"issued_utc,omitempty" xml:"issued_utc,omitempty"`
		Failed      string `json:"failed,omitempty" xml:"failed,omitempty"`
		FailedEpoch int64  `json:"failed_epoch,omitempty" xml:"failed_epoch,omitempty"`
		FailedUTC   string `json:"failed_utc,omitempty" xml:"failed_utc,omitempty"`
	} `json:"command" xml:"command"`
}

type UserLocation struct {
	Location UserInfo `json:"location" xml:"location"`
}

type UserInfo struct {
	DateTime      string `json:"date_time,omitempty" xml:"date_time,omitempty"`
	DateTimeEpoch int64  `json:"date_time_epoch,omitempty" xml:"date_time_epoch,omitempty"`
	DateTimeUTC   string `json:"date_time_utc,omitempty" xml:"date_time_utc,omitempty"`
	Username      string `json:"username,omitempty" xml:"username,omitempty"`
	FullName      string `json:"full_name,omitempty" xml:"full_name,omitempty"`
	EmailAddress  string `json:"email_address,omitempty" xml:"email_address,omitempty"`
	PhoneNumber   string `json:"phone_number,omitempty" xml:"phone_number,omitempty"`
	Department    string `json:"department,omitempty" xml:"department,omitempty"`
	Building      string `json:"building,omitempty" xml:"building,omitempty"`
	Room          int    `json:"room,omitempty" xml:"room,omitempty"`
	Position      string `json:"position,omitempty" xml:"position,omitempty"`
}

type AppStatuses struct {
	Installed []AppInfo `json:"installed,omitempty" xml:"installed,omitempty"`
	Pending   []AppInfo `json:"pending,omitempty" xml:"pending,omitempty"`
	Failed    []AppInfo `json:"failed,omitempty" xml:"failed,omitempty"`
}

type AppInfo struct {
	App ApplicationDetails `json:"app" xml:"app"`
}

type ApplicationDetails struct {
	Name            string `json:"name,omitempty" xml:"name,omitempty"`
	Version         string `json:"version,omitempty" xml:"version,omitempty"`
	SizeMB          int    `json:"size_mb,omitempty" xml:"size_mb,omitempty"`
	Deployed        string `json:"deployed,omitempty" xml:"deployed,omitempty"`
	DeployedEpoch   int64  `json:"deployed_epoch,omitempty" xml:"deployed_epoch,omitempty"`
	DeployedUTC     string `json:"deployed_utc,omitempty" xml:"deployed_utc,omitempty"`
	LastUpdate      string `json:"last_update,omitempty" xml:"last_update,omitempty"`
	LastUpdateEpoch int64  `json:"last_update_epoch,omitempty" xml:"last_update_epoch,omitempty"`
	LastUpdateUTC   string `json:"last_update_utc,omitempty" xml:"last_update_utc,omitempty"`
}

// CRUD Methods

// GetComputerHistoryByID retrieves the historical information of a computer given its ID.
func (c *Client) GetComputerHistoryByID(id int) (*ResourceComputerHistory, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriComputerHistory, id)

	var computerHistory ResourceComputerHistory
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &computerHistory)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "computer histroy", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &computerHistory, nil
}
