package main

import (
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

	// Example: Fetch the resource history of a building by ID
	buildingID := "" // Replace with a real building ID
	history, err := client.GetBuildingResourceHistoryByID(buildingID, "")
	if err != nil {
		log.Fatalf("Error fetching building resource history: %v", err)
	}

	// Print the resource history
	fmt.Printf("Resource History for Building ID %s:\n", buildingID)
	for _, record := range history.Results {
		fmt.Printf("ID: %d, Username: %s, Date: %s, Note: %s, Details: %s\n",
			record.ID, record.Username, record.Date, record.Note, record.Details)
	}
}
