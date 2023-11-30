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
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug // Adjust log level as needed

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

	// Prepare the attribute to update
	attributeToUpdate := &jamfpro.ResponseUserExtensionAttribute{
		Name:        "User Attributes",                    // Updated name
		Description: "Text field for logging custom data", // Updated description
		DataType:    "String",                             // Updated data type
		InputType: jamfpro.UserExtensionInputType{
			Type: "Text Field", // Updated input type
		},
	}

	// Example ID of the user extension attribute to rupdate
	attributeName := "User Attributes"

	// Update the attribute
	updatedAttribute, err := client.UpdateUserExtensionAttributeByName(attributeName, attributeToUpdate) // Use the correct name
	if err != nil {
		log.Fatalf("Error updating user extension attribute: %v", err)
	}

	// Print the updated attribute
	fmt.Printf("Updated Attribute: %+v\n", updatedAttribute)
}
