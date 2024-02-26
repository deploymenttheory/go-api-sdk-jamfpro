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

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	siteID := 1 // Replace 1 with the actual site ID

	// Fetch the site by ID
	site, err := client.GetSiteByID(siteID)
	if err != nil {
		log.Fatalf("Error fetching site by ID: %v", err)
	}

	fmt.Printf("Site ID: %d, Name: %s\n", site.ID, site.Name)
}
