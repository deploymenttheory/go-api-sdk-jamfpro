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

	// Example ID of the directory binding to delete
	bindingID := 1 // Assuming an existing binding ID

	// Delete directory binding by ID
	err = client.DeleteDirectoryBindingByID(bindingID)
	if err != nil {
		fmt.Println("Error deleting directory binding by ID:", err)
		return
	}
	fmt.Println("Successfully deleted Directory Binding by ID")
}
