package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}
	// Define the name of the group you want to update
	groupID := 1 // Replace with the actual group name

	// Define the updated group data with mobile device deletions
	updatedGroup := &jamfpro.ResourceMobileDeviceGroup{
		MobileDeviceDeletions: []jamfpro.MobileDeviceGroupSubsetDeviceItem{
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

	// Call the UpdateMobileDeviceGroupByName function
	updatedGroupResponse, err := client.UpdateMobileDeviceGroupByID(groupID, updatedGroup)
	if err != nil {
		log.Fatalf("Error updating mobile device group: %s\n", err)
	}

	// Marshal and print the updated group
	updatedGroupXML, err := xml.MarshalIndent(updatedGroupResponse, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling updated group data: %v", err)
	}
	fmt.Println("Updated Mobile Device Group:\n", string(updatedGroupXML))
}
