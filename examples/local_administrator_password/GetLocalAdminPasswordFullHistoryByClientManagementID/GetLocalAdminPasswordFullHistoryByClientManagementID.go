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

	// Example value for clientManagementID
	clientManagementID := "2db90ebf-ce9c-4078-b508-034c8ee3a060" // Replace with actual client management ID

	// Call Function to get full LAPS history
	fullHistory, err := client.GetLocalAdminPasswordFullHistoryByClientManagementID(clientManagementID)
	if err != nil {
		log.Fatalf("Error fetching full LAPS history: %v", err)
	}

	// Pretty print the JSON
	response, err := json.MarshalIndent(fullHistory, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling full LAPS history data: %v", err)
	}
	fmt.Printf("Fetched full LAPS history for device %s:\n%s\n",
		clientManagementID, string(response))
}
