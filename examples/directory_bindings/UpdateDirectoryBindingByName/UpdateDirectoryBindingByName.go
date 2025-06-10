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

	// Directory binding data to update
	updateBinding := &jamfpro.ResponseDirectoryBinding{
		Name:       "Updated Binding",
		Priority:   1,
		Domain:     "updated.example.com",
		Username:   "user@updated.com",
		Password:   "newpassword",
		ComputerOU: "CN=UpdatedComputers,DC=updated,DC=example,DC=com",
		Type:       "Active Directory",
	}

	// Update directory binding by Name
	bindingName := "New Binding" // Assuming an existing binding name
	updatedBindingByName, err := client.UpdateDirectoryBindingByName(bindingName, updateBinding)
	if err != nil {
		fmt.Println("Error updating directory binding by Name:", err)
		return
	}
	updatedBindingByNameXML, _ := xml.MarshalIndent(updatedBindingByName, "", "    ")
	fmt.Printf("Updated Directory Binding by Name:\n%s\n", string(updatedBindingByNameXML))
}
