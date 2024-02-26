package main

import (
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

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// ID of the software update server to delete
	serverID := 1

	// Call DeleteSoftwareUpdateServerByID
	err = client.DeleteSoftwareUpdateServerByID(serverID)
	if err != nil {
		log.Fatalf("Error deleting software update server by ID: %v", err)
	}

	fmt.Printf("Successfully deleted software update server with ID %d\n", serverID)
}
