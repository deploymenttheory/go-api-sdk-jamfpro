package main

import (
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

	// Replace with the actual invitation ID you want to fetch
	invitationID := "1"

	invitation, err := client.GetComputerInvitationByInvitationID(invitationID)
	if err != nil {
		log.Fatalf("Error fetching computer invitation by ID: %s", err)
	}

	fmt.Printf("Invitation Details: %+v\n", invitation)
}
