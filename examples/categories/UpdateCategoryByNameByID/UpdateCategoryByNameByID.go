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
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the category name you want to update and the updated category details
	categoryName := "Existing Category Name" // Replace with the actual category name you want to update
	updatedCategory := &jamfpro.ResourceCategory{
		Name:     "Updated Category Name", // Replace with the updated name
		Priority: 10,                      // Replace with the updated priority
	}

	// Call UpdateCategoryByNameByID function
	updatedCategoryResult, err := client.UpdateCategoryByName(categoryName, updatedCategory)
	if err != nil {
		log.Fatalf("Error updating category: %v", err)
	}

	// Pretty print the updated category in JSON
	categoryJSON, err := json.MarshalIndent(updatedCategoryResult, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling updated category data: %v", err)
	}
	fmt.Println("Updated Category:\n", string(categoryJSON))
}
