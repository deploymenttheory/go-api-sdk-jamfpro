package main

import (
	"encoding/json"
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

	// Call GetReturnToService function
	returnToService, err := client.GetReturnToServiceConfigurations()
	if err != nil {
		log.Fatalf("Error fetching return to service  properties: %v", err)
	}

	// Pretty print the return to service  files in JSON
	response, err := json.MarshalIndent(returnToService, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling return to service data: %v", err)
	}
	fmt.Println("Fetched return to service properties:\n", string(response))
}
