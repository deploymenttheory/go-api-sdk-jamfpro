package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

const (
	configFilePath = "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"
	attributeName  = "Battery Cycle Count" // replace this with the name of the attribute you want to delete
)

func main() {
	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		DebugMode:    true,
		Logger:       jamfpro.NewDefaultLogger(),
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Call DeleteComputerExtensionAttributeByName function
	err = client.DeleteComputerExtensionAttributeByNameByID(attributeName)
	if err != nil {
		log.Fatalf("Error deleting Computer Extension Attribute by name: %v", err)
	}

	fmt.Println("Successfully deleted Computer Extension Attribute:", attributeName)
}
