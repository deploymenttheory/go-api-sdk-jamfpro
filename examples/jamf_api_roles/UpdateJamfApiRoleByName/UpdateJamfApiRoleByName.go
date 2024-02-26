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

	// Define the name of the Jamf API Role you want to update
	roleName := "One Role to Rule them all"

	// Updated role data
	updatedRole := &jamfpro.ResourceAPIRole{
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
