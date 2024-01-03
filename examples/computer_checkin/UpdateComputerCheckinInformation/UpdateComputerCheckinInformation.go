package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client" // Import http_client for logging
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName:       authConfig.InstanceName,
		OverrideBaseDomain: authConfig.OverrideBaseDomain,
		LogLevel:           logLevel,
		Logger:             logger,
		ClientID:           authConfig.ClientID,
		ClientSecret:       authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
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
