package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-http-client/logger"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

const (
	profileName = "accessibility"
)

func main() {
	// Define the path to the JSON configuration file inside the main function
	configFilePath := "/Users/dafyddwatkins/localtesting/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logLevel := logger.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := httpclient.Config{
		InstanceName: authConfig.InstanceName,
		Auth: httpclient.AuthConfig{
			ClientID:     authConfig.ClientID,
			ClientSecret: authConfig.ClientSecret,
		},
		LogLevel: logLevel,
	}

	// Create a new jamfpro client instanceclient,
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Call GetMacOSConfigurationProfileByName function
	profile, err := client.GetMacOSConfigurationProfileByName(profileName)
	if err != nil {
		log.Fatalf("Error fetching macOS Configuration Profile by Name: %v", err)
	}

	// Pretty print the profile in XML
	profileXML, err := xml.MarshalIndent(profile, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling macOS Configuration Profile data: %v", err)
	}
	fmt.Println("Fetched macOS Configuration Profile:\n", string(profileXML))
}
