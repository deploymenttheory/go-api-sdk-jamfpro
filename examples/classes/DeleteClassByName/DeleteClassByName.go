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

	// The name of the class you want to delete
	classNameToDelete := "Math 101" // Replace with the actual class name

	// Call the delete function with the class name
	err = client.DeleteClassByName(classNameToDelete)
	if err != nil {
		log.Fatalf("Error deleting class by name: %v", err)
	} else {
		fmt.Printf("Class with name %s deleted successfully.\n", classNameToDelete)
	}
}
