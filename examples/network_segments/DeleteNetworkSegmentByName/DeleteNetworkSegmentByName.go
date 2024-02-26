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

	segmentName := "NY Office" // Replace with the actual name of the network segment to delete

	err = client.DeleteNetworkSegmentByName(segmentName)
	if err != nil {
		log.Fatalf("Error deleting network segment by name: %v", err)
	} else {
		fmt.Printf("Network segment '%s' successfully deleted.\n", segmentName)
	}
}
