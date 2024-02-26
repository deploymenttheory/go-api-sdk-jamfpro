package main

import (
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

	// Define the computer group name for deletion
	groupName := "NewGroupName" // Replace with the actual name

	// Call DeleteComputerGroupByName function
	err = client.DeleteComputerGroupByName(groupName)
	if err != nil {
		log.Fatalf("Error deleting Computer Group by Name: %v", err)
	}

	log.Println("Successfully deleted the Computer Group by name.")
}
