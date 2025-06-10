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

	segmentID := "1" // Replace with actual ID of the network segment to delete

	err = client.DeleteNetworkSegmentByID(segmentID)
	if err != nil {
		log.Fatalf("Error deleting network segment by ID: %v", err)
	} else {
		fmt.Printf("Network segment with ID %d successfully deleted.\n", segmentID)
	}
}
