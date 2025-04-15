package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

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

	// Fetch API roles
	apiRoles, err := client.GetJamfAPIRoles(url.Values{})
	if err != nil {
		log.Fatalf("Error fetching API roles: %v", err)
	}

	// Pretty print the fetched API roles using JSON marshaling
	rolesJSON, err := json.MarshalIndent(apiRoles, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling API roles data: %v", err)
	}
	fmt.Println("Fetched API Roles:", string(rolesJSON))
}
