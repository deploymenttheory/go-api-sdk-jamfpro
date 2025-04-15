package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

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

	// Call GetJamfAppCatalogAppInstallerTitles function
	appInstallers, err := client.GetJamfAppCatalogAppInstallerTitles(url.Values{})
	if err != nil {
		log.Fatalf("Error fetching Jamf App Catalog App Installer list: %v", err)
	}

	// Pretty print the app Installer in JSON
	appInstallerJSON, err := json.MarshalIndent(appInstallers, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Jamf App Catalog App Installer list data: %v", err)
	}
	fmt.Println("Fetched Jamf App Catalog App Installer list:\n", string(appInstallerJSON))
}
