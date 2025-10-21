package main

import (
	"encoding/json"
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

	// Define a Static Mobile Device Group name for testing
	groupName := "Excluded Devices" // Replace with actual group name

	// Call GetStaticMobileDeviceGroupByName function
	group, err := client.GetStaticMobileDeviceGroupByNameV1(groupName)
	if err != nil {
		log.Fatalf("Error fetching Static Mobile Device Group by name: %v", err)
	}

	// Pretty print the group in JSON
	groupJSON, err := json.MarshalIndent(group, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Static Mobile Device Group data: %v", err)
	}
	fmt.Println("Fetched Static Mobile Device Group:\n", string(groupJSON))
}
