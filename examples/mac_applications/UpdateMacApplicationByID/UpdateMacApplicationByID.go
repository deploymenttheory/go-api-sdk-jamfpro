package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-http-client/logger"
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
	logLevel := logger.LogLevelWarn // LogLevelNone / LogLevelDebug / LogLevelInfo / LogLevelError

	// Configuration for the jamfpro
	config := httpclient.Config{
		InstanceName: authConfig.InstanceName,
		Auth: httpclient.AuthConfig{
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

	// Define the Mac Application to be updated
	updateMacApp := jamfpro.ResourceMacApplications{
		General: jamfpro.MacApplicationsSubsetGeneral{
			Name:     "TextWrangler.app",
			Version:  "5.5.2",
			IsFree:   true,
			BundleID: "com.barebones.textwrangler",
			URL:      "https://itunes.apple.com/us/app/textwrangler/id404010395?mt=12&uo=4",
			Category: jamfpro.SharedResourceCategory{ID: -1, Name: "Unknown"},
			Site:     jamfpro.SharedResourceSite{ID: -1, Name: "None"},
		},
		Scope: jamfpro.MacApplicationsSubsetScope{
			AllComputers: false,
			AllJSSUsers:  false,
		},
		SelfService: jamfpro.MacAppSubsetSelfService{
			InstallButtonText:           "Install",
			SelfServiceDescription:      "Installs the TextWrangler application",
			ForceUsersToViewDescription: true,
			SelfServiceIcon:             jamfpro.SharedResourceSelfServiceIcon{},
			FeatureOnMainPage:           true,
			SelfServiceCategories:       []jamfpro.MacAppSubsetSelfServiceCategories{},
			Notification:                "string",
			NotificationSubject:         "TextWrangler is Available to Install",
			NotificationMessage:         "You can install TextWrangler by clicking this link or going to Self Service",
			VPP: jamfpro.MacAppSubsetSelfServiceVPP{
				AssignVPPDeviceBasedLicenses: false,
				VPPAdminAccountID:            -1,
			},
		},
	}

	macAppId := 1 // Replace with your Mac application ID

	// Call UpdateMacApplicationByID
	updatedMacApp, err := client.UpdateMacApplicationByID(macAppId, updateMacApp) // Replace 123 with the actual ID
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
