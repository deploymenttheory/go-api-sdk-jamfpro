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

	loginCustomization, err := client.GetLoginCustomization()
	if err != nil {
		log.Fatalf("Error fetching Mac Applications: %v", err)
	}

	// Pretty print the login customization details
	responseJSON, err := json.MarshalIndent(loginCustomization, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling login customization data: %v", err)
	}
	fmt.Println("Fetched login customization Details:", string(responseJSON))
}
