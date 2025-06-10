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

	// Define new SMTP server settings
	newSMTPSettings := &jamfpro.ResourceSMTPServer{
		Enabled:            true,
		AuthenticationType: "BASIC", // Can be NONE, BASIC, GRAPH_API, or GOOGLE_MAIL

		ConnectionSettings: &jamfpro.ResourceSMTPServerConnectionSettings{
			Host:              "smtp.sendgrid.net",
			Port:              587,
			EncryptionType:    "TLS_1_2",
			ConnectionTimeout: 5,
		},

		SenderSettings: &jamfpro.ResourceSMTPServerSenderSettings{
			DisplayName:  "Jamf Pro Server",
			EmailAddress: "user@company.com",
		},

		// Include the appropriate credentials based on AuthenticationType
		BasicAuthCredentials: &jamfpro.ResourceSMTPServerBasicAuthCredentials{
			Username: "sample-username",
			Password: "password",
		},
	}

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
	fmt.Println("Updated SMTP Server Settings:\n", string(smtpInfoJSON))

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
