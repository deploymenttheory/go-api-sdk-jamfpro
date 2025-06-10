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

	AttributeIDToDelete := "1" // Replace 1 with the actual attribute ID

	// Deleting the attribute with specified ID
	err = client.DeleteComputerExtensionAttributeByID(AttributeIDToDelete)
	if err != nil {
		log.Fatalf("Error deleting Computer Extension Attribute by ID: %v", err)
	}

	fmt.Printf("Successfully deleted Computer Extension Attribute with ID: %d\n", AttributeIDToDelete)
}
