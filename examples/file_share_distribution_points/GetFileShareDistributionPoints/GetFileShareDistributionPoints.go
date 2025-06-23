package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/lloyds/Documents/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Call GetDistributionPoints function
	distributionPoints, err := client.GetDistributionPoints()
	if err != nil {
		log.Fatalf("Error fetching distribution points: %v", err)
	}

	// Pretty print the distribution points in XML
	distributionPointsXML, err := json.Marshal(distributionPoints) // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling distribution points data: %v", err)
	}
	fmt.Println("Fetched Distribution Points:\n", string(distributionPointsXML))
}
