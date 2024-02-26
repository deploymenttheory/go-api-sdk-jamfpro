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

	// Name of the distribution point to delete
	distributionPointName := "New York Share" // Replace with the actual name

	// Call DeleteDistributionPointByName function
	err = client.DeleteDistributionPointByName(distributionPointName)
	if err != nil {
		log.Fatalf("Error deleting distribution point by name '%s': %v", distributionPointName, err)
	}

	fmt.Printf("Distribution Point '%s' deleted successfully\n", distributionPointName)
}
