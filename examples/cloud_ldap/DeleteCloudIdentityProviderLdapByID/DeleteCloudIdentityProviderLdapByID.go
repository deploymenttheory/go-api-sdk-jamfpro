package main

import (
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

	// Define the ID of the Cloud LDAP configuration you want to delete
	ldapID := "1022" // Replace with your actual LDAP configuration ID

	// Delete the Cloud LDAP configuration
	err = client.DeleteCloudIdentityProviderLdapByID(ldapID)
	if err != nil {
		log.Fatalf("Error deleting cloud LDAP configuration: %v", err)
	}

	fmt.Printf("Successfully deleted Cloud LDAP configuration with ID: %s\n", ldapID)
}
