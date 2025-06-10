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

	// Define group ID to delete
	groupID := "1"

	// Call function
	err = client.DeleteSmartComputerGroupByID(groupID)
	if err != nil {
		log.Fatalf("Error deleting smart computer group: %v", err)
	}

	fmt.Printf("Successfully deleted smart computer group with ID %s\n", groupID)
}
