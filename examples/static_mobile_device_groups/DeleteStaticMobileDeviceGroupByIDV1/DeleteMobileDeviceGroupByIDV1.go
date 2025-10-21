package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/Shared/GitHub/go-api-sdk-jamfpro/localtesting/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Define group ID to delete
	groupID := "103"

	// Call function
	err = client.DeleteStaticMobileDeviceGroupByIDV1(groupID)
	if err != nil {
		log.Fatalf("Error deleting static mobile device group: %v", err)
	}

	fmt.Printf("Successfully deleted static mobile device group with ID %s\n", groupID)
}
