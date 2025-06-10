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

	// The ID of the class you want to delete
	classIDToDelete := "123" // Replace with the actual ID

	// Call the delete function with the class ID
	err = client.DeleteClassByID(classIDToDelete)
	if err != nil {
		log.Fatalf("Error deleting class by ID: %v", err)
	} else {
		fmt.Printf("Class with ID %s deleted successfully.\n", classIDToDelete)
	}

}
