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

	// Define a new mobile device application
	newApp := &jamfpro.ResponseMobileDeviceApplication{
		General: jamfpro.MobileDeviceApplicationGeneral{
			Name:        "Jamf Self Service",
			DisplayName: "Jamf Self Service",
			Description: "Jamf Self Service empowers you to be more productive...",
			BundleID:    "com.jamfsoftware.selfservice",
			Version:     "10.10.6",
			InternalApp: true,
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

	createdApp, err := client.CreateMobileDeviceApplication(newApp)
	if err != nil {
		fmt.Println("Error creating mobile device application:", err)
	} else {
		fmt.Println("Created Mobile Device Application:", createdApp)
	}
}
