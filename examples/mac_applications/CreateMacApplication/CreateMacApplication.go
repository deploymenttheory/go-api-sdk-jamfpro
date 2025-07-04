package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/Shared/GitHub/go-api-sdk-jamfpro/localtesting/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Define a new Mac Application
	newMacApp := jamfpro.ResourceMacApplications{
		General: jamfpro.MacApplicationsSubsetGeneral{
			Name:           "TextWrangler.app",
			Version:        "5.5.2",
			IsFree:         jamfpro.BoolPtr(true),
			BundleID:       "com.barebones.textwrangler",
			URL:            "https://itunes.apple.com/us/app/textwrangler/id404010395?mt=12&uo=4",
			Category:       &jamfpro.SharedResourceCategory{ID: -1, Name: "Unknown"},
			Site:           &jamfpro.SharedResourceSite{ID: -1, Name: "None"},
			DeploymentType: "Install Automatically/Prompt Users to Install",
		},
		Scope: jamfpro.MacApplicationsSubsetScope{
			AllComputers: jamfpro.BoolPtr(false),
			AllJSSUsers:  jamfpro.BoolPtr(false),
		},
		SelfService: jamfpro.MacAppSubsetSelfService{
			InstallButtonText:           "Install",
			SelfServiceDescription:      "Installs the TextWrangler application",
			ForceUsersToViewDescription: jamfpro.BoolPtr(true),
			SelfServiceIcon:             jamfpro.SharedResourceSelfServiceIcon{},
			FeatureOnMainPage:           jamfpro.BoolPtr(true),
			SelfServiceCategories:       []jamfpro.MacAppSubsetSelfServiceCategories{},
			Notification:                "string",
			NotificationSubject:         "TextWrangler is Available to Install",
			NotificationMessage:         "You can install TextWrangler by clicking this link or going to Self Service",
		},
		VPP: jamfpro.MacAppSubsetVPP{
			AssignVPPDeviceBasedLicenses: jamfpro.BoolPtr(false),
			VPPAdminAccountID:            -1,
		},
	}

	// Call CreateMacApplication
	createdMacApp, err := client.CreateMacApplication(&newMacApp)
	if err != nil {
		log.Fatalf("Error creating Mac Application: %v", err)
	}

	// Pretty print the created Mac Application in XML
	macAppXML, err := xml.MarshalIndent(createdMacApp, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Mac Application data: %v", err)
	}
	fmt.Println("Created Mac Application:\n", string(macAppXML))
}
