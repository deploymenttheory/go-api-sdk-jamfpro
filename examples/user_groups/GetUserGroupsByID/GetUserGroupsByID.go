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

	// Replace with the actual user group ID you want to fetch
	userGroupID := "1"

	// Call GetUserGroupsByID to fetch details of a specific user group
	userGroupDetail, err := client.GetUserGroupByID(userGroupID)
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
