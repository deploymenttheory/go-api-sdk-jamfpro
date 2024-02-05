package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

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

	// Define updated attributes
	updatedAttribute := &jamfpro.ResourceMobileExtensionAttribute{
		Name:             "New Attribute",
		Description:      "This is a test attribute",
		DataType:         "String",
		InventoryDisplay: "General",
		InputType: jamfpro.MobileExtensionAttributeSubsetInputType{
			Type: "Text Field",
		},
	}

	updatedAttributeName := "New Attribute"
	// Update the attribute
	attribute, err := client.UpdateMobileExtensionAttributeByName(updatedAttributeName, updatedAttribute)
	if err != nil {
		log.Fatalf("Error updating mobile extension attribute by name: %v", err)
	}

	fmt.Printf("Updated Attribute: %+v\n", attribute)
}
