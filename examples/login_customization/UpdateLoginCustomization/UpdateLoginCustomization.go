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

	// Define new login customization settings
	newLoginCustomization := &jamfpro.ResourceLoginCustomization{
		RampInstance:            true,
		IncludeCustomDisclaimer: true,
		DisclaimerHeading:       "Updated Disclaimer Header",
		DisclaimerMainText:      "Updated disclaimer main text",
		ActionText:              "Accept",
	}

	// Update the login customization settings
	updatedLoginCustomization, err := client.UpdateLoginCustomization(newLoginCustomization)
	if err != nil {
		log.Fatalf("Error updating login customization: %v", err)
	}

	// Pretty print the updated login customization settings
	updatedResponseJSON, err := json.MarshalIndent(updatedLoginCustomization, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling updated login customization data: %v", err)
	}
	fmt.Println("Updated login customization details:", string(updatedResponseJSON))
}
