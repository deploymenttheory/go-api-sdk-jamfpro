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

	// Define the name of the macOS Configuration Profile you want to delete
	VPPMacApplicationName := "TextWrangler.app"

	// Call the DeleteMacApplicationByName function
	err = client.DeleteMacApplicationByName(VPPMacApplicationName)
	if err != nil {
		log.Fatalf("Failed to delete VPP Mac Application with name '%s': %v", VPPMacApplicationName, err)
	}

	fmt.Printf("VPP Mac Application with name '%s' deleted successfully\n", VPPMacApplicationName)
}
