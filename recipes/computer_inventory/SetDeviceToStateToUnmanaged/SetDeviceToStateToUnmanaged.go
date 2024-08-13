package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

// Define the device ID
var deviceID = "8675309"

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Retrieve the computer inventory by ID
	computerInventory, err := client.GetComputerInventoryByID(deviceID)
	if err != nil {
		log.Fatalf("Failed to retrieve computer inventory: %v", err)
	}

	// Set the managed state to false
	computerInventory.General.RemoteManagement.Managed = false

	// Update the computer inventory with the modified managed state
	updatedInventory, err := client.UpdateComputerInventoryByID(deviceID, computerInventory)
	if err != nil {
		log.Fatalf("Failed to update computer inventory: %v", err)
	}

	// Print the updated managed state to confirm the change
	fmt.Printf("Updated managed state: %v\n", updatedInventory.General.RemoteManagement.Managed)
}
