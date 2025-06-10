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

	// Example computer name
	computerName := "MyMacBook" // Replace with actual computer name

	// Get computer inventory
	inventory, err := client.GetComputerInventoryByName(computerName)
	if err != nil {
		log.Fatalf("Failed to get computer inventory for %s: %v", computerName, err)
	}

	// Extract management ID
	managementID := inventory.General.ManagementId
	if managementID == "" {
		log.Fatalf("No management ID found for computer %s", computerName)
	}

	// Get LAPS capable accounts using the management ID
	capableAccounts, err := client.GetLocalAdminPasswordCapableAccountsByClientManagementID(managementID)
	if err != nil {
		log.Fatalf("Error fetching LAPS capable accounts: %v", err)
	}

	// Pretty print the JSON
	response, err := json.MarshalIndent(capableAccounts, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling LAPS capable accounts data: %v", err)
	}

	// Print results
	fmt.Printf("Computer Name: %s\n", computerName)
	fmt.Printf("Management ID: %s\n", managementID)
	fmt.Printf("LAPS Capable Accounts:\n%s\n", string(response))
}
