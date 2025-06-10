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

	// Create update request
	updateRequest := &jamfpro.ResourceDeviceEnrollmentUpdate{
		Name:                  "JAMF - LBGBSTAGING Updated", // Example name update
		SupervisionIdentityId: "-1",                         // Optional
		SiteId:                "-1",                         // Optional
	}

	// Update device enrollment
	deviceEnrollmentID := "1" // Using the known device enrollment ID
	response, err := client.UpdateDeviceEnrollmentMetadataByID(deviceEnrollmentID, updateRequest)
	if err != nil {
		log.Fatalf("Error updating device enrollment: %v", err)
	}

	// Pretty print the updated device enrollment using JSON marshaling
	responseJSON, err := json.MarshalIndent(response, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling updated device enrollment data: %v", err)
	}
	fmt.Println("Updated Device Enrollment:", string(responseJSON))
}
