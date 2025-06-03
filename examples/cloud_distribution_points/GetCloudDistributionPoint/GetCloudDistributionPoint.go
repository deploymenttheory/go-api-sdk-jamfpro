package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/Shared/GitHub/go-api-sdk-jamfpro/localtesting/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Call GetCloudDistributionPoint function
	groups, err := client.GetCloudDistributionPoint()
	if err != nil {
		log.Fatalf("Error fetching Cloud Distribution Point Details: %v", err)
	}

	// Pretty print the details in JSON
	cloudDistributionPointJSON, err := json.MarshalIndent(groups, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Cloud Distribution Point Details data: %v", err)
	}
	fmt.Println("Fetched Cloud Distribution Point Details:\n", string(cloudDistributionPointJSON))
}
