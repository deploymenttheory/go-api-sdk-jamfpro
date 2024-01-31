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

	// Fetch all printers
	printers, err := client.GetPrinters()
	if err != nil {
		log.Fatalf("Error fetching printers: %v", err)
	}

	fmt.Println("Printers fetched. Starting deletion process:")

	// Iterate over each printer and delete
	for _, printer := range printers.Printer {
		fmt.Printf("deleting printer ID: %d, Name: %s\n", printer.ID, printer.Name)

		err = client.DeletePrinterByID(printer.ID)
		if err != nil {
			log.Printf("error deleting printer ID %d: %v\n", printer.ID, err)
			continue // Move to the next printer if there's an error
		}

		fmt.Printf("printer ID %d deleted successfully.\n", printer.ID)
	}

	fmt.Println("Printer deletion process completed.")
}
