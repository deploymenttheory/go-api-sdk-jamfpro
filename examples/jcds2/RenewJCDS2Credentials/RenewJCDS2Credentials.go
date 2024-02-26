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

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Call RenewJCDS2Credentials function
	JCDS2Creds, err := client.RenewJCDS2Credentials()
	if err != nil {
		log.Fatalf("Error fetching JCDS 2 files: %v", err)
	}

	// Pretty print the JCDS 2 files in JSON
	response, err := json.MarshalIndent(JCDS2Creds, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling JCDS 2 credentials: %v", err)
	}
	fmt.Println("Fetched JCDS 2 credentials:\n", string(response))
}
