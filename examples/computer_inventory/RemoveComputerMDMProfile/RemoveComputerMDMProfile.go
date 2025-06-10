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

	// Example: Remove the MDM profile for a computer with a specific ID
	computerID := "9" // Replace with the actual computer ID

	response, err := client.RemoveComputerMDMProfile(computerID)
	if err != nil {
		fmt.Println("Error removing MDM profile:", err)
	} else {
		fmt.Printf("MDM profile removed successfully. Device ID: %s, Command UUID: %s\n", response.DeviceID, response.CommandUUID)
	}
}
