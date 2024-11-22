package main

import (
	"encoding/json"
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

	// Set sorting filter (optional)
	sortFilter := "name:asc"

	// Call function
	groups, err := client.GetSmartComputerGroupsV2(sortFilter)
	if err != nil {
		log.Fatalf("Error fetching smart computer groups v2: %v", err)
	}

	// Pretty print the JSON
	response, err := json.MarshalIndent(groups, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling groups data: %v", err)
	}
	fmt.Println("Fetched Smart Computer Groups V2:\n", string(response))
}
