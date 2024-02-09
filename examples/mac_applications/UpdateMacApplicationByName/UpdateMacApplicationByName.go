package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	loadedConfig, err := jamfpro.LoadClientConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	//logLevel := logger.LogLevelInfo // LogLevelNone / LogLevelDebug / LogLevelInfo / LogLevelError

	// Configuration for the HTTP client
	config := httpclient.ClientConfig{
		Auth: httpclient.AuthConfig{
			ClientID:     loadedConfig.Auth.ClientID,
			ClientSecret: loadedConfig.Auth.ClientSecret,
		},
		Environment: httpclient.EnvironmentConfig{
			APIType:      loadedConfig.Environment.APIType,
			InstanceName: loadedConfig.Environment.InstanceName,
		},
		ClientOptions: httpclient.ClientOptions{
			LogLevel:          loadedConfig.ClientOptions.LogLevel,
			HideSensitiveData: loadedConfig.ClientOptions.HideSensitiveData,
			LogOutputFormat:   loadedConfig.ClientOptions.LogOutputFormat,
		},
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
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

	macAppName := "TextWrangler.app" // Replace with your vpp Mac application name
	// Call UpdateMacApplicationByName
	updatedMacApp, err := client.UpdateMacApplicationByName(macAppName, updateMacApp)
	if err != nil {
		log.Fatalf("Error updating Mac Application by Name: %v", err)
	}

	// Print the updated Mac Application details
	updatedMacAppXML, err := xml.MarshalIndent(updatedMacApp, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling updated Mac Application data: %v", err)
	}
	fmt.Println("Updated Mac Application:\n", string(updatedMacAppXML))

}
