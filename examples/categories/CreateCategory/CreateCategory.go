package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-http-client/logger"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logLevel := logger.LogLevelWarn // LogLevelNone / LogLevelDebug / LogLevelInfo / LogLevelError

	// Configuration for the jamfpro
	config := httpclient.Config{
		InstanceName: authConfig.InstanceName,
		Auth: httpclient.AuthConfig{
			ClientID:     authConfig.ClientID,
			ClientSecret: authConfig.ClientSecret,
		},
		LogLevel: logLevel,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the new category you want to create
	newCategory := &jamfpro.ResourceCategory{
		Name:     "Applications",
		Priority: 9,
	}

	// Call CreateCategory function
	createdCategory, err := client.CreateCategory(newCategory)
	if err != nil {
		log.Fatalf("Error creating category: %v", err)
	}

	// Pretty print the created category in JSON
	categoryJSON, err := json.MarshalIndent(createdCategory, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling category data: %v", err)
	}
	fmt.Println("Created Category:\n", string(categoryJSON))
}
