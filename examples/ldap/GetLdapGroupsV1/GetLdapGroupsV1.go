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

	// Optional: provide text to perform a contains search on LDAP group names
	searchTerm := "Admin"

	// Call GetLdapGroupsV1 to retrieve Jamf Pro LDAP groups
	groups, err := client.GetLdapGroupsV1(searchTerm)
	if err != nil {
		log.Fatalf("Error fetching LDAP groups: %v", err)
	}

	// Pretty print the group list
	groupsJSON, err := json.MarshalIndent(groups, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling LDAP groups data: %v", err)
	}

	fmt.Println("Fetched LDAP Groups:\n", string(groupsJSON))
}
