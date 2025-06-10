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

	iBeaconName := "Room 1 Beacon" // Replace with the actual iBeacon Name

	err = client.DeleteIBeaconByName(iBeaconName)
	if err != nil {
		log.Fatalf("Error deleting iBeacon by ID: %v", err)
	}

	fmt.Println("iBeacon successfully deleted by ID")
}
