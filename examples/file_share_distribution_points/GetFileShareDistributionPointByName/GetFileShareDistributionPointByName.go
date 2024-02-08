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
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Name of the distribution point to fetch
	distributionPointName := "New York Share" // Replace with the actual name

	// Call GetDistributionPointByName function
	distributionPoint, err := client.GetDistributionPointByName(distributionPointName)
	if err != nil {
		log.Fatalf("Error fetching distribution point: %v", err)
	}

	// Pretty print the distribution point in XML
	distributionPointXML, err := xml.MarshalIndent(distributionPoint, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling distribution point data: %v", err)
	}
	fmt.Println("Fetched Distribution Point:\n", string(distributionPointXML))
}
