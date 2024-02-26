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
	// Define the ID of the Jamf API Role you want to update
	roleID := "13" // For example, use the ID "1"

	// Define the new data for the role
	updatedRole := &jamfpro.ResourceAPIRole{
		DisplayName: "Updated Role Display Name",
		Privileges:  []string{"Update eBooks", "Update User"},
	}

	// Call UpdateJamfApiRoleByID function
	role, err := client.UpdateJamfApiRoleByID(roleID, updatedRole)
	if err != nil {
		log.Fatalf("Error updating Jamf API role by ID: %v", err)
	}

	// Pretty print the updated role in XML
	roleXML, err := xml.MarshalIndent(role, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Jamf API role data: %v", err)
	}
	fmt.Println("Updated Jamf API Role:\n", string(roleXML))
}
