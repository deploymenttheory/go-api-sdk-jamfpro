package main

import (
	"encoding/xml"
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

	// Example ID of the user extension attribute to retrieve
	attributeID := 1

	// Fetch user extension attribute by ID
	attribute, err := client.GetUserExtensionAttributeByID(attributeID)
	if err != nil {
		log.Fatalf("Error fetching user extension attribute by ID: %v", err)
	}

	// Pretty print the attribute details in XML
	attributeXML, err := xml.MarshalIndent(attribute, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling attribute data: %v", err)
	}
	fmt.Println("User Extension Attribute Details:\n", string(attributeXML))
}
