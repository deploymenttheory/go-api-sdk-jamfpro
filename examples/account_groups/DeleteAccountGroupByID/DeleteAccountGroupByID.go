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

	// Let's assume you want to delete an account group with ID 42.
	accountGroupID := 1

	// Call DeleteAccountGroupByID function
	err = client.DeleteAccountGroupByID(accountGroupID)
	if err != nil {
		log.Fatalf("Error deleting account group by ID: %v", err)
	}

	// Print the success message
	fmt.Println("Deleted Account Group Successfully.")
}
