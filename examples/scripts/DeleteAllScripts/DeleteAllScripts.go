package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client" // Import http_client for logging
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
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

	// Fetch all scripts
	scripts, err := client.GetScripts("")
	if err != nil {
		log.Fatalf("Error fetching scripts: %v", err)
	}

	fmt.Println("Scripts fetched. Starting deletion process:")

	// Iterate over each script and delete
	for _, script := range scripts.Results {
		fmt.Printf("Deleting script ID: %s, Name: %s\n", script.ID, script.Name)

		err = client.DeleteScriptByID(script.ID)
		if err != nil {
			log.Printf("Error deleting script ID %s: %v\n", script.ID, err)
			continue // Move to the next script if there's an error
		}

		fmt.Printf("Script ID %s deleted successfully.\n", script.ID)
	}

	fmt.Println("Script deletion process completed.")
}
