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

	// Specify the enrollment customization ID and panel ID
	customizationID := "22" // Replace with your actual customization ID
	paneID := "19"          // Replace with your actual pane ID

	// Get the LDAP prestage pane
	ldapPane, err := client.GetLDAPPrestagePaneByID(customizationID, paneID)
	if err != nil {
		log.Fatalf("Failed to get LDAP prestage pane: %v", err)
	}

	// Pretty print the result in JSON
	prettyJSON, err := json.MarshalIndent(ldapPane, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling result: %v", err)
	}
	fmt.Println("LDAP Prestage Pane Details:\n", string(prettyJSON))
}
