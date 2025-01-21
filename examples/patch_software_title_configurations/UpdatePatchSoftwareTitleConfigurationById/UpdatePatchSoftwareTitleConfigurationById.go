package main

import (
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"

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

	version := "134.0"                                                // Version to be patched
	pkgFilePath := "/Users/dafyddwatkins/Downloads/Firefox 134.0.pkg" // Path to the package file
	softwareTitleConfigurationId := "14"

	// Create a new PatchSoftwareTitleConfiguration
	newConfig := &jamfpro.ResourcePatchSoftwareTitleConfiguration{
		CategoryID:         "-1",
		SiteID:             "-1",
		SoftwareTitleID:    "14",
		UiNotifications:    true, // UI Notifications enabled (match Python behavior)
		EmailNotifications: true, // Email Notifications enabled (match Python behavior)
		Packages: []jamfpro.PatchSoftwareTitleConfigurationSubsetPackage{
			{
				PackageId:   "57",                       // Package ID from uploaded package
				Version:     version,                    // Version of the package
				DisplayName: filepath.Base(pkgFilePath), // Package display name
			},
		},
	}

	// Call the CreatePatchSoftwareTitleConfiguration method
	response, err := client.UpdatePatchSoftwareTitleConfigurationById(softwareTitleConfigurationId, *newConfig)
	if err != nil {
		log.Fatalf("Error creating Patch Software Title Configuration: %v", err)
	}

	// Pretty print the network segments in JSON
	softwareTitleJSON, err := json.MarshalIndent(response, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling network segments data: %v", err)
	}
	fmt.Println("Network Segments:\n", string(softwareTitleJSON))
}
