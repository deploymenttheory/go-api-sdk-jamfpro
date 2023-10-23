package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		DebugMode:    true,
		Logger:       jamfpro.NewDefaultLogger(),
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	attributeToUpdate := &jamfpro.ComputerExtensionAttributeResponse{
		Name:             "Battery Cycle Count",
		Description:      "Number of charge cycles logged on the current battery",
		DataType:         "String",
		InputType:        jamfpro.ComputerExtensionAttributeInputType{Type: "Text Field"},
		InventoryDisplay: "General",
		ReconDisplay:     "Extension Attributes",
	}

	// Assuming you're updating the attribute with ID = 1
	updatedAttribute, err := client.UpdateComputerExtensionAttributeByID(1, attributeToUpdate)
	if err != nil {
		log.Fatalf("Error updating Computer Extension Attribute by ID: %v", err)
	}

	attributeXML, err := xml.MarshalIndent(updatedAttribute, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling updated Computer Extension Attribute data: %v", err)
	}
	fmt.Println("Updated Computer Extension Attribute by ID:\n", string(attributeXML))
}
