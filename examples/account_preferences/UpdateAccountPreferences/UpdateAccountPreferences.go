package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-http-client/logger"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logLevel := logger.LogLevelWarn // LogLevelNone / LogLevelDebug / LogLevelInfo / LogLevelError

	// Configuration for the jamfpro
	config := httpclient.Config{
		InstanceName: authConfig.InstanceName,
		Auth: httpclient.AuthConfig{
			ClientID:     authConfig.ClientID,
			ClientSecret: authConfig.ClientSecret,
		},
		LogLevel: logLevel,
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
