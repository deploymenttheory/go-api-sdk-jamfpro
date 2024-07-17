// classicapi_computer_history.go
// Jamf Pro Classic Api - Computer History
// api reference: https://developer.jamf.com/jamf-pro/reference/computerhistory
// Classic API requires the structs to support an XML data structure.

package jamfpro

import "fmt"

const uriComputerHistory = "/JSSResource/computerhistory"

// ResourceComputerHistory represents the root structure of the computer history resource.
type ResourceComputerHistory struct {
	General           ComputerHistorySubsetGeneralInfo     `json:"general" xml:"general"`
	ComputerUsageLogs []ComputerHistorySubsetUsageLog      `json:"computer_usage_logs,omitempty" xml:"computer_usage_logs,omitempty"`
	Audits            []ComputerHistorySubsetAudit         `json:"audits,omitempty" xml:"audits,omitempty"`
	PolicyLogs        []ComputerHistorySubsetPolicyLog     `json:"policy_logs,omitempty" xml:"policy_logs,omitempty"`
	CasperRemoteLogs  []ComputerHistorySubsetCasperRemote  `json:"casper_remote_logs,omitempty" xml:"casper_remote_logs,omitempty"`
	ScreenSharingLogs []ComputerHistorySubsetScreenSharing `json:"screen_sharing_logs,omitempty" xml:"screen_sharing_logs,omitempty"`
	CasperImagingLogs []ComputerHistorySubsetCasperImaging `json:"casper_imaging_logs,omitempty" xml:"casper_imaging_logs,omitempty"`
	Commands          ComputerHistorySubsetCommands        `json:"commands,omitempty" xml:"commands,omitempty"`
	UserLocation      []ComputerHistorySubsetLocation      `json:"user_location,omitempty" xml:"user_location,omitempty"`
	MacAppStoreApps   ComputerHistorySubsetAppStoreApps    `json:"mac_app_store_applications,omitempty" xml:"mac_app_store_applications,omitempty"`
}

// ComputerHistorySubsetGeneralInfo stores general information about the computer.
type ComputerHistorySubsetGeneralInfo struct {
	ID           int    `json:"id,omitempty" xml:"id,omitempty"`
	Name         string `json:"name,omitempty" xml:"name,omitempty"`
	UDID         string `json:"udid,omitempty" xml:"udid,omitempty"`
	SerialNumber string `json:"serial_number,omitempty" xml:"serial_number,omitempty"`
	MacAddress   string `json:"mac_address,omitempty" xml:"mac_address,omitempty"`
}

// ComputerHistorySubsetUsageLog stores logs related to computer usage.
type ComputerHistorySubsetUsageLog struct {
	UsageLog ComputerHistorySubsetEventDetails `json:"usage_log,omitempty" xml:"usage_log,omitempty"`
}

// ComputerHistorySubsetAudit stores audit logs.
type ComputerHistorySubsetAudit struct {
	Audit ComputerHistorySubsetEventDetails `json:"audit,omitempty" xml:"audit,omitempty"`
}

// ComputerHistorySubsetPolicyLog stores logs related to policies.
type ComputerHistorySubsetPolicyLog struct {
	PolicyLog ComputerHistorySubsetPolicyDetails `json:"policy_log,omitempty" xml:"policy_log,omitempty"`
}

// ComputerHistorySubsetCasperRemote stores logs for Casper remote actions.
type ComputerHistorySubsetCasperRemote struct {
	CasperRemoteLog ComputerHistorySubsetEventStatus `json:"casper_remote_log" xml:"casper_remote_log"`
}

// ComputerHistorySubsetScreenSharing stores logs related to screen sharing.
type ComputerHistorySubsetScreenSharing struct {
	ScreenSharingLog ComputerHistorySubsetScreenSharingDetails `json:"screen_sharing_log,omitempty" xml:"screen_sharing_log,omitempty"`
}

// ComputerHistorySubsetCasperImaging stores logs for Casper imaging actions.
type ComputerHistorySubsetCasperImaging struct {
	CasperImagingLog ComputerHistorySubsetEventStatus `json:"casper_imaging_log,omitempty" xml:"casper_imaging_log,omitempty"`
}

// ComputerHistorySubsetCommands groups completed, pending, and failed commands.
type ComputerHistorySubsetCommands struct {
	Completed []ComputerHistorySubsetCommand `json:"completed,omitempty" xml:"completed,omitempty"`
	Pending   []ComputerHistorySubsetCommand `json:"pending,omitempty" xml:"pending,omitempty"`
	Failed    []ComputerHistorySubsetCommand `json:"failed,omitempty" xml:"failed,omitempty"`
}

// ComputerHistorySubsetLocation stores location data related to a user.
type ComputerHistorySubsetLocation struct {
	Location ComputerHistorySubsetUserLocation `json:"location,omitempty" xml:"location,omitempty"`
}

