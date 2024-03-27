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
		Enabled:                true,
		Server:                 "smtp.sendgrid.net",
		Port:                   587,
		EncryptionType:         "", // Specify the encryption type if applicable, for example "SSL" or "TLS"
		ConnectionTimeout:      5,
		SenderDisplayName:      "Jamf Pro Server",
		SenderEmailAddress:     "user@company.com",
		RequiresAuthentication: true,
		Username:               "sample-username",
	}

	// Call the UpdateSMTPServerInformation function
	err = client.UpdateSMTPServerInformation(newSMTPSettings)
	if err != nil {
		log.Fatalf("Error updating SMTP server information: %v", err)
	}

	// Pretty print the details in XML
	smtpInfoJSON, err := json.MarshalIndent(newSMTPSettings, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling server data: %v", err)
	}
	fmt.Println("Created Script Details:\n", string(smtpInfoJSON))
}
