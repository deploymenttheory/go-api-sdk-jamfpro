package main

import (
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

	// Construct the mobile device object
	mobileDevice := &jamfpro.ResourceMobileDevice{
		General: jamfpro.MobileDeviceSubsetGeneral{
			DisplayName:  "iPad",
			DeviceName:   "iPad",
			Name:         "iPad",
			AssetTag:     "string",
			SerialNumber: "C02Q7KHTGFW2",
			UDID:         "270aae10800b6e61a2ee2bbc285eb967050b6112",
		},
	}

	// Create the mobile device in Jamf Pro
	responseDevice, err := client.CreateMobileDevice(mobileDevice)
	if err != nil {
		log.Fatalf("Failed to create mobile device: %v", err)
	}

	// Print the response
	fmt.Printf("Created Mobile Device: %+v\n", responseDevice)
}
