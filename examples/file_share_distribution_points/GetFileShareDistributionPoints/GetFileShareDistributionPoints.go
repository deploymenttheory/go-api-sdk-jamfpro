package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/work/Documents/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Call GetDistributionPointByID function
	distributionPoint, err := client.GetDistributionPoints()
	if err != nil {
		log.Fatalf("Error fetching distribution point: %v", err)
	}

	// Pretty print the newly created distribution point in XML
	distributionPointJSON, err := json.MarshalIndent(distributionPoint, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created distribution point data: %v", err)
	}
	fmt.Println("Retrieved Distribution Point:\n", string(distributionPointJSON))

	fmt.Println("Distribution Point retrieved successfully")
}
