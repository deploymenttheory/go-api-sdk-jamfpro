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

	// Define the new ADUE access group details
	newGroup := jamfpro.ResourceAccountDrivenUserEnrollmentAccessGroup{
		GroupID:                            "123", // Example Group ID, replace with actual
		LdapServerID:                       "456", // Example LDAP Server ID, replace with actual
		Name:                               "New Access Group",
		SiteID:                             "789", // Example Site ID, replace with actual
		EnterpriseEnrollmentEnabled:        true,
		PersonalEnrollmentEnabled:          false,
		AccountDrivenUserEnrollmentEnabled: true,
		RequireEula:                        false,
	}

	// Create the new ADUE access group using the client
	response, err := client.CreateAccountDrivenUserEnrollmentAccessGroup(&newGroup)
	if err != nil {
		log.Fatalf("Failed to create ADUE access group: %v", err)
	}

	// Pretty print the created script details in XML
	JSON, err := json.MarshalIndent(response, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling created script data: %v", err)
	}
	fmt.Println("Created Script Details:\n", string(JSON))
}
