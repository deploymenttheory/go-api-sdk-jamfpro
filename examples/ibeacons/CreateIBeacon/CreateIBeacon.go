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

	// Define the new iBeacon details
	newBeacon := &jamfpro.ResourceIBeacons{
		Name:  "Room 1 Beacon",
		UUID:  "55900BDC-347C-58B1-D249-F32244B11D30",
		Major: -1,
		Minor: -1,
	}

	// Call the CreateIBeacon function
	createdBeacon, err := client.CreateIBeacon(newBeacon)
	if err != nil {
		log.Fatalf("Error creating iBeacon: %v", err)
	}

	fmt.Printf("Created iBeacon: %+v\n", createdBeacon)
}
