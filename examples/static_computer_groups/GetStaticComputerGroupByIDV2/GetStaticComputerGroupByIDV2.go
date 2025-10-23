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

	// Define a Static Computer Group ID for testing
	groupID := "8" // Replace with actual group ID

	// Call GetStaticComputerGroupByID function
	group, err := client.GetStaticComputerGroupByIDV2(groupID)
	if err != nil {
		log.Fatalf("Error fetching Static Computer Group by ID: %v", err)
	}

	// Pretty print the group in JSON
	groupJSON, err := json.MarshalIndent(group, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Static Computer Group data: %v", err)
	}
	fmt.Println("Fetched Static Computer Group:\n", string(groupJSON))
}
