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

	// Generate a new OIDC keystore
	err = client.GenerateKeystoreForOIDCMessages()
	if err != nil {
		log.Fatalf("Error generating OIDC keystore: %v", err)
	}

	fmt.Println("Successfully generated a new OIDC keystore.")
}
