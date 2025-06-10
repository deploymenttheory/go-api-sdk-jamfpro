package main

import (
	"encoding/xml"
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

	// Use the client to call the GetComputerInvitations function.
	invitations, err := client.GetComputerInvitations()
	if err != nil {
		log.Fatalf("Error retrieving computer invitations: %v", err)
	}

	// Output the invitations in a human-readable format.
	// Here we're simply printing the XML representation of the invitations.
	invitationsXML, err := xml.MarshalIndent(invitations, "", "    ")
	if err != nil {
		log.Fatalf("Error marshalling invitations to XML: %v", err)
	}
	fmt.Printf("Computer Invitations:\n%s\n", invitationsXML)
}
