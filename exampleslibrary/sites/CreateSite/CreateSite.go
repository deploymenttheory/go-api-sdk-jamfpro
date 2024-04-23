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

	// Define the site to be created
	newSite := &jamfpro.SharedResourceSite{
		Name: "NewSiteName", // Replace "NewSiteName" with the actual name for the new site
	}

	// Create the site
	createdSite, err := client.CreateSite(newSite)
	if err != nil {
		log.Fatalf("Error creating site: %v", err)
	}

	fmt.Printf("Site created successfully. ID: %d, Name: %s\n", createdSite.ID, createdSite.Name)
}
