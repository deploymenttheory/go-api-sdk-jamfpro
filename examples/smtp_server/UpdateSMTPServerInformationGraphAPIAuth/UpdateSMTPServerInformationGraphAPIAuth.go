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

	// Define new SMTP server settings for Graph API
	newSMTPSettings := &jamfpro.ResourceSMTPServer{
		Enabled:            true,
		AuthenticationType: "GRAPH_API",
		SenderSettings: &jamfpro.ResourceSMTPServerSenderSettings{
			EmailAddress: "noreply@yourdomain.onmicrosoft.com",
		},
		GraphApiCredentials: &jamfpro.ResourceSMTPServerGraphApiCredentials{
			TenantId:     "c84b7b82-c277-411b-975d-7431b4ce40ac",
			ClientId:     "5294f9d1-f723-419c-93db-ff040bf7c947",
			ClientSecret: "password",
		},
	}

	// Print the request JSON before sending
	requestJSON, err := json.MarshalIndent(newSMTPSettings, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling request data: %v", err)
	}
	fmt.Println("Request Payload:\n", string(requestJSON))

	// Call the UpdateSMTPServerInformation function
	updatedSettings, err := client.UpdateSMTPServerInformation(newSMTPSettings)
	if err != nil {
		log.Fatalf("Error updating SMTP server information: %v", err)
	}

	// Pretty print the updated settings in JSON
	smtpInfoJSON, err := json.MarshalIndent(updatedSettings, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling server data: %v", err)
	}
	fmt.Println("\nUpdated SMTP Server Settings:\n", string(smtpInfoJSON))

	// Get and display current SMTP settings
	currentSettings, err := client.GetSMTPServerInformation()
	if err != nil {
		log.Fatalf("Error getting SMTP server information: %v", err)
	}

	// Pretty print the current settings
	currentSettingsJSON, err := json.MarshalIndent(currentSettings, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling current settings: %v", err)
	}
	fmt.Println("\nCurrent SMTP Server Settings:\n", string(currentSettingsJSON))
}
