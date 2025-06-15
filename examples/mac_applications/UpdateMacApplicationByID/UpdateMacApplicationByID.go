package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Define the Mac Application to be updated
	updateMacApp := jamfpro.ResourceMacApplications{
		General: jamfpro.MacApplicationsSubsetGeneral{
			Name:     "TextWrangler.app",
			Version:  "5.5.2",
			IsFree:   jamfpro.BoolPtr(true),
			BundleID: "com.barebones.textwrangler",
			URL:      "https://itunes.apple.com/us/app/textwrangler/id404010395?mt=12&uo=4",
			Category: &jamfpro.SharedResourceCategory{ID: -1, Name: "Unknown"},
			Site:     &jamfpro.SharedResourceSite{ID: -1, Name: "None"},
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
			VPP: jamfpro.MacAppSubsetSelfServiceVPP{
				AssignVPPDeviceBasedLicenses: jamfpro.BoolPtr(false),
				VPPAdminAccountID:            -1,
			},
		},
	}

	macAppID := "1" // Replace with your Mac application ID

	// Call UpdateMacApplicationByID
	updatedMacApp, err := client.UpdateMacApplicationByID(macAppID, updateMacApp) // Replace 123 with the actual ID
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
