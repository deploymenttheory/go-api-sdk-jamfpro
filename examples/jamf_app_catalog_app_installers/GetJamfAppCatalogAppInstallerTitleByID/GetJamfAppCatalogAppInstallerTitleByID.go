package main

import (
	"encoding/json"
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

	// Define a app installer ID for testing
	appInstallerID := "518"

	// Call GetJamfAppCatalogAppInstallerTitleByID function
	appInstaller, err := client.GetJamfAppCatalogAppInstallerTitleByID(appInstallerID)
	if err != nil {
		log.Fatalf("Error fetching JamfAppCatalogAppInstallerTitle by ID: %v", err)
	}

	// Pretty print the app Installer in JSON
	appInstallerJSON, err := json.MarshalIndent(appInstaller, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Jamf App Catalog App Title data: %v", err)
	}
	fmt.Println("Fetched Jamf App Catalog App Title:\n", string(appInstallerJSON))
}
