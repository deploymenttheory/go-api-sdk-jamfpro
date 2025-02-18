// jamfproapi_account_preferences.go
// Jamf Pro Api - Account Preferences
// api reference: undocumented
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
)

// Responses

const uriAccountPreferences = "/api/v2/account-preferences"

// Resource

type ResourceAccountPreferences struct {
	Language                             string `json:"language"`
	DateFormat                           string `json:"dateFormat"`
	Timezone                             string `json:"timezone"`
	DisableRelativeDates                 bool   `json:"disableRelativeDates"`
	DisablePageLeaveCheck                bool   `json:"disablePageLeaveCheck"`
	DisableShortcutsTooltips             bool   `json:"disableShortcutsTooltips"`
	DisableTablePagination               bool   `json:"disableTablePagination"`
	ConfigProfilesSortingMethod          string `json:"configProfilesSortingMethod"`
	ResultsPerPage                       int    `json:"resultsPerPage"`
	UserInterfaceDisplayTheme            string `json:"userInterfaceDisplayTheme"`
	ComputerSearchMethod                 string `json:"computerSearchMethod"`
	ComputerApplicationSearchMethod      string `json:"computerApplicationSearchMethod"`
	ComputerApplicationUsageSearchMethod string `json:"computerApplicationUsageSearchMethod"`
	ComputerFontSearchMethod             string `json:"computerFontSearchMethod"`
	ComputerPluginSearchMethod           string `json:"computerPluginSearchMethod"`
	ComputerLocalUserAccountSearchMethod string `json:"computerLocalUserAccountSearchMethod"`
	ComputerSoftwareUpdateSearchMethod   string `json:"computerSoftwareUpdateSearchMethod"`
	ComputerPackageReceiptSearchMethod   string `json:"computerPackageReceiptSearchMethod"`
	ComputerPrinterSearchMethod          string `json:"computerPrinterSearchMethod"`
	ComputerPeripheralSearchMethod       string `json:"computerPeripheralSearchMethod"`
	ComputerServiceSearchMethod          string `json:"computerServiceSearchMethod"`
	MobileDeviceSearchMethod             string `json:"mobileDeviceSearchMethod"`
	MobileDeviceAppSearchMethod          string `json:"mobileDeviceAppSearchMethod"`
	UserSearchMethod                     string `json:"userSearchMethod"`
	UserAllContentSearchMethod           string `json:"userAllContentSearchMethod"`
	UserMobileDeviceAppSearchMethod      string `json:"userMobileDeviceAppSearchMethod"`
	UserMacAppStoreAppSearchMethod       string `json:"userMacAppStoreAppSearchMethod"`
	UserEbookSearchMethod                string `json:"userEbookSearchMethod"`
}

// GetAccountPreferences retrieves the jamf pro account settings.
func (c *Client) GetAccountPreferences() (*ResourceAccountPreferences, error) {
	endpoint := uriAccountPreferences

	var accountPreferences ResourceAccountPreferences
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &accountPreferences)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "Account Preferences", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &accountPreferences, nil
}

// UpdateAccountPreferences updates the jamf pro account settings.
func (c *Client) UpdateAccountPreferences(updatedSettings ResourceAccountPreferences) (*ResourceAccountPreferences, error) {
	endpoint := uriAccountPreferences
	var out ResourceAccountPreferences

	resp, err := c.HTTP.DoRequest("PATCH", endpoint, updatedSettings, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdate, "Account Preferences", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}
