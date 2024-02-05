package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

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

	// Update a profile by Name (assuming name is known)
	updatedProfileByName, err := client.UpdateMobileDeviceConfigurationProfileByName("WiFi", &newProfile)
	if err != nil {
		fmt.Println("Error updating profile by Name:", err)
	} else {
		fmt.Println("Updated Profile by Name:", updatedProfileByName)
	}
}
