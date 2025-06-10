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

	// Define the feature toggle payload
	payload := &jamfpro.ResourceManagedSoftwareUpdateFeatureToggle{
		Toggle: true, // Set to true or false based on the desired state
	}

	// Update the feature toggle
	response, err := client.UpdateManagedSoftwareUpdateFeatureToggle(payload)
	if err != nil {
		log.Fatalf("Error updating managed software update feature toggle: %v", err)
	}

	// Pretty print the response
	responseJSON, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling response data: %v", err)
	}

	fmt.Printf("Managed software update feature toggle updated successfully. Response: \n%s\n", string(responseJSON))
}
