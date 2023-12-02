package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

//go:embed payload.json
var fsys embed.FS

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug // Set the desired log level

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName:       authConfig.InstanceName,
		OverrideBaseDomain: authConfig.OverrideBaseDomain,
		LogLevel:           logLevel,
		Logger:             logger,
		ClientID:           authConfig.ClientID,
		ClientSecret:       authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
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

	var payload jamfpro.ResponseComputerInventory
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
