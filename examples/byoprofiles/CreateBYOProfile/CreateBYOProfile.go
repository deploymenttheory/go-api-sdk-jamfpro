package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logLevel := http_client.LogLevelWarning // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := http_client.Config{
		InstanceName: authConfig.InstanceName,
		Auth: http_client.AuthConfig{
			ClientID:     authConfig.ClientID,
			ClientSecret: authConfig.ClientSecret,
		},
		LogLevel: logLevel,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Create a BYOProfile structure to send
	newProfile := jamfpro.ResourceBYOProfile{
		General: jamfpro.BYOProfileSubsetGeneral{
			Name:        "Personal Device Profile with jamf pro sdk",
			Site:        jamfpro.SharedResourceSite{ID: -1, Name: "None"},
			Enabled:     true,
			Description: "Used for Android or iOS BYO device enrollments",
		},
	}

	// Convert the profile to XML to see the output (optional, for debug purposes)
	xmlData, err := xml.MarshalIndent(newProfile, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling XML: %v", err)
	}
	fmt.Printf("XML Request: %s\n", xmlData)

	// Now call the create function
	createdProfile, err := client.CreateBYOProfile(&newProfile)
	if err != nil {
		log.Fatalf("Error creating BYO Profile: %v", err)
	}
	fmt.Printf("Created BYO Profile: %+v\n", createdProfile)
}
