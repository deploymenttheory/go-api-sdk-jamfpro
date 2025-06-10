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

	// Let's assume you want to delete an account with ID 31.
	accountID := "388"

	// Call DeleteAccountByID function
	err = client.DeleteAccountByID(accountID)
	if err != nil {
		log.Fatalf("Error deleting account by ID: %v", err)
	}

	// Print the success message
	fmt.Println("Deleted Account Successfully.")
}
