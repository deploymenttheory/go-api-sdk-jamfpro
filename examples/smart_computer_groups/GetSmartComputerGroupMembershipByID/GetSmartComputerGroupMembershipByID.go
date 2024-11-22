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

	// Define the group ID
	groupID := "1"

	// Call function
	membership, err := client.GetSmartComputerGroupMembershipByID(groupID)
	if err != nil {
		log.Fatalf("Error fetching smart computer group membership: %v", err)
	}

	// Pretty print the JSON
	response, err := json.MarshalIndent(membership, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling membership data: %v", err)
	}
	fmt.Println("Fetched Smart Computer Group Membership:\n", string(response))
}
