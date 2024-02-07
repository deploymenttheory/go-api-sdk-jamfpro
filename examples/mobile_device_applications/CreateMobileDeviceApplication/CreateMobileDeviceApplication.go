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

	// Define a new mobile device application
	newApp := &jamfpro.ResourceMobileDeviceApplication{
		General: jamfpro.MobileDeviceApplicationSubsetGeneral{
			Name:        "Jamf Self Service",
			DisplayName: "Jamf Self Service",
			Description: "Jamf Self Service empowers you to be more productive...",
			BundleID:    "com.jamfsoftware.selfservice",
			Version:     "10.10.6",
			InternalApp: true,
			OsType:      "iOS", // iOS or tvOS
			Category: jamfpro.SharedResourceCategory{
				ID:   -1,
				Name: "No category assigned",
			},
			IPA: jamfpro.MobileDeviceApplicationSubsetGeneralIPA{
				Name: "IPAName",
				URI:  "http://example.com/ipa",
				Data: "Base64EncodedString",
			},
			Icon: jamfpro.MobileDeviceApplicationSubsetIcon{
				ID:   27,
				Name: "1024x1024bb.png",
				URI:  "http://example.com/icon.png",
				Data: "Base64EncodedString",
			},
			// Populate other fields as necessary...
		},
	}

	createdApp, err := client.CreateMobileDeviceApplication(newApp)
	if err != nil {
		fmt.Println("Error creating mobile device application:", err)
	} else {
		fmt.Println("Created Mobile Device Application:", createdApp)
	}
}
