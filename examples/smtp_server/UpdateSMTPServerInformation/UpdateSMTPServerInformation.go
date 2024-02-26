package main

import (
	"encoding/xml"
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
		Enabled:               true,
		Host:                  "smtp.sendgrid.net",
		Port:                  587,
		Timeout:               5,
		AuthorizationRequired: true,
		Username:              "sample-username",
		Password:              "sample-password",
		SSL:                   false,
		TLS:                   false,
		SendFromName:          "Jamf Pro Server",
		SendFromEmail:         "user@company.com",
	}

	// Call the UpdateSMTPServerInformation function
	err = client.UpdateSMTPServerInformation(newSMTPSettings)
	if err != nil {
		log.Fatalf("Error updating SMTP server information: %v", err)
	}

	// Pretty print the details in XML
	smtpInfoXML, err := xml.MarshalIndent(newSMTPSettings, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling server data: %v", err)
	}
	fmt.Println("Created Script Details:\n", string(smtpInfoXML))
}
