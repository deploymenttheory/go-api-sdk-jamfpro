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

	// Specify the enrollment customization ID to get prestage panes for
	customizationID := "14" // Replace with your actual customization ID

	// Get all prestage panes for the specified enrollment customization
	prestagePanes, err := client.GetPrestagePanes(customizationID)
	if err != nil {
		log.Fatalf("Failed to get prestage panes: %v", err)
	}

	// Pretty print the prestage panes details in JSON
	prettyJSON, err := json.MarshalIndent(prestagePanes, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling prestage panes data: %v", err)
	}
	fmt.Println("Prestage Panes List:\n", string(prettyJSON))

	// Print total count
	fmt.Printf("\nTotal number of prestage panes: %d\n", len(prestagePanes.Panels))

	// Iterate through the prestage panes to display their details
	fmt.Println("\nPrestage Pane Details:")
	for _, pane := range prestagePanes.Panels {
		fmt.Printf("- %s (ID: %d, Type: %s, Rank: %d)\n",
			pane.DisplayName,
			pane.ID,
			pane.Type,
			pane.Rank)
	}
}
