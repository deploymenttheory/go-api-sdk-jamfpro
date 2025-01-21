package main

import (
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

	// Create a new PatchSoftwareTitleConfiguration
	newConfig := &jamfpro.ResourcePatchSoftwareTitleConfiguration{
		// Required fields
		DisplayName:     "Adobe After Effects CC", // %patch_software_title_name%
		SoftwareTitleID: "2B8",                    // %patch_softwaretitle_id%

		// Optional fields with defaults
		CategoryID:         "-1", // Optional, defaults to "-1"
		SiteID:             "-1", // Optional, defaults to "-1"
		UiNotifications:    true, // Optional, defaults to false
		EmailNotifications: true, // Optional, defaults to false
		ExtensionAttributes: []jamfpro.PatchSoftwareTitleConfigurationSubsetExtensionAttribute{
			{
				EaID:     "jamf-patch-adobe-after-effects-cc", // Required if extension attributes are included
				Accepted: false,                               // Optional, defaults to false
			},
		},
	}

	// Call the CreatePatchSoftwareTitleConfiguration method
	response, err := client.CreatePatchSoftwareTitleConfiguration(*newConfig)
	if err != nil {
		log.Fatalf("Error creating Patch Software Title Configuration: %v", err)
	}

	// Print the response
	fmt.Printf("Created Patch Software Title Configuration with ID: %s, Href: %s\n", response.ID, response.Href)
}
