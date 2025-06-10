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

	// Define the name of the macOS Configuration Profile you want to delete
	profileName := "WiFi Test"

	// Call the DeleteMacOSConfigurationProfileByName function
	err = client.DeleteMacOSConfigurationProfileByName(profileName)
	if err != nil {
		log.Fatalf("Failed to delete macOS Configuration Profile with name '%s': %v", profileName, err)
	}

	fmt.Printf("Profile with name '%s' deleted successfully\n", profileName)
}
