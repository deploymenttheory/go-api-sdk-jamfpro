package main

import (
	"encoding/json"
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

	// Define a Smart Computer Group name for testing
	groupName := "Operating System Version like 15" // Replace with actual group name

	// Call GetSmartComputerGroupByName function
	group, err := client.GetSmartComputerGroupByName(groupName)
	if err != nil {
		log.Fatalf("Error fetching Smart Computer Group by name: %v", err)
	}

	// Pretty print the group in JSON
	groupJSON, err := json.MarshalIndent(group, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Smart Computer Group data: %v", err)
	}
	fmt.Println("Fetched Smart Computer Group:\n", string(groupJSON))
}
