package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file inside the main function
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		DebugMode:    true,
		Logger:       jamfpro.NewDefaultLogger(),
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new jamfpro client instanceclient,
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the new Jamf API Role you want to create
	newRole := &jamfpro.APIRole{
		DisplayName: "One Role to Rule them all",
		Privileges:  []string{"View License Serial Numbers"},
	}

	// Call CreateJamfApiRole function
	createdRole, err := client.CreateJamfApiRole(newRole)
	if err != nil {
		log.Fatalf("Error creating Jamf API role: %v", err)
	}

	// Pretty print the created role in XML
	roleXML, err := xml.MarshalIndent(createdRole, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Jamf API role data: %v", err)
	}
	fmt.Println("Created Jamf API Role:\n", string(roleXML))
}
