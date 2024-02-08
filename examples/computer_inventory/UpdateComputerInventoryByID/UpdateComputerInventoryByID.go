package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-http-client/logger"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

//go:embed payload.json
var fsys embed.FS

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logLevel := logger.LogLevelDebug // Set the desired log level

	// Configuration for the jamfpro
	config := httpclient.Config{
		InstanceName: authConfig.InstanceName,
		Auth: httpclient.AuthConfig{
			ClientID:     authConfig.ClientID,
			ClientSecret: authConfig.ClientSecret,
		},
		LogLevel: logLevel,
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
