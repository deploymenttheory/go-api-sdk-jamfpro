package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logLevel := http_client.LogLevelWarning // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := http_client.Config{
		InstanceName: authConfig.InstanceName,
		Auth: http_client.AuthConfig{
			ClientID:     authConfig.ClientID,
			ClientSecret: authConfig.ClientSecret,
		},
		LogLevel: logLevel,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
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
	bindingID := 1 // Assuming an existing binding ID
	updatedBindingByID, err := client.UpdateDirectoryBindingByID(bindingID, updateBinding)
	if err != nil {
		fmt.Println("Error updating directory binding by ID:", err)
		return
	}
	updatedBindingByIDXML, _ := xml.MarshalIndent(updatedBindingByID, "", "    ")
	fmt.Printf("Updated Directory Binding by ID:\n%s\n", string(updatedBindingByIDXML))
}
