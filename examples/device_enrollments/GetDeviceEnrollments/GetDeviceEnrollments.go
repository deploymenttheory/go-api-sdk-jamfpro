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

	// Fetch Device Enrollments
	response, err := client.GetDeviceEnrollments(url.Values{})
	if err != nil {
		log.Fatalf("Error fetching Device enrollment instances: %v", err)
	}

	// Pretty print the fetched Device enrollment instances using JSON marshaling
	responseJSON, err := json.MarshalIndent(response, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Device enrollment instances data: %v", err)
	}
	fmt.Println("Fetched Device enrollment instances:", string(responseJSON))
}
