package main

import (
	"encoding/xml"
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

	// Define the Mac Application to be updated
	updateMacApp := jamfpro.ResponseMacApplications{
		General: jamfpro.MacAppDataSubsetGeneral{
			Name:     "TextWrangler.app",
			Version:  "5.5.2",
			IsFree:   true,
			BundleID: "com.barebones.textwrangler",
			URL:      "https://itunes.apple.com/us/app/textwrangler/id404010395?mt=12&uo=4",
			Category: jamfpro.MacAppCategory{ID: -1, Name: "Unknown"},
			Site:     jamfpro.MacAppSite{ID: -1, Name: "None"},
		},
		Scope: jamfpro.MacAppDataSubsetScope{
			AllComputers: false,
			AllJSSUsers:  false,
		},
		SelfService: jamfpro.MacAppDataSubsetSelfService{
			InstallButtonText:           "Install",
			SelfServiceDescription:      "Installs the TextWrangler application",
			ForceUsersToViewDescription: true,
			SelfServiceIcon:             jamfpro.MacAppSelfServiceIcon{},
			FeatureOnMainPage:           true,
			SelfServiceCategories:       []jamfpro.MacAppSelfServiceCategory{},
			Notification:                "string",
			NotificationSubject:         "TextWrangler is Available to Install",
			NotificationMessage:         "You can install TextWrangler by clicking this link or going to Self Service",
			VPP: jamfpro.MacAppVPP{
				AssignVPPDeviceBasedLicenses: false,
				VPPAdminAccountID:            -1,
			},
		},
	}

	id := 1 // Replace with your Mac application ID
	// Call UpdateMacApplicationByID
	updatedMacApp, err := client.UpdateMacApplicationByID(id, updateMacApp) // Replace 123 with the actual ID
	if err != nil {
		log.Fatalf("Error updating Mac Application by ID: %v", err)
	}

	// Print the updated Mac Application details
	updatedMacAppXML, err := xml.MarshalIndent(updatedMacApp, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling updated Mac Application data: %v", err)
	}
	fmt.Println("Updated Mac Application:\n", string(updatedMacAppXML))

}