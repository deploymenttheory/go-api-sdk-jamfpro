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

	// Update directory binding by ID
	bindingID := "1" // Assuming an existing binding ID
	updatedBindingByID, err := client.UpdateDirectoryBindingByID(bindingID, updateBinding)
	if err != nil {
		fmt.Println("Error updating directory binding by ID:", err)
		return
	}
	updatedBindingByIDXML, _ := xml.MarshalIndent(updatedBindingByID, "", "    ")
	fmt.Printf("Updated Directory Binding by ID:\n%s\n", string(updatedBindingByIDXML))
}
