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

	// Example software title configuration ID
	softwareTitleID := "13" // Example ID

	// Get the patch software title extension attributes
	attributes, err := client.GetPatchSoftwareTitleExtensionAttributes(softwareTitleID)
	if err != nil {
		log.Fatalf("Error getting patch software title extension attributes: %v", err)
	}

	// Pretty print the JSON response
	response, err := json.MarshalIndent(attributes, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling patch software title extension attributes data: %v", err)
	}
	fmt.Printf("Found patch software title extension attributes for ID %s:\n%s\n", softwareTitleID, string(response))
}
