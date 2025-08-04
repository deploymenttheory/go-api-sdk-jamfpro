package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/Shared/GitHub/go-api-sdk-jamfpro/localtesting/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Define updated attributes
	updatedAttribute := &jamfpro.ResourceMobileDeviceExtensionAttribute{
		Name:                 "New Attribute",
		Description:          "This is a test attribute",
		DataType:             "String",
		InventoryDisplayType: "Hardware",
		InputType:            "TEXT",
	}

	updatedAttributeID := "2"

	// Update the attribute
	attribute, err := client.UpdateMobileDeviceExtensionAttributeByID(updatedAttributeID, updatedAttribute)
	if err != nil {
		log.Fatalf("Error updating mobile device extension attribute by ID: %v", err)
	}

	fmt.Printf("Updated Attribute: %+v\n", attribute)
}
