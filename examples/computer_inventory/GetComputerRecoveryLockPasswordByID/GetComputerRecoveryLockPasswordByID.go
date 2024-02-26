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

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the ID of the computer
	computerID := "8"

	// Call the GetComputerRecoveryLockPasswordByID function
	recoveryLockPasswordResponse, err := client.GetComputerRecoveryLockPasswordByID(computerID)
	if err != nil {
		log.Fatalf("Error fetching Recovery Lock password by ID: %v", err)
	}

	// Print the Recovery Lock password
	fmt.Printf("Recovery Lock Password: %s\n", recoveryLockPasswordResponse.RecoveryLockPassword)
}
