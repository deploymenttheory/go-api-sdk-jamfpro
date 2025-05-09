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

	// Define the ID of the Cloud LDAP configuration you want to retrieve
	cloudLdapID := "1019" // Replace with your actual Cloud LDAP ID

	// Get the Cloud LDAP configuration by ID
	cloudLdap, err := client.GetCloudIdentityProviderLdapByID(cloudLdapID)
	if err != nil {
		log.Fatalf("Error retrieving cloud LDAP configuration: %v", err)
	}

	// Pretty print the response in JSON
	responseJSON, err := json.MarshalIndent(cloudLdap, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling response data: %v", err)
	}
	fmt.Printf("Cloud LDAP Configuration (ID: %s):\n%s\n", cloudLdapID, string(responseJSON))
}
