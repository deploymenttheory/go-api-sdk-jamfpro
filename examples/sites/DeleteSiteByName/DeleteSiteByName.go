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

	siteNameToDelete := "Site Name" // Replace with the actual site name

	err = client.DeleteSiteByName(siteNameToDelete)
	if err != nil {
		log.Fatalf("Error deleting site by name: %v", err)
	}

	fmt.Printf("Site with name '%s' deleted successfully.\n", siteNameToDelete)
}
