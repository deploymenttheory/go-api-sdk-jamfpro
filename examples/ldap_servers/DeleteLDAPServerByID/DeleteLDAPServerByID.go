package main

import (
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

	// Delete LDAP server by ID
	id := "1" // Replace with actual LDAP server ID
	err = client.DeleteLDAPServerByID(id)
	if err != nil {
		log.Fatalf("Error deleting LDAP server by ID: %v", err)
	}

	fmt.Println("LDAP Server deleted successfully")
}
