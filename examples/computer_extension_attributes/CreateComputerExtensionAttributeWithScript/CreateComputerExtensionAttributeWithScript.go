package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func loadScriptFromFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Load the script from a file
	scriptPath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/examples/support_files/computer_extensioin_attribute.sh"
	scriptContent, err := loadScriptFromFile(scriptPath)
	if err != nil {
		log.Fatalf("Failed to load script from file: %v", err)
	}

	// Define the new computer extension attribute
	attribute := &jamfpro.ResourceComputerExtensionAttribute{
		Name:                          "Computer Extension Attribute Script Test",
		Description:                   "Computer Extension Attribute Script Test",
		DataType:                      "STRING",
		Enabled:                       true,
		InventoryDisplayType:          "GENERAL",
		InputType:                     "SCRIPT",
		ScriptContents:                scriptContent,
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
