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

	updatedBeacon := &jamfpro.ResourceIBeacons{
		Name:  "Room 1 Beacon Updated",
		UUID:  "55900BDC-347C-58B1-D249-F32244B11D30",
		Major: -1,
		Minor: -1,
	}

	result, err := client.UpdateIBeaconByName("Room 1 Beacon", updatedBeacon)
	if err != nil {
		log.Fatalf("Error updating iBeacon: %v", err)
	}

	fmt.Printf("Updated iBeacon: %+v\n", result)
}
