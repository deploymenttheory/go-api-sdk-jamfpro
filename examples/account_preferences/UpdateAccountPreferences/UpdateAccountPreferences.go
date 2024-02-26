package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define new settings to update
	newSettings := jamfpro.ResourceAccountPreferences{
		Language:                             "en",
		DateFormat:                           "MM/dd/yyyy", // MM/dd/yyyy /
		Timezone:                             "America/Chicago",
		DisableRelativeDates:                 false,
		DisablePageLeaveCheck:                true,
		DisableShortcutsTooltips:             false,
		DisableTablePagination:               true,
		ConfigProfilesSortingMethod:          "ALPHABETICALLY", // ALPHABETICALLY / STANDARD
		ResultsPerPage:                       20,
		UserInterfaceDisplayTheme:            "DARK",     // DARK / LIGHT / SYSTEM
		ComputerSearchMethod:                 "CONTAINS", // Exact Match / Starts with / Contains
		ComputerApplicationSearchMethod:      "CONTAINS",
		ComputerApplicationUsageSearchMethod: "CONTAINS",
		ComputerFontSearchMethod:             "CONTAINS",
		ComputerPluginSearchMethod:           "CONTAINS",
		ComputerLocalUserAccountSearchMethod: "CONTAINS",
		ComputerSoftwareUpdateSearchMethod:   "CONTAINS",
		ComputerPackageReceiptSearchMethod:   "CONTAINS",
		ComputerPrinterSearchMethod:          "CONTAINS",
		ComputerPeripheralSearchMethod:       "CONTAINS",
		ComputerServiceSearchMethod:          "CONTAINS",
		MobileDeviceSearchMethod:             "CONTAINS",
		MobileDeviceAppSearchMethod:          "CONTAINS",
		UserSearchMethod:                     "CONTAINS",
		UserAllContentSearchMethod:           "CONTAINS",
		UserMobileDeviceAppSearchMethod:      "CONTAINS",
		UserMacAppStoreAppSearchMethod:       "CONTAINS",
		UserEbookSearchMethod:                "CONTAINS",
	}

	// Update account preferences
	updatedSettings, err := client.UpdateAccountPreferences(newSettings)
	if err != nil {
		log.Fatalf("Error updating Account Preferences: %v", err)
	}

	// Pretty print the updated settings in JSON
	updatedSettingsJSON, err := json.MarshalIndent(updatedSettings, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling updated Account Preferences data: %v", err)
	}
	fmt.Println("Updated Account Preferences:\n", string(updatedSettingsJSON))
}
