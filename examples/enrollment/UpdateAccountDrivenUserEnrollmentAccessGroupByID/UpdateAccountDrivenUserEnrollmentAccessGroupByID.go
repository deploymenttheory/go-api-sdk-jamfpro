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
	groupID := "12"

	groupUpdate := &jamfpro.ResourceAccountDrivenUserEnrollmentAccessGroup{
		ID:                                 groupID,
		GroupID:                            "cccdad37-ec0e-4738-956c-3f8c0203dace", // Example Group ID
		LdapServerID:                       "1006",                                 // Example LDAP Server ID
		Name:                               "some-group-name",
		SiteID:                             "-1", // Example Site ID (global)
		EnterpriseEnrollmentEnabled:        false,
		PersonalEnrollmentEnabled:          false,
		AccountDrivenUserEnrollmentEnabled: false,
		RequireEula:                        false,
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
