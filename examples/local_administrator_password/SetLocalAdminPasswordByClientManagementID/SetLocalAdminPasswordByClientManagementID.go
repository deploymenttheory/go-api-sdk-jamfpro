package main

import (
	"encoding/json"
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

	// Example management ID
	clientManagementID := "2db90ebf-ce9c-4078-b508-034c8ee3a060" // Replace with actual ID

	// Create password list
	passwordList := &jamfpro.ResourceLapsPasswordList{
		LapsUserPasswordList: []jamfpro.LapsUserPassword{
			{
				Username: "admin",
				Password: "NewSecurePassword123!",
			},
			{
				Username: "ladmin",
				Password: "AnotherSecurePassword456!",
			},
		},
	}

	// Set the new passwords
	response, err := client.SetLocalAdminPasswordByClientManagementID(clientManagementID, passwordList)
	if err != nil {
		log.Fatalf("Error setting LAPS passwords: %v", err)
	}

	// Pretty print the response
	prettyResponse, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling response: %v", err)
	}

	fmt.Printf("Successfully set LAPS passwords. Response:\n%s\n", string(prettyResponse))
}
