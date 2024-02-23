package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"
	// Load the client OAuth credentials from the configuration file
	loadedConfig, err := jamfpro.LoadClientConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for the HTTP client
	config := httpclient.ClientConfig{
		Auth: httpclient.AuthConfig{
			ClientID:     loadedConfig.Auth.ClientID,
			ClientSecret: loadedConfig.Auth.ClientSecret,
		},
		Environment: httpclient.EnvironmentConfig{
			APIType:      loadedConfig.Environment.APIType,
			InstanceName: loadedConfig.Environment.InstanceName,
		},
		ClientOptions: httpclient.ClientOptions{
			LogLevel:            loadedConfig.ClientOptions.LogLevel,
			LogOutputFormat:     loadedConfig.ClientOptions.LogOutputFormat,
			LogConsoleSeparator: loadedConfig.ClientOptions.LogConsoleSeparator,
			HideSensitiveData:   loadedConfig.ClientOptions.HideSensitiveData,
		},
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Prepare the attribute to update
	attributeToUpdate := &jamfpro.ResourceUserExtensionAttribute{
		Name:        "User Attributes",                    // Updated name
		Description: "Text field for logging custom data", // Updated description
		DataType:    "String",                             // Updated data type
		InputType: jamfpro.ResourceUserExtensionAttributeSubsetInputType{
			Type: "Text Field", // Updated input type
		},
	}

	// Example ID of the user extension attribute to rupdate
	attributeID := 1

	// Update the attribute
	updatedAttribute, err := client.UpdateUserExtensionAttributeByID(attributeID, attributeToUpdate) // Use the correct ID
	if err != nil {
		log.Fatalf("Error updating user extension attribute: %v", err)
	}

	// Print the updated attribute
	fmt.Printf("Updated Attribute: %+v\n", updatedAttribute)
}
