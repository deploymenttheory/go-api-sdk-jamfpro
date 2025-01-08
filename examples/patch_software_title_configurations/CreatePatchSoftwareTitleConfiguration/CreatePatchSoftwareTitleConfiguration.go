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
		DisplayName:            "Adobe Audition CC 2015",
		CategoryID:             "-1",
		SiteID:                 "-1",
		UiNotifications:        true,
		EmailNotifications:     true,
		SoftwareTitleNameId:    "048",
		SoftwareTitleName:      "Adobe Audition CC 2015",
		SoftwareTitlePublisher: "Adobe",
		JamfOfficial:           true,
		PatchSourceName:        "Jamf",
	}

	// Call the CreatePatchSoftwareTitleConfiguration method
	response, err := client.CreatePatchSoftwareTitleConfiguration(*newConfig)
	if err != nil {
		fmt.Println("Error creating Patch Software Title Configuration:", err)
		return
	}

	// Print the response
	fmt.Printf("Created Patch Software Title Configuration with ID: %s, Href: %s\n", response.ID, response.Href)
}
