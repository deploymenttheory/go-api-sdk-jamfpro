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
	ComputerUsageLogs []ComputerHistorySubsetUsageLog      `json:"computer_usage_logs" xml:"computer_usage_logs"`
	Audits            []ComputerHistorySubsetAudit         `json:"audits" xml:"audits"`
	PolicyLogs        []ComputerHistorySubsetPolicyLog     `json:"policy_logs" xml:"policy_logs"`
	CasperRemoteLogs  []ComputerHistorySubsetCasperRemote  `json:"casper_remote_logs" xml:"casper_remote_logs"`
	ScreenSharingLogs []ComputerHistorySubsetScreenSharing `json:"screen_sharing_logs" xml:"screen_sharing_logs"`
	CasperImagingLogs []ComputerHistorySubsetCasperImaging `json:"casper_imaging_logs" xml:"casper_imaging_logs"`
	Commands          ComputerHistorySubsetCommands        `json:"commands" xml:"commands"`
	UserLocation      []ComputerHistorySubsetLocation      `json:"user_location" xml:"user_location"`
	MacAppStoreApps   ComputerHistorySubsetAppStoreApps    `json:"mac_app_store_applications" xml:"mac_app_store_applications"`
}

// ComputerHistorySubsetGeneralInfo stores general information about the computer.
type ComputerHistorySubsetGeneralInfo struct {
	ID           int    `json:"id" xml:"id"`
	Name         string `json:"name" xml:"name"`
	UDID         string `json:"udid" xml:"udid"`
	SerialNumber string `json:"serial_number" xml:"serial_number"`
	MacAddress   string `json:"mac_address" xml:"mac_address"`
}

// ComputerHistorySubsetUsageLog stores logs related to computer usage.
type ComputerHistorySubsetUsageLog struct {
	UsageLog ComputerHistorySubsetEventDetails `json:"usage_log" xml:"usage_log"`
}

// ComputerHistorySubsetAudit stores audit logs.
type ComputerHistorySubsetAudit struct {
	Audit ComputerHistorySubsetEventDetails `json:"audit" xml:"audit"`
}

// ComputerHistorySubsetPolicyLog stores logs related to policies.
type ComputerHistorySubsetPolicyLog struct {
	PolicyLog ComputerHistorySubsetPolicyDetails `json:"policy_log" xml:"policy_log"`
}

// ComputerHistorySubsetCasperRemote stores logs for Casper remote actions.
type ComputerHistorySubsetCasperRemote struct {
	CasperRemoteLog ComputerHistorySubsetEventStatus `json:"casper_remote_log" xml:"casper_remote_log"`
}

// ComputerHistorySubsetScreenSharing stores logs related to screen sharing.
type ComputerHistorySubsetScreenSharing struct {
	ScreenSharingLog ComputerHistorySubsetScreenSharingDetails `json:"screen_sharing_log" xml:"screen_sharing_log"`
}

// ComputerHistorySubsetCasperImaging stores logs for Casper imaging actions.
type ComputerHistorySubsetCasperImaging struct {
	CasperImagingLog ComputerHistorySubsetEventStatus `json:"casper_imaging_log" xml:"casper_imaging_log"`
}

// ComputerHistorySubsetCommands groups completed, pending, and failed commands.
type ComputerHistorySubsetCommands struct {
	Completed []ComputerHistorySubsetCommand `json:"completed" xml:"completed"`
	Pending   []ComputerHistorySubsetCommand `json:"pending" xml:"pending"`
	Failed    []ComputerHistorySubsetCommand `json:"failed" xml:"failed"`
}

// ComputerHistorySubsetLocation stores location data related to a user.
type ComputerHistorySubsetLocation struct {
	Location ComputerHistorySubsetUserLocation `json:"location" xml:"location"`
}

// ComputerHistorySubsetAppStoreApps groups installed, pending, and failed applications from the Mac App Store.
type ComputerHistorySubsetAppStoreApps struct {
	Installed []ComputerHistorySubsetApp `json:"installed" xml:"installed"`
	Pending   []ComputerHistorySubsetApp `json:"pending" xml:"pending"`
	Failed    []ComputerHistorySubsetApp `json:"failed" xml:"failed"`
}

// ComputerHistorySubsetEventDetails defines the structure for logging events with timestamps and user information.
type ComputerHistorySubsetEventDetails struct {
	Event         string `json:"event" xml:"event"`
	Username      string `json:"username" xml:"username"`
	DateTime      string `json:"date_time" xml:"date_time"`
	DateTimeEpoch int64  `json:"date_time_epoch" xml:"date_time_epoch"`
	DateTimeUTC   string `json:"date_time_utc" xml:"date_time_utc"`
}

