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

	roleName := "API Access"

	// Call GetJamfApiRolesNameById function
	role, err := client.GetJamfApiRoleByName(roleName)
	if err != nil {
		log.Fatalf("Error fetching Jamf API role by name: %v", err)
	}

	// Pretty print the role in JSON
	roleJSON, err := json.MarshalIndent(role, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Jamf API role data: %v", err)
	}
	fmt.Println("Fetched Jamf API role:\n", string(roleJSON))
}
