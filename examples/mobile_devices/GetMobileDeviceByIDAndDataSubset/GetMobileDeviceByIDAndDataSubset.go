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

	// Example device ID and subset
	deviceID := 1       // Replace with an actual device ID
	subset := "General" // Replace with the desired subset

	// Get mobile device by ID and subset
	deviceSubset, err := client.GetMobileDeviceByIDAndDataSubset(deviceID, subset)
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
