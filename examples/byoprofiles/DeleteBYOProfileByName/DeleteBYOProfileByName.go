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

	profileName := "Personal Device Profile" // Use the actual name of the profile to be deleted

	err = client.DeleteBYOProfileByName(profileName)
	if err != nil {
		log.Fatalf("Error deleting BYO Profile by name: %v", err)
	} else {
		fmt.Println("BYO Profile deleted successfully by name")
	}
}
