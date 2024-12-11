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

	// Define the language IDs to delete
	languageIds := []string{"en", "fr", "es"}

	// Call DeleteMultipleEnrollmentMessagesByLanguageIDs function
	err = client.DeleteMultipleEnrollmentMessagesByLanguageIDs(languageIds)
	if err != nil {
		log.Fatalf("Error deleting enrollment language messages: %v", err)
	}

	fmt.Printf("Successfully deleted enrollment messages for languages: %v\n", languageIds)
}
