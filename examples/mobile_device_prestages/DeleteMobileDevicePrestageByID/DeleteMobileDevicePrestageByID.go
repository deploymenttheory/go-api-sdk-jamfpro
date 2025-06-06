package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/neilmartin/GitHub/go-api-sdk-jamfpro/client_auth.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// The ID of the mobile device prestage you want to delete
	prestageID := "12"

	// Call DeleteMobileDevicePrestageByID to delete the prestage
	err = client.DeleteMobileDevicePrestageByID(prestageID)
	if err != nil {
		log.Fatalf("Error deleting mobile device prestage: %v", err)
	}

	// Print a confirmation message
	fmt.Println("Mobile device prestage deleted successfully.")
}
