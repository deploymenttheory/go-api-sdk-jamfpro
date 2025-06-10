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

	profileID := "8" // Use the actual ID of the profile to be deleted

	err = client.DeleteBYOProfileByID(profileID)
	if err != nil {
		log.Fatalf("Error deleting BYO Profile by ID: %v", err)
	} else {
		fmt.Println("BYO Profile deleted successfully by ID")
	}
}
