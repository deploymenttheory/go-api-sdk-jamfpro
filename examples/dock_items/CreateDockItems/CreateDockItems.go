package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file for OAuth credentials
	configFilePath := "D:\\github\\go-api-sdk-jamfpro\\go-api-sdk-jamfpro\\clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelInfo // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName:       authConfig.InstanceName,
		OverrideBaseDomain: authConfig.OverrideBaseDomain,
		LogLevel:           logLevel,
		Logger:             logger,
		ClientID:           authConfig.ClientID,
		ClientSecret:       authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the details of the dock item to create
	newDockItem := &jamfpro.ResourceDockItem{
		Name:     "Safari6",
		Type:     "App",
		Path:     "file://localhost/Applications/Safari.app/",
		Contents: "string",
	}

	// Call the CreateDockItem function
	result, err := client.CreateDockItem(newDockItem)
	if err != nil {
		log.Fatalf("Error creating dock item: %v", err)
	}

	// Output the result
	fmt.Printf("Created Dock Item: %+v\n", result)
}
