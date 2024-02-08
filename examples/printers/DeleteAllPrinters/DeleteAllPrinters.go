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
