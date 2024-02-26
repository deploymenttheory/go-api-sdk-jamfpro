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

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Update the activation code
	err = client.UpdateActivationCode("Organization Name", "ABCD-1234-A1B2-AAAA-BBBB-CCCC-DDDD-EEEE")
	if err != nil {
		fmt.Printf("Error updating activation code: %v\n", err)
		return
	}

	fmt.Println("Activation code updated successfully")
}
