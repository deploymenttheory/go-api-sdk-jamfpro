package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

const AttributeIDToDelete = 123 // Replace 123 with the ID of the attribute you wish to delete

func main() {
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		DebugMode:    true,
		Logger:       jamfpro.NewDefaultLogger(),
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Deleting the attribute with specified ID
	err = client.DeleteComputerExtensionAttributeByID(AttributeIDToDelete)
	if err != nil {
		log.Fatalf("Error deleting Computer Extension Attribute by ID: %v", err)
	}

	fmt.Printf("Successfully deleted Computer Extension Attribute with ID: %d\n", AttributeIDToDelete)
}
