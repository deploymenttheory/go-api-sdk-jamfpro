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

	// Define a Smart Mobile Device Group ID for testing
	groupID := "5" // Replace with actual group ID

	// Call GetSmartMobileDeviceGroupByID function
	group, err := client.GetSmartMobileDeviceGroupByIDV1(groupID)
	if err != nil {
		log.Fatalf("Error fetching Smart Mobile Device Group by ID: %v", err)
	}

	// Pretty print the group in JSON
	groupJSON, err := json.MarshalIndent(group, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Smart Mobile Device Group data: %v", err)
	}
	fmt.Println("Fetched Smart Mobile Device Group:\n", string(groupJSON))
}
