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

	// Define the ID of the group to update
	groupID := "1"

	groupUpdate := &jamfpro.ResourceAccountDrivenUserEnrollmentAccessGroup{
		ID:                                 groupID,
		GroupID:                            "12345",
		LdapServerID:                       "1",
		Name:                               "Updated ADUE Access Group",
		SiteID:                             "-1",
		EnterpriseEnrollmentEnabled:        true,
		PersonalEnrollmentEnabled:          false,
		AccountDrivenUserEnrollmentEnabled: true,
		RequireEula:                        true,
	}

	updatedGroup, err := client.UpdateAccountDrivenUserEnrollmentAccessGroupByID(groupID, groupUpdate)
	if err != nil {
		log.Fatalf("Error updating ADUE access group: %v", err)
	}

	// Pretty print the updated group in JSON
	JSON, err := json.MarshalIndent(updatedGroup, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling updated group data: %v", err)
	}
	fmt.Printf("Updated ADUE Access Group:\n%s\n", string(JSON))
}
