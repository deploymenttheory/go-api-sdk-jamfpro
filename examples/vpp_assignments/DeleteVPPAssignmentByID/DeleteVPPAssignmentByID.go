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

	// Specify the ID of the VPP assignment to delete
	vppAssignmentID := "1" // Replace with the actual ID

	// Call the DeleteVPPAssignmentByID function
	err = client.DeleteVPPAssignmentByID(vppAssignmentID)
	if err != nil {
		log.Fatalf("Error deleting VPP Assignment by ID: %v", err)
	}

	fmt.Println("VPP Assignment deleted successfully.")
}
