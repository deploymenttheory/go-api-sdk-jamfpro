package main

import (
	"encoding/xml"
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

	// Call the GetLDAPServerByNameAndGroupDataSubset function
	ldapServerName := "exampleLDAPServerName"
	group := "exampleGroup"
	ldapServer, err := client.GetLDAPServerByNameAndGroupDataSubset(ldapServerName, group)
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
