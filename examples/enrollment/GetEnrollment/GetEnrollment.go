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

	// Call GetEnrollment function
	enrollment, err := client.GetEnrollment()
	if err != nil {
		log.Fatalf("Error getting enrollment configuration: %v", err)
	}

	// Pretty print the enrollment configuration in JSON
	JSON, err := json.MarshalIndent(enrollment, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling enrollment data: %v", err)
	}
	fmt.Println("Current Enrollment Configuration:\n", string(JSON))
}
