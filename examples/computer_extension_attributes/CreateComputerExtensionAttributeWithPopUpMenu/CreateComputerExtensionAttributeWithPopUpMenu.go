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

	// Define the choices for the popup menu
	choices := []string{"Choice 1", "Choice 2", "Choice 3"}

	// Define the new computer extension attribute
	attribute := &jamfpro.ResourceComputerExtensionAttribute{
		Name:                          "Pop Up Menu Test",
		Description:                   "Pop Up Menu Test",
		DataType:                      "STRING",
		Enabled:                       true,
		InventoryDisplayType:          "GENERAL",
		InputType:                     "POP_UP_MENU",
		PopupMenuChoices:              choices,
		LDAPAttributeMapping:          "",
		LDAPExtensionAttributeAllowed: false,
	}

	// Call CreateComputerExtensionAttribute function
	createdAttribute, err := client.CreateComputerExtensionAttribute(attribute)
	if err != nil {
		log.Fatalf("Error creating Computer Extension Attribute: %v", err)
	}

	// Pretty print the created attribute in JSON
	createdAttributeJSON, err := json.MarshalIndent(createdAttribute, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling created Computer Extension Attribute data: %v", err)
	}
	fmt.Println("Created Computer Extension Attribute:\n", string(createdAttributeJSON))
}
