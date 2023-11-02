// This uses a seperate function to the other delete operations, so it is intentionally
// included as a seperate example.

package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for the Jamf Pro client
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		DebugMode:    true,
		Logger:       jamfpro.NewDefaultLogger(),
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new Jamf Pro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the category IDs to delete as strings
	categoryIDsToDelete := []string{"3", "4"} // Replace with the actual category IDs you want to delete

	// Call DeleteMultipleCategoriesByID function
	err = client.DeleteMultipleCategoriesByID(categoryIDsToDelete)
	if err != nil {
		log.Fatalf("Error deleting multiple categories: %v", err)
	}

	fmt.Println("Categories deleted successfully")
}
