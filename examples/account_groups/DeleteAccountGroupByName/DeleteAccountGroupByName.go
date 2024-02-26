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

	// Let's assume you want to delete an account group with the name "Administrators modified group".
	accountGroupName := "Administrators modified"

	// Call DeleteAccountGroupByName function
	err = client.DeleteAccountGroupByName(accountGroupName)
	if err != nil {
		log.Fatalf("Error deleting account group by name: %v", err)
	}

	// Print the success message
	fmt.Println("Deleted Account Group Successfully.")
}