// ComputerHistorySubsetPolicyDetails defines the details for policy logs.
type ComputerHistorySubsetPolicyDetails struct {
	PolicyID      int    `json:"policy_id" xml:"policy_id"`
	PolicyName    string `json:"policy_name" xml:"policy_name"`
	Username      string `json:"username" xml:"username"`
	DateTime      string `json:"date_time" xml:"date_time"`
	DateTimeEpoch int64  `json:"date_time_epoch" xml:"date_time_epoch"`
	DateTimeUTC   string `json:"date_time_utc" xml:"date_time_utc"`
	Status        string `json:"status" xml:"status"`
}

// ComputerHistorySubsetEventStatus defines a simple structure for logs with status and timestamps.
type ComputerHistorySubsetEventStatus struct {
	DateTime      string `json:"date_time" xml:"date_time"`
	DateTimeEpoch int64  `json:"date_time_epoch" xml:"date_time_epoch"`
	DateTimeUTC   string `json:"date_time_utc" xml:"date_time_utc"`
	Status        string `json:"status" xml:"status"`
}

// ComputerHistorySubsetScreenSharingDetails extends event status with details specific to screen sharing.
type ComputerHistorySubsetScreenSharingDetails struct {
	ComputerHistorySubsetEventStatus
	Details string `json:"details" xml:"details"`
}

// ComputerHistorySubsetCommand details a command with its issue and completion status.
type ComputerHistorySubsetCommand struct {
	Name           string `json:"name" xml:"name"`
	Status         string `json:"status" xml:"status"`
	Issued         string `json:"issued" xml:"issued"`
	IssuedEpoch    int64  `json:"issued_epoch" xml:"issued_epoch"`
	IssuedUTC      string `json:"issued_utc" xml:"issued_utc"`
	LastPush       string `json:"last_push" xml:"last_push"`
	LastPushEpoch  int64  `json:"last_push_epoch" xml:"last_push_epoch"`
	LastPushUTC    string `json:"last_push_utc" xml:"last_push_utc"`
	Username       string `json:"username" xml:"username"`
	Completed      string `json:"completed,omitempty" xml:"completed,omitempty"`
	CompletedEpoch int64  `json:"completed_epoch,omitempty" xml:"completed_epoch,omitempty"`
	CompletedUTC   string `json:"completed_utc,omitempty" xml:"completed_utc,omitempty"`
	Failed         string `json:"failed,omitempty" xml:"failed,omitempty"`
	FailedEpoch    int64  `json:"failed_epoch,omitempty" xml:"failed_epoch,omitempty"`
	FailedUTC      string `json:"failed_utc,omitempty" xml:"failed_utc,omitempty"`
}

// ComputerHistorySubsetUserLocation defines the detailed information about a user's location.
type ComputerHistorySubsetUserLocation struct {
	DateTime      string `json:"date_time" xml:"date_time"`
	DateTimeEpoch int64  `json:"date_time_epoch" xml:"date_time_epoch"`
	DateTimeUTC   string `json:"date_time_utc" xml:"date_time_utc"`
	Username      string `json:"username" xml:"username"`
	FullName      string `json:"full_name" xml:"full_name"`
	EmailAddress  string `json:"email_address" xml:"email_address"`
	PhoneNumber   string `json:"phone_number" xml:"phone_number"`
	Department    string `json:"department" xml:"department"`
	Building      string `json:"building" xml:"building"`
	Room          int    `json:"room" xml:"room"`
	Position      string `json:"position" xml:"position"`
}

// ComputerHistorySubsetApp defines the structure for application details in the Mac App Store context.
type ComputerHistorySubsetApp struct {
	Name            string `json:"name" xml:"name"`
	Version         string `json:"version" xml:"version"`
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
func (c *Client) GetComputerHistoryByComputerID(id int) (*ResourceComputerHistory, error) {
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

// GetComputerHistoryByComputerIDAndDataSubset retrieves a subset of the historical information of a computer given its ID and subset name.
func (c *Client) GetComputerHistoryByComputerIDAndDataSubset(id int, subset string) (*ResourceComputerHistory, error) {
	endpoint := fmt.Sprintf("%s/id/%d/subset/%s", uriComputerHistory, id, subset)

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
