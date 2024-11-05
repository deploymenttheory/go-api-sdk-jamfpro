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

	// Example values for clientManagementID and username
	clientManagementID := "2db90ebf-ce9c-4078-b508-034c8ee3a060" // Replace with device management ID
	username := "admin"                                          // Replace with actual username

	// Call Function to get current LAPS password
	currentPassword, err := client.GetCurrentLocalAdminPasswordForSpecifiedUsernameByClientManagementID(clientManagementID, username)
	if err != nil {
		log.Fatalf("Error fetching current LAPS password: %v", err)
	}

	// Pretty print the JSON
	response, err := json.MarshalIndent(currentPassword, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling current LAPS password data: %v", err)
	}
	fmt.Printf("Fetched current LAPS password for user %s on device %s:\n%s\n",
		username, clientManagementID, string(response))
}
