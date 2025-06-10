package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "./clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	originalSiteName := "Site Name"        // Replace with the actual site name
	updatedSiteName := "Updated Site Name" // Replace with the updated site name

	// Define the updated site data
	updatedSite := &jamfpro.SharedResourceSite{
		Name: updatedSiteName,
	}

	// Update the site by name
	updated, err := client.UpdateSiteByName(originalSiteName, updatedSite)
	if err != nil {
		log.Fatalf("Error updating site by name: %v", err)
	}

	fmt.Printf("Site updated successfully. ID: %d, Name: %s\n", updated.ID, updated.Name)
}
