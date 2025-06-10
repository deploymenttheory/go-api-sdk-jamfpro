package main

import (
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

	// Construct the invitation object based on the example XML provided
	invitation := &jamfpro.ResourceComputerInvitation{
		InvitationType:              "DEFAULT",
		ExpirationDate:              "2024-12-07 11:13:35",
		SSHUsername:                 "jamfadmin",
		SSHPassword:                 "accountpassword",
		MultipleUsersAllowed:        false,
		CreateAccountIfDoesNotExist: false,
		HideAccount:                 false,
		LockDownSSH:                 false,
		EnrollIntoSite: jamfpro.ComputerInvitationSubsetEnrollIntoState{
			ID:   -1,
			Name: "None",
		},
		KeepExistingSiteMembership: false,
		Site: jamfpro.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
	}

	// Call the CreateComputerInvitation function with the constructed invitation
	createdInvitation, err := client.CreateComputerInvitation(invitation)
	if err != nil {
		fmt.Printf("Error creating computer invitation: %s\n", err)
	} else {
		fmt.Printf("Created computer invitation: %+v\n", createdInvitation)
	}
}
