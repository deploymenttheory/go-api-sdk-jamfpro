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

	// Define the new computer extension attribute
	attribute := &jamfpro.ResourceComputerExtensionAttribute{
		Name:             "Battery Cycle Count",
		Description:      "Number of charge cycles logged on the current battery",
		DataType:         "String",                                                              // String / Integer / Date (YYYY-MM-DD hh:mm:ss)
		InputType:        jamfpro.ComputerExtensionAttributeSubsetInputType{Type: "Text Field"}, //  Text Field / Pop Up Menu / Script
		InventoryDisplay: "General",                                                             // General / Hardware / Operating System / User and Location / Purchasing / Extension Attribute
		ReconDisplay:     "Extension Attributes",
	}

	// Call CreateComputerExtensionAttribute function
	createdAttribute, err := client.CreateComputerExtensionAttribute(attribute)
	if err != nil {
		log.Fatalf("Error creating Computer Extension Attribute: %v", err)
	}

	// Pretty print the created attribute in XML
	createdAttributeXML, err := xml.MarshalIndent(createdAttribute, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling created Computer Extension Attribute data: %v", err)
	}
	fmt.Println("Created Computer Extension Attribute:\n", string(createdAttributeXML))
}
