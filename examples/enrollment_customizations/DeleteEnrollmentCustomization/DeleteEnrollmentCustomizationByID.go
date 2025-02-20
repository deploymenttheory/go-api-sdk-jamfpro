package main

import (
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

	// Define the ID of the enrollment customization to delete
	customizationID := "2" // Replace with your actual ID

	// Delete the enrollment customization
	err = client.DeleteEnrollmentCustomizationByID(customizationID)
	if err != nil {
		log.Fatalf("Failed to delete enrollment customization: %v", err)
	}

	fmt.Printf("Successfully deleted enrollment customization with ID: %s\n", customizationID)
}
