package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logLevel := http_client.LogLevelWarning // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := http_client.Config{
		InstanceName: authConfig.InstanceName,
		Auth: http_client.AuthConfig{
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

	// Define the ID of the group you want to update
	groupName := "Mobile Smart Group" // Replace with the actual group ID

	// Define the updated group data for a static group
	updatedStaticGroup := &jamfpro.ResourceMobileDeviceGroup{
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

	// Call the UpdateMobileDeviceGroupByID function
	updatedGroup, err := client.UpdateMobileDeviceGroupByName(groupName, updatedStaticGroup)
	if err != nil {
		log.Fatalf("Error updating mobile device group: %s\n", err)
	}

	// Marshal and print the updated group
	updatedGroupXML, err := xml.MarshalIndent(updatedGroup, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling updated group data: %v", err)
	}
	fmt.Println("Updated Mobile Device Group:\n", string(updatedGroupXML))
}
