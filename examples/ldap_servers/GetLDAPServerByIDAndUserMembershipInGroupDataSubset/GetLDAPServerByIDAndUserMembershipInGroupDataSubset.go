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

	// Call the GetLDAPServerByIDAndUserMembershipInGroupDataSubset function
	ldapServerID := 1 // Replace with actual LDAP server ID
	group := "exampleGroup"
	user := "exampleUser1,exampleUser2" // Multiple users separated by a comma
	ldapServer, err := client.GetLDAPServerByIDAndUserMembershipInGroupDataSubset(ldapServerID, group, user)
	if err != nil {
		log.Fatalf("Error retrieving LDAP server user membership data: %v", err)
	}

	// Print the response
	ldapServerXML, err := xml.MarshalIndent(ldapServer, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling LDAP server data: %v", err)
	}
	fmt.Println("Fetched LDAP Server by ID:", string(ldapServerXML))
}
