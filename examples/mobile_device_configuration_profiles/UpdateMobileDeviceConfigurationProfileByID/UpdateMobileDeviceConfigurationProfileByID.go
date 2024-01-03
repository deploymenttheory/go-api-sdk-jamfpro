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

	// Define a new profile
	newProfile := jamfpro.ResourceMobileDeviceConfigurationProfile{
		General: jamfpro.MobileDeviceConfigurationProfileSubsetGeneral{
			Name: "WiFi",
			Site: jamfpro.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
			Category: jamfpro.SharedResourceCategory{
				ID:   -1,
				Name: "No category assigned",
			},
			DeploymentMethod: "Install Automatically",
			Payloads:         "<plist version=\"1\"><dict>...</dict></plist>", // Replace with actual XML payload
		},
		Scope: jamfpro.MobileDeviceConfigurationProfileSubsetScope{
			AllMobileDevices: false,
			AllJSSUsers:      false,
		},
		SelfService: jamfpro.MobileDeviceConfigurationProfileSubsetSelfService{
			// Fill in self service details if needed
		},
	}

	// Update a profile by ID (assuming ID is known)
	updatedProfileByID, err := client.UpdateMobileDeviceConfigurationProfileByID(2, &newProfile)
	if err != nil {
		fmt.Println("Error updating profile by ID:", err)
	} else {
		fmt.Println("Updated Profile by ID:", updatedProfileByID)
	}
}
