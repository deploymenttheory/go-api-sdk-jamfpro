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

	// Call GetEnrollmentMessages to retrieve all configured language messages
	messages, err := client.GetEnrollmentMessages()
	if err != nil {
		log.Fatalf("Error retrieving configured enrollment language messages: %v", err)
	}

	// Pretty print the full list of language messages
	JSON, err := json.MarshalIndent(messages, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling configured messages: %v", err)
	}

	fmt.Println("Configured Enrollment Language Messages:\n", string(JSON))
}
