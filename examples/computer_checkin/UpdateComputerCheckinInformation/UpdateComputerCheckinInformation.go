package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
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
	logLevel := http_client.LogLevelWarning // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := http_client.Config{
		InstanceName: authConfig.InstanceName,
		Auth: http_client.AuthConfig{
			ClientID:     authConfig.ClientID,
			ClientSecret: authConfig.ClientSecret,
		},
		LogLevel: logLevel,
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
