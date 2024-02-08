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

	// Example device ID and subset
	deviceName := "iPad" // Replace with an actual device name
	subset := "General"  // Replace with the desired subset

	// Get mobile device by ID and subset
	deviceSubset, err := client.GetMobileDeviceByNameAndDataSubset(deviceName, subset)
	if err != nil {
		log.Fatalf("Error fetching mobile device by ID and subset: %v", err)
	}

	// Pretty print the device subset data in XML
	deviceSubsetXML, err := xml.MarshalIndent(deviceSubset, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling device subset data: %v", err)
	}
	fmt.Println("Device Subset Data:\n", string(deviceSubsetXML))
}
