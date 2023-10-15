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

	// Define the name of the Jamf API Role you want to update
	roleName := "One Role to Rule them all"

	// Updated role data
	updatedRole := &jamfpro.Role{
		DisplayName: "Updated Role Name",
		Privileges:  []string{"Update eBooks", "Update User"},
	}

	// Call UpdateJamfApiRoleByName function
	role, err := client.UpdateJamfApiRoleByName(roleName, updatedRole)
	if err != nil {
		log.Fatalf("Error updating Jamf API role by name: %v", err)
	}

	// Pretty print the updated role in XML
	roleXML, err := xml.MarshalIndent(role, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling updated Jamf API role data: %v", err)
	}
	fmt.Println("Updated Jamf API Role:\n", string(roleXML))
}
