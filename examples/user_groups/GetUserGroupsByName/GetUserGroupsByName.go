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
	// Replace with the actual name of the user group you want to fetch
	userGroupName := "Teachers"

	// Call GetUserGroupsByName to fetch details of a specific user group
	userGroupDetail, err := client.GetUserGroupByName(userGroupName)
	if err != nil {
		fmt.Println("Error fetching user group details:", err)
		return
	}

	// Pretty print the user group details in XML
	userGroupDetailXML, err := xml.MarshalIndent(userGroupDetail, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling user group data: %v", err)
	}
	fmt.Println("User Group Details:\n", string(userGroupDetailXML))
}
