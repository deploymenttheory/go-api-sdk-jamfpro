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

	// Define the updated group data for a static group
	newGroup := &jamfpro.ResourceMobileDeviceGroup{
		Name:    "Static Group",
		IsSmart: false,
		Site: jamfpro.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
		MobileDevices: []jamfpro.MobileDeviceGroupSubsetDeviceItem{
			{
				ID:             38,
				Name:           "Test Device",
				MacAddress:     "18:E7:F4:35:2F:99",
				UDID:           "141f86e409a5a614a7ef691cd3c6b0821e7d9a98",
				WifiMacAddress: "18:E7:F4:35:2F:99",
				SerialNumber:   "C02Q7KHTGFWF",
			},
			{
				ID:             44,
				Name:           "Device Name",
				MacAddress:     "E0:AC:CB:97:36:G4",
				UDID:           "270aae10800b6e61a2ee2bbc285eb977050b5989",
				WifiMacAddress: "E0:AC:CB:97:36:G4",
				SerialNumber:   "C02Q7KHTGFWF",
			},
		},
	}

	createdGroup, err := client.CreateMobileDeviceGroup(newGroup)
	if err != nil {
		log.Fatalf("Error creating mobile device group: %s\n", err)
	}

	createdGroupXML, err := xml.MarshalIndent(createdGroup, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created group data: %v", err)
	}
	fmt.Println("Created Mobile Device Group:\n", string(createdGroupXML))
}
