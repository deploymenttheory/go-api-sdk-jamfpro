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

	// Specify the enrollment customization ID and panel ID
	customizationID := "22" // Replace with your actual customization ID
	paneID := "19"          // Replace with your actual panel ID

	// Prepare the updated SSO pane settings
	updatedSSOPane := jamfpro.ResourceEnrollmentCustomizationSSOPane{
		DisplayName:                    "Updated SSO Pane",
		Rank:                           2,
		IsGroupEnrollmentAccessEnabled: true,
		GroupEnrollmentAccessName:      "Enrollment Users",
		IsUseJamfConnect:               true,
		ShortNameAttribute:             "sAMAccountName",
		LongNameAttribute:              "displayName",
	}

	// Update the SSO prestage pane
	result, err := client.UpdateSSOPrestagePaneByID(customizationID, paneID, updatedSSOPane)
	if err != nil {
		log.Fatalf("Failed to update SSO prestage pane: %v", err)
	}

	// Pretty print the result in JSON
	prettyJSON, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling result: %v", err)
	}
	fmt.Println("Updated SSO Prestage Pane:\n", string(prettyJSON))
}
