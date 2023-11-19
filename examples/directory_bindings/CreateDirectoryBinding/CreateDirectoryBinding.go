package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client" // Import http_client for logging
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName:       authConfig.InstanceName,
		OverrideBaseDomain: authConfig.OverrideBaseDomain,
		LogLevel:           logLevel,
		Logger:             logger,
		ClientID:           authConfig.ClientID,
		ClientSecret:       authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// New directory binding data
	newBinding := &jamfpro.ResponseDirectoryBinding{
		Name:       "New Binding",
		Priority:   1,
		Domain:     "example.com",
		Username:   "user@example.com",
		Password:   "password",
		ComputerOU: "CN=Computers,DC=example,DC=com",
		Type:       "Active Directory",
	}

	// Create new directory binding
	createdBinding, err := client.CreateDirectoryBinding(newBinding)
	if err != nil {
		fmt.Println("Error creating directory binding:", err)
		return
	}

	// Pretty print the created directory binding in xml
	createdBindingXML, err := xml.MarshalIndent(createdBinding, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created binding data: %v", err)
	}
	fmt.Printf("Created Directory Binding:\n%s\n", string(createdBindingXML))
}
