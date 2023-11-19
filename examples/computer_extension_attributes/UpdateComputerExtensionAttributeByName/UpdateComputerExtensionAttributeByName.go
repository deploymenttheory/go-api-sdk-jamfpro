package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName:       authConfig.InstanceName,
		OverrideBaseDomain: authConfig.OverrideBaseDomain,
		LogLevel:           logLevel,
		Logger:             logger,
		ClientID:           authConfig.ClientID,
		ClientSecret:       authConfig.ClientSecret,
	}

	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	attributeToUpdate := &jamfpro.ResponseComputerExtensionAttribute{
		Name:             "Battery Cycle Count Updated", // Notice the "Updated" suffix for demonstration
		Description:      "Number of charge cycles logged on the current battery",
		DataType:         "String",
		InputType:        jamfpro.ComputerExtensionAttributeInputType{Type: "Text Field"},
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
