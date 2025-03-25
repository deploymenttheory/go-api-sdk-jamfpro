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

	// Call GetJamfAppCatalogAppInstallerGlobalSettings function
	appInstallers, err := client.GetJamfAppCatalogAppInstallerGlobalSettings()
	if err != nil {
		log.Fatalf("Error fetching Jamf App Catalog App Installer Global Settings: %v", err)
	}

	// Pretty print the app Installer in JSON
	appInstallerJSON, err := json.MarshalIndent(appInstallers, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Jamf App Catalog App Installer Global Settings: %v", err)
	}
	fmt.Println("Fetched Jamf App Catalog App Installer Global Settings:\n", string(appInstallerJSON))
}
