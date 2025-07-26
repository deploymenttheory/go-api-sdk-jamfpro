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

	// Define a new mobile device extension attribute
	newAttribute := &jamfpro.ResourceMobileDeviceExtensionAttribute{
		Name:                 "New Attribute",
		Description:          "This is a test attribute",
		DataType:             "String",
		InventoryDisplayType: "General",
		InputType:            "TEXT",
	}

	// Create the new attribute
	createdAttribute, err := client.CreateMobileDeviceExtensionAttribute(newAttribute)
	if err != nil {
		log.Fatalf("Error creating mobile device extension attribute: %v", err)
	}

	fmt.Printf("Created Attribute: %+v\n", createdAttribute)
}
