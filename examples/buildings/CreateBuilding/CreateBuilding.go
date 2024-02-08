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
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Building details to be created
	newBuilding := &jamfpro.ResourceBuilding{
		Name:           "Apple Park",
		StreetAddress1: "The McIntosh Tree",
		StreetAddress2: "One Apple Park Way",
		City:           "Cupertino",
		StateProvince:  "California",
		ZipPostalCode:  "95014",
		Country:        "The United States of America",
	}

	// Create the building
	createdBuilding, err := client.CreateBuilding(newBuilding)
	if err != nil {
		log.Fatalf("Error creating building: %v", err)
	}

	// Print the details of the created building
	fmt.Printf("Created Building: %+v\n", createdBuilding)
}
