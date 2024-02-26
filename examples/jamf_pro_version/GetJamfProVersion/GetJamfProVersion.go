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

	// Call GetJamfProVersion function
	version, err := client.GetJamfProVersion()
	if err != nil {
		log.Fatalf("Error fetching Jamf Pro version: %v", err)
	}

	// Print the fetched Jamf Pro version
	fmt.Println("Current Jamf Pro Version:", *version.Version)
}
