package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "./clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	macAddressID := "1" // Replace with the actual MAC address ID

	err = client.DeleteRemovableMACAddressByID(macAddressID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("MAC Address successfully deleted.")
}
