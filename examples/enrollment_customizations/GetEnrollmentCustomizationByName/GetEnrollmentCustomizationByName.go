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

	// Define the name you want to search for
	customizationName := "test" // Replace with your actual name

	// Get the enrollment customization by name
	customization, err := client.GetEnrollmentCustomizationByName(customizationName)
	if err != nil {
		log.Fatalf("Failed to get enrollment customization by name: %v", err)
	}

	// Pretty print the enrollment customization details in JSON
	JSON, err := json.MarshalIndent(customization, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling enrollment customization data: %v", err)
	}
	fmt.Println("Enrollment Customization Details:\n", string(JSON))
}
