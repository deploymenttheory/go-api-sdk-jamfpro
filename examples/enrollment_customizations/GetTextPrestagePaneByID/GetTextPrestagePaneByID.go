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

	// Specify the enrollment customization ID and panel ID
	customizationID := "22" // Replace with your actual customization ID
	paneID := "17"          // Replace with your actual pane ID

	// Get the text prestage pane
	textPane, err := client.GetTextPrestagePaneByID(customizationID, paneID)
	if err != nil {
		log.Fatalf("Failed to get text prestage pane: %v", err)
	}

	// Pretty print the result in JSON
	prettyJSON, err := json.MarshalIndent(textPane, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling result: %v", err)
	}
	fmt.Println("Text Prestage Pane Details:\n", string(prettyJSON))

}
