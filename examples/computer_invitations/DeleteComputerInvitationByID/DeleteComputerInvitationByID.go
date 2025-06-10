package main

import (
	"fmt"
	"log"
	"os"

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

	// The ID of the computer invitation you wish to delete.
	invitationID := "1" // Replace with the actual ID.

	// Call the function to delete the computer invitation by ID.
	err = client.DeleteComputerInvitationByID(invitationID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error deleting computer invitation by ID: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Computer invitation deleted successfully.")
}
