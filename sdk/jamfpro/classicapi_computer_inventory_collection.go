// classicapi_computer_inventory_collection.go
// Jamf Pro Classic Api - Computer Inventory Collection
// api reference: https://developer.jamf.com/jamf-pro/reference/computerinventorycollection
// Jamf Pro Classic Api requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriComputerInventoryCollection = "/JSSResource/computerinventorycollection"

// Resource

// ResourceComputerInventoryCollection represents the detailed information of inventory collection settings.

type ResourceComputerInventoryCollection struct {
	LocalUserAccounts             bool               `json:"local_user_accounts" xml:"local_user_accounts"`
	HomeDirectorySizes            bool               `json:"home_directory_sizes" xml:"home_directory_sizes"`
	HiddenAccounts                bool               `json:"hidden_accounts" xml:"hidden_accounts"`
	Printers                      bool               `json:"printers" xml:"printers"`
	ActiveServices                bool               `json:"active_services" xml:"active_services"`
	MobileDeviceAppPurchasingInfo bool               `json:"mobile_device_app_purchasing_info" xml:"mobile_device_app_purchasing_info"`
	ComputerLocationInformation   bool               `json:"computer_location_information" xml:"computer_location_information"`
	PackageReceipts               bool               `json:"package_receipts" xml:"package_receipts"`
	AvailableSoftwareUpdates      bool               `json:"available_software_updates" xml:"available_software_updates"`
	InclueApplications            bool               `json:"inclue_applications" xml:"inclue_applications"`
	InclueFonts                   bool               `json:"inclue_fonts" xml:"inclue_fonts"`
	IncluePlugins                 bool               `json:"inclue_plugins" xml:"inclue_plugins"`
	Applications                  []ApplicationEntry `json:"applications,omitempty" xml:"applications,omitempty"`
	Fonts                         []FontEntry        `json:"fonts,omitempty" xml:"fonts,omitempty"`
	Plugins                       []PluginEntry      `json:"plugins,omitempty" xml:"plugins,omitempty"`
}

type ApplicationEntry struct {
	Application Application `json:"application,omitempty" xml:"application,omitempty"`
}

type Application struct {
	Path     string `json:"path,omitempty" xml:"path,omitempty"`
	Platform string `json:"platform,omitempty" xml:"platform,omitempty"`
}

type FontEntry struct {
	Font Font `json:"font,omitempty" xml:"font,omitempty"`
}

type Font struct {
	Path     string `json:"path,omitempty" xml:"path,omitempty"`
	Platform string `json:"platform,omitempty" xml:"platform,omitempty"`
}

type PluginEntry struct {
	Plugin Plugin `json:"plugin,omitempty" xml:"plugin,omitempty"`
}

type Plugin struct {
	Path     string `json:"path,omitempty" xml:"path,omitempty"`
	Platform string `json:"platform,omitempty" xml:"platform,omitempty"`
}

// CRUD

// GetComputerInventoryCollection gets the jamf pro inventory collection settings
func (c *Client) GetComputerInventoryCollectionInformation() (*ResourceComputerInventoryCollection, error) {
	endpoint := uriComputerInventoryCollection

	var inventoryCollection ResourceComputerInventoryCollection
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &inventoryCollection)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "computer inventory collection settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &inventoryCollection, nil
}

// UpdateComputerInventoryCollectionInformation updates the jamf pro computer check-in settings
func (c *Client) UpdateComputerInventoryCollectionInformation(settings *ResourceComputerInventoryCollection) error {
	endpoint := uriComputerInventoryCollection

	requestBody := struct {
		XMLName xml.Name `xml:"computer_inventory_collection"`
		*ResourceComputerInventoryCollection
	}{
		ResourceComputerInventoryCollection: settings,
	}

	var handleResponse struct{}

	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &handleResponse)
	if err != nil {
		return fmt.Errorf(errMsgFailedUpdate, "computer inventory collection settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
