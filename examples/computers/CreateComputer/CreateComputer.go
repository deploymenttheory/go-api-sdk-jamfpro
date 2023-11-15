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
	logLevel := http_client.LogLevelDebug

	// Configuration for the jamfpro client
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		LogLevel:     logLevel,
		Logger:       logger,
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Create a new computer configuration
	// Create a new computer configuration
	computer := jamfpro.Computer{
		General: jamfpro.ComputerDataSubsetGeneral{
			Name:         "admin’s MacBook Pro",
			SerialNumber: "C02Q7KHTGFWD",
			UDID:         "EBBFF74D-C6B7-5589-93A9-19E8BDFEDE33",
		},
		// Other fields like Location, Purchasing, etc. can be set similarly
	}

	// Call CreateComputer function
	createdComputer, err := client.CreateComputer(computer)
	if err != nil {
		log.Fatalf("Error creating computer: %v", err)
	}

	// Output the result
	fmt.Printf("Created Computer: %+v\n", createdComputer)
}
