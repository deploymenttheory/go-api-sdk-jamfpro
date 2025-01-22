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

	// Fetch Device Enrollment by Name
	deviceEnrollmentName := "Example Device Enrollment Instance" // Replace with actual name
	response, err := client.GetDeviceEnrollmentByName(deviceEnrollmentName)
	if err != nil {
		log.Fatalf("Error fetching device enrollment by name: %v", err)
	}

	// Pretty print the fetched device enrollment using JSON marshaling
	responseJSON, err := json.MarshalIndent(response, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling device enrollment data: %v", err)
	}
	fmt.Println("Fetched Device Enrollment:", string(responseJSON))
}
