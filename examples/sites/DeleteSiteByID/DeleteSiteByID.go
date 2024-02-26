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

	siteIDToDelete := 1 // Replace 1 with the actual site ID

	err = client.DeleteSiteByID(siteIDToDelete)
	if err != nil {
		log.Fatalf("Error deleting site by ID: %v", err)
	}

	fmt.Printf("Site with ID %d deleted successfully.\n", siteIDToDelete)
}