// ComputerHistorySubsetAppStoreApps groups installed, pending, and failed applications from the Mac App Store.
type ComputerHistorySubsetAppStoreApps struct {
	Installed []ComputerHistorySubsetApp `json:"installed,omitempty" xml:"installed,omitempty"`
	Pending   []ComputerHistorySubsetApp `json:"pending,omitempty" xml:"pending,omitempty"`
	Failed    []ComputerHistorySubsetApp `json:"failed,omitempty" xml:"failed,omitempty"`
}

// ComputerHistorySubsetEventDetails defines the structure for logging events with timestamps and user information.
type ComputerHistorySubsetEventDetails struct {
	Event         string `json:"event,omitempty" xml:"event,omitempty"`
	Username      string `json:"username,omitempty" xml:"username,omitempty"`
	DateTime      string `json:"date_time,omitempty" xml:"date_time,omitempty"`
	DateTimeEpoch int64  `json:"date_time_epoch,omitempty" xml:"date_time_epoch,omitempty"`
	DateTimeUTC   string `json:"date_time_utc,omitempty" xml:"date_time_utc,omitempty"`
}

// ComputerHistorySubsetPolicyDetails defines the details for policy logs.
type ComputerHistorySubsetPolicyDetails struct {
	PolicyID      int    `json:"policy_id,omitempty" xml:"policy_id,omitempty"`
	PolicyName    string `json:"policy_name,omitempty" xml:"policy_name,omitempty"`
	Username      string `json:"username,omitempty" xml:"username,omitempty"`
	DateTime      string `json:"date_time,omitempty" xml:"date_time,omitempty"`
	DateTimeEpoch int64  `json:"date_time_epoch,omitempty" xml:"date_time_epoch,omitempty"`
	DateTimeUTC   string `json:"date_time_utc,omitempty" xml:"date_time_utc,omitempty"`
	Status        string `json:"status,omitempty" xml:"status,omitempty"`
}

// ComputerHistorySubsetEventStatus defines a simple structure for logs with status and timestamps.
type ComputerHistorySubsetEventStatus struct {
	DateTime      string `json:"date_time,omitempty" xml:"date_time,omitempty"`
	DateTimeEpoch int64  `json:"date_time_epoch,omitempty" xml:"date_time_epoch,omitempty"`
	DateTimeUTC   string `json:"date_time_utc,omitempty" xml:"date_time_utc,omitempty"`
	Status        string `json:"status,omitempty" xml:"status,omitempty"`
}

// ComputerHistorySubsetScreenSharingDetails extends event status with details specific to screen sharing.
type ComputerHistorySubsetScreenSharingDetails struct {
	ComputerHistorySubsetEventStatus
	Details string `json:"details,omitempty" xml:"details,omitempty"`
}

// ComputerHistorySubsetCommand details a command with its issue and completion status.
type ComputerHistorySubsetCommand struct {
	Name           string `json:"name,omitempty" xml:"name,omitempty"`
	Status         string `json:"status,omitempty" xml:"status,omitempty"`
	Issued         string `json:"issued,omitempty" xml:"issued,omitempty"`
	IssuedEpoch    int64  `json:"issued_epoch,omitempty" xml:"issued_epoch,omitempty"`
	IssuedUTC      string `json:"issued_utc,omitempty" xml:"issued_utc,omitempty"`
	LastPush       string `json:"last_push,omitempty" xml:"last_push,omitempty"`
	LastPushEpoch  int64  `json:"last_push_epoch,omitempty" xml:"last_push_epoch,omitempty"`
	LastPushUTC    string `json:"last_push_utc,omitempty" xml:"last_push_utc,omitempty"`
	Username       string `json:"username,omitempty" xml:"username,omitempty"`
	Completed      string `json:"completed,omitempty" xml:"completed,omitempty"`
	CompletedEpoch int64  `json:"completed_epoch,omitempty" xml:"completed_epoch,omitempty"`
	CompletedUTC   string `json:"completed_utc,omitempty" xml:"completed_utc,omitempty"`
	Failed         string `json:"failed,omitempty" xml:"failed,omitempty"`
	FailedEpoch    int64  `json:"failed_epoch,omitempty" xml:"failed_epoch,omitempty"`
	FailedUTC      string `json:"failed_utc,omitempty" xml:"failed_utc,omitempty"`
}

