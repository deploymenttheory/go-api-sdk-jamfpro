package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client" // Import http_client for logging
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file for OAuth credentials
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for the jamfpro client
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		LogLevel:     http_client.LogLevelDebug,
		Logger:       http_client.NewDefaultLogger(),
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Fetch LDAP server group details by ID
	ldapServerID := 1 // Replace with actual ID
	group := "exampleGroup"
	ldapServer, err := client.GetLDAPServerByIDAndGroupDataSubset(ldapServerID, group)
	if err != nil {
		log.Fatalf("Error retrieving LDAP server group data: %v", err)
	}

	// Print the response
	ldapServerXML, err := xml.MarshalIndent(ldapServer, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling LDAP server data: %v", err)
	}
	fmt.Println("Fetched LDAP Server by ID:", string(ldapServerXML))
}
