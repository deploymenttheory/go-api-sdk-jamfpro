package main

import (
	"encoding/xml"
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

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define new SMTP server settings
	newSMTPSettings := &jamfpro.ResourceSMTPServer{
		Enabled:               true,
		Host:                  "smtp.sendgrid.net",
		Port:                  587,
		Timeout:               5,
		AuthorizationRequired: true,
		Username:              "sample-username",
		Password:              "sample-password",
		SSL:                   false,
		TLS:                   false,
		SendFromName:          "Jamf Pro Server",
		SendFromEmail:         "user@company.com",
	}

	// Call the UpdateSMTPServerInformation function
	err = client.UpdateSMTPServerInformation(newSMTPSettings)
	if err != nil {
		log.Fatalf("Error updating SMTP server information: %v", err)
	}

	// Pretty print the details in XML
	smtpInfoXML, err := xml.MarshalIndent(newSMTPSettings, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling server data: %v", err)
	}
	fmt.Println("Created Script Details:\n", string(smtpInfoXML))
}
