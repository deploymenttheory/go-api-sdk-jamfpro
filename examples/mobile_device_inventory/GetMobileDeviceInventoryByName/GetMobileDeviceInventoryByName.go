package main

import (
	"encoding/json"
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

	deviceName := "John's iPhone"

	mobileDeviceInventory, err := client.GetMobileDeviceInventoryByName(deviceName)
	if err != nil {
		log.Fatalf("Error fetching mobile device inventory by name: %v", err)
	}

	fmt.Printf("Device: %s (ID: %s)\n", mobileDeviceInventory.General.DisplayName, mobileDeviceInventory.MobileDeviceId)
	fmt.Printf("Serial: %s\n", mobileDeviceInventory.Hardware.SerialNumber)
	fmt.Printf("Model: %s\n\n", mobileDeviceInventory.Hardware.Model)

	// Pretty print the response
	prettyJSON, err := json.MarshalIndent(mobileDeviceInventory, "", "    ")
	if err != nil {
		log.Fatalf("Failed to generate pretty JSON: %v", err)
	}
	fmt.Printf("%s\n", prettyJSON)
}
