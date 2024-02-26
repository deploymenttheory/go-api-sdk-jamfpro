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

	// ID of the distribution point to delete
	distributionPointID := 1 // Replace with the actual ID

	// Call DeleteDistributionPointByID function
	err = client.DeleteDistributionPointByID(distributionPointID)
	if err != nil {
		log.Fatalf("Error deleting distribution point: %v", err)
	}

	fmt.Println("Distribution Point deleted successfully")
}
