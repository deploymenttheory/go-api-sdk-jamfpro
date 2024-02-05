package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	// Import the http_client package for logging
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/joseph/github/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logLevel := http_client.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := http_client.Config{
		InstanceName: "lbgsandbox",
		Auth: http_client.AuthConfig{
			InstanceName: "lbgsandbox",
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

	fmt.Println(client.HTTP.Token)

	departmentID := "24979"

	department, err := client.GetDepartmentByID(departmentID)
	if err != nil {
		log.Fatalf("Error fetching department by ID: %v", err)
	}

	departmentJSON, err := json.MarshalIndent(department, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling department data: %v", err)
	}
	fmt.Println("Fetched Department:\n", string(departmentJSON))
}
