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

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}
	// New user extension attribute details
	newAttribute := &jamfpro.ResourceUserExtensionAttribute{
		Name:        "User Attributes",
		Description: "Text field for logging custom data",
		DataType:    "String",
		InputType: jamfpro.ResourceUserExtensionAttributeSubsetInputType{
			Type: "Text Field",
		},
	}

	// Create the new user extension attribute
	createdAttribute, err := client.CreateUserExtensionAttribute(newAttribute)
	if err != nil {
		log.Fatalf("Error creating user extension attribute: %v", err)
	}

	// Print the created attribute details in XML
	attributeXML, err := xml.MarshalIndent(createdAttribute, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling created attribute data: %v", err)
	}
	fmt.Println("Created User Extension Attribute Details:\n", string(attributeXML))
}
