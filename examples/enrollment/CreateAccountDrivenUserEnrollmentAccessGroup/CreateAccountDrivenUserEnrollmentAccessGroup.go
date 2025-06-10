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

	// Define the new ADUE access group details
	newGroup := jamfpro.ResourceAccountDrivenUserEnrollmentAccessGroup{
		GroupID:                            "cccdad37-ec0e-4738-956c-3f8c0203dace", // Example Group ID
		LdapServerID:                       "1006",                                 // Example LDAP Server ID
		Name:                               "some-group-name",
		SiteID:                             "-1", // Example Site ID (global)
		EnterpriseEnrollmentEnabled:        false,
		PersonalEnrollmentEnabled:          false,
		AccountDrivenUserEnrollmentEnabled: false,
		RequireEula:                        false,
	}

	payload, err := json.MarshalIndent(newGroup, "", "    ")
	if err != nil {
		log.Fatalf("Failed to marshal request payload: %v", err)
	}
	fmt.Println("Request Payload to be sent to Jamf Pro API:\n", string(payload))

	response, err := client.CreateAccountDrivenUserEnrollmentAccessGroup(&newGroup)
	if err != nil {
		log.Fatalf("Failed to create ADUE access group: %v", err)
	}

	JSON, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created group data: %v", err)
	}
	fmt.Println("Created Group Response:\n", string(JSON))
}
