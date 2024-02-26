package main

import (
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

	// Example device Name to delete
	deviceName := "iPad" // Replace with an actual device Name

	// Delete mobile device by Name
	err = client.DeleteMobileDeviceByName(deviceName)
	if err != nil {
		log.Fatalf("Error deleting mobile device by Name: %v", err)
	} else {
		log.Printf("Mobile device with Name %s has been successfully deleted.\n", deviceName)
	}
}
