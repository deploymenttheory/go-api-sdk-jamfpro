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
		JamfOfficial:           true,
		DisplayName:            "Adobe AIR",
		CategoryID:             "-1",
		SiteID:                 "-1",
		UiNotifications:        true,
		EmailNotifications:     true,
		SoftwareTitleID:        "11",
		SoftwareTitleName:      "Adobe AIR",
		SoftwareTitleNameId:    "0AE",
		SoftwareTitlePublisher: "HARMAN",
		PatchSourceName:        "Jamf",
		PatchSourceEnabled:     true,
		Packages: []jamfpro.PatchSoftwareTitleConfigurationSubsetPackage{
			{
				PackageId:   "38",
				DisplayName: "Firefox 133.0.3.pkg",
				Version:     "51.1.2.2",
			},
		},
		ExtensionAttributes: []jamfpro.PatchSoftwareTitleConfigurationSubsetExtensionAttribute{
			{
				Accepted: false,
				EaID:     "jamf-patch-adobe-air",
			},
		},
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
