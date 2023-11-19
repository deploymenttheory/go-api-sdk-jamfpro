package main

import (
	"encoding/xml"
	"fmt"
	"log"

	// Import http_client for logging
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
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

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName:       authConfig.InstanceName,
		OverrideBaseDomain: authConfig.OverrideBaseDomain,
		LogLevel:           logLevel,
		Logger:             logger,
		ClientID:           authConfig.ClientID,
		ClientSecret:       authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Fetch LDAP server details by ID and user data subset
	ldapServerID := 1 // Replace with actual ID
	user := "exampleUser"
	ldapServer, err := client.GetLDAPServerByIDAndUserDataSubset(ldapServerID, user)
	if err != nil {
		log.Fatalf("Error retrieving LDAP server: %v", err)
	}

	// Print the response
	ldapServerXML, err := xml.MarshalIndent(ldapServer, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling LDAP server data: %v", err)
	}
	fmt.Println("Fetched LDAP Server by ID:", string(ldapServerXML))
}
