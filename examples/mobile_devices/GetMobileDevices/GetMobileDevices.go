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

	// Fetch mobile devices
	mobileDevices, err := client.GetMobileDevices()
	if err != nil {
		log.Fatalf("Error fetching mobile devices: %s\n", err)
	}

	// Print the mobile devices in a formatted XML
	mobileDevicesXML, err := xml.MarshalIndent(mobileDevices, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling mobile devices data: %v", err)
	}
	fmt.Println("Mobile Devices:\n", string(mobileDevicesXML))
}
