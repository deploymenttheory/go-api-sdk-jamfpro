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

	// Define the client management ID and key for which to get a specific DDM status item
	clientManagementID := "10"      // Replace with the actual client management ID
	key := "device.identifier.udid" // Replace with the actual status item key

	// Call GetDDMStatusItem function
	statusItem, err := client.GetDDMStatusItem(clientManagementID, key)
	if err != nil {
		log.Fatalf("Error retrieving DDM status item: %v", err)
	}

	fmt.Printf("Successfully retrieved DDM status item for client management ID: %s and key: %s\n", clientManagementID, key)
	fmt.Printf("Key: %s, Value: %s, Last Update Time: %s\n", statusItem.Key, statusItem.Value, statusItem.LastUpdateTime)
}
