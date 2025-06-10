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

	// Define the client management ID for which to get DDM status items
	clientManagementID := "10" // Replace with the actual client management ID

	// Call GetDDMStatusItems function
	statusItemsResponse, err := client.GetDDMStatusItems(clientManagementID)
	if err != nil {
		log.Fatalf("Error retrieving DDM status items: %v", err)
	}

	fmt.Printf("Successfully retrieved DDM status items for client management ID: %s\n", clientManagementID)
	for _, item := range statusItemsResponse.StatusItems {
		fmt.Printf("Key: %s, Value: %s, Last Update Time: %s\n", item.Key, item.Value, item.LastUpdateTime)
	}
}
