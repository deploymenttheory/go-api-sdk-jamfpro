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

	// Call the GetLDAPServers function
	ldapServers, err := client.GetLDAPServers()
	if err != nil {
		log.Fatalf("Error retrieving LDAP servers: %v", err)
	}

	// Process and print the response
	ldapServersXML, err := xml.MarshalIndent(ldapServers, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling LDAP servers data: %v", err)
	}
	fmt.Println("Fetched LDAP Servers List:", string(ldapServersXML))
}
