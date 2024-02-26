package main

import (
	"encoding/xml"
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

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}
	// Call GetUserGroups to fetch user groups
	userGroups, err := client.GetUserGroups()
	if err != nil {
		fmt.Println("Error fetching user groups:", err)
		return
	}

	// Pretty print the user groups details in XML
	userGroupsXML, err := xml.MarshalIndent(userGroups, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling user groups data: %v", err)
	}
	fmt.Println("User Groups Details:\n", string(userGroupsXML))
}
