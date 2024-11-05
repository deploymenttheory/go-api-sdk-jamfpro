// jamfproapi_local_admin_password.go
// Jamf Pro Api - JAMF local administrator password (LAPS)
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-pending-rotations
// Jamf Pro API requires the structs to support an JSON data structure.

package jamfpro

import "fmt"

const uriLocalAdminPassword = "/api/v2/local-admin-password"

// Resource
type ResourceLocalAdminPasswordSettings struct {
	AutoDeployEnabled        bool `json:"autoDeployEnabled"`
	PasswordRotationTime     int  `json:"passwordRotationTime"`
	AutoRotateEnabled        bool `json:"autoRotateEnabled"`
	AutoRotateExpirationTime int  `json:"autoRotateExpirationTime"`
}

type ResourceLapsPasswordList struct {
	LapsUserPasswordList []LapsUserPassword `json:"lapsUserPasswordList"`
}

type LapsUserPassword struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// Response

// Response structs for pending rotations
type ResponseLocalAdminPasswordPendingRotations struct {
	TotalCount int                                 `json:"totalCount"`
	Results    []LocalAdminPasswordPendingRotation `json:"results"`
}

type LocalAdminPasswordPendingRotation struct {
	LapsUser    LocalAdminPasswordUser `json:"lapsUser"`
	CreatedDate string                 `json:"createdDate"`
}

type LocalAdminPasswordUser struct {
	ClientManagementID string `json:"clientManagementId"`
	GUID               string `json:"guid"`
	Username           string `json:"username"`
	UserSource         string `json:"userSource"`
}

// Response structs for password history
type ResponseLocalAdminPasswordHistory struct {
	TotalCount int                             `json:"totalCount"`
	Results    []LocalAdminPasswordHistoryItem `json:"results"`
}

type LocalAdminPasswordHistoryItem struct {
	Password       string                    `json:"password"`
	DateLastSeen   string                    `json:"dateLastSeen"`
	ExpirationTime string                    `json:"expirationTime"`
	Audits         []LocalAdminPasswordAudit `json:"audits"`
}

type LocalAdminPasswordAudit struct {
	ViewedBy string `json:"viewedBy"`
	DateSeen string `json:"dateSeen"`
}

// Resoponse struct for current password
type ResponseLocalAdminCurrentPassword struct {
	Password string `json:"password"`
}

// Response structs for full LAPS history
type ResponseLocalAdminPasswordFullHistory struct {
	TotalCount int                           `json:"totalCount"`
	Results    []LocalAdminPasswordFullEvent `json:"results"`
}

type LocalAdminPasswordFullEvent struct {
	Username   string `json:"username"`
	EventType  string `json:"eventType"`
	EventTime  string `json:"eventTime"`
	ViewedBy   string `json:"viewedBy"`
	UserSource string `json:"userSource"`
}

// Response structs for LAPS capable accounts
type ResponseLocalAdminPasswordCapableAccounts struct {
	TotalCount int                         `json:"totalCount"`
	Results    []LocalAdminPasswordAccount `json:"results"`
}

type LocalAdminPasswordAccount struct {
	ClientManagementID string `json:"clientManagementId"`
	GUID               string `json:"guid"`
	Username           string `json:"username"`
	UserSource         string `json:"userSource"`
}

// ResponseLapsPasswordSet represents the response after setting LAPS passwords
type ResponseLapsPasswordSet struct {
	LapsUserPasswordList []LapsUserPasswordResponse `json:"lapsUserPasswordList"`
}

type LapsUserPasswordResponse struct {
	Username string `json:"username"`
}

