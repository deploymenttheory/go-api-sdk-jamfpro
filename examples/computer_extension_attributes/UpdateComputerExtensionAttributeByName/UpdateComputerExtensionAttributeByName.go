package main

import (
	"encoding/xml"
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
			LogLevel:          loadedConfig.ClientOptions.LogLevel,
			HideSensitiveData: loadedConfig.ClientOptions.HideSensitiveData,
			LogOutputFormat:   loadedConfig.ClientOptions.LogOutputFormat,
		},
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}
	attributeToUpdate := &jamfpro.ResourceComputerExtensionAttribute{
		Name:             "Battery Cycle Count Updated", // Notice the "Updated" suffix for demonstration
		Description:      "Number of charge cycles logged on the current battery",
		DataType:         "String",
		InputType:        jamfpro.ComputerExtensionAttributeSubsetInputType{Type: "Text Field"},
		InventoryDisplay: "General",
		ReconDisplay:     "Extension Attributes",
	}

	// Updating the attribute with name "Battery Cycle Count" to "Battery Cycle Count Updated"
	updatedAttribute, err := client.UpdateComputerExtensionAttributeByName("Battery Cycle Count", attributeToUpdate)
	if err != nil {
		log.Fatalf("Error updating Computer Extension Attribute by Name: %v", err)
	}

	attributeXML, err := xml.MarshalIndent(updatedAttribute, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling updated Computer Extension Attribute data: %v", err)
	}
	fmt.Println("Updated Computer Extension Attribute by Name:\n", string(attributeXML))
}
