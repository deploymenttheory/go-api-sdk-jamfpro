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
	newSettings := jamfpro.ResourceComputerCheckin{
		CheckInFrequency:                 15, // Values can be 60, 30, 15 and  5
		CreateStartupScript:              true,
		LogStartupEvent:                  true,
		CheckForPoliciesAtStartup:        true,
		ApplyComputerLevelManagedPrefs:   true,
		EnsureSSHIsEnabled:               false,
		CreateLoginLogoutHooks:           true,
		LogUsername:                      true,
		CheckForPoliciesAtLoginLogout:    true,
		ApplyUserLevelManagedPreferences: true,
		HideRestorePartition:             false,
		PerformLoginActionsInBackground:  true,
		DisplayStatusToUser:              false,
	}

	// Update computer check-in settings
	err = client.UpdateComputerCheckinInformation(&newSettings)
	if err != nil {
		fmt.Printf("Error updating computer check-in settings: %s\n", err)
		return
	}

	fmt.Println("Computer check-in settings updated successfully.")
}
