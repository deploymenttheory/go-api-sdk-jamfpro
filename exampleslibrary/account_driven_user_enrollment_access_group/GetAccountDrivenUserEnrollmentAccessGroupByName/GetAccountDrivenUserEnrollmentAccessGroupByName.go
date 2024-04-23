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

	// Define a ADUE Account Group ID for testing
	ADUEAccountGroupName := "All Directory Service Users"

	// Call GetScriptsByID function
	script, err := client.GetAccountDrivenUserEnrollmentAccessGroupByName(ADUEAccountGroupName)
	if err != nil {
		log.Fatalf("Error fetching script by ID: %v", err)
	}

	// Pretty print the script details in XML
	JSON, err := json.MarshalIndent(script, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling script details data: %v", err)
	}
	fmt.Println("Fetched Script Details:\n", string(JSON))
}
