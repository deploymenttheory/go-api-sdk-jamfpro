package main

import (
	"encoding/xml"
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

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}
	// Call the GetLDAPServerByNameAndUserMembershipInGroupDataSubset function
	ldapServerName := "exampleLDAPServerName"
	groupName := "exampleGroupName"
	userName := "exampleUserName"
	ldapServer, err := client.GetLDAPServerByNameAndUserMembershipInGroupDataSubset(ldapServerName, groupName, userName)
	if err != nil {
		log.Fatalf("Error retrieving LDAP server user membership in group data: %v", err)
	}

	// Process and print the response
	ldapServersXML, err := xml.MarshalIndent(ldapServer, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling LDAP servers data: %v", err)
	}
	fmt.Println("Fetched LDAP Servers List:", string(ldapServersXML))
}
