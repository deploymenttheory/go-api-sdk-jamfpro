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

	customPath := &jamfpro.ResourceComputerInventoryCollectionSettingsCustomPath{
		Scope: "PLUGIN", // APP / FONT / PLUGIN
		Path:  "/Example/Path/",
	}

	// Create a custom path for computer inventory collection settings
	response, err := client.CreateComputerInventoryCollectionSettingsCustomPath(customPath)
	if err != nil {
		log.Fatalf("Error creating custom path for Computer Inventory Collection Settings: %s", err)
	}

	fmt.Printf("Custom path created: %+v\n", response)
}
