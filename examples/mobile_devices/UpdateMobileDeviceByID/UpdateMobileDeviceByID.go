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

	// Construct the mobile device object
	mobileDevice := &jamfpro.ResourceMobileDevice{
		General: jamfpro.MobileDeviceSubsetGeneral{
			DisplayName:  "iPad",
			DeviceName:   "iPad",
			Name:         "iPad",
			AssetTag:     "string",
			SerialNumber: "C02Q7KHTGFWD",
			UDID:         "270aae10800b6e61a2ee2bbc285eb967050b6984",
		},
	}

	mobileDeviceID := 1

	// Create the mobile device in Jamf Pro
	responseDevice, err := client.UpdateMobileDeviceByID(mobileDeviceID, mobileDevice)
	if err != nil {
		log.Fatalf("Failed to create mobile device: %v", err)
	}

	// Print the response
	fmt.Printf("Created Mobile Device: %+v\n", responseDevice)
}
