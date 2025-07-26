package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/Shared/GitHub/go-api-sdk-jamfpro/localtesting/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Replace with an actual ID
	id := "2"

	extensionAttributes, err := client.GetMobileDeviceExtensionAttributeByID(id)
	if err != nil {
		log.Fatalf("Error fetching mobile device extension attribute by ID: %v", err)
	}

	// Pretty print the extension attribute details in JSON
	extensionAttributesJSON, err := json.MarshalIndent(extensionAttributes, "", "    ") // Indent with 4 spaces and use '='
	if err != nil {
		log.Fatalf("Error marshaling mobile device extension attribute details data: %v", err)
	}
	fmt.Println("Mobile Device Extension Attribute Details:\n", string(extensionAttributesJSON))
}
