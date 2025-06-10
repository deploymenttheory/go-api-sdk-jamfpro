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

	// Call Function
	JCDS2Files, err := client.GetJCDS2Packages()
	if err != nil {
		log.Fatalf("Error fetching JCDS 2 packages: %v", err)
	}

	// Pretty print the JSON
	response, err := json.MarshalIndent(JCDS2Files, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling JCDS 2 package data: %v", err)
	}
	fmt.Println("Fetched JCDS 2 packages:\n", string(response))
}
