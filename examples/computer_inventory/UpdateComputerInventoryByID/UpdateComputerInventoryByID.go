package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

//go:embed payload.json
var fsys embed.FS

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
	// Define the computer ID you want to update
	computerID := "8" // Replace with the actual computer ID

	// Read the content of the embedded file
	data, err := fsys.ReadFile("payload.json")
	if err != nil {
		fmt.Println("Error reading embedded file:", err)
		return
	}

	var payload jamfpro.ResourceComputerInventory
	if err := json.Unmarshal(data, &payload); err != nil {
		fmt.Fprintf(os.Stderr, "Error decoding payload: %v\n", err)
		return
	}

	// Call the UpdateComputerInventoryByID function
	updatedInventory, err := client.UpdateComputerInventoryByID(computerID, &payload)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error updating computer inventory: %v\n", err)
		return
	}

	// Print the updated inventory
	fmt.Printf("Updated Inventory: %+v\n", updatedInventory)
}
