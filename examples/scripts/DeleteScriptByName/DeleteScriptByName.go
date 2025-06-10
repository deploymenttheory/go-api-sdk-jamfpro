package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "./clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	scriptNameToDelete := "Script Name" // Replace with the actual script name

	err = client.DeleteScriptByName(scriptNameToDelete)
	if err != nil {
		log.Fatalf("Error deleting script by name: %v", err)
	}

	fmt.Printf("Script with name '%s' was successfully deleted.\n", scriptNameToDelete)
}
