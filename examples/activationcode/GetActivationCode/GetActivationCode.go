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

	// Get the activation code
	activationCode, err := client.GetActivationCode()
	if err != nil {
		fmt.Printf("Error getting activation code: %v\n", err)
		return
	}

	fmt.Printf("Organization Name: %s\n", activationCode.OrganizationName)
	fmt.Printf("Activation Code: %s\n", activationCode.Code)
}
