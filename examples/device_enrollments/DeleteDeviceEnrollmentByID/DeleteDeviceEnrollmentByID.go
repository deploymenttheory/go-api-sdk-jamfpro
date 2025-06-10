package main

import (
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

	// Delete device enrollment
	deviceEnrollmentID := "1" // Using the known device enrollment ID
	err = client.DeleteDeviceEnrollmentByID(deviceEnrollmentID)
	if err != nil {
		log.Fatalf("Error deleting device enrollment: %v", err)
	}

	fmt.Printf("Successfully deleted device enrollment with ID: %s\n", deviceEnrollmentID)
}
