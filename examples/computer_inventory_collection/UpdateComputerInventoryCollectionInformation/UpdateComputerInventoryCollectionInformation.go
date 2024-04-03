package main

import (
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

	// Set up new settings
	newSettings := jamfpro.ResourceComputerInventoryCollection{
		LocalUserAccounts:             true,
		HomeDirectorySizes:            true,
		HiddenAccounts:                true,
		Printers:                      true,
		ActiveServices:                true,
		MobileDeviceAppPurchasingInfo: false,
		ComputerLocationInformation:   false,
		PackageReceipts:               true,
		AvailableSoftwareUpdates:      true,
		InclueApplications:            true, // Note: Assuming "InclueApplications" is a typo, it should be "IncludeApplications" if following XML structure.
		InclueFonts:                   true,
		IncluePlugins:                 true,
		Applications: []jamfpro.ApplicationEntry{
			{
				Application: jamfpro.Application{
					Path:     "/Library/Application Support/Applications",
					Platform: "Mac",
				},
			},
		},
		Fonts: []jamfpro.FontEntry{
			{
				Font: jamfpro.Font{
					Path:     "~/Downloads/",
					Platform: "Mac",
				},
			},
		},
		Plugins: []jamfpro.PluginEntry{
			{
				Plugin: jamfpro.Plugin{
					Path:     "~/Library/Internet Plug-Ins/",
					Platform: "Mac",
				},
			},
		},
	}

	// Update computer check-in settings
	err = client.UpdateComputerInventoryCollectionInformation(&newSettings)
	if err != nil {
		fmt.Printf("Error updating computer check-in settings: %s\n", err)
		return
	}

	fmt.Println("Computer check-in settings updated successfully.")
}
