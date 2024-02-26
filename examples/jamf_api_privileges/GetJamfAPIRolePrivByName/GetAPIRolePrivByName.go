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

	// Fetch API role privileges by name
	name := "Read API Roles" // Replace with the privilege name you want to search for
	limit := 15              // Replace with your desired limit for results
	apiPrivileges, err := client.GetJamfAPIPrivilegesByName(name, limit)
	if err != nil {
		log.Fatalf("Error fetching API role privileges by name: %v", err)
	}

	// Pretty print the fetched API role privileges using JSON marshaling
	privilegesJSON, err := json.MarshalIndent(apiPrivileges, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling API role privileges data: %v", err)
	}
	fmt.Println("Fetched API Role Privileges by Name:", string(privilegesJSON))
}
