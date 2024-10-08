package main

import (
	"encoding/json"
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

	attributeToUpdate := &jamfpro.ResourceComputerExtensionAttribute{
		Name:                          "Battery Cycle Count - Updated",
		Description:                   "Number of charge cycles logged on the current battery",
		DataType:                      "STRING",
		Enabled:                       jamfpro.BoolPtr(true),
		InventoryDisplayType:          "GENERAL",
		InputType:                     "TEXT_FIELD",
		LDAPAttributeMapping:          "",
		LDAPExtensionAttributeAllowed: jamfpro.BoolPtr(false),
	}

	// Assuming you're updating the attribute with ID = 1
	updatedAttribute, err := client.UpdateComputerExtensionAttributeByID("1", attributeToUpdate)
	if err != nil {
		log.Fatalf("Error updating Computer Extension Attribute by ID: %v", err)
	}

	attributeJSON, err := json.MarshalIndent(updatedAttribute, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling updated Computer Extension Attribute data: %v", err)
	}
	fmt.Println("Updated Computer Extension Attribute by ID:\n", string(attributeJSON))
}
