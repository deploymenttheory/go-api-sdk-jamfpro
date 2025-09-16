package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/Shared/GitHub/go-api-sdk-jamfpro/localtesting/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Disable SSO
	fmt.Println("Disabling SSO...")
	err = client.DisableSso()
	if err != nil {
		fmt.Printf("Error disabling SSO: %v\n", err)
		return
	}

	fmt.Println("SSO has been successfully disabled.")
}