// ComputerHistorySubsetUserLocation defines the detailed information about a user's location.
type ComputerHistorySubsetUserLocation struct {
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

// ComputerHistorySubsetApp defines the structure for application details in the Mac App Store context.
type ComputerHistorySubsetApp struct {
	Name            string `json:"name,omitempty" xml:"name,omitempty"`
	Version         string `json:"version,omitempty" xml:"version,omitempty"`
	SizeMB          int    `json:"size_mb,omitempty" xml:"size_mb,omitempty"`
	Status          string `json:"status,omitempty" xml:"status,omitempty"`
	Deployed        string `json:"deployed,omitempty" xml:"deployed,omitempty"`
	DeployedEpoch   int64  `json:"deployed_epoch,omitempty" xml:"deployed_epoch,omitempty"`
	DeployedUTC     string `json:"deployed_utc,omitempty" xml:"deployed_utc,omitempty"`
	LastUpdate      string `json:"last_update,omitempty" xml:"last_update,omitempty"`
	LastUpdateEpoch int64  `json:"last_update_epoch,omitempty" xml:"last_update_epoch,omitempty"`
	LastUpdateUTC   string `json:"last_update_utc,omitempty" xml:"last_update_utc,omitempty"`
}

// CRUD Methods

// GetComputerHistoryByComputerID retrieves the historical information of a computer given its ID.
func (c *Client) GetComputerHistoryByComputerID(id string) (*ResourceComputerHistory, error) {
	endpoint := fmt.Sprintf("%s/id/%s", uriComputerHistory, id)

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

// GetComputerHistoryByComputerIDAndDataSubset retrieves a subset of the historical information of a computer given its ID and subset name.
func (c *Client) GetComputerHistoryByComputerIDAndDataSubset(id string, subset string) (*ResourceComputerHistory, error) {
	endpoint := fmt.Sprintf("%s/id/%s/subset/%s", uriComputerHistory, id, subset)

	var computerHistory ResourceComputerHistory
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &computerHistory)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "computer history with data subset", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &computerHistory, nil
}

// GetComputerHistoryByComputerName retrieves the historical information of a computer given its name.
func (c *Client) GetComputerHistoryByComputerName(name string) (*ResourceComputerHistory, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriComputerHistory, name)

	var computerHistory ResourceComputerHistory
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &computerHistory)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve computer history by computer name '%s': %v", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &computerHistory, nil
}

// GetComputerHistoryByComputerNameAndDataSubset retrieves a subset of the historical information of a computer given its name and subset name.
func (c *Client) GetComputerHistoryByComputerNameAndDataSubset(name string, subset string) (*ResourceComputerHistory, error) {
	endpoint := fmt.Sprintf("%s/name/%s/subset/%s", uriComputerHistory, name, subset)

	var computerHistory ResourceComputerHistory
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &computerHistory)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "computer history with data subset", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &computerHistory, nil
}

// GetComputerHistoryByComputerUDID retrieves the historical information of a computer by it's UDID.
func (c *Client) GetComputerHistoryByComputerUDID(udid string) (*ResourceComputerHistory, error) {
	endpoint := fmt.Sprintf("%s/udid/%s", uriComputerHistory, udid)

	var computerHistory ResourceComputerHistory
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &computerHistory)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve computer history by computer udid '%s': %v", udid, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &computerHistory, nil
}

// GetComputerHistoryByComputerUDIDAndDataSubset retrieves a subset of the historical information of a computer given its udid and subset name.
func (c *Client) GetComputerHistoryByComputerUDIDAndDataSubset(udid string, subset string) (*ResourceComputerHistory, error) {
	endpoint := fmt.Sprintf("%s/udid/%s/subset/%s", uriComputerHistory, udid, subset)

	var computerHistory ResourceComputerHistory
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &computerHistory)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve computer history by computer udid '%s': %v", udid, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &computerHistory, nil
}

// GetComputerHistoryByComputerSerialNumber retrieves the historical information of a computer by it's serial number
func (c *Client) GetComputerHistoryByComputerSerialNumber(serial string) (*ResourceComputerHistory, error) {
	endpoint := fmt.Sprintf("%s/serialnumber/%s", uriComputerHistory, serial)

	var computerHistory ResourceComputerHistory
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &computerHistory)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "computer history with computer serial number", serial, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &computerHistory, nil
}

// GetComputerHistoryByComputerSerialNumberAndDataSubset retrieves a subset of the historical information of a computer by it's serial number and data subset name.
func (c *Client) GetComputerHistoryByComputerSerialNumberAndDataSubset(udid string, subset string) (*ResourceComputerHistory, error) {
	endpoint := fmt.Sprintf("%s/serialnumber/%s/subset/%s", uriComputerHistory, udid, subset)

	var computerHistory ResourceComputerHistory
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &computerHistory)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve computer history by computer serial number '%s': %v", udid, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &computerHistory, nil
}

// GetComputerHistoryByComputerMACAddress retrieves the historical information of a computer by it's MAC Address
func (c *Client) GetComputerHistoryByComputerMACAddress(MACAddress string) (*ResourceComputerHistory, error) {
	endpoint := fmt.Sprintf("%s/macaddress/%s", uriComputerHistory, MACAddress)

	var computerHistory ResourceComputerHistory
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &computerHistory)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "computer history with computer MAC Address", MACAddress, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &computerHistory, nil
}

// GetComputerHistoryByComputerMACAddressAndDataSubset retrieves a subset of the historical information of a computer by it's serial number and data subset name.
func (c *Client) GetComputerHistoryByComputerMACAddressAndDataSubset(MACAddress string, subset string) (*ResourceComputerHistory, error) {
	endpoint := fmt.Sprintf("%s/macaddress/%s/subset/%s", uriComputerHistory, MACAddress, subset)

	var computerHistory ResourceComputerHistory
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &computerHistory)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve computer history by computer MAC Address '%s': %v", MACAddress, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &computerHistory, nil
}
