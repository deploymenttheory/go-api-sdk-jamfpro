package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

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

	// Fetch Device Enrollment History
	deviceEnrollmentID := "1" // Using the known device enrollment ID from the system
	response, err := client.GetDeviceEnrollmentHistory(deviceEnrollmentID, url.Values{})
	if err != nil {
		log.Fatalf("Error fetching device enrollment history: %v", err)
	}

	// Pretty print the fetched device enrollment history using JSON marshaling
	responseJSON, err := json.MarshalIndent(response, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling device enrollment history data: %v", err)
	}
	fmt.Println("Fetched Device Enrollment History:", string(responseJSON))
}
