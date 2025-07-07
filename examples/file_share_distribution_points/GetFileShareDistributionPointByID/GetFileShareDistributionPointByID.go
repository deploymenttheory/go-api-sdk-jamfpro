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

	// ID of the distribution point to delete
	distributionPointID := "135" // Replace with the actual ID

	fmt.Print("\nhello???")
	// Call GetDistributionPointByID function
	distributionPoint, err := client.GetDistributionPointByID(distributionPointID)
	if err != nil {
		log.Fatalf("Error fetching distribution point: %v", err)
	}

	// Pretty print the newly created distribution point in XML
	distributionPointJSON, err := json.MarshalIndent(distributionPoint, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created distribution point data: %v", err)
	}
	fmt.Println("Created Distribution Point:\n", string(distributionPointJSON))

	fmt.Println("Distribution Point retrieved successfully")
}