// GetListOfPendingLapsRotations retrieves a list of devices and usernames with pending LAPS rotations
func (c *Client) GetListOfPendingLapsRotations() (*ResponseLocalAdminPasswordPendingRotations, error) {
	endpoint := fmt.Sprintf("%s/pending-rotations", uriLocalAdminPassword)

	var out ResponseLocalAdminPasswordPendingRotations
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "pending LAPS rotations", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// GetLocalAdminPasswordSettings retrieves current Jamf Pro LAPS settings
func (c *Client) GetLocalAdminPasswordSettings() (*ResourceLocalAdminPasswordSettings, error) {
	endpoint := uriLocalAdminPassword + "/settings"
	var out ResourceLocalAdminPasswordSettings
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "LAPS settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// UpdateLocalAdminPasswordSettings updates the current Jamf Pro LAPS settings
func (c *Client) UpdateLocalAdminPasswordSettings(settings *ResourceLocalAdminPasswordSettings) error {
	endpoint := uriLocalAdminPassword + "/settings"

	var handleResponse struct{}

	resp, err := c.HTTP.DoRequest("PUT", endpoint, &settings, &handleResponse)
	if err != nil {
		return fmt.Errorf(errMsgFailedUpdate, "LAPS settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// GetLocalAdminPasswordViewedHistory retrieves the password view history for a specific username on a target device.
// History will include password, who viewed the password and when it was viewed.
func (c *Client) GetLocalAdminPasswordViewedHistoryByClientManagementID(clientManagementID string, username string) (*ResponseLocalAdminPasswordHistory, error) {
	endpoint := fmt.Sprintf("%s/%s/account/%s/audit", uriLocalAdminPassword, clientManagementID, username)

	var out ResponseLocalAdminPasswordHistory
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, fmt.Sprintf("LAPS history for user %s on device %s", username, clientManagementID), err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// GetCurrentLocalAdminPasswordForSpecifiedUsernameByClientManagementID retrieves the current LAPS password for a specific username on a target device.
// Note: Once viewed, the password will be rotated based on rotation time settings.
func (c *Client) GetCurrentLocalAdminPasswordForSpecifiedUsernameByClientManagementID(clientManagementID string, username string) (*ResponseLocalAdminCurrentPassword, error) {
	endpoint := fmt.Sprintf("%s/%s/account/%s/password", uriLocalAdminPassword, clientManagementID, username)

	var out ResponseLocalAdminCurrentPassword
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, fmt.Sprintf("current LAPS password for user %s on device %s", username, clientManagementID), err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// GetLocalAdminPasswordFullHistory retrieves the complete history of all local admin passwords for all accounts
// on a specific device, including both viewing and rotation history.
func (c *Client) GetLocalAdminPasswordFullHistoryByClientManagementID(clientManagementID string) (*ResponseLocalAdminPasswordFullHistory, error) {
	endpoint := fmt.Sprintf("%s/%s/history", uriLocalAdminPassword, clientManagementID)

	var out ResponseLocalAdminPasswordFullHistory
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, fmt.Sprintf("LAPS full history for device %s", clientManagementID), err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// GetLocalAdminPasswordCapableAccounts retrieves a list of all admin accounts that are LAPS capable for a specific device.
// Capable accounts are returned in the AutoSetupAdminAccounts from QueryResponses.
func (c *Client) GetLocalAdminPasswordCapableAccountsByClientManagementID(clientManagementID string) (*ResponseLocalAdminPasswordCapableAccounts, error) {
	endpoint := fmt.Sprintf("%s/%s/accounts", uriLocalAdminPassword, clientManagementID)

	var out ResponseLocalAdminPasswordCapableAccounts
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, fmt.Sprintf("LAPS capable accounts for device %s", clientManagementID), err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// SetLocalAdminPasswordByClientManagementID sets LAPS passwords for all capable accounts on a device
// The passwords are provided as a list of username/password pairs
func (c *Client) SetLocalAdminPasswordByClientManagementID(clientManagementID string, passwordList *ResourceLapsPasswordList) (*ResponseLapsPasswordSet, error) {
	endpoint := fmt.Sprintf("%s/%s/set-password", uriLocalAdminPassword, clientManagementID)

	var response ResponseLapsPasswordSet
	resp, err := c.HTTP.DoRequest("PUT", endpoint, passwordList, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdate, fmt.Sprintf("LAPS passwords for device %s", clientManagementID), err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}
