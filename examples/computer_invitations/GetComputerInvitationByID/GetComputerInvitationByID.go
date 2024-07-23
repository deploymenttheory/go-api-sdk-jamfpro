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

	invitationID := "1" // Replace with the actual invitation ID you want to retrieve

	invitation, err := client.GetComputerInvitationByID(invitationID)
	if err != nil {
		fmt.Printf("Error fetching computer invitation by ID: %s\n", err)
		return
	}

	fmt.Printf("Invitation Details: %+v\n", invitation)
}
