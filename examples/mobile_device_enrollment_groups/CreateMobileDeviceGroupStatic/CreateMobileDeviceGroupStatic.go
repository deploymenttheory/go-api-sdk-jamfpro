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
