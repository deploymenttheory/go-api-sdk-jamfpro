package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client" // Import http_client for logging
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/path/to/your/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro client
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		LogLevel:     logLevel,
		Logger:       logger,
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the bundle ID of the mobile device application you want to update
	bundleID := "com.jamfsoftware.selfservice"

	// Define the mobile device application data for update
	updateApp := &jamfpro.ResponseMobileDeviceApplication{
		General: jamfpro.MobileDeviceApplicationGeneral{
			Name:        "Jamf Self Service",
			DisplayName: "Jamf Self Service",
			Description: "Jamf Self Service empowers you to be more productive...",
			BundleID:    "com.jamfsoftware.selfservice",
			Version:     "10.10.6",
			InternalApp: false,
			OsType:      "iOS", // iOS or tvOS
			Category: jamfpro.MobileDeviceApplicationCategory{
				ID:   -1,
				Name: "No category assigned",
			},
			IPA: jamfpro.MobileDeviceApplicationIPA{
				// Set IPA details here
			},
			Icon: jamfpro.MobileDeviceApplicationIcon{
				ID:   27,
				Name: "1024x1024bb.png",
				URI:  "string",
				Data: "Base64EncodedString",
			},
			// ... other fields ...
		},
		Scope: jamfpro.MobileDeviceApplicationScope{
			// Populate the Scope details
		},
		SelfService: jamfpro.MobileDeviceApplicationSelfService{
			SelfServiceDescription: "Jamf Self Service empowers you...",
			// ... other fields ...
		},
		VPP: jamfpro.MobileDeviceApplicationVPP{
			// Populate the VPP details
		},
		AppConfiguration: jamfpro.MobileDeviceApplicationConfiguration{
			Preferences: "",
		},
	}

	// Perform the update
	updatedApp, err := client.UpdateMobileDeviceApplicationByApplicationBundleID(bundleID, updateApp)
	if err != nil {
		fmt.Println("Error updating mobile device application:", err)
	} else {
		fmt.Println("Updated Mobile Device Application:", updatedApp)
	}
}
