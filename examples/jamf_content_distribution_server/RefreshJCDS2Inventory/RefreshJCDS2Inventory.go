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

	// Call Function to refresh JCDS2 inventory
	err = client.RefreshJCDS2Inventory()
	if err != nil {
		log.Fatalf("Error refreshing JCDS 2.0 inventory: %v", err)
	}

	fmt.Println("Successfully refreshed JCDS 2.0 inventory")
}
