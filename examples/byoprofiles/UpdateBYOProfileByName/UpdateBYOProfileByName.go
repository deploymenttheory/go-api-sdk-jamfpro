package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client" // Import http_client for logging
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

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
	logLevel := http_client.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

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

	// Create a BYOProfile structure to send
	updatedProfile := jamfpro.ResponseBYOProfile{
		General: jamfpro.BYOProfileGeneralInfo{
			Name:        "Personal Device Profile Updated",
			Site:        jamfpro.BYOProfileSiteInfo{ID: -1, Name: "None"},
			Enabled:     true,
			Description: "Updated description for BYO device enrollments",
		},
	}

	profileName := "Personal Device Profile" // Use the actual name of the profile to be updated

	// Convert the profile to XML to see the output (optional, for debug purposes)
	xmlData, err := xml.MarshalIndent(updatedProfile, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling XML: %v", err)
	}
	fmt.Printf("XML Request: %s\n", xmlData)

	// Now call the update function
	updatedProfileResp, err := client.UpdateBYOProfileByName(profileName, &updatedProfile)
	if err != nil {
		log.Fatalf("Error updating BYO Profile by Name: %v", err)
	}
	fmt.Printf("Updated BYO Profile: %+v\n", updatedProfileResp)
}
