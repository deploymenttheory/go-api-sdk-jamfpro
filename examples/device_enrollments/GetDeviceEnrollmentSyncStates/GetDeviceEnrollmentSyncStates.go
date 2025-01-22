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

	// Fetch Device Enrollment Sync States
	deviceEnrollmentID := "1" // Using the known device enrollment ID from the system
	response, err := client.GetDeviceEnrollmentSyncStates(deviceEnrollmentID)
	if err != nil {
		log.Fatalf("Error fetching device enrollment sync states: %v", err)
	}

	// Pretty print the fetched sync states using JSON marshaling
	responseJSON, err := json.MarshalIndent(response, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling device enrollment sync states data: %v", err)
	}
	fmt.Println("Fetched Device Enrollment Sync States:", string(responseJSON))
}
