package main

import (
	"encoding/json"
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

	// Call Function
	HealthCheck, err := client.GetHealthCheck()
	if err != nil {
		log.Fatalf("Error fetching Health Check properties: %v", err)
	}

	// Pretty print the Health Check files in XML
	response, err := json.MarshalIndent(HealthCheck, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Health Check package data: %v", err)
	}
	fmt.Println("Fetched Health Check properties:\n", string(response))
}
