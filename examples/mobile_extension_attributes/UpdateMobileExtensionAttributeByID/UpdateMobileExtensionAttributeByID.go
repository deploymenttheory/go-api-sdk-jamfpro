package main

import (
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

	updatedAttributeID := 1

	// Update the attribute
	attribute, err := client.UpdateMobileExtensionAttributeByID(updatedAttributeID, updatedAttribute)
	if err != nil {
		log.Fatalf("Error updating mobile extension attribute by ID: %v", err)
	}

	fmt.Printf("Updated Attribute: %+v\n", attribute)
}
